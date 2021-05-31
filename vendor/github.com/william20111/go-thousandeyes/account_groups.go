package thousandeyes

import "fmt"

// AccountGroups - list of account groups
type AccountGroups []AccountGroup

// AccountGroup - An account within a ThousandEyes organization
type AccountGroup struct {
	AccountGroupName string `json:"accountGroupName,omitempty"`
	AID              int    `json:"aid,omitempty"`
}

// SharedWithAccount describes accounts with which a resource is shared.
// This is separate from the AccountGroup above only due to the difference
// in JSON object names.
type SharedWithAccount struct {
	AccountGroupName string `json:"name,omitempty"`
	AID              int    `json:"aid,omitempty"`
}

// GetAccountGroups - Get third party and webhook integrations
func (c *Client) GetAccountGroups() (*[]SharedWithAccount, error) {
	resp, err := c.get("/account-groups")
	if err != nil {
		return nil, err
	}
	var target map[string][]AccountGroup

	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	if _, ok := target["accountGroups"]; !ok {
		return nil, fmt.Errorf("'accountGroups' not found in JSON response")
	}

	// Since the use of this is for the SharedWithAccount configuration,
	// we will return that type.
	var accountGroups []SharedWithAccount
	for _, v := range target["accountGroups"] {
		account := SharedWithAccount{
			AccountGroupName: v.AccountGroupName,
			AID:              v.AID,
		}
		accountGroups = append(accountGroups, account)
	}
	return &accountGroups, nil
}
