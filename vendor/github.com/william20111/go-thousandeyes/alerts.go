package thousandeyes

import (
	"fmt"
	"log"
)

// Alerts - list of alerts
type Alerts []Alert

// Alert - An alert
type Alert struct {
	AlertID        int      `json:"alertId,omitempty"`
	TestID         int      `json:"testId,omitempty"`
	TestName       string   `json:"testName,omitempty"`
	Active         int      `json:"active,omitempty"`
	RuleExpression string   `json:"ruleExpression,omitempty"`
	DateStart      string   `json:"dateStart,omitempty"`
	DateEnd        string   `json:"dateEnd,omitempty"`
	ViolationCount int      `json:"violationCount,omitempty"`
	RuleName       string   `json:"ruleName,omitempty"`
	Permalink      string   `json:"permalink,omitempty"`
	Type           string   `json:"type,omitempty"`
	Agents         Agents   `json:"agents,omitempty"`
	Monitors       Monitors `json:"monitors,omitempty"`
	APILinks       APILinks `json:"apiLinks,omitempty"`
}

// AlertRules - list of alert rules
type AlertRules []AlertRule

// AlertRule - An alert rule
type AlertRule struct {
	AlertRuleID             int    `json:"alertRuleId,omitempty"`
	AlertType               string `json:"alertType,omitempty"`
	Default                 int    `json:"default,omitempty"`
	Direction               string `json:"direction,omitempty"`
	Expression              string `json:"expression,omitempty"`
	IncludeCoveredPrefixes  int    `json:"includeCoveredPrefixes,omitempty"`
	MinimumSources          int    `json:"minimumSources,omitempty"`
	MinimumSourcesPct       int    `json:"minimumSourcesPct,omitempty"`
	NotifyOnClear           int    `json:"notifyOnClear,omitempty"`
	RoundsViolatingMode     string `json:"roundsViolatingMode,omitempty"`
	RoundsViolatingOutOf    int    `json:"roundsViolatingOutOf,omitempty"`
	RoundsViolatingRequired int    `json:"roundsViolatingRequired,omitempty"`
	RuleID                  int    `json:"ruleId,omitempty"`
	RuleName                string `json:"ruleName,omitempty"`
	TestIds                 []int  `json:"testIds,omitempty"`
}

// CreateAlertRule - Create alert rule
func (c Client) CreateAlertRule(a AlertRule) (*AlertRule, error) {
	resp, err := c.post("/alert-rules/new", a, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to create alert rule, response code %d", resp.StatusCode)
	}
	var target AlertRule
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}

	// Set RuleID, because on creation the V6 API returns this as alertRuleId instead.
	// We'll also UNset AlertRuleID so that it isn't seen as a change when it isn't
	// present in other API calls.
	target.RuleID = target.AlertRuleID
	target.AlertRuleID = 0

	return &target, nil
}

//GetAlertRules - Get alert rules
func (c Client) GetAlertRules() (*AlertRules, error) {
	resp, err := c.get(fmt.Sprintf("/alert-rules"))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get alert rule, response code %d", resp.StatusCode)
	}

	var target map[string]AlertRules

	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	alertRules := target["alertRules"]

	return &alertRules, nil
}

// GetAlertRule - Get single alert rule by ID
func (c *Client) GetAlertRule(id int) (*AlertRule, error) {
	log.Printf("[INFO] Getting Alert Rule %v", id)
	resp, err := c.get(fmt.Sprintf("/alert-rules/%d", id))
	if err != nil {
		return &AlertRule{}, err
	}
	var target map[string][]AlertRule
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	if len(target["alertRules"]) < 1 {
		return nil, fmt.Errorf("Could not get alert rule %v", id)
	}
	return &target["alertRules"][0], nil
}

//DeleteAlertRule - delete alert rule
func (c Client) DeleteAlertRule(id int) error {
	resp, err := c.post(fmt.Sprintf("/alert-rules/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete alert rule, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateAlertRule - update alert rule
func (c Client) UpdateAlertRule(id int, a AlertRule) (*AlertRule, error) {
	resp, err := c.post(fmt.Sprintf("/alert-rules/%d/update", id), a, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to update alert rule, response code %d", resp.StatusCode)
	}
	var target AlertRule
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &target, nil
}
