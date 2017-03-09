package libwimark

type Document map[string]interface{}

type UUID string

type SecurityType string

const (
	WPA2Personal   = SecurityType("wpa2personal")
	WPA2Enterprise = SecurityType("wpa2enterprise")
)

func SecurityTypeFromString(v string) *SecurityType {
	switch v {
	case "wpa2personal":
		var v = WPA2Personal
		return &v
	case "wpa2enterprise":
		var v = WPA2Enterprise
		return &v
	default:
		return nil
	}
}

type InterfaceType string

const (
	I2_4 = InterfaceType("2.4")
	I5_0 = InterfaceType("5.0")
)

type SecuritySuite string

const (
	AES  = SecuritySuite("aes")
	TKIP = SecuritySuite("tkip")
)

func (self *SecuritySuite) UnmarshalJSON(b []byte) error {
	var data string
	var err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	switch data {
	case "aes":
		var v = AES
		(*self) = v
	case "tkip":
		var v = TKIP
		(*self) = v
	default:
		return errors.New("Unknown security suite")
	}

	return nil
}

type WLAN struct {
	Name     string `json:"name"`
	Power    int    `json:"power"`
	Security *struct {
		T    SecurityType `json:"type"`
		Data UUID         `json:"data"`
	} `json:"security"`
}

type InterfaceConfiguration struct {
	WLANs []UUID
}

type Configuration struct {
	IfaceConfigs map[InterfaceType]InterfaceConfiguration
}

type CPE struct {
	Name   string
	Macs   []string
	Config Configuration
}
