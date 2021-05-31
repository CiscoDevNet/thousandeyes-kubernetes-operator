package thousandeyes

import (
	"fmt"
	"strconv"
	"strings"
)

// AgentServer  - Agent to server test
type AgentServer struct {
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
	Agents                Agents       `json:"agents,omitempty"`
	BandwidthMeasurements int          `json:"bandwidthMeasurements,omitempty"`
	BGPMeasurements       int          `json:"bgpMeasurements,omitempty"`
	BGPMonitors           []BGPMonitor `json:"bgpMonitors,omitempty"`
	Interval              int          `json:"interval,omitempty"`
	MTUMeasurements       int          `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements   int          `json:"networkMeasurements,omitempty"`
	NumPathTraces         int          `json:"numPathTraces,omitempty"`
	PathTraceMode         string       `json:"pathTraceMode,omitempty"`
	Port                  int          `json:"port,omitempty"`
	ProbeMode             string       `json:"probeMode,omitempty"`
	Protocol              string       `json:"protocol,omitempty"`
	Server                string       `json:"server,omitempty"`
	UsePublicBGP          int          `json:"usePublicBgp,omitempty"`
}

// extractPort - Set Server and Port fields if they are combined in the Server field.
func extractPort(test AgentServer) (AgentServer, error) {
	// Unfortunately, the V6 API returns the server value with the port,
	// rather than having them in separate values as the API requires for
	// submissions.  Not required for ICMP tests.
	var err error
	if test.Protocol != "ICMP" && strings.Index(test.Server, ":") != -1 {
		serverParts := strings.Split(test.Server, ":")
		test.Server = serverParts[0]
		test.Port, err = strconv.Atoi(serverParts[1])
		if err != nil {
			err = fmt.Errorf("Invalid port in server declaration")
		}
	}
	return test, nil
}

// AddAgent - Add agent to server test
func (t *AgentServer) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *AgentServer) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: id}
	t.AlertRules = append(t.AlertRules, alertRule)
}

// GetAgentServer - Get agent to server test
func (c *Client) GetAgentServer(id int) (*AgentServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &AgentServer{}, err
	}
	var target map[string][]AgentServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	test := target["test"][0]
	test, err = extractPort(test)
	if err != nil {
		return nil, err
	}
	return &test, nil
}

// CreateAgentServer  - Create agent to server test
func (c Client) CreateAgentServer(t AgentServer) (*AgentServer, error) {
	resp, err := c.post("/tests/agent-to-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create agent server, response code %d", resp.StatusCode)
	}
	var target map[string][]AgentServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	test := target["test"][0]
	test, err = extractPort(test)
	if err != nil {
		return nil, err
	}
	return &test, nil
}

// DeleteAgentServer  - Delete agent to server test
func (c *Client) DeleteAgentServer(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/agent-to-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete agent server, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateAgentServer  - Update agent to server test
func (c *Client) UpdateAgentServer(id int, t AgentServer) (*AgentServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/agent-to-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update agent server, response code %d", resp.StatusCode)
	}
	var target map[string][]AgentServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
