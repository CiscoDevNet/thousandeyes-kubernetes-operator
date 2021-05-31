package thousandeyes

import "fmt"

// DNSPDomain - DNSP domain test
type DNSPDomain struct {
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
	Domain   string `json:"domain,omitempty"`
	Interval int    `json:"interval,omitempty"`
}

// AddAlertRule - Adds an alert to agent test
func (t *DNSPDomain) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: id}
	t.AlertRules = append(t.AlertRules, alertRule)
}

// GetDNSPDomain - get dnsp domain test
func (c *Client) GetDNSPDomain(id int) (*DNSPDomain, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &DNSPDomain{}, err
	}
	var target map[string][]DNSPDomain
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// CreateDNSPDomain - Create dnsp domain test
func (c Client) CreateDNSPDomain(t DNSPDomain) (*DNSPDomain, error) {
	resp, err := c.post("/tests/dnsp-domain/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to dnsp domain create test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSPDomain
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteDNSPDomain - delete dnsp domain test
func (c *Client) DeleteDNSPDomain(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/dnsp-domain/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete dnsp domain test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateDNSPDomain - update dnsp domain test
func (c *Client) UpdateDNSPDomain(id int, t DNSPDomain) (*DNSPDomain, error) {
	resp, err := c.post(fmt.Sprintf("/tests/dnsp-domain/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]DNSPDomain
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
