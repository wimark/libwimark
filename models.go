package libwimark

import (
	"github.com/globalsign/mgo/bson"
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
	IsPortal  bool   `json:"is_portal"`
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
	Hotspot20Profile     UUID   `json:"hotspot20_profile" bson:"hotspot20_profile"`
}

type WPAPersonalData struct {
	WPACommon `bson:",inline"`
	PSK       string `json:"psk"`
}

type WPAEnterpriseData struct {
	WPACommon            `bson:",inline"`
	NasID                string `json:"nasid"`
	RadiusAuthentication []UUID `json:"radiusauthentication"`
	Hotspot20Profile     UUID   `json:"hotspot20_profile" bson:"hotspot20_profile"`
}

type WMMCategoryConfig struct {
	CliCwMin int  `json:"cli_cw_min" bson:"cli_cw_min"`
	CliCwMax int  `json:"cli_cw_max" bson:"cli_cw_max"`
	CliAifs  int  `json:"cli_aifs" bson:"cli_aifs"`
	CliTxOp  int  `json:"cli_txop" bson:"cli_txop"`
	CliAcm   bool `json:"cli_acm" bson:"cli_acm"`

	AcCwMin int     `json:"ac_cw_min" bson:"ac_cw_min"`
	AcCwMax int     `json:"ac_cw_max" bson:"ac_cw_max"`
	AcAifs  int     `json:"ac_aifs" bson:"ac_aifs"`
	AcBurst float32 `json:"ac_burst" bson:"ac_burst"`
}
type WMMConfig struct {
	Categories map[WMMAccessCategory]WMMCategoryConfig `json:"categories" bson:"categories"`
	Disabled   bool                                    `json:"disabled" bson:"disabled"`
	Uapsd      bool                                    `json:"uapsd" bson:"uapsd"`
}

type WLAN struct {
	Name               string               `json:"name"`
	SSID               string               `json:"ssid"`
	Description        string               `json:"description"`
	Security           EnumSecurity         `json:"security"`
	VLAN               int                  `json:"vlan"`
	Hidden             bool                 `json:"hidden"`
	NasID              *string              `json:"nas_id"`
	NasPortID          string               `json:"nas_port_id" bson:"nas_port_id"`
	RadiusAcctServers  []UUID               `json:"radius_acct_servers"`
	RadiusAcctInterval int                  `json:"radius_acct_interval"`
	WhiteList          []string             `json:"whitelist"`
	BlackList          []string             `json:"blacklist"`
	FilterMode         MacFilterType        `json:"filtermode"`
	L2Isolate          bool                 `json:"l2isolate"`
	PMKCaching         bool                 `json:"pmkcaching"`
	Roaming80211r      bool                 `json:"roam80211r"`
	Tunneling          bool                 `json:"tunneling"`
	DefaultTunnel      string               `json:"default_tunnel"`
	Firewall           FireWallSettings     `json:"firewall"`
	GuestControl       GuestControlSettings `json:"guest_control"`
	WMMConfig          WMMConfig            `json:"wmm" bson:"wmm"`
	NAT                bool                 `json:"nat" bson:"nat"`
	NATNetwork         IPAddress            `json:"nat_network" bson:"nat_network"`
}

type WLANCompact struct {
	Name        string `json:"name" bson:"name"`
	SSID        string `json:"ssid" bson:"ssid"`
	Description string `json:"description" bson:"description"`
	// Security    EnumSecurity `json:"security" bson:"security"`
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
	LogIP     string `json:"log_ip" bson:"log_ip"`
	LogPrefix string `json:"log_prefix" bson:"log_prefix"`
	LogProto  string `json:"log_proto" bson:"log_proto"`
	LogRemote bool   `json:"enabled" bson:"log_remote"`
	LogPort   int    `json:"log_port" bson:"log_port"`
}

type DHCPCapConfig struct {
	Enabled       bool     `json:"enabled"`
	MsgTypeFilter []string `json:"msgtypefilter"`
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

type PowerConfig struct {
	Range  []int `json:"range"`
	IsAuto bool  `json:"auto"`
}

type WiFiConfig struct {
	BandMode    string      `json:"bandmode"`
	Bandwidth   string      `json:"bandwidth"`
	TxPower     string      `json:"txpower"`
	MinTxPower  string      `json:"mintxpower"`
	Power       PowerConfig `json:"power"`
	WLANs       []UUID      `json:"wlans"`
	Channels    []int       `json:"channels"`
	Country     string      `json:"country"`
	MaxClients  int         `json:"maxclients"`
	ScanConfig  ScanConfig  `json:"scanningconfig"`
	RequireMode MCSRequire  `json:"require_mode"`
	Frequency   string      `json:"frequency"`
}

type WiFiConfigs map[string]WiFiConfig
type wifiCfg struct {
	Id       string     `bson:"_id"`
	Contents WiFiConfig `bson:",inline"`
}

func (self WiFiConfigs) GetBSON() (interface{}, error) {
	var out = make([]wifiCfg, len(self))
	var index = 0
	for k, v := range self {
		out[index] = wifiCfg{k, v}
		index += 1
	}
	return out, nil
}

func (self *WiFiConfigs) SetBSON(raw bson.Raw) error {
	var in = []wifiCfg{}
	if err := raw.Unmarshal(&in); err != nil {
		return err
	}
	var out = make(WiFiConfigs, len(in))
	for _, v := range in {
		out[v.Id] = v.Contents
	}
	*self = out
	return nil
}

// ---- Tunnel config ----

type TunnelConfig struct {
	CPETunnelId         int    `json:"cpe_tunnel_id"`
	CPESessionId        int    `json:"cpe_session_id"`
	CPEInterfaceName    string `json:"cpe_interface_name"`
	HostTunnelId        int    `json:"host_tunnel_id"`
	HostSessionId       int    `json:"host_session_id"`
	HostInterfaceName   string `json:"host_interface_name"`
	HostL2InterfaceName string `json:"host_l2interface_name"`
}
type TunnelConfigs map[string]TunnelConfig

// ---- Wired switch config ----

type WiredVlanConfig struct {
	Vlan         int                  `json:"vlan" bson:"vlan"`
	Ports        []string             `json:"ports" bson:"ports"`
	Tunnel       string               `json:"tunnel" bson:"tunnel"`
	FakeWlan     UUID                 `json:"fake_wlan" bson:"fake_wlan"`
	Accounting   bool                 `json:"acct" bson:"acct"`
	Interface    string               `json:"interface" bson:"interface"`
	GuestControl GuestControlSettings `json:"guest_control" bson:"guest_control"`
}
type WiredConfig struct {
	PrimaryVlan int               `json:"primary_vlan" bson:"primary_vlan"`
	Vlans       []WiredVlanConfig `json:"vlans" bson:"vlans"`
}
type WiredConfigs map[string]WiredConfig
type wiredCfg struct {
	Id       string      `bson:"_id"`
	Contents WiredConfig `bson:",inline"`
}

func (self WiredConfigs) GetBSON() (interface{}, error) {
	var out = make([]wiredCfg, len(self))
	var index = 0
	for k, v := range self {
		out[index] = wiredCfg{k, v}
		index += 1
	}
	return out, nil
}

func (self *WiredConfigs) SetBSON(raw bson.Raw) error {
	var in = []wiredCfg{}
	if err := raw.Unmarshal(&in); err != nil {
		return err
	}
	var out = make(WiredConfigs, len(in))
	for _, v := range in {
		out[v.Id] = v.Contents
	}
	*self = out
	return nil
}

// ---- CPE config ----

type CPEConfig struct {
	Wifi             WiFiConfigs      `json:"wifi" bson:"wifi"`
	Wired            WiredConfigs     `json:"wired" bson:"wired"`
	LbsConfig        LBSConfig        `json:"lbs_config" bson:"lbs_config"`
	StatisticsConfig StatisticsConfig `json:"stats_config" bson:"stats_config"`
	LogConfig        LogConfig        `json:"log_config" bson:"log_config"`
	DHCPCapConfig    DHCPCapConfig    `json:"dhcpcap_config" bson:"dhcpcap_config"`
	Firewall         FireWallSettings `json:"firewall" bson:"firewall"`
	Firmware         FirmwareConfig   `json:"firmware" bson:"firmware"`
	Tunnels          TunnelConfigs    `json:"tunnels" bson:"tunnels"`
	Beeline          BeelineConfig    `json:"beeline_config" bson:"beeline_config"`
}

// ---- Beeline config ----

type BeelineConfig struct {
	NASIP string `json:"nas_ip" bson:"nas_ip"`
}

type CPEConfigItemCompact struct {
	Enabled bool `json:"enabled" bson:"enabled"`
}

type CPEConfigCompact struct {
	LbsConfig CPEConfigItemCompact `json:"lbs_config" bson:"lbs_config"`
}

// ---- Service states ----

type FirmwareState struct {
	HasUpdate  bool               `json:"has_update" bson:"has_update"`
	CurrentMd5 string             `json:"current_md5" bson:"current_md5"`
	Version    Version            `json:"version" bson:"version"`
	Packages   map[string]Version `json:"packages" bson:"packages"`
	Statics    map[string]Version `json:"statics" bson:"statics"`
}

// ---- General tunnel state ----

type L2TPState struct {
	Enabled        bool   `json:"enabled" bson:"enabled"`
	HostId         UUID   `json:"host" bson:"host"`
	HostAddr       string `json:"host_addr" bson:"host_addr"`
	HostTunnelId   int    `json:"host_tunnel" bson:"host_tunnel"`
	TunnelType     string `json:"tunnel_type" bson:"tunnel_type"`
	LocalAddr      string `json:"local_addr" bson:"local_addr"`
	LocalInterface string `json:"local_iface" bson:"local_iface"`
	LocalTunnelId  int    `json:"local_tunnel" bson:"local_tunnel"`
}

// ---- Wifi state ----

type WlanState struct {
	State        CPEInterfaceState `json:"state"`
	VirtualIface string            `json:"virtual_iface" bson:"virtual_iface"`
	BSSID        string            `json:"bssid" bson:"bssid"`
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
	var out = make([]wifiStat, len(self))
	var index = 0
	for k, v := range self {
		out[index] = wifiStat{k, v}
		index += 1
	}
	return out, nil
}

func (self *WiFiStates) SetBSON(raw bson.Raw) error {
	var in = []wifiStat{}
	if err := raw.Unmarshal(&in); err != nil {
		return err
	}
	var out = make(WiFiStates, len(in))
	for _, v := range in {
		out[v.Id] = v.Contents
	}
	*self = out
	return nil
}

// ---- Wired state ----

type WiredStates map[string]WiredState
type WiredState struct {
	CableIn     bool        `json:"cable_in" json:"cable_in"`
	SwitchVlans []VlanState `json:"vlans" json:"vlans"`
}
type VlanState struct {
	Vid    int      `json:"vid" bson:"vid"`
	Vlan   int      `json:"vlan" bson:"vlan"`
	Ports  []string `json:"ports" bson:"ports"`
	System bool     `json:"system" bson:"system"`
	Switch string   `json:"device" bson:"device"`
}

// ---- WAN state ----

type WanState struct {
	Interface string `json:"iface" bson:"iface"`
	Protocol  string `json:"proto" bson:"proto"`
}

// ---- Net state ----

type NetworkState struct {
	IPAddr  string      `json:"ipaddr"`
	MACAddr string      `json:"macaddr"`
	IPAddrs []IPAddress `json:"ipaddrs"`
	Gateway string      `json:"gateway"`
}

type NetworkStateCompact struct {
	IPAddr  string `json:"ipaddr" bson:"ipaddr"`
	MACAddr string `json:"macaddr" bson:"macaddr"`
}

// ---- CPE state ----

type CPEState struct {
	Wifi      WiFiStates    `json:"wifi,omitempty"`
	Wired     WiredStates   `json:"wired,omitempty"`
	Firmware  FirmwareState `json:"firmware,omitempty"`
	Wan       WanState      `json:"wan"`
	L2TPState L2TPState     `json:"l2tp_state" bson:"l2tp_state"`
	Network   NetworkState  `json:"network" bson:"network"`
	Tunnels   TunnelConfigs `json:"tunnels" bson:"tunnels"`
}

type CPEStateCompact struct {
	Network NetworkStateCompact `json:"network" bson:"network"`
}

// ---- CPE itself ----

type CPEModelLink struct {
	Id        UUID   `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short"`
}

type IPAddress struct {
	Addr    string `json:"ipaddr"`
	NetMask string `json:"netmask"`
}
type CPE struct {
	Name         string              `json:"name"`
	Connected    bool                `json:"connected"`
	Description  string              `json:"description"`
	Model        CPEModelLink        `json:"model"`
	ConfigStatus ConfigurationStatus `json:"config_status"`
	LastError    ModelError          `json:"last_error" bson:"last_error"`

	Config CPEConfig `json:"config"`
	State  CPEState  `json:"state"`
}

type CPECompact struct {
	Name         string              `json:"name" bson:"name"`
	Connected    bool                `json:"connected" bson:"connected"`
	Description  string              `json:"description" bson:"description"`
	Model        CPEModelLink        `json:"model" bson:"model"`
	ConfigStatus ConfigurationStatus `json:"config_status" bson:"configstatus"`

	Config CPEConfigCompact `json:"config" bson:"config"`
	State  CPEStateCompact  `json:"state" bson:"state"`
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

// ==== L2TP objects ====

type VPNHost struct {
	HostName   string       `json:"hostname"`
	OSUUID     UUID         `json:"os_uuid"`
	Interfaces []string     `json:"interfaces"`
	State      ServiceState `json:"state"`
}

// ==== CPE model ====

type CPEModel struct {
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Caps        Capabilities  `json:"caps" bson:"caps"`
	Firmwares   []CPEFirmware `json:"firmwares" bson:"firmwares"`
	Version     Version       `json:"version" bson:"version"`
}

// ==== Config template ====

type ConfigRule struct {
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Model       UUID      `json:"model" bson:"model"`
	CPEs        []UUID    `json:"cpes" bson:"cpes"`
	MacPrefix   string    `json:"mac_prefix" bson:"mac_prefix"`
	Subnet      IPAddress `json:"subnet" bson:"subnet"`
	Template    struct {
		WLANs     []UUID    `json:"wlans" bson:"wlans"`
		CpeConfig CPEConfig `json:"cpe_config_template" bson:"cpe_config_template"`
		Tags      []string  `json:"tags" bson:"tags"`
		Location  UUID      `json:"location" bson:"location"`
	} `json:"template" bson:"template"`

	Is_auto   bool `json:"is_auto" bson:"is_auto"`
	Is_always bool `json:"is_always" bson:"is_always"`
}

// ==== CPE firmware ====

type CPEFirmware struct {
	Name    string `json:"name" bson:"name"`
	Version string `json:"version" bson:"version"`
	URL     string `json:"url" bson:"url"`
}

// ==== RRM template ====

const COLLECTION_RRM_GROUPS = "rrm_groups"

type RRMGroup struct {
	Name string        `json:"name" bson:"name"`
	CPEs []UUID        `json:"cpes" bson:"cpes"`
	Algo RRMAlgoObject `json:"algo" bson:"algo"`
}
