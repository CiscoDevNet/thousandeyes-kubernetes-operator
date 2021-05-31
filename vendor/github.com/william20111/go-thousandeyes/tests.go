package thousandeyes

import "fmt"

// GenericTest - GenericTest struct to represent all test types
type GenericTest struct {
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
	Agents []Agent `json:"agents,omitempty"`
}

// GetTests  - get all tests
func (c *Client) GetTests() (*[]GenericTest, error) {
	resp, err := c.get("/tests")
	if err != nil {
		return nil, err
	}
	var target map[string][]GenericTest
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	tests := target["test"]
	return &tests, nil
}

// GetTest - Get test
func (c *Client) GetTest(id int) (*GenericTest, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return nil, err
	}
	var target map[string][]GenericTest
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	test := target["test"][0]
	return &test, nil
}
