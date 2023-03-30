package libwimark

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

type Document map[string]interface{}

type UUID string

// ==== Radius ====

type Radius struct {
	Name             string `json:"name"`
	Hostname         string `json:"hostname"`
	Auth_port        string `json:"auth_port"`
	Acc_port         string `json:"acc_port"`
	Secret           string `json:"secret"`
	Is_local         bool   `json:"is_local"`
	IsPortal         bool   `json:"is_portal"`
	DaeClient        string `json:"dae_client"`
	DaeSecret        string `json:"dae_secret"`
	DaePort          string `json:"dae_port"`
	ACLAuthorized    string `json:"acl_authorized"`
	ACLNotAuthorized string `json:"acl_not_authorized"`
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

type SpeedConfig struct {
	Value int       `json:"value" bson:"value"`
	Type  SpeedType `json:"type" bson:"type"`
}

func (sc *SpeedConfig) String() string {
	return fmt.Sprintf("%d%s", sc.Value, sc.Type.String())
}

type WLAN struct {
	Name        string       `json:"name"`
	SSID        string       `json:"ssid"`
	Description string       `json:"description"`
	Security    EnumSecurity `json:"security"`

	// RADIUS section
	NasID               *string `json:"nas_id"`
	NasPortID           string  `json:"nas_port_id" bson:"nas_port_id"`
	NasIPAddress        string  `json:"nas_ip_address" bson:"nas_ip_address"`
	RadiusAcctServers   []UUID  `json:"radius_acct_servers"`
	RadiusAcctInterval  int     `json:"radius_acct_interval"`
	RadiusAcctMirroring bool    `json:"radius_acct_mirroring" bson:"radius_acct_mirroring"`

	// ACL section
	FilterMode MacFilterType    `json:"filtermode"`
	WhiteList  []string         `json:"whitelist"`
	BlackList  []string         `json:"blacklist"`
	Firewall   FireWallSettings `json:"firewall"`

	// WLAN related specifics
	Hidden    bool      `json:"hidden"`
	L2Isolate bool      `json:"l2isolate"`
	WMMConfig WMMConfig `json:"wmm" bson:"wmm"`

	// traffic shaping
	SpeedUpload   SpeedConfig `json:"speed_upload" bson:"speed_upload"`
	SpeedDownload SpeedConfig `json:"speed_download" bson:"speed_download"`

	// network section
	// remote tunneling
	Tunneling     bool       `json:"tunneling"`      // turn on tunneling
	Proto         TunnelType `json:"proto"`          // l2tp or gretap or ??
	DefaultTunnel string     `json:"default_tunnel"` // network name on server (only for l2tp)
	PeerAddress   string     `json:"peer_address"`   // peer address of server

	// local vlan (if 0 than untag)
	VLAN int `json:"vlan"`

	// local nat
	NAT        bool      `json:"nat" bson:"nat"`
	NATNetwork IPAddress `json:"nat_network" bson:"nat_network"`

	// Portal authorization section
	GuestControl GuestControlSettings `json:"guest_control"`

	// clients specifics section
	BeelineAccountingType string `json:"beeline_accountng_type"`

	// 802.11r
	PMKCaching    bool `json:"pmkcaching"`
	Roaming80211r bool `json:"roam80211r"`
	FTOverDS      bool `json:"ft_over_ds"`
	// generate NAS ID (for roaming - will be generated from bssid)
	NASGenerate bool `json:"nas_generate"`

	// 802.11k
	RoamingIEEE80211k bool `json:"ieee80211k"`
	RrmNeighborReport bool `json:"rrm_neighbor_report"`
	RrmBeaconReport   bool `json:"rrm_beacon_report"`

	// 802.11v
	RoamingIEEE80211v bool `json:"ieee80211v"`
	WnmSleepMode      bool `json:"wnm_sleep_mode"`
	BssTransition     bool `json:"bss_transition"`

	// roaming marketing
	RSSIThreshold int  `json:"rssi_threshold"`
	BandSteering  bool `json:"band_steering"`
	LoadBalancing bool `json:"load_balancing"`

	// for wmwdisd - rssi based disconnector
	SignalConnect    int `json:"signal_connect"`
	SignalStay       int `json:"signal_stay"`
	SignalStrikes    int `json:"signal_strikes"`
	SignalPollTime   int `json:"signal_poll_time"`
	SignalDropReason int `json:"signal_drop_reason"`

	//option 82
	Option82 Option82 `json:"option82"`
}

type Option82 struct {
	Enable bool `json:"option82_enable"`
	//todo; check ascii data
	CircuetID string `json:"option82_cid"`
	RemoteID  string `json:"option82_rid"`
}

type WLANCompact struct {
	Name        string       `json:"name" bson:"name"`
	SSID        string       `json:"ssid" bson:"ssid"`
	Description string       `json:"description" bson:"description"`
	Security    EnumSecurity `json:"security"`

	// network section
	// remote tunneling
	Tunneling     bool       `json:"tunneling"`      // turn on tunneling
	Proto         TunnelType `json:"proto"`          // l2tp or gretap or ??
	DefaultTunnel string     `json:"default_tunnel"` // network name on server (only for l2tp)
	PeerAddress   string     `json:"peer_address"`   // peer address of server
}

// ==== CPE ====

// ---- Service configs ----

type LBSConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds

	ReportPeriod  int `json:"reportperiod"`
	ClientTimeout int `json:"clienttimeout"`

	MaxQuiet      int  `json:"maxquiet"`
	MaxCacheQueue int  `json:"maxcachequeue"`
	EmptyWatcher  bool `json:"emptywatcher"`

	WhiteList  []string      `json:"whitelist"`
	BlackList  []string      `json:"blacklist"`
	FilterMode MacFilterType `json:"filtermode"`
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
	LogPort   int    `json:"log_port" bson:"log_port"`
	LogRemote bool   `json:"enabled" bson:"log_remote"`
}

type DHCPCapConfig struct {
	MsgTypeFilter []string `json:"msgtypefilter"`
	Enabled       bool     `json:"enabled"`
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
	BandMode       string      `json:"bandmode"`
	Bandwidth      string      `json:"bandwidth"`
	TxPower        string      `json:"txpower"`
	MinTxPower     string      `json:"mintxpower"`
	Power          PowerConfig `json:"power"`
	WLANs          []UUID      `json:"wlans"`
	Channels       []int       `json:"channels"`
	Country        string      `json:"country"`
	MaxClients     int         `json:"maxclients"`
	ScanConfig     ScanConfig  `json:"scanningconfig"`
	RequireMode    MCSRequire  `json:"require_mode"`
	Frequency      string      `json:"frequency"`
	BasicRate      string      `json:"basic_rate"`
	SupportedRates interface{} `json:"supported_rates"`
	LegacyRates    string      `json:"legacy_rates"`
	LogLevel       string      `json:"log_level"`
	MaxInactivity  int         `json:"max_inactivity"`
	CellDensity    int         `json:"cell_density"`
}

type WiFiConfigs map[string]WiFiConfig
type wifiCfg struct {
	Id       string     `bson:"_id"`
	Contents WiFiConfig `bson:",inline"`
}

func (wc WiFiConfigs) GetBSON() (interface{}, error) {
	var out = make([]wifiCfg, len(wc))
	var index = 0
	for k, v := range wc {
		out[index] = wifiCfg{k, v}
		index += 1
	}
	return out, nil
}

func (wc *WiFiConfigs) SetBSON(raw bson.Raw) error {
	var in = []wifiCfg{}
	if err := raw.Unmarshal(&in); err != nil {
		return err
	}
	var out = make(WiFiConfigs, len(in))
	for _, v := range in {
		out[v.Id] = v.Contents
	}
	*wc = out
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

type WiredSpeedConfig struct {
	Max  int       `json:"max" bson:"max"`
	Min  int       `json:"min" bson:"min"`
	Type SpeedType `json:"type" bson:"type"`
}

type WiredVlanConfig struct {
	Vlan  int      `json:"vlan" bson:"vlan"`
	Vid   int      `json:"vid" bson:"vid"`
	Ports []string `json:"ports" bson:"ports"`

	TunnelProto TunnelType `json:"tunnel_proto" bson:"tunnel_proto"`
	Tunnel      string     `json:"tunnel" bson:"tunnel"`
	Interface   string     `json:"interface" bson:"interface"`
	PeerAddr    string     `json:"peer_addr" bson:"peer_addr"`

	Accounting   bool                 `json:"acct" bson:"acct"`
	FakeWlan     UUID                 `json:"fake_wlan" bson:"fake_wlan"`
	GuestControl GuestControlSettings `json:"guest_control" bson:"guest_control"`

	NAT        bool      `json:"nat" bson:"nat"`
	NATNetwork IPAddress `json:"nat_network" bson:"nat_network"`
	NatAccess  bool      `json:"nat_access" bson:"nat_access"`

	SpeedDownload WiredSpeedConfig `json:"speed_download" bson:"speed_download"`
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

func (wc WiredConfigs) GetBSON() (interface{}, error) {
	var out = make([]wiredCfg, len(wc))
	var index = 0
	for k, v := range wc {
		out[index] = wiredCfg{k, v}
		index += 1
	}
	return out, nil
}

func (wc *WiredConfigs) SetBSON(raw bson.Raw) error {
	var in = []wiredCfg{}
	if err := raw.Unmarshal(&in); err != nil {
		return err
	}
	var out = make(WiredConfigs, len(in))
	for _, v := range in {
		out[v.Id] = v.Contents
	}
	*wc = out
	return nil
}

// Network Manual interface

// ---- CPE config ----

type CPEConfig struct {
	Name             string           `json:"name" bson:"name"`
	Description      string           `json:"description" bson:"description"`
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
	Wmsnmpd          WMSNMPDConfig    `json:"wmsnmpd" bson:"wmsnmpd"`

	NTPServerConfig   NTPServerConfig `json:"ntp_config" bson:"ntp_config"`
	WiFiLock          bool            `json:"wifi_lock" bson:"wifi_lock"`
	GrePeerAddrConfig string          `json:"gre_peer_addr" bson:"gre_peer_addr"`

	NetManual  NetManual  `json:"net_manual" bson:"net_manual"`
	WifiManual WifiManual `json:"wifi_manual" bson:"wifi_manual"`
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

type NetManual map[string]UciNetIface
type WifiManual map[string]UciWifiWlan

// ---- Service states ----

type FirmwareState struct {
	HasUpdate  bool     `json:"has_update" bson:"has_update"`
	CurrentMd5 string   `json:"current_md5" bson:"current_md5"`
	Version    Version  `json:"version" bson:"version"`
	Features   []string `json:"features" bson:"features"`
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
// -->> state
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

func (wc WiFiStates) GetBSON() (interface{}, error) {
	var out = make([]wifiStat, len(wc))
	var index = 0
	for k, v := range wc {
		out[index] = wifiStat{k, v}
		index += 1
	}
	return out, nil
}

func (wc *WiFiStates) SetBSON(raw bson.Raw) error {
	var in = []wifiStat{}
	if err := raw.Unmarshal(&in); err != nil {
		return err
	}
	var out = make(WiFiStates, len(in))
	for _, v := range in {
		out[v.Id] = v.Contents
	}
	*wc = out
	return nil
}

// ---- Wired state ----

type WiredStates map[string]WiredState
type WiredState struct {
	CableIn     bool        `json:"cable_in" bson:"cablein"`
	SwitchVlans []VlanState `json:"vlans" bson:"switchvlans"`
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

	NetManual  NetManual  `json:"net_manual" bson:"net_manual"`
	WifiManual WifiManual `json:"wifi_manual" bson:"wifi_manual"`
}

type CPEStateCompact struct {
	Network NetworkStateCompact `json:"network" bson:"network"`
}

// ---- CPE itwc ----

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
	Name        string       `json:"name"`
	Connected   bool         `json:"connected"`
	Description string       `json:"description"`
	Model       CPEModelLink `json:"model"`

	ConfigStatus ConfigurationStatus `json:"config_status"`
	LastError    ModelError          `json:"last_error" bson:"last_error"`

	Config CPEConfig `json:"config"`
	State  CPEState  `json:"state"`

	FirstConnection   int64 `json:"first_connection" bson:"first_connection"`
	LastConnection    int64 `json:"last_connection" bson:"last_connection"`
	LastDisconnection int64 `json:"last_disconnection" bson:"last_disconnection"`

	// reconfiguration
	ConfigNotSend bool `json:"config_not_send" bson:"config_not_send"`

	Latitude  float64 `json:"latitude"  bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
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
	HostName   string           `json:"hostname"`
	OSUUID     UUID             `json:"os_uuid"`
	Interfaces []string         `json:"interfaces"`
	State      ServiceState     `json:"state"`
	Ifaces     []LinkDescriptor `json:"interfaces_details"`
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

type RRMGroup struct {
	Name  string        `json:"name"  bson:"name"`
	CPEs  []UUID        `json:"cpes"  bson:"cpes"`
	Algo  RRMAlgoObject `json:"algo"  bson:"algo"`
	Force bool          `json:"force" bson:"force"`
}

// ==== WSNMPD config =====

type WMSNMPDConfig struct {
	Default WMSNMPDDefault `json:"default"`
}

type WMSNMPDDefault struct {
	Enabled         bool     `json:"enabled"`
	Community       string   `json:"community"`
	Location        string   `json:"location"`
	ListenInterface string   `json:"listen_interface,omitempty"`
	Interfaces      []string `json:"interfaces"`
}

type NTPServerConfig struct {
	Type    string   `json:"type" bson:"type"`
	Enable  bool     `json:"enable" bson:"enable"`
	Servers []string `json:"servers" bson:"servers"`
}
