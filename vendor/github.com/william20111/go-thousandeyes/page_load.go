package thousandeyes

import (
	"fmt"
)

// PageLoad - a page log struct
type PageLoad struct {
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
	BGPMonitors           []BGPMonitor  `json:"bgpMonitors,omitempty"`
	ContentRegex          string        `json:"contentRegex,omitempty"`
	CustomHeaders         CustomHeaders `json:"customHeaders,omitempty"`
	FollowRedirects       int           `json:"followRedirects,omitempty"`
	HTTPInterval          int           `json:"httpInterval,omitempty"`
	HTTPTargetTime        int           `json:"httpTargetTime,omitempty"`
	HTTPTimeLimit         int           `json:"httpTimeLimit,omitempty"`
	HTTPVersion           int           `json:"httpVersion,omitempty"`
	IncludeHeaders        int           `json:"includeHeaders,omitempty"`
	Interval              int           `json:"interval,omitempty"`
	MTUMeasurements       int           `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements   int           `json:"networkMeasurements,omitempty"`
	NumPathTraces         int           `json:"numPathTraces,omitempty"`
	PageLoadTargetTime    int           `json:"pageLoadTargetTime,omitempty"`
	PageLoadTimeLimit     int           `json:"pageLoadTimeLimit,omitempty"`
	Password              string        `json:"password,omitempty"`
	PathTraceMode         string        `json:"pathTraceMode,omitempty"`
	ProbeMode             string        `json:"probeMode,omitempty"`
	Protocol              string        `json:"protocol,omitempty"`
	SSLVersion            string        `json:"sslVersion,omitempty"`
	SSLVersionID          int           `json:"sslVersionId,omitempty"`
	Subinterval           int           `json:"subinterval,omitempty"`
	URL                   string        `json:"url,omitempty"`
	UseNTLM               int           `json:"useNtlm,omitempty"`
	UsePublicBGP          int           `json:"usePublicBgp,omitempty"`
	UserAgent             string        `json:"userAgent,omitempty"`
	Username              string        `json:"username,omitempty"`
	VerifyCertificate     int           `json:"verifyCertificate,omitempty"`
}

// AddAgent  - add an aget
func (t *PageLoad) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

//GetPageLoad - get page load test
func (c *Client) GetPageLoad(id int) (*PageLoad, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &PageLoad{}, err
	}
	var target map[string][]PageLoad
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//CreatePageLoad - create pager load test
func (c Client) CreatePageLoad(t PageLoad) (*PageLoad, error) {
	resp, err := c.post("/tests/page-load/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create test, response code %d", resp.StatusCode)
	}
	var target map[string][]PageLoad
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// DeletePageLoad - Delete page load tes
func (c *Client) DeletePageLoad(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/page-load/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete page load, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdatePageLoad - Upload page load
func (c *Client) UpdatePageLoad(id int, t PageLoad) (*PageLoad, error) {
	resp, err := c.post(fmt.Sprintf("/tests/page-load/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]PageLoad
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
