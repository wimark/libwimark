package libwimark

type Document map[string]interface{}

type UUID string

type SecurityType float64

const (
	WPA2Personal   = SecurityType(0)
	WPA2Enterprise = SecurityType(1)
	RADIUS         = SecurityType(2)
)

type InterfaceType float64

const (
	I2_4 = InterfaceType(0)
	I5_0 = InterfaceType(1)
)

type WLAN struct {
	Name     string `json:"name"`
	Power    int    `json:"power"`
	Security *struct {
		T    SecurityType `json:"type"`
		Data UUID         `json:"data"`
	} `json:"security"`
}

type InterfaceConfiguration struct {
	WLANs map[UUID]bool // This is a set
}

type Configuration struct {
	IfaceConfigs map[InterfaceType]InterfaceConfiguration
}

type CPE struct {
	Name   string
	Macs   []string
	Config Configuration
}
