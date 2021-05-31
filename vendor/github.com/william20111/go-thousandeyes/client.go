package thousandeyes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	apiEndpoint = "https://api.thousandeyes.com/v6"
)

var orgRate RateLimit
var instantTestRate RateLimit

// RateLimit contains data representing rate limit headers returned in
// ThousandEyes API responses.  int64 everywhere for ease of interacting
// with time values.
type RateLimit struct {
	Limit              int64
	Remaining          int64
	Reset              int64
	LastRemaining      int64
	ConcurrentMessages []time.Time
}

// APILinks - List of APILink
type APILinks []APILink

// APILink - an api link
type APILink struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"rel,omitempty"`
}

type errorObject struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Limiter - Rate limiter interface
type Limiter interface {
	Wait()
}

//HTTPClient - an http client
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientOptions - Thousandeyes client options for accountID, AuthToken & rate limiter
type ClientOptions struct {
	Limiter   Limiter
	AccountID string
	AuthToken string
	Timeout   time.Duration
}

// Client wraps http client
type Client struct {
	AuthToken      string
	AccountGroupID string
	APIEndpoint    string
	HTTPClient     http.Client
	Limiter        Limiter
}

// DefaultLimiter -  thousandeyes rate limit is 240 per minute
type DefaultLimiter struct{}

// Wait - Satisfying the Limiter interface and wait on 300ms to avoid TE 240 per minute default
func (l DefaultLimiter) Wait() {
	time.Sleep(time.Millisecond * 300)
}

// NewClient creates an API client
func NewClient(opts *ClientOptions) *Client {
	// Set default timeout if a custom duration is 0 or unset (since we
	// can't tell the difference without using an additional value).
	// Overriding a default value of 0 has the side effect of preventing
	// use of http.Client.Timeout behavior of using 0 to mean "no timeout".
	var timeout time.Duration
	if opts.Timeout != 0 {
		timeout = opts.Timeout
	} else {
		// Default timeout
		timeout = time.Second * 20
	}

	return &Client{
		AuthToken:      opts.AuthToken,
		AccountGroupID: opts.AccountID,
		APIEndpoint:    apiEndpoint,
		HTTPClient: http.Client{
			Timeout: timeout,
		},
		Limiter: opts.Limiter,
	}
}

func (c *Client) delete(path string) (*http.Response, error) {
	return c.do("DELETE", path, nil, nil)
}

func (c *Client) put(path string, payload interface{}, headers *map[string]string) (*http.Response, error) {
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		return c.do("PUT", path, bytes.NewBuffer(data), headers)
	}
	return c.do("PUT", path, nil, headers)
}

func (c *Client) post(path string, payload interface{}, headers *map[string]string) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.do("POST", path, bytes.NewBuffer(data), headers)
}

func (c *Client) get(path string) (*http.Response, error) {
	return c.do("GET", path, nil, nil)
}

func (c *Client) do(method, path string, body io.Reader, headers *map[string]string) (*http.Response, error) {
	if c.Limiter != nil {
		c.Limiter.Wait()
	}
	endpoint := c.APIEndpoint + path + ".json"
	req, _ := http.NewRequest(method, endpoint, body)
	if c.AccountGroupID != "" {
		q := req.URL.Query()
		q.Add("aid", c.AccountGroupID)
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	req.Header.Set("content-type", "application/json")
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}

	// Perform any delays required by previously observed rate headers
	delay := setDelay(req, nil, time.Now())
	time.Sleep(delay)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Store reported rate limit status
	storeLimits(req, resp, time.Now())

	// If request was rate limited, back off and retry.
	// We shouldn't typically need to do this, because the above delays should
	// prevent us from hitting the limit, but there may be other users in an
	// org who might have triggered the limiting.
	if resp.StatusCode == 429 {
		delay := setDelay(req, resp, time.Now())
		time.Sleep(delay)
		resp, err = c.HTTPClient.Do(req)
	}

	return c.checkResponse(resp, err)
}

func (c *Client) decodeJSON(resp *http.Response, payload interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(payload)
}

func (c *Client) checkResponse(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return resp, fmt.Errorf("Error calling the API endpoint: %v", err)
	}
	if 199 >= resp.StatusCode || 300 <= resp.StatusCode {
		var eo *errorObject
		var getErr error
		if eo, getErr = c.getErrorFromResponse(resp); getErr != nil {
			return resp, fmt.Errorf("Response did not contain formatted error: %s. HTTP response code: %v. Raw response: %+v", getErr, resp.StatusCode, resp)
		}
		return resp, fmt.Errorf("Failed call API endpoint. HTTP response code: %v. Error: %v", resp.StatusCode, eo)
	}
	return resp, nil
}

func (c *Client) getErrorFromResponse(resp *http.Response) (*errorObject, error) {
	var result errorObject
	if err := c.decodeJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", err)
	}
	return &result, nil
}

// setDelay determines the pause time needed to prevent invoking rate limiting
func setDelay(req *http.Request, resp *http.Response, now time.Time) time.Duration {
	// Choose which rate limit applies
	var delay time.Duration
	var rate RateLimit
	if resp == nil {
		resp = &http.Response{}
	}
	instantTest := isInstantTest(req)
	if instantTest {
		rate = instantTestRate
	} else {
		rate = orgRate
	}

	// If the limit is 0, this is either our first request or we are not receiving
	// rate limit data in the headers
	if rate.Limit == 0 {
		return 0
	}

	// If this is the first time we've sent this particular request and we
	// aren't at the end of our remaining requests for the period...
	if rate.Remaining > 1 && resp.StatusCode != 429 {
		baseDelay := 1.0 / float64(rate.Limit) * float64(time.Minute.Nanoseconds())
		// The rate limit is per minute, so if there was a zero response time
		// then the ideal delay would be the one minute divided by the rate.
		// To account for potential other users, we will multiply by the
		// difference between the remaining count and our last seen remaining
		// count.
		delta := rate.LastRemaining - rate.Remaining
		if delta < 1 {
			delta = 1
		}

		// It's possible that these calls could be made concurrently, in which
		// case the pacing delay would effectively be divided by the batch size.
		// To account for this, we track messages sent for this session and
		// account for any that have delays which have not expired.
		for i, t := range rate.ConcurrentMessages {
			if t.Sub(now) >= time.Duration(0) {
				rate.ConcurrentMessages = rate.ConcurrentMessages[i:]
				break
			}
		}

		delta += int64(len(rate.ConcurrentMessages))
		delay = time.Duration(baseDelay * float64(delta))
		rate.ConcurrentMessages = append(rate.ConcurrentMessages, now.Add(delay))
		log.Printf("[INFO] %v of %v requests / min remain.  Sleeping %v to prevent rate limiting.",
			rate.Remaining, rate.Limit, delay)
	} else {
		// else calculate delay until resume time.
		// Assume our clock is roughly in sync with the clock setting the resume time.
		delay = time.Duration((rate.Reset - now.Unix() + 1) * time.Second.Nanoseconds())
		// ThousandEyes rates reset within one minute (but not guaranteed).
		// If we exceed a minute wait time, something may be wrong.
		if delay > time.Minute {
			delay = time.Minute
		}
		log.Printf("[INFO] Rate Limited: Sleeping %v before resubmitting\n", delay)
	}
	if instantTest {
		instantTestRate.ConcurrentMessages = rate.ConcurrentMessages
	} else {
		orgRate.ConcurrentMessages = rate.ConcurrentMessages
	}
	return delay
}

// storeLimits assigns the global variables to track current rate limit data
func storeLimits(req *http.Request, resp *http.Response, now time.Time) {
	// We discard errors, because an error or blank result also return 0
	if resp.Header != nil {
		if v := resp.Header.Get("X-Organization-Rate-Limit-Limit"); v != "" {
			orgRate.Limit, _ = strconv.ParseInt(v, 10, 64)
		}
		if v := resp.Header.Get("X-Organization-Rate-Limit-Remaining"); v != "" {
			orgRate.Remaining, _ = strconv.ParseInt(v, 10, 64)
		}
		if v := resp.Header.Get("X-Organization-Rate-Limit-Reset"); v != "" {
			orgRate.Reset, _ = strconv.ParseInt(v, 10, 64)
		}
		if v := resp.Header.Get("X-Instant-Test-Rate-Limit-Limit"); v != "" {
			instantTestRate.Limit, _ = strconv.ParseInt(v, 10, 64)
		}
		if v := resp.Header.Get("X-Instant-Test-Rate-Limit-Remaining"); v != "" {
			instantTestRate.Remaining, _ = strconv.ParseInt(v, 10, 64)
		}
		if v := resp.Header.Get("X-Instant-Test-Rate-Limit-Reset"); v != "" {
			instantTestRate.Reset, _ = strconv.ParseInt(v, 10, 64)
		}
	}
}

func isInstantTest(req *http.Request) bool {
	return strings.HasPrefix(req.URL.Path, "/v6/instant") == true || strings.HasPrefix(req.URL.Path, "/v6/endpoint-instant")
}
