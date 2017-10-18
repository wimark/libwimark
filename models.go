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
	Is_local  bool   `json:"is_local"`
}

type InterfaceType string

const (
	I2_4 = InterfaceType("2.4")
	I5_0 = InterfaceType("5.0")
)

type WPACommon struct {
	Suites []SecuritySuite `json:"suites"`
}

type WPA2PersonalData struct {
	WPACommon `bson:",inline"`
	PSK       string `json:"psk"`
}

type WPA2EnterpriseData struct {
	WPACommon            `bson:",inline"`
	NasID                string `json:"nasid"`
	PMKCaching           bool   `json:"pmkcaching"`
	RadiusAuthentication []UUID `json:"radiusauthentication"`
}

type WPAPersonalData struct {
	WPACommon `bson:",inline"`
	PSK       string `json:"psk"`
}

type WPAEnterpriseData struct {
	WPACommon            `bson:",inline"`
	NasID                string `json:"nasid"`
	RadiusAuthentication []UUID `json:"radiusauthentication"`
}

type SecuritySettings interface {
	is_security_settings()
}

type WLAN struct {
	Name               string        `json:"name"`
	SSID               string        `json:"ssid"`
	Description        string        `json:"description"`
	Security           *EnumSecurity `json:"security"`
	VLAN               int           `json:"vlan"`
	Hidden             bool          `json:"hidden"`
	NasID              *string       `json:"nas_id"`
	RadiusAcctServers  []UUID        `json:"radius_acct_servers"`
	RadiusAcctInterval int           `json:"radius_acct_interval"`
	WhiteList          []string      `json:"whitelist"`
	BlackList          []string      `json:"blacklist"`
	FilterMode         MacFilterType `json:"filtermode"`
	L2Isolate          bool          `json:"l2isolate"`
	PMKCaching         bool          `json:"pmkcaching"`
	Roaming80211r      bool          `json:"roam80211r"`
}

type WiredConfig struct {
}

type WiredState struct {
}

type WiredData struct {
	Name   string      `json:"name"`
	Config WiredConfig `json:"config,omitempty"`
	State  WiredState  `json:"state,omitempty"`
}

type WiFiConfig struct {
	BandMode   string `json:"bandmode,omitempty"`
	Bandwidth  string `json:"bandwidth,omitempty"`
	TxPower    string `json:"txpower,omitempty"`
	WLANs      []UUID `json:"wlans,omitempty"`
	Channels   []int  `json:"channels,omitempty"`
	Country    string `json:"country,omitempty"`
	MaxClients int    `json:"maxclients,omitempty"`
}

type WlanState struct {
	State CPEInterfaceState `json:"state,omitempty"`
}

type WiFiState struct {
	Frequency  string             `json:"frequency,omitempty"`
	BandMode   string             `json:"bandmode,omitempty"`
	Bandwidth  string             `json:"bandwidth,omitempty"`
	Channel    string             `json:"channel,omitempty"`
	TxPower    string             `json:"txpower,omitempty"`
	Enabled    bool               `json:"enabled"`
	WLANStates map[UUID]WlanState `json:"wlanstates,omitempty"`
}

type WiFiData struct {
	Name   string     `json:"name,omitempty"`
	Config WiFiConfig `json:"config,omitempty"`
	State  WiFiState  `json:"state,omitempty"`
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

func (self *CPEInterface) GetBSON() (interface{}, error) {
	var out = bson.M{}

	var obj_b []byte
	{
		var err error
		obj_b, err = bson.Marshal(self.CPEInterfaceInfo)
		if err != nil {
			return nil, err
		}
	}

	var obj bson.M
	{
		var err error
		err = bson.Unmarshal(obj_b, &obj)
		if err != nil {
			return nil, err
		}
	}
	out = obj
	out["capabilities"] = self.Capabilities
	out["addr"] = self.Addr

	return out, nil
}

func (self *CPEInterface) SetBSON(raw bson.Raw) error {
	var in = map[string]bson.Raw{}
	{
		if err := raw.Unmarshal(&in); err != nil {
			return err
		}
	}

	//addr
	{
		var v, k_found = in["addr"]
		if !k_found {
			return errors.New("No addr found")
		}
		if err := v.Unmarshal(&self.Addr); err != nil {
			return err
		}

		delete(in, "addr")
	}

	//capabilities
	{
		var v, k_found = in["capabilities"]
		if !k_found {
			return errors.New("No subject_id found")
		}
		if err := v.Unmarshal(&self.Capabilities); err != nil {
			return err
		}

		delete(in, "capabilities")
	}

	var obj_b, mErr = bson.Marshal(in)
	if mErr != nil {
		return mErr
	}

	if err := bson.Unmarshal(obj_b, &self.CPEInterfaceInfo); err != nil {
		return err
	}
	return nil
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
		obj_m, err := v.GetBSON()
		if err != nil {
			return nil, err
		}
		obj_m.(bson.M)["_id"] = k
		out = append(out, obj_m.(bson.M))
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
	IPAddr      string `json:"ipaddr"`
	MACAddr     string `json:"macaddr"`
	NetMask     string `json:"netmask"`
	Gateway     string `json:"gateway"`
	Model       struct {
		Id   UUID   `json:"id"`
		Name string `json:"name"`
	} `json:"model"`
	Interfaces       CPEInterfaces       `json:"interfaces"`
	ConfigStatus     ConfigurationStatus `json:"config_status"`
	LbsConfig        LBSConfig           `json:"lbs_config"`
	StatisticsConfig StatisticsConfig    `json:"stats_config"`
	LogConfig        LogConfig           `json:"log_config"`
	DHCPCapConfig    DHCPCapConfig       `json:"dhcpcap_config"`
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

type LBSConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod  int           `json:"reportperiod,omitempty"`
	ClientTimeout int           `json:"clienttimeout,omitempty"`
	WhiteList     []string      `json:"whitelist"`
	BlackList     []string      `json:"blacklist"`
	FilterMode    MacFilterType `json:"filtermode"`
}

type StatisticsConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod int `json:"reportperiod,omitempty"`
}

type LogConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod int `json:"reportperiod,omitempty"`
}

type DHCPCapConfig struct {
	Enabled       bool     `json:"enabled"`
	MsgTypeFilter []string `json:"msgtypefilter"`
}

type L2TPTunnel struct {
	Host     UUID   `json:"host"`
	CPE      UUID   `json:"cpe"`
	HostIP   string `json:"hostip"`
	CPEIP    string `json:"cpeip"`
	TunnelId int    `json:"tunid"`
}
