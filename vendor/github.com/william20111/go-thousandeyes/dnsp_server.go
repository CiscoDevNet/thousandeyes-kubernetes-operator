package thousandeyes

import "fmt"

// DNSPServer - DNSP server test
type DNSPServer struct {
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
	// Fields unique to this test
	Interval int    `json:"interval,omitempty"`
	Server   string `json:"server,omitempty"`
}

// AddAlertRule - Adds an alert to agent test
func (t *DNSPServer) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: id}
	t.AlertRules = append(t.AlertRules, alertRule)
}

// GetDNSPServer - get dnsp server test
func (c *Client) GetDNSPServer(id int) (*DNSPServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &DNSPServer{}, err
	}
	var target map[string][]DNSPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// CreateDNSPServer - Create dnsp server test
func (c Client) CreateDNSPServer(t DNSPServer) (*DNSPServer, error) {
	resp, err := c.post("/tests/dnsp-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to dnsp server create test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteDNSPServer - delete dnsp server test
func (c *Client) DeleteDNSPServer(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/dnsp-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete dnsp server test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateDNSPServer - update dnsp server test
func (c *Client) UpdateDNSPServer(id int, t DNSPServer) (*DNSPServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/dnsp-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
