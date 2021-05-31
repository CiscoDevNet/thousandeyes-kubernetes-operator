package thousandeyes

import (
	"fmt"
)

// AgentAgent - test
type AgentAgent struct {
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
	Agents                 []Agent      `json:"agents,omitempty"`
	BGPMeasurements        int          `json:"bgpMeasurements,omitempty"`
	BGPMonitors            []BGPMonitor `json:"bgpMonitors,omitempty"`
	Direction              string       `json:"direction,omitempty"`
	DSCP                   string       `json:"dscp,omitempty"`
	DSCPID                 int          `json:"dscpId"`
	Interval               int          `json:"interval,omitempty"`
	MSS                    int          `json:"mss,omitempty"`
	NetworkMeasurements    int          `json:"networkMeasurements,omitempty"`
	MTUMeasurements        int          `json:"mtuMeasurements,omitempty"`
	NumPathTraces          int          `json:"numPathTraces,omitempty"`
	PathTraceMode          string       `json:"pathTraceMode,omitempty"`
	Port                   int          `json:"port,omitempty"`
	Protocol               string       `json:"protocol,omitempty"`
	TargetAgentID          int          `json:"targetAgentId,omitempty"`
	ThroughputDuration     int          `json:"throughputDuration,omitempty"`
	ThroughputMeasurements int          `json:"throughputMeasurements,omitempty"`
	ThroughputRate         int          `json:"throughputRate,omitempty"`
	UsePublicBGP           int          `json:"usePublicBgp,omitempty"`
}

// AddAgent - Adds an agent to agent test
func (t *AgentAgent) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *AgentAgent) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: id}
	t.AlertRules = append(t.AlertRules, alertRule)
}

// GetAgentAgent - Get an agent to agent test
func (c *Client) GetAgentAgent(id int) (*AgentAgent, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &AgentAgent{}, err
	}
	var target map[string][]AgentAgent
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// CreateAgentAgent - Create an agent to agent test
func (c Client) CreateAgentAgent(t AgentAgent) (*AgentAgent, error) {
	resp, err := c.post("/tests/agent-to-agent/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create agent test, response code %d", resp.StatusCode)
	}
	var target map[string][]AgentAgent
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteAgentAgent - delete agent to agent test
func (c *Client) DeleteAgentAgent(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/agent-to-agent/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete agent test, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateAgentAgent - update agent to agent test
func (c *Client) UpdateAgentAgent(id int, t AgentAgent) (*AgentAgent, error) {
	resp, err := c.post(fmt.Sprintf("/tests/agent-to-agent/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update agent test, response code %d", resp.StatusCode)
	}
	var target map[string][]AgentAgent
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
