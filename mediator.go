package libwimark

// SSHAccess for ssh creds data
type SSHAccess struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Key      string `json:"key"`
	Port     int    `json:"port"`
}

// Controller for integration with external controller (like Cisco)
type Controller struct {
	Status ControllerStatusType `json:"status"`
	Enable bool                 `json:"enable"`

	Name        string `json:"name"`
	Description string `json:"description"`

	MAC      string `json:"mac"`
	Serial   string `json:"serial"`
	Vendor   string `json:"vendor"`
	Firmware string `json:"fw_version"`

	IPAddr string    `json:"ip_addr"`
	Access SSHAccess `json:"access"`
}

// ExtAccessPoint for represent External access point
type ExtAccessPoint struct {
	Status ConfigurationStatus `json:"status"`

	Name        string `json:"name"`
	Description string `json:"description"`

	MAC      string `json:"mac"`
	Serial   string `json:"serial"`
	Vendor   string `json:"vendor"`
	Firmware string `json:"fw_version"`

	IPAddr string    `json:"ip_addr"`
	Access SSHAccess `json:"access"`
}
