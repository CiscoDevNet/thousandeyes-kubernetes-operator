package v1alpha1

type Agent struct {
	AgentID               int      `json:"agentId,omitempty"`
	AgentName             string   `json:"agentName,omitempty"`
	AgentType             string   `json:"agentType,omitempty"`
	CountryID             string   `json:"countryId,omitempty"`
	IPAddresses           []string `json:"ipAddresses,omitempty"`
	Location              string   `json:"location,omitempty"`
	Hostname              string   `json:"hostname,omitempty"`
	Prefix                string   `json:"prefix,omitempty"`
	Enabled               int      `json:"enabled,omitempty"`
	Network               string   `json:"network,omitempty"`
	CreatedDate           string   `json:"createdDate,omitempty"`
	LastSeen              string   `json:"lastSeen,omitempty"`
	AgentState            string   `json:"agentState,omitempty"`
	VerifySslCertificates int      `json:"verifySslCertificate,omitempty"`
	KeepBrowserCache      int      `json:"keepBrowserCache,omitempty"`
	Utilization           int      `json:"utilization,omitempty"`
	Ipv6Policy            string   `json:"IPV6Policy,omitempty"`
	TargetForTests        string   `json:"targetForTests,omitempty"`
}
