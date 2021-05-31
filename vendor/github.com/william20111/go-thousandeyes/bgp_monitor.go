package thousandeyes

import "fmt"

// BGPMonitors - list of bgp montors
type BGPMonitors []BGPMonitor

// BGPMonitor - BGPMonitor struct
type BGPMonitor struct {
	MonitorID   int    `json:"monitorId,omitempty"`
	IPAddress   string `json:"ipAddress,omitempty"`
	Network     string `json:"network,omitempty"`
	MonitorType string `json:"monitorType,omitempty"`
	MonitorName string `json:"monitorName,omitempty"`
}

// GetBPGMonitors - Get bgp monitors
func (c *Client) GetBPGMonitors() (*BGPMonitors, error) {
	resp, err := c.get("/bgp-monitors")
	if err != nil {
		return &BGPMonitors{}, err
	}
	var target map[string]BGPMonitors
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	monitors := target["bgpMonitors"]
	return &monitors, nil
}
