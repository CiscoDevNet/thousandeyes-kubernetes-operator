package thousandeyes

import "fmt"

// Integration - Integration struct
type Integration struct {
	AuthMethod      string `json:"authMethod,omitempty"`
	AuthUser        string `json:"authUser,omitempty"`
	AuthToken       string `json:"authToken,omitempty"`
	Channel         string `json:"channel,omitempty"`
	IntegrationID   string `json:"integrationId,omitempty"`
	IntegrationName string `json:"integrationName,omitempty"`
	IntegrationType string `json:"integrationType,omitempty"`
	Target          string `json:"target,omitempty"`
}

// GetIntegrations - Get third party and webhook integrations
func (c *Client) GetIntegrations() (*[]Integration, error) {
	resp, err := c.get("/integrations")
	if err != nil {
		return nil, err
	}
	var target map[string]map[string][]Integration

	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	var integrations []Integration
	integrations = append(integrations, target["integrations"]["thirdParty"]...)
	integrations = append(integrations, target["integrations"]["webhook"]...)
	return &integrations, nil
}
