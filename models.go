package libwimark

import (
	"gopkg.in/mgo.v2/bson"
)

type Document map[string]interface{}

type UUID string

// ==== Radius ====

type Radius struct {
	Name      string `json:"name"`
	Hostname  string `json:"hostname"`
	Auth_port string `json:"auth_port"`
	Acc_port  string `json:"acc_port"`
	Secret    string `json:"secret"`
	Is_local  bool   `json:"is_local"`
}

// ==== WLANs ====

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
	Name               string           `json:"name"`
	SSID               string           `json:"ssid"`
	Description        string           `json:"description"`
	Security           EnumSecurity     `json:"security"`
	VLAN               int              `json:"vlan"`
	Hidden             bool             `json:"hidden"`
	NasID              *string          `json:"nas_id"`
	RadiusAcctServers  []UUID           `json:"radius_acct_servers"`
	RadiusAcctInterval int              `json:"radius_acct_interval"`
	WhiteList          []string         `json:"whitelist"`
	BlackList          []string         `json:"blacklist"`
	FilterMode         MacFilterType    `json:"filtermode"`
	L2Isolate          bool             `json:"l2isolate"`
	PMKCaching         bool             `json:"pmkcaching"`
	Roaming80211r      bool             `json:"roam80211r"`
	Tunneling          bool             `json:"tunneling"`
	DefaultTunnel      string           `json:"default_tunnel"`
	Firewall           FireWallSettings `json:"firewall"`
}

// ==== CPE ====

// ---- Service configs ----

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

type ScanConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod int `json:"reportperiod"`
	ScanTimeout  int `json:"scantimeout"`
	ScanNumber   int `json:"scannumber"`
}

type FirmwareConfig struct {
	FileUrl      string             `json:"file" bson:"file"`
	StorageUrl   string             `json:"storage" bson:"storage"`
	ChecksumUrl  string             `json:"checksum" bson:"checksum"`
	CheckTimeout int                `json:"timeout" bson:"timeout"`
	Mode         FirmwareUpdateMode `json:"mode" bson:"mode"`
}

// ---- Wifi config ----

type WlanConfig struct {
	L2TPEnabled         bool   `json:"l2tp_enabled"`
	CPETunnelId         int    `json:"cpe_tunnel_id"`
	CPESessionId        int    `json:"cpe_session_id"`
	CPEInterfaceName    string `json:"cpe_interface_name"`
	HostTunnelId        int    `json:"host_tunnel_id"`
	HostSessionId       int    `json:"host_session_id"`
	HostInterfaceName   string `json:"host_interface_name"`
	HostL2InterfaceName string `json:"host_l2interface_name"`
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

type WiFiConfigs map[string]WiFiConfig
type wifiCfg struct {
	Id       string     `bson:"_id"`
	Contents WiFiConfig `bson:",inline"`
}

func (self WiFiConfigs) GetBSON() (interface{}, error) {
	out := []wifiCfg{}
	for k, v := range self {
		out = append(out, wifiCfg{k, v})
	}
	return out, nil
}

func (self *WiFiConfigs) SetBSON(raw bson.Raw) error {
	var in = []wifiCfg{}
	var out = WiFiConfigs{}
	if err := raw.Unmarshal(&in); err != nil {
		return err
	}
	for _, v := range in {
		out[v.Id] = v.Contents
	}
	*self = out
	return nil
}

// ---- Wired switch config ----

type WiredVlanConfig struct {
	Vlan     int      `json:"vlan" bson:"vlan"`
	Ports    []string `json:"ports" bson:"ports"`
	Tunnel   string   `json:"tunnel" bson:"tunnel"`
	FakeWlan UUID     `json:"fake_wlan" bson:"fake_wlan"`
}
type WiredConfig struct {
	Vlans []WiredVlanConfig `json:"vlans" bson:"vlans"`
}
type WiredConfigs map[string]WiredConfig

// ---- CPE config ----

type CPEConfig struct {
	Wifi             WiFiConfigs      `json:"wifi" bson:"wifi"`
	Wired            WiredConfigs     `json:"wired" bson:"wired"`
	LbsConfig        LBSConfig        `json:"lbs_config" bson:"lbs_config"`
	StatisticsConfig StatisticsConfig `json:"stats_config" bson:"stats_config"`
	LogConfig        LogConfig        `json:"log_config" bson:"log_config"`
	DHCPCapConfig    DHCPCapConfig    `json:"dhcpcap_config" bson:"dhcpcap_config"`
	L2TPConfig       L2TPConfig       `json:"l2tp_config" bson:"l2tp_config"`
	Firewall         FireWallSettings `json:"firewall" bson:"firewall"`
	Firmware         FirmwareConfig   `json:"firmware" bson:"firmware"`
}

// ---- Service states ----

type FirmwareState struct {
	HasUpdate  bool    `json:"has_update" bson:"has_update"`
	CurrentMd5 string  `json:"current_md5" bson:"current_md5"`
	Version    Version `json:"version" bson:"version"`
}

// ---- Wifi state ----

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

type WiFiStates map[string]WiFiState
type wifiStat struct {
	Id       string    `bson:"_id"`
	Contents WiFiState `bson:",inline"`
}

func (self WiFiStates) GetBSON() (interface{}, error) {
	out := []wifiStat{}
	for k, v := range self {
		out = append(out, wifiStat{k, v})
	}
	return out, nil
}

func (self *WiFiStates) SetBSON(raw bson.Raw) error {
	var in = []wifiStat{}
	var out = WiFiStates{}
	if err := raw.Unmarshal(&in); err != nil {
		return err
	}
	for _, v := range in {
		out[v.Id] = v.Contents
	}
	*self = out
	return nil
}

// ---- Wired state ----

type WiredStates map[string]WiredState
type WiredState struct {
	CableIn bool `json:"cable_in" json:"cable_in"`
}

// ---- CPE state ----

type CPEState struct {
	Wifi     WiFiStates    `json:"wifi,omitempty"`
	Wired    WiredStates   `json:"wired,omitempty"`
	Firmware FirmwareState `json:"firmware,omitempty"`
}

// ---- CPE itself ----

type CPEModelLink struct {
	Id   UUID   `json:"id"`
	Name string `json:"name"`
}

type IPAddress struct {
	Addr    string `json:"ipaddr"`
	NetMask string `json:"netmask"`
}
type CPE struct {
	Name         string              `json:"name"`
	Connected    bool                `json:"connected"`
	Description  string              `json:"description"`
	IPAddr       string              `json:"ipaddr"`
	MACAddr      string              `json:"macaddr"`
	NetMask      string              `json:"netmask"`
	IPAddrs      []IPAddress         `json:"ipaddrs"`
	Gateway      string              `json:"gateway"`
	Model        CPEModelLink        `json:"model"`
	ConfigStatus ConfigurationStatus `json:"config_status"`

	Config CPEConfig `json:"config"`
	State  CPEState  `json:"state"`
}

// ==== Capabilities ====

type WiredPortCaps struct {
	Index  string `json:"index" bson:"index"`
	Number int    `json:"num" bson:"num"`
	Role   string `json:"role" bson:"role"`
	Type   string `json:"type" bson:"type"`
}
type WiredCapabilities struct {
	Switch string          `json:"switch" bson:"switch"`
	Ports  []WiredPortCaps `json:"ports" bson:"ports"`
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

type WifiCapabilities struct {
	TxPowers  []CapTxPower    `json:"txpwrlist"`
	HTModes   map[string]bool `json:"htmodelist"`
	HWModes   map[string]bool `json:"hwmodelist"`
	Channels  []CapChannel    `json:"freqlist"`
	TxOffset  int             `json:"txpower_offset"`
	Frequency string          `json:"frequency"`
}

type Capabilities struct {
	Wifi  map[string]WifiCapabilities  `json:"wifi" bson:"wifi"`
	Wired map[string]WiredCapabilities `json:"wired" bson:"wired"`
}

// ==== OBSOLETE Capabilities ====

type CPECapabilities map[string]WifiCapabilities

// ==== L2TP objects ====

type VPNHost struct {
	HostName   string       `json:"hostname"`
	OSUUID     UUID         `json:"os_uuid"`
	IpAddr     string       `json:"ipaddr"`
	Interfaces []string     `json:"interfaces"`
	State      ServiceState `json:"state"`
}

// ==== CPE model ====

type CPEModel struct {
	Name         string          `json:"name" bson:"name"`
	Description  string          `json:"description" bson:"description"`
	Capabilities CPECapabilities `json:"capabilities" bson:"capabilities"`
	Caps         Capabilities    `json:"caps" bson:"caps"`
	Firmwares    []CPEFirmware   `json:"firmwares" bson:"firmwares"`
	Version      Version         `json:"version" bson:"version"`
}

// ==== Config template ====

type ConfigRule struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Model       UUID   `json:"model" bson:"model"`
	CPEs        []UUID `json:"cpes" bson:"cpes"`
	Template    struct {
		WLANs     []UUID    `json:"wlans" bson:"wlans"`
		CpeConfig CPEConfig `json:"cpe_config_template" bson:"cpe_config_template"`
		Tags      []string  `json:"tags" bson:"tags"`
		Location  UUID      `json:"location" bson:"location"`
	} `json:"template" bson:"template"`

	Is_auto bool `json:"is_auto" bson:"is_auto"`
}

// ==== CPE firmware ====

type CPEFirmware struct {
	Name    string `json:"name" bson:"name"`
	Version string `json:"version" bson:"version"`
	URL     string `json:"url" bson:"url"`
}
