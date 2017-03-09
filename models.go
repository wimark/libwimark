package libwimark

import (
	"encoding/json"
	"errors"
)

type Document map[string]interface{}

func (self *Document) ToValue(factory func() interface{}) (interface{}, error) {
	var s, merr = json.Marshal(self)
	if merr != nil {
		return nil, merr
	}

	var v = factory()
	var umerr = json.Unmarshal(s, v)
	if umerr != nil {
		return nil, umerr
	}

	return v, nil
}

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

type StatisticsData struct {
	CPE       UUID
	Timestamp int64
	CPU       float64
	Mem       float64
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

type WPA2Common struct {
	Suite SecuritySuite `json:"suite"`
}

type WPA2PersonalData struct {
	WPA2Common `json:,inline`
	PSK        string `json:"psk"`
}

type WPA2EnterpriseData struct {
	WPA2Common           `json:,inline`
	RadiusAuthentication []string `json:"radius_auth"`
}

type SecuritySettings interface {
	is_security_settings()
}

func (*WPA2PersonalData) is_security_settings()   {}
func (*WPA2EnterpriseData) is_security_settings() {}

type EnumSecurity struct {
	T SecurityType
	D SecuritySettings
}

func (self *EnumSecurity) UnmarshalJSON(b []byte) error {
	var data Document
	var err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	if data == nil {
		return nil
	}
	var t_erased, t_found = data["type"]
	var secdata_erased, secdata_found = data["data"]
	if !t_found || !secdata_found {
		return nil
	}
	var t_s = t_erased.(string)
	var t = SecurityTypeFromString(t_s)
	if t == nil {
		return errors.New("Invalid security type")
	}
	var secdata_m = Document(secdata_erased.(map[string]interface{}))
	var secdata interface{}
	var secdata_err error

	switch *t {
	case WPA2Personal:
		secdata, secdata_err = secdata_m.ToValue(func() interface{} { return &WPA2PersonalData{} })
	case WPA2Enterprise:
		secdata, secdata_err = secdata_m.ToValue(func() interface{} { return &WPA2EnterpriseData{} })
	}
	if secdata_err != nil {
		return secdata_err
	}

	self.T = *t
	self.D = secdata.(SecuritySettings)

	return nil
}

type WLAN struct {
	Name             string        `json:"name"`
	SSID             string        `json:"ssid"`
	Description      string        `json:"description"`
	Security         *EnumSecurity `json:"security"`
	VLAN             int           `json:"vlan"`
	RadiusAccounting []UUID        `json:"radius_accounting"`
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
