package libwimark

type SSHAccess struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Key      string `json:"key"`
}

type Controller struct {
	Status      ControllerStatusType `json:"status"`
	Enable      bool                 `json:"enable"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	MAC         string               `json:"mac"`
	Serial      string               `json:"serial"`
	Vendor      string               `json:"vendor"`
	Firmware    string               `json:"fw_version"`
	Access      SSHAccess            `json:"access"`
	IPAddr      string               `json:"ip_addr"`
}
