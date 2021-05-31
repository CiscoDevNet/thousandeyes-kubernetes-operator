package thousandeyes

import (
	"fmt"
)

// HTTPServerResponse - a http server response
type HTTPServerResponse struct {
	tests []HTTPServer
}

// CustomHeaders represents the JSON object exchanged for specifying
// custom HTTP headers in HTTP Server, Page Load, and Web Transaction tests.
type CustomHeaders struct {
	Root    map[string]string            `json:"root,omitempty"`
	All     map[string]string            `json:"all,omitempty"`
	Domains map[string]map[string]string `json:"domains,omitempty"`
}

// HTTPServer - a http server test
type HTTPServer struct {
	// Common test fields
	AlertsEnabled      int                 `json:"alertsEnabled,omitempty"`
	AlertRules         []AlertRule         `json:"alertRules,omitempty"`
	APILinks           []APILink           `json:"apiLinks,omitempty"`
	CreatedBy          string              `json:"createdBy,omitempty"`
	CreatedDate        string              `json:"createdDate,omitempty"`
	Description        string              `json:"description,omitempty"`
	Enabled            int                 `json:"enabled,omitempty"`
	Groups             []GroupLabel        `json:"groups,omitempty"`
	ModifiedBy         string              `json:"modifiedBy,omitempty"`
	ModifiedDate       string              `json:"modifiedDate,omitempty"`
	SavedEvent         int                 `json:"savedEvent,omitempty"`
	SharedWithAccounts []SharedWithAccount `json:"sharedWithAccounts,omitempty"`
	TestID             int                 `json:"testId,omitempty"`
	TestName           string              `json:"testName,omitempty"`
	Type               string              `json:"type,omitempty"`
	// LiveShare is common to all tests except DNS+
	LiveShare int `json:"liveShare,omitempty"`
	// Fields unique to this test
	Agents                Agents        `json:"agents,omitempty"`
	AuthType              string        `json:"authType,omitempty"`
	BandwidthMeasurements int           `json:"bandwidthMeasurements,omitempty"`
	BGPMeasurements       int           `json:"bgpMeasurements,omitempty"`
	BGPMonitors           []Monitor     `json:"bgpMonitors,omitempty"`
	ClientCertificate     string        `json:"clientCertificate,omitempty"`
	ContentRegex          string        `json:"contentRegex,omitempty"`
	CustomHeaders         CustomHeaders `json:"customHeaders,omitempty"`
	DesiredStatusCode     string        `json:"desiredStatusCode,omitempty"`
	DownloadLimit         string        `json:"downloadLimit,omitempty"`
	DNSOverride           string        `json:"dnsOverride,omitempty"`
	FollowRedirects       int           `json:"followRedirects,omitempty"`
	Headers               []string      `json:"headers,omitempty"`
	HTTPVersion           int           `json:"httpVersion,omitempty"`
	HTTPTargetTime        int           `json:"httpTargetTime,omitempty"`
	HTTPTimeLimit         int           `json:"httpTimeLimit,omitempty"`
	Interval              int           `json:"interval,omitempty"`
	MTUMeasurements       int           `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements   int           `json:"networkMeasurements,omitempty"`
	NumPathTraces         int           `json:"numPathTraces,omitempty"`
	Password              string        `json:"password,omitempty"`
	PathTraceMode         string        `json:"pathTraceMode,omitempty"`
	PostBody              string        `json:"postBody,omitempty"`
	ProbeMode             string        `json:"probeMode,omitempty"`
	Protocol              string        `json:"protocol,omitempty"`
	SSLVersion            string        `json:"sslVersion,omitempty"`
	SSLVersionID          int           `json:"sslVersionId,omitempty"`
	URL                   string        `json:"url,omitempty"`
	UseNTLM               int           `json:"useNtlm,omitempty"`
	UsePublicBGP          int           `json:"usePublicBgp,omitempty"`
	UserAgent             string        `json:"userAgent,omitempty"`
	Username              string        `json:"username,omitempty"`
	VerifyCertificate     int           `json:"verifyCertificate,omitempty"`
}

// AddAgent - add an agent
func (t *HTTPServer) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

//GetHTTPServer - Get an HTTP Server test
func (c *Client) GetHTTPServer(id int) (*HTTPServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &HTTPServer{}, err
	}
	var target map[string][]HTTPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//CreateHTTPServer - create a http server
func (c Client) CreateHTTPServer(t HTTPServer) (*HTTPServer, error) {
	resp, err := c.post("/tests/http-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create http server, response code %d", resp.StatusCode)
	}
	var target map[string][]HTTPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteHTTPServer - delete an http server
func (c *Client) DeleteHTTPServer(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/http-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete http server, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateHTTPServer - Update an http server test
func (c *Client) UpdateHTTPServer(id int, t HTTPServer) (*HTTPServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/http-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update http server, response code %d", resp.StatusCode)
	}
	var target map[string][]HTTPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
