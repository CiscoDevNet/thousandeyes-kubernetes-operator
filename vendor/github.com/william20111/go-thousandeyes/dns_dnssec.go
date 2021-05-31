package thousandeyes

import "fmt"

// DNSSec - DNSSec test
type DNSSec struct {
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
	Agents   []Agent `json:"agents,omitempty"`
	Domain   string  `json:"domain,omitempty"`
	Interval int     `json:"interval,omitempty"`
}

// AddAgent - Add agent to DNSSec test
func (t *DNSSec) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *DNSSec) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: id}
	t.AlertRules = append(t.AlertRules, alertRule)
}

// GetDNSSec - get DNSSec test
func (c *Client) GetDNSSec(id int) (*DNSSec, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &DNSSec{}, err
	}
	var target map[string][]DNSSec
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// CreateDNSSec - Create DNSSec test
func (c Client) CreateDNSSec(t DNSSec) (*DNSSec, error) {
	resp, err := c.post("/tests/dns-dnssec/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create dns dnssec test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSSec
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// DeleteDNSSec - delete DNSSec test
func (c *Client) DeleteDNSSec(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/dns-dnssec/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete dnsp domain test, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateDNSSec - update DNSSec test
func (c *Client) UpdateDNSSec(id int, t DNSSec) (*DNSSec, error) {
	resp, err := c.post(fmt.Sprintf("/tests/dns-dnssec/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSSec
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
