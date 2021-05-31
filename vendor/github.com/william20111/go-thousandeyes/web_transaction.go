package thousandeyes

import (
	"fmt"
)

// WebTransaction - a web transcation test
type WebTransaction struct {
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
	ContentRegex          string        `json:"contentRegex,omitempty"`
	Credentials           []int         `json:"credentials,omitempty"`
	CustomHeaders         CustomHeaders `json:"customHeaders,omitempty"`
	DesiredStatusCode     string        `json:"desiredStatusCode,omitempty"`
	HTTPTargetTime        int           `json:"httpTargetTime,omitempty"`
	HTTPTimeLimit         int           `json:"httpTimeLimit,omitempty"`
	HTTPVersion           int           `json:"httpVersion,omitempty"`
	IncludeHeaders        int           `json:"includeHeaders,omitempty"`
	Interval              int           `json:"interval,omitempty"`
	MTUMeasurements       int           `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements   int           `json:"networkMeasurements,omitempty"`
	NumPathTraces         int           `json:"numPathTraces,omitempty"`
	Password              string        `json:"password,omitempty"`
	PathTraceMode         string        `json:"pathTraceMode,omitempty"`
	ProbeMode             string        `json:"probeMode,omitempty"`
	Protocol              string        `json:"protocol,omitempty"`
	SSLVersionID          int           `json:"sslVersionId,omitempty"`
	SubInterval           int           `json:"subinterval,omitempty"`
	TargetTime            int           `json:"targetTime,omitempty"`
	TimeLimit             int           `json:"timeLimit,omitempty"`
	TransactionScript     string        `json:"transactionScript,omitempty"`
	URL                   string        `json:"url,omitempty"`
	UseNTLM               int           `json:"useNtlm,omitempty"`
	UserAgent             string        `json:"userAgent,omitempty"`
	Username              string        `json:"username,omitempty"`
	VerifyCertificate     int           `json:"verifyCertificate,omitempty"`
}

// CreateWebTransaction - Create a web transaction test
func (c Client) CreateWebTransaction(t WebTransaction) (*WebTransaction, error) {
	resp, err := c.post("/tests/web-transactions/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create web transaction, response code %d", resp.StatusCode)
	}
	var target map[string][]WebTransaction
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//GetWebTransaction - get a web transactiont test
func (c *Client) GetWebTransaction(id int) (*WebTransaction, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &WebTransaction{}, err
	}
	var target map[string][]WebTransaction
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteWebTransaction - delete a web transactiont est
func (c *Client) DeleteWebTransaction(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/web-transactions/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete http server, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateWebTransaction - update a web transaction test
func (c *Client) UpdateWebTransaction(id int, t WebTransaction) (*WebTransaction, error) {
	resp, err := c.post(fmt.Sprintf("/tests/web-transactions/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to web transaction, response code %d", resp.StatusCode)
	}
	var target map[string][]WebTransaction
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
