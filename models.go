package libwimark

import (
	"errors"

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
	Tunneling          bool          `json:"tunneling"`
	DefaultTunnel      string        `json:"default_tunnel"`
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

type WlanConfig struct {
	L2TPSession *UUID `json:"l2tpsession"`
}

type ScanConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod int `json:"reportperiod,omitempty"`
	ScanTimeout  int `json:"scantimeout,omitempty"`
	ScanNumber   int `json:"scannumber,omitempty"`
}

type WiFiConfig struct {
	BandMode   string              `json:"bandmode,omitempty"`
	Bandwidth  string              `json:"bandwidth,omitempty"`
	TxPower    string              `json:"txpower,omitempty"`
	WLANs      []UUID              `json:"wlans,omitempty"`
	WLANConfig map[UUID]WlanConfig `json:"wlanconfig,omitempty"`
	Channels   []int               `json:"channels,omitempty"`
	Country    string              `json:"country,omitempty"`
	MaxClients int                 `json:"maxclients,omitempty"`
	ScanConfig ScanConfig          `json:"scanningconfig,omitempty"`
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

type CPEInterfaces map[string]CPEInterfaceInfo

func (self CPEInterfaces) GetBSON() (interface{}, error) {

	out := []bson.M{}

	for k, v := range self {
		obj_m := bson.M{}

		obj_m["_id"] = k
		obj_m["type"] = v.Type
		obj_m["data"] = v.Data
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

		var obj CPEInterfaceInfo
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
	L2TPConfig       L2TPConfig          `json:"l2tp_config"`
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
	TxPowers  []CapTxPower    `json:"txpwrlist"`
	HTModes   map[string]bool `json:"htmodelist"`
	HWModes   map[string]bool `json:"hwmodelist"`
	Channels  []CapChannel    `json:"freqlist"`
	TxOffset  int             `json:"txpower_offset"`
	Frequency string          `json:"frequency"`
}

type CPECapabilities map[string]Capabilities

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

type L2TPConfig struct {
	Enabled           bool   `json:"enabled"`
	VPNHost           UUID   `json:"host,omitempty"`
	LocalVPNAddress   string `json:"local_vpn_addr,omitempty"`
	LocalVPNInterface string `json:"local_vpn_iface,omitempty"`
	HostTunnelId      int    `json:"host_tunnel,omitempty"`
	LocalTunnelId     int    `json:"local_tunnel,omitempty"`
}

// L2TP objects

type VPNHost struct {
	HostName   string       `json:"hostname"`
	OSUUID     UUID         `json:"os_uuid"`
	IpAddr     string       `json:"ipaddr"`
	Interfaces []string     `json:"interfaces"`
	State      ServiceState `json:"state"`
}

type L2TPTunnelSession struct {
	CPE                 UUID   `json:"cpe"`
	CPETunnelId         int    `json:"cpe_tunnel_id"`
	CPESessionId        int    `json:"cpe_session_id"`
	CPEInterfaceName    string `json:"cpe_interface_name"`
	Host                UUID   `json:"host"`
	HostTunnelId        int    `json:"host_tunnel_id"`
	HostSessionId       int    `json:"host_session_id"`
	HostInterfaceName   string `json:"host_interface_name"`
	HostL2InterfaceName string `json:"host_l2interface_name"`
}

// CPE models

type CPEConfigTemplate struct {
	Description      string                `json:"description" bson:"description"`
	Wifi             map[string]WiFiConfig `json:"wifi" bson:"wifi"`
	LbsConfig        LBSConfig             `json:"lbs_config" bson:"lbs_config"`
	StatisticsConfig StatisticsConfig      `json:"stats_config" bson:"stats_config"`
	LogConfig        LogConfig             `json:"log_config" bson:"log_config"`
	DHCPCapConfig    DHCPCapConfig         `json:"dhcpcap_config" bson:"dhcpcap_config"`
	L2TPConfig       L2TPConfig            `json:"l2tp_config" bson:"l2tp_config"`
}

type CPEModel struct {
	Name         string          `json:"name" bson:"name"`
	Description  string          `json:"description" bson:"description"`
	Capabilities CPECapabilities `json:"capabilities" bson:"capabilities"`
}

type ConfigRule struct {
	Name     string `json:"name" bson:"name"`
	Model    UUID   `json:"model" bson:"model"`
	CPEs     []UUID `json:"cpes" bson:"cpes"`
	Template struct {
		WLANs     []UUID            `json:"wlans" bson:"wlans"`
		CpeConfig CPEConfigTemplate `json:"cpe_config_template" bson:"cpe_config_template"`
		Tag       string            `json:"tag" bson:"tag"`
		Location  string            `json:"location" bson:"location"`
	} `json:"template" bson:"template"`

	Is_auto bool `json:"is_auto" bson:"is_auto"`
}
