package libwimark

type UUID string

type SecurityType float64

const (
	None           = SecurityType(0)
	WPA2Personal   = SecurityType(1)
	WPA2Enterprise = SecurityType(2)
	RADIUS         = SecurityType(3)
)

type InterfaceType float64

const (
	I2_4 = InterfaceType(1)
	I5_0 = InterfaceType(0)
)

type WLAN struct {
	Name     string
	Power    int
	Security struct {
		T    SecurityType
		Data *UUID
	}
}

type InterfaceConfiguration struct {
	WLANs map[UUID]bool // This is a set
}

type Configuration struct {
	iconfigs map[InterfaceType]InterfaceConfiguration
}

type CPE struct {
	Name   string
	Macs   []string
	Config UUID
}
