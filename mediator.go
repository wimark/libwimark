package libwimark

type SSHAccess struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Key      string `json:"key"`
}

type Controller struct {
	// Mac: string,
	// SN: string,
	// Name: string,
	// Description: string,
	// IP: string,
	// SshKey: string,
	// Connected: boolean,
	// Config_status: string,
	// Fw_version: string,
	// Vendor/Model: string
	Status      ControllerStatusType `json:"status"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	MAC         string               `json:"mac"`
	Serial      string               `json:"serial"`
	Vendor      string               `json:"vendor"`
	Firmware    string               `json:"fw_version"`
	Access      SSHAccess            `json:"access"`
	IPAddr      string               `json:"ip_addr"`
}
