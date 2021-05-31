package thousandeyes

import (
	"fmt"
)

// BGP - BGP trace test
type BGP struct {
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
	BGPMonitors            []BGPMonitor `json:"bgpMonitors,omitempty"`
	IncludeCoveredPrefixes int          `json:"includeCoveredPrefixes,omitempty"`
	Prefix                 string       `json:"prefix,omitempty"`
	UsePublicBGP           int          `json:"usePublicBgp,omitempty"`
}

// AddAlertRule - Adds an alert to agent test
func (t *BGP) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: id}
	t.AlertRules = append(t.AlertRules, alertRule)
}

// GetBGP  - get bgp test
func (c *Client) GetBGP(id int) (*BGP, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &BGP{}, err
	}
	var target map[string][]BGP
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//CreateBGP - Create bgp test
func (c Client) CreateBGP(t BGP) (*BGP, error) {
	resp, err := c.post("/tests/bgp/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create test, response code %d", resp.StatusCode)
	}
	var target map[string][]BGP
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteBGP - delete bgp test
func (c *Client) DeleteBGP(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/bgp/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete bgp test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateBGP - - Update bgp trace test
func (c *Client) UpdateBGP(id int, t BGP) (*BGP, error) {
	resp, err := c.post(fmt.Sprintf("/tests/bgp/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]BGP
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
