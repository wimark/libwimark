package libwimark

import (
	"encoding/json"
	"errors"
	"github.com/vorot93/goutil"
)

type Document goutil.Document

type UUID string

type Radius struct {
	Type     RadiusType `json:"type"`
	Hostname string     `json:"hostname"`
	Secret   string     `json:"secret"`
}

type InterfaceType string

const (
	I2_4 = InterfaceType("2.4")
	I5_0 = InterfaceType("5.0")
)

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

type WiredData struct {
	Name string
	Mac  string
}

type WiFiData struct {
	Name      string  `json:"name,omitempty"`
	Mac       string  `json:"mac,omitempty"`
	Frequency float32 `json:"frequency,omitempty"`
	BandMode  string  `json:"band_mode,omitempty"`
	Bandwidth string  `json:"bandwidth,omitempty"`
	Channel   string  `json:"channel,omitempty"`
	TxPower   int     `json:"tx_power,omitempty"`
	WLANs     []UUID  `json:"wlans,omitempty"`
}

type CPEInterface struct {
	CPEInterfaceInfo `json:",inline"`
	Addr             string `json:"addr"`
}

func (self *CPEInterface) UnmarshalJSON(b []byte) error {
	var doc Document
	var err = json.Unmarshal(b, &doc)
	if err != nil {
		return err
	}

	if doc == nil {
		return nil
	}

	var addrErased, addrExists = doc["addr"]
	if addrExists {
		var addr, addrOk = addrErased.(string)
		if addrOk {
			self.Addr = addr
		} else {
			return errors.New("Address must be a string")
		}
	}

	delete(doc, "addr")

	var v, _ = json.Marshal(doc)

	return self.CPEInterfaceInfo.UnmarshalJSON(v)
}

type CPE struct {
	Name         string                  `json:"name"`
	Description  string                  `json:"description"`
	Model        UUID                    `json:"model"`
	Interfaces   map[string]CPEInterface `json:"interfaces"`
	ConfigStatus ConfigurationStatus     `json:"config_status"`
}

type Stat struct {
	Timestamp    int64   `json:"time"`
	CPU          float64 `json:"cpu"`
	RAM          float64 `json:"ram"`
	Storage      float64 `json:"storage"`
	ProcActive   uint64  `json:"processes_active"`
	ProcSleeping uint64  `json:"processes_sleeping"`
	Interfaces   map[string]struct {
		Name string  `json:"name"`
		Tx   float64 `json:"tx"`
		Rx   float64 `json:"rx"`
	} `json:"interfaces"`
}
