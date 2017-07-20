package libwimark

import (
	"encoding/json"
	"errors"
	"sort"

	"github.com/vorot93/goutil"
	"gopkg.in/mgo.v2/bson"
)

type Document goutil.Document

type UUID string

type Radius struct {
	Name      string `json:"name"`
	Hostname  string `json:"hostname"`
	Auth_port string `json:"auth_port"`
	Acc_port  string `json:"acc_port"`
	Secret    string `json:"secret"`
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
	WPA2Common `bson:",inline"`
	PSK        string `json:"psk"`
}

type WPA2EnterpriseData struct {
	WPA2Common           `bson:",inline"`
	NasID                string `json:"nasid"`
	PMKCaching           bool   `json:"pmkcaching"`
	RadiusAuthentication []UUID `json:"radiusauthentication"`
}

type SecuritySettings interface {
	is_security_settings()
}

type WLAN struct {
	Name              string        `json:"name"`
	SSID              string        `json:"ssid"`
	Description       string        `json:"description"`
	Security          *EnumSecurity `json:"security"`
	VLAN              int           `json:"vlan"`
	Hidden            bool          `json:"hidden"`
	NasID             *string       `json:"nas_id"`
	RadiusAcctServers []UUID        `json:"radius_acct_servers"`
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
	Name      string `json:"name,omitempty"`
	Mac       string `json:"mac,omitempty"`
	Frequency string `json:"frequency,omitempty"`
	BandMode  string `json:"bandmode,omitempty"`
	Bandwidth string `json:"bandwidth,omitempty"`
	Channel   string `json:"channel,omitempty"`
	TxPower   string `json:"txpower,omitempty"`
	WLANs     []UUID `json:"wlans,omitempty"`
}

type CPEInterface struct {
	CPEInterfaceInfo
	Addr         string
	Capabilities Capabilities
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
	doc["capabilities"] = self.Capabilities

	return json.Marshal(doc)
}

func (self *CPEInterface) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	var err = json.Unmarshal(b, &doc)
	if err != nil {
		return err
	}

	if doc == nil {
		return nil
	}

	var addrRaw, addrExists = doc["addr"]
	if addrExists {
		var addr string
		var addrErr = json.Unmarshal(addrRaw, &addr)
		if addrErr == nil {
			self.Addr = addr
		} else {
			return addrErr
		}
	}

	delete(doc, "addr")

	var capsRaw, capsExists = doc["capabilities"]
	if capsExists {
		var caps Capabilities
		var capsErr = json.Unmarshal(capsRaw, &caps)
		if capsErr == nil {
			self.Capabilities = caps
		} else {
			return capsErr
		}
	}

	delete(doc, "capabilities")

	var v, _ = json.Marshal(doc)

	return self.CPEInterfaceInfo.UnmarshalJSON(v)
}

type CPEInterfaces map[string]CPEInterface

func (self CPEInterfaces) GetBSON() (interface{}, error) {
	var out = []bson.M{}
	var keys = []string{}
	for k, _ := range self {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := self[k]
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
	Name        string `json:"name"`
	Connected   bool   `json:"connected"`
	Description string `json:"description"`
	Model       struct {
		Id   UUID   `json:"id"`
		Name string `json:"name"`
	} `json:"model"`
	Interfaces   CPEInterfaces       `json:"interfaces"`
	ConfigStatus ConfigurationStatus `json:"config_status"`
	LbsConfig    LBSConfig           `json:"lbs_config"`
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
		Type string  `json:"type"`
		Tx   float64 `json:"tx"`
		Rx   float64 `json:"rx"`
	} `json:"interfaces"`
}

type ClientStat struct {
	AcctStatusType      ClientStatPacketType `json:"Acct-Status-Type"`
	CPE                 UUID                 `json:"cpeid"`
	CallingStationId    string               `json:"Calling-Station-Id"`
	UserName            string               `json:"User-Name"`
	AcctDelayTime       int                  `json:"Acct-Delay-Time"`
	AcctSessionId       string               `json:"Acct-Session-Id"`
	AcctInputGigawords  *int                 `json:"Acct-Input-Gigawords"`
	CalledStationId     string               `json:"Called-Station-Id"`
	AcctOutputGigawords *int                 `json:"Acct-Output-Gigawords"`
	AcctOutputOctets    *int                 `json:"Acct-Output-Octets"`
	AcctInputOctets     *int                 `json:"Acct-Input-Octets"`
	AcctSessionTime     *int                 `json:"Acct-Session-Time"`
	AcctInputPackets    *int                 `json:"Acct-Input-Packets"`
	AcctOutputPackets   *int                 `json:"Acct-Output-Packets"`
	Timestamp           int                  `json:"Timestamp"`
}

type CapTxPower struct {
	DBelMw    int `json:"dbm"`
	MilliWatt int `json:"mw"`
}

type CapChannel struct {
	Restricted bool       `json:"restricted"`
	Freq       int        `json:"mhz"`
	Channel    int        `json:"channel"`
	MaxPower   CapTxPower `json:"max_txpower"`
}

type Capabilities struct {
	TxPowers []CapTxPower    `json:"txpwrlist"`
	HTModes  map[string]bool `json:"htmodelist"`
	HWModes  map[string]bool `json:"hwmodelist"`
	Channels []CapChannel    `json:"freqlist"`
	TxOffset int             `json:"txpower_offset"`
}

type CPEAgentResponse struct {
	Result json.RawMessage `json:"result"`
	Status CPEAgentError   `json:"error"`
}

type LimitBetween struct {
	Upper float64 `json:"upper"`
	Lower float64 `json:"lower"`
}

type LimitBetweenOptional struct {
	Upper *float64 `json:"upper"`
	Lower *float64 `json:"lower"`
}

type CameraClient struct {
	Rtsp_stream string `json:"rtsp_stream"`
}

type OtherClient struct {
	Description string `json:"description"`
}

type CPEPollSettings struct {
	Rules []UUID `json:"rules"`
}

type StatEventRule struct {
	StatEventRuleObject `json:",inline" bson:",inline"`
	Name                string `json:"name"`
}

// Events
type SystemEvent struct {
	SystemEventObject `bson:",inline" json:",inline"`
	Timestamp         int64            `json:"timestamp"`
	Subject_id        string           `json:"subject_id"`
	Level             SystemEventLevel `json:"level"`
}

type CPEConnectedData struct {
	Cpeagent_version string `json:"cpeagent_version"`
	Os_version       string `json:"os_version"`
}

type MonitorRuleViolationData struct {
	Cpe_id UUID `json:"cpe_id"`
}

type ClientConnectedData struct {
	Mac     string `json:"mac"`
	Freq    string `json:"freq"`
	Cpe_id  UUID   `json:"cpe_id"`
	Wlan_id UUID   `json:"wlan_id"`
}

type ClientDisconnectedData struct {
	Mac     string `json:"mac"`
	Freq    string `json:"freq"`
	Cpe_id  UUID   `json:"cpe_id"`
	Wlan_id UUID   `json:"wlan_id"`
}

type CPEConfigurationErrorData struct {
	Description string `json:"description"`
}

type ServiceFatalErrorData struct {
	Sw_version  string `json:"sw_version"`
	Description string `json:"description"`
}

type ServiceConnectedData struct {
	Sw_version string `json:"sw_version"`
}

type LBSClientData struct {
	Timestamp int64  `json:"timestamp"`
	CPE       UUID   `json:"cpe"`
	Radio     string `json:"radio"`
	ClientMac string `json:"client_mac"`
	RSSI      int    `json:"rssi"`
}

type LBSCPEInfo struct {
	Group UUID    `json:"group"`
	CPE   UUID    `json:"cpe"`
	Name  string  `json:"name"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Z     float64 `json:"z"`
}

type LBSClientCoords struct {
	Timestamp int64   `json:"timestamp"`
	Group     UUID    `json:"group"`
	Mac       string  `json:"mac"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

type LBSConfig struct {
	Enabled bool `json:"enabled"`
	// in milliseconds
	ReportPeriod  int `json:"reportperiod,omitempty"`
	ClientTimeout int `json:"clienttimeout,omitempty"`
}

type SimpleMask struct {
	UUID []UUID `json:"uuid"`
}

type TimestampMask struct {
	UUID  []UUID `json:"uuid"`
	Start *int64 `json:"start"`
	Stop  *int64 `json:"stop"`
}

type EventMask struct {
	TimestampMask
	Type       []SystemEventObjectType `json:"type"`
	Subject_id []string                `json:"subject_id"`
	Level      []SystemEventLevel      `json:"level"`
}

type ClientStatMask struct {
	TimestampMask
	CPE              []UUID
	CallingStationId []string
}

type CPEMask struct {
	UUID      []UUID `json:"uuid"`
	HasWLANs  []UUID `json:"has_wlans"`
	Connected *bool  `json:"connected"`
}

type WLANMask struct {
	UUID      []UUID `json:"uuid"`
	HasRadius []UUID `json:"has_radius"`
}

type StatsMask struct {
	UUID    []UUID `json:"uuid"`
	CPEUUID []UUID `json:"cpe"`
	Start   *int64 `json:"start"`
	Stop    *int64 `json:"stop"`
}

type LBSClientDataMask struct {
	TimestampMask
	CPE       []UUID   `json:"cpe"`
	Radio     []string `json:"radio"`
	ClientMac []string `json:"client_mac"`
	RSSI      []int    `json:"rssi"`
}

type LBSCPEInfoMask struct {
	SimpleMask
	Group []UUID               `json:"group"`
	CPE   []UUID               `json:"cpe"`
	Name  []string             `json:"name"`
	X     LimitBetweenOptional `json:"x"`
	Y     LimitBetweenOptional `json:"y"`
	Z     LimitBetweenOptional `json:"z"`
}

type LBSClientCoordsMask struct {
	TimestampMask
	Group []UUID               `json:"group"`
	Mac   []string             `json:"mac"`
	X     LimitBetweenOptional `json:"x"`
	Y     LimitBetweenOptional `json:"y"`
	Z     LimitBetweenOptional `json:"z"`
}

type WirelessClient struct {
	WirelessClientObject `json:",inline" bson:",inline"`
	Mac                  string `json:"mac"`
	Ip                   string `json:"ip"`
	Net_mask             string `json:"net_mask"`
	Wlan_id              string `json:"wlan_id"`
	Cpe_id               string `json:"cpe_id"`
	Channel              int    `json:"channel"`
	Rssi                 int    `json:"rssi"`
	Timestamp            int    `json:"timestamp"`
}
