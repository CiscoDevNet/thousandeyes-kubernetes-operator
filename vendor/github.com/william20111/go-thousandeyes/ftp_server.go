package thousandeyes

import "fmt"

// FTPServer - ftp server test
type FTPServer struct {
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
	// LiveSVhare is common to all tests except DNS+
	LiveShare int `json:"liveShare,omitempty"`
	// Fields unique to this test
	Agents              []Agent `json:"agents,omitempty"`
	BGPMeasurements     int     `json:"bgpMeasurements,omitempty"`
	DownloadLimit       int     `json:"downloadLimit,omitempty"`
	FTPTargetTime       int     `json:"ftpTargetTime,omitempty"`
	FTPTimeLimit        int     `json:"ftpTimeLimit,omitempty"`
	Interval            int     `json:"interval,omitempty"`
	MTUMeasurements     int     `json:"mtuMeasurements,omitempty"`
	NetworkMeasurements int     `json:"networkMeasurements,omitempty"`
	NumPathTraces       int     `json:"numPathTraces,omitempty"`
	Password            string  `json:"password,omitempty"`
	PathTraceMode       string  `json:"pathTraceMode,omitempty"`
	ProbeMode           string  `json:"probeMode,omitempty"`
	Protocol            string  `json:"protocol,omitempty"`
	RequestType         string  `json:"requestType,omitempty"`
	URL                 string  `json:"url,omitempty"`
	UseActiveFTP        int     `json:"useActiveFtp,omitempty"`
	UseExplicitFTPS     int     `json:"useExplicitFtps,omitempty"`
	Username            string  `json:"username,omitempty"`
}

// AddAgent - Add ftp server test
func (t *FTPServer) AddAgent(id int) {
	agent := Agent{AgentID: id}
	t.Agents = append(t.Agents, agent)
}

// AddAlertRule - Adds an alert to agent test
func (t *FTPServer) AddAlertRule(id int) {
	alertRule := AlertRule{RuleID: id}
	t.AlertRules = append(t.AlertRules, alertRule)
}

// GetFTPServer - get ftp server test
func (c *Client) GetFTPServer(id int) (*FTPServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &FTPServer{}, err
	}
	var target map[string][]FTPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// CreateFTPServer - Create ftp server test
func (c Client) CreateFTPServer(t FTPServer) (*FTPServer, error) {
	resp, err := c.post("/tests/ftp-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create ftp test, response code %d", resp.StatusCode)
	}
	var target map[string][]FTPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

// DeleteFTPServer - delete ftp server test
func (c *Client) DeleteFTPServer(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/ftp-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete ftp server test, response code %d", resp.StatusCode)
	}
	return nil
}

// UpdateFTPServer - - Update ftp server test
func (c *Client) UpdateFTPServer(id int, t FTPServer) (*FTPServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/ftp-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update ftp test, response code %d", resp.StatusCode)
	}
	var target map[string][]FTPServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
