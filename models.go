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

type WLAN struct {
	Name               string        `json:"name"`
	SSID               string        `json:"ssid"`
	Description        string        `json:"description"`
	Security           EnumSecurity  `json:"security"`
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
	Config WiredConfig `json:"config"`
	State  WiredState  `json:"state"`
}

type WlanConfig struct {
	L2TPSession *UUID `json:"l2tpsession"`
}

type ScanConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod int `json:"reportperiod"`
	ScanTimeout  int `json:"scantimeout"`
	ScanNumber   int `json:"scannumber"`
}

type WiFiConfig struct {
	BandMode    string              `json:"bandmode"`
	Bandwidth   string              `json:"bandwidth"`
	TxPower     string              `json:"txpower"`
	WLANs       []UUID              `json:"wlans"`
	WLANConfig  map[UUID]WlanConfig `json:"wlanconfig"`
	Channels    []int               `json:"channels"`
	Country     string              `json:"country"`
	MaxClients  int                 `json:"maxclients"`
	ScanConfig  ScanConfig          `json:"scanningconfig"`
	RequireMode MCSRequire          `json:"require_mode"`
}

type WlanState struct {
	State CPEInterfaceState `json:"state"`
}

type WiFiState struct {
	Frequency  string             `json:"frequency"`
	BandMode   string             `json:"bandmode"`
	Bandwidth  string             `json:"bandwidth"`
	Channel    string             `json:"channel"`
	TxPower    string             `json:"txpower"`
	Enabled    bool               `json:"enabled"`
	WLANStates map[UUID]WlanState `json:"wlanstates"`
}

type WiFiData struct {
	Name   string     `json:"name"`
	Config WiFiConfig `json:"config"`
	State  WiFiState  `json:"state"`
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

type CPEState struct {
	Wifi map[string]WiFiState `json:"wifi,omitempty"`
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

	Config CPEConfigTemplate `json:"config"`
	State  CPEState          `json:"state"`
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
	ReportPeriod  int           `json:"reportperiod"`
	ClientTimeout int           `json:"clienttimeout"`
	WhiteList     []string      `json:"whitelist"`
	BlackList     []string      `json:"blacklist"`
	FilterMode    MacFilterType `json:"filtermode"`
}

type StatisticsConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod int `json:"reportperiod"`
}

type LogConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod int `json:"reportperiod"`
}

type DHCPCapConfig struct {
	Enabled       bool     `json:"enabled"`
	MsgTypeFilter []string `json:"msgtypefilter"`
}

type L2TPConfig struct {
	Enabled           bool   `json:"enabled"`
	VPNHost           UUID   `json:"host"`
	LocalVPNAddress   string `json:"local_vpn_addr"`
	LocalVPNInterface string `json:"local_vpn_iface"`
	HostTunnelId      int    `json:"host_tunnel"`
	LocalTunnelId     int    `json:"local_tunnel"`
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
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Model       UUID   `json:"model" bson:"model"`
	CPEs        []UUID `json:"cpes" bson:"cpes"`
	Template    struct {
		WLANs     []UUID            `json:"wlans" bson:"wlans"`
		CpeConfig CPEConfigTemplate `json:"cpe_config_template" bson:"cpe_config_template"`
		Tags      []string          `json:"tags" bson:"tags"`
		Location  UUID              `json:"location" bson:"location"`
	} `json:"template" bson:"template"`

	Is_auto bool `json:"is_auto" bson:"is_auto"`
}
