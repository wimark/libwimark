package libwimark

import (
	"encoding/json"
	"errors"
	"github.com/vorot93/goutil"
	"gopkg.in/mgo.v2/bson"
)

type Document goutil.Document

type UUID string

type Radius struct {
	Type     RadiusType `json:"type"`
	Name     string     `json:"name"`
	Hostname string     `json:"hostname"`
	Port     uint16     `json:"port"`
	Secret   string     `json:"secret"`
}

type InterfaceType string

const (
	I2_4 = InterfaceType("2.4")
	I5_0 = InterfaceType("5.0")
)

type WPA2Common struct {
	Suites []SecuritySuite `json:"suites"`
}

type WPA2PersonalData struct {
	WPA2Common `json:,inline`
	PSK        string `json:"psk"`
}

type WPA2EnterpriseData struct {
	WPA2Common           `json:,inline`
	NasID                string   `json:"nas_id"`
	PMKCaching           bool     `json:"pmk_caching"`
	RadiusAuthentication []UUID   `json:"radius_auth"`
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
	Hidden           bool          `json:"hidden"`
	NasID            *string       `json:"nas_id"`
	RadiusAcctServers []UUID       `json:"radius_acct_servers"`
}

type InterfaceConfiguration struct {
	WLANs []UUID
}

type Configuration struct {
	IfaceConfigs map[InterfaceType]InterfaceConfiguration
}

type WiredData struct {
	Name string `json:"name"`
	Mac  string `json:"mac"`
	VLAN int    `json:"vlan"`
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
	CPEInterfaceInfo
	Addr string
}

func (self CPEInterface) MarshalJSON() ([]byte, error) {
	var b []byte
	{
		var err error
		b, err = json.Marshal(self.CPEInterfaceInfo)
		if err != nil {
			return nil, err
		}
	}

	var doc Document
	{
		var err error
		err = json.Unmarshal(b, &doc)
		if err != nil {
			return nil, err
		}
	}

	doc["addr"] = self.Addr

	return json.Marshal(doc)
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

type CPEInterfaces map[string]CPEInterface

func (self CPEInterfaces) GetBSON() (interface{}, error) {
	var out = []bson.M{}
	for k, v := range self {
		var b []byte
		{
			var err error
			b, err = bson.Marshal(v)
			if err != nil {
				return nil, err
			}
		}

		var obj_m bson.M
		{
			var err error
			err = bson.Unmarshal(b, &obj_m)
			if err != nil {
				return nil, err
			}
		}

		obj_m["_id"] = k
		out = append(out, obj_m)
	}

	return out, nil
}

func (self *CPEInterfaces) SetBSON(raw bson.Raw) error {
	var in = []bson.M{}
	{
		var err error
		err = raw.Unmarshal(&in)
		if err != nil {
			return err
		}
	}

	var out = CPEInterfaces{}

	for _, v := range in {
		var obj_b []byte
		{
			var err error
			obj_b, err = bson.Marshal(v)
			if err != nil {
				return err
			}
		}

		var obj CPEInterface
		{
			var err error
			err = bson.Unmarshal(obj_b, &obj)
			if err != nil {
				return err
			}
		}
		var k_any, k_found = v["_id"]
		if !k_found {
			return errors.New("No ID found")
		}

		var k, k_correct = k_any.(string)
		if !k_correct {
			return errors.New("Invalid key type")
		}

		out[k] = obj
	}

	*self = out

	return nil
}

type CPE struct {
	Name         string              `json:"name"`
	Connected    bool                `json:"connected"`
	Description  string              `json:"description"`
	Model        UUID                `json:"model"`
	Interfaces   CPEInterfaces       `json:"interfaces"`
	ConfigStatus ConfigurationStatus `json:"config_status"`
}

type Stat struct {
	CPE       UUID    `json:"cpeid"`
	Timestamp int64   `json:"time"`
	CPU       float64 `json:"cpu"`
	RAM       struct {
		Total    float64 `json:"total"`
		Buffered float64 `json:"buffered"`
		Shared   float64 `json:"shared"`
		Free     float64 `json:"free"`
	} `json:"memory"`
	Storage      float64 `json:"storage"`
	ProcActive   uint64  `json:"processes_active"`
	ProcSleeping uint64  `json:"processes_sleeping"`
	Interfaces   map[string]struct {
		Name string  `json:"name"`
		Tx   float64 `json:"tx"`
		Rx   float64 `json:"rx"`
	} `json:"interfaces"`
}

type CPEAgentResponse struct {
	Result json.RawMessage `json:"result"`
	Status CPEAgentError   `json:"error"`
}

type CPEStatSettings struct {
	Rules       []StatEventRule `json:"rules"`
	RuleTimeout int64           `json:"rule-timeout"`
}

type StatDaemonSettings struct {
	CPEList  map[UUID]CPEStatSettings `json:"cpe-list"`
	Interval int64                    `json:"interval"`
}

type LimitBetween struct {
	Upper float64 `json:"upper"`
	Lower float64 `json:"lower"`
}

type CPEEventData struct {
	ID       UUID            `json:"id"`
	Settings CPEStatSettings `json:"settings"`
}

type Event struct {
	Time    int64     `json:"time"`
	Payload EventData `json:"payload"`
}

type SimpleMask struct {
	UUID []UUID `json:"uuid"`
}

type CPEMask struct {
	UUID      []UUID `json:"uuid"`
	HasWLANs  []UUID `json:"has_wlans"`
	Connected *bool  `json:"connected"`
}
