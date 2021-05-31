package thousandeyes

import (
	"fmt"
	"time"
)

// User - a user
type User struct {
	Name                 string             `json:"name,omitempty"`
	Email                string             `json:"email,omitempty"`
	UID                  int                `json:"uid,omitempty"`
	LastLogin            *time.Time         `json:"lastLogin,omitempty"`
	DateRegistered       *time.Time         `json:"dateRegistered,omitempty"`
	LoginAccountGroup    AccountGroup       `json:"loginAccountGroup,omitempty"`
	AccountGroupRoles    []AccountGroupRole `json:"accountGroupRoles,omitempty"`
	AllAccountGroupRoles []AccountGroupRole `json:"allAccountGroupRoles,omitempty"`
}

// GetUsers - get users
func (c *Client) GetUsers() (*[]User, error) {
	resp, err := c.get("/users")
	if err != nil {
		return nil, err
	}
	var target map[string][]User
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	users := target["users"]
	return &users, nil
}

// GetUser - get user
func (c *Client) GetUser(id int) (*User, error) {
	resp, err := c.get(fmt.Sprintf("/users/%d", id))
	if err != nil {
		return nil, err
	}
	var target map[string][]User
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	user := target["users"][0]
	return &user, nil
}

// DeleteUser - delete user
func (c *Client) DeleteUser(id int) error {
	resp, err := c.post(fmt.Sprintf("/users/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete alert rule, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateUser - update user
func (c *Client) UpdateUser(id int, user User) (*User, error) {
	resp, err := c.post(fmt.Sprintf("/users/%d/update", id), user, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to update alert rule, response code %d", resp.StatusCode)
	}
	var target User
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &target, nil
}

// CreateUser - create user
func (c *Client) CreateUser(user User) (*User, error) {
	resp, err := c.post("/users/new", user, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to update alert rule, response code %d", resp.StatusCode)
	}
	var target User
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &target, nil
}
