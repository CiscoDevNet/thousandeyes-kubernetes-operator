package thousandeyes

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// SIPServer - SIPServer trace test
type SIPServer struct {
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
	Agents                []Agent     `json:"agents,omitempty"`
	BandwidthMeasurements int         `json:"bandwidthMeasurements,omitempty"`
	BGPMeasurements       int         `json:"bgpMeasurements,omitempty"`
	Interval              int         `json:"interval,omitempty"`
	MTUMeasurements       int         `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements   int         `json:"networkMeasurements,omitempty"`
	NumPathTraces         int         `json:"numPathTraces,omitempty"`
	OptionsRegex          string      `json:"options_regex,omitempty"`
	PathTraceMode         string      `json:"pathTraceMode,omitempty"`
	ProbeMode             string      `json:"probeMode,omitempty"`
	RegisterEnabled       int         `json:"registerEnabled,omitempty"`
	SIPTargetTime         int         `json:"sipTargetTime,omitempty"`
	SIPTimeLimit          int         `json:"sipTimeLimit,omitempty"`
	TargetSIPCredentials  SIPAuthData `json:"targetSipCredentials,omitempty"`
	UsePublicBGP          int         `json:"usePublicBgp,omitempty"`
}

// AddAgent - Add agemt to sip server  test
func (t *SIPServer) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *SIPServer) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: id}
	t.AlertRules = append(t.AlertRules, alertRule)
}

// GetSIPServer  - get sip server test
func (c *Client) GetSIPServer(id int) (*SIPServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &SIPServer{}, err
	}

	// Duplicate http response so we can read JSON directly
	// and still use the normal client interaction to process
	// the http response.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not decode HTTP response: %v", err)
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))

	var target map[string][]SIPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}

	// A design flaw in ThousandEyes V6 API results in field on sip-server tests which
	// should be part of a targetSipCredentials object (matching the behavior of voice-call
	// tests) are instead part of the sip-server test object itself.
	// As this is not intended to be fixed until V7, The solution will be to have a
	// separate struct for reads, which will be converted before being passed.
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))
	var sipTarget map[string][]SIPAuthData
	if dErr := c.decodeJSON(resp, &sipTarget); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	for i := range target["test"] {
		sipAuth := SIPAuthData{
			AuthUser:     sipTarget["test"][i].AuthUser,
			Password:     sipTarget["test"][i].Password,
			Port:         sipTarget["test"][i].Port,
			Protocol:     sipTarget["test"][i].Protocol,
			SIPProxy:     sipTarget["test"][i].SIPProxy,
			SIPRegistrar: sipTarget["test"][i].AuthUser,
			User:         sipTarget["test"][i].User,
		}
		target["test"][i].TargetSIPCredentials = sipAuth
	}
	return &target["test"][0], nil
}

//CreateSIPServer - Create sip server test
func (c Client) CreateSIPServer(t SIPServer) (*SIPServer, error) {
	resp, err := c.post("/tests/sip-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create sip-server test, response code %d", resp.StatusCode)
	}
	var target map[string][]SIPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteSIPServer - delete sip server test
func (c *Client) DeleteSIPServer(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/sip-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete sip test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateSIPServer - - update sip server test
func (c *Client) UpdateSIPServer(id int, t SIPServer) (*SIPServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/sip-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]SIPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
