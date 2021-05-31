package thousandeyes

import (
	"fmt"
)

// RTP Stream, labeled "voice"

// RTPStream - RTPStream trace test
type RTPStream struct {
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
	Agents          []Agent      `json:"agents,omitempty"`
	BGPMeasurements int          `json:"bgpMeasurements,omitempty"`
	BGPMonitors     []BGPMonitor `json:"bgpMonitors,omitempty"`
	Codec           string       `json:"codec,omitempty"`
	CodecID         int          `json:"codecId,omitempty"`
	DSCP            string       `json:"dscp,omitempty"`
	DSCPID          int          `json:"dscpId,omitempty"`
	Duration        int          `json:"duration,omitempty"`
	Interval        int          `json:"interval,omitempty"`
	JitterBuffer    int          `json:"jitterBuffer,omitempty"`
	MTUMeasurements int          `json:"mtuMeasurements,omitempty"`
	NumPathTraces   int          `json:"numPathTraces,omitempty"`
	TargetAgentID   int          `json:"targetAgentId,omitempty"`
	UsePublicBGP    int          `json:"usePublicBgp,omitempty"`
	// server field is present in response, but we should not track it.
	//Server          string       `json:"server,omitempty"`
}

// AddAgent - Add agent to voice call  test
func (t *RTPStream) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

// GetRTPStream - get voice call test
func (c *Client) GetRTPStream(id int) (*RTPStream, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &RTPStream{}, err
	}
	var target map[string][]RTPStream
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//CreateRTPStream - Create voice call test
func (c Client) CreateRTPStream(t RTPStream) (*RTPStream, error) {
	resp, err := c.post("/tests/voice/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create voice test, response code %d", resp.StatusCode)
	}
	var target map[string][]RTPStream
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteRTPStream - delete voice call test
func (c *Client) DeleteRTPStream(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/voice/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete voice test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateRTPStream - update voice call test
func (c *Client) UpdateRTPStream(id int, t RTPStream) (*RTPStream, error) {
	resp, err := c.post(fmt.Sprintf("/tests/voice/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]RTPStream
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
