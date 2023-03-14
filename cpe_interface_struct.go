package libwimark

import (
	"time"
)

//============= UCI config contents ================

//------------- Wimark config ----------------

type UciCpeagentCfg struct {
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	TunnelIFace  string `json:"tunneliface,omitempty"`
	HostUUID     UUID   `json:"tunneluuid,omitempty"`
	TunnelType   string `json:"tunnel,omitempty"`
	IpsecEncrypt string `json:"ipsec_disable_encryption,omitempty"`
}
type UciBrokerCfg struct {
	Host string `json:"host"`
}
type UciLbsCfg struct {
	Enabled        string   `json:"enabled"`
	ReportPeriod   string   `json:"report_timeout"`
	ClearPeriod    string   `json:"clear_timeout"`
	ClientTimeout  string   `json:"clear_delta"`
	ReceiverModule string   `json:"receiver"`
	ReceiverId     string   `json:"receiver_id"`
	ReqOperation   string   `json:"crud_op"`
	ModelName      string   `json:"model_name"`
	MacList        []string `json:"maclist"`
	MacFilter      string   `json:"macfilter"`
	MaxQuiet       string   `json:"max_quiet"`
	MaxCacheQueue  string   `json:"max_cache_queue"`
	EmptyWatcher   string   `json:"empty_watcher"`
	Type           string   `json:".type,omitempty"`
}
type UciStatCfg struct {
	Enabled        string `json:"enabled"`
	ReportPeriod   string `json:"timeout"`
	ReceiverModule string `json:"receiver"`
	ReceiverId     string `json:"receiver_id"`
	ReqOperation   string `json:"crud_op"`
	ModelName      string `json:"model_name"`
	Tag            string `json:"tag"`
	Type           string `json:".type,omitempty"`
}
type UciDhcpcapCfg struct {
	Enabled        string   `json:"enabled"`
	MsgTypeFilter  []string `json:"typefilter,omitempty"`
	ReceiverModule string   `json:"receiver"`
	ReceiverId     string   `json:"receiver_id"`
	ReqOperation   string   `json:"crud_op"`
	ModelName      string   `json:"model_name"`
	Type           string   `json:".type,omitempty"`
}
type UciScanningCfg struct {
	Enabled        string `json:"enabled"`
	ReportPeriod   string `json:"report_timeout"`
	ScanTimeout    string `json:"scan_timeout"`
	ScanNumber     string `json:"scan_times"`
	ReceiverModule string `json:"receiver,omitempty"`
	ReceiverId     string `json:"receiver_id,omitempty"`
	ReqOperation   string `json:"crud_op,omitempty"`
	ModelName      string `json:"model_name,omitempty"`
	Tag            string `json:"tag,omitempty"`
	Type           string `json:".type,omitempty"`
}
type UciFirmwareCfg struct {
	FileUrl      string `json:"file"`
	StorageUrl   string `json:"storage"`
	ChecksumUrl  string `json:"checksum"`
	CheckTimeout string `json:"timeout"`
	Mode         string `json:"mode"`
	CurrentMD5   string `json:"current_md5,omitempty"`
	AvailableMD5 string `json:"available_md5,omitempty"`
}

type UciWsnmpd struct {
	Default UciWsnmpdDefautl `json:"default,omitempty"`
}

type UciWsnmpdDefautl struct {
	Enabled         string   `json:"enabled"`
	Community       string   `json:"community"`
	Location        string   `json:"location"`
	ListenInterface string   `json:"listen_interface"`
	Interfaces      []string `json:"interfaces"`
}

type UciEthernetAcct struct {
	Type       string   `json:".type,omitempty"`
	Interfaces []string `json:"interface,omitempty"`
}
type UciEthernetStat struct {
	Type       string   `json:".type,omitempty"`
	Enabled    string   `json:"enabled"`
	Interfaces []string `json:"ifaces,omitempty"`
	Timeout    string   `json:"timeout"`
}
type UciRedirectCfg struct {
	Id           UUID     `json:"id"`
	FormattedUrl string   `json:"redirect_url_format"`
	MacWhiteList []string `json:"whitelist"`
	PreAuthList  []string `json:"preauth"`
	DnsResolve   []string `json:"resolve"`
	LanPortal    string   `json:"lan_portal"`
	EnableHttps  string   `json:"enable_https"`
	Type         string   `json:".type,omitempty"`
}
type innerUciWimark struct {
	CpeAgent         UciCpeagentCfg            `json:"broker"`
	Lbs              UciLbsCfg                 `json:"lbs"`
	Stat             UciStatCfg                `json:"statistic"`
	DhcpCap          UciDhcpcapCfg             `json:"dhcpcap"`
	EthernetAcct     UciEthernetAcct           `json:"-" inline:"yes,.type:eth_acct,unique,omitempty,eth_acct"`
	EthernetStat     UciEthernetStat           `json:"-" inline:"yes,.type:wsmd,unique,omitempty,wsmd"`
	Firmware         UciFirmwareCfg            `json:"-" inline:"yes,.type:firmware,unique,.firmware"`
	ScanSettings     map[string]UciScanningCfg `json:"-" inline:"yes,.type:scanning"`
	RedirectSettings map[string]UciRedirectCfg `json:"-" inline:"yes,.type:redirect"`
	Brokers          map[string]UciBrokerCfg   `json:"-" inline:"yes,.type:broker"`
}
type UciWimark innerUciWimark

//------------- Wireless config ----------------

type UciWifiWlan struct {
	Type              string      `json:".type,omitempty"`
	SSID              string      `json:"ssid,omitempty"`
	UUID              string      `json:"uuid,omitempty"`
	Iface             string      `json:"device,omitempty"`
	Hidden            string      `json:"hidden,omitempty"`
	WifiMode          string      `json:"mode,omitempty"`
	PMKCaching        string      `json:"auth_cache,omitempty"`
	IEEE80211r        string      `json:"ieee80211r,omitempty"`
	IEEE80211k        string      `json:"ieee80211k,omitempty"`
	IEEE80211v        string      `json:"ieee80211v,omitempty"`
	RrmNeighborReport string      `json:"rrm_neighbor_report,omitempty"`
	RrmBeaconReport   string      `json:"rrm_beacon_report,omitempty"`
	WnmSleepMode      string      `json:"wnm_sleep_mode,omitempty"`
	BssTransition     string      `json:"bss_transition,omitempty"`
	MobDomain         string      `json:"mobility_domain,omitempty"`
	R0KH              interface{} `json:"r0kh,omitempty"`
	R1KH              interface{} `json:"r1kh,omitempty"`
	VLAN              interface{} `json:"network,omitempty"`
	MacList           interface{} `json:"maclist,omitempty"`
	MacFilter         interface{} `json:"macfilter,omitempty"`
	L2Isolate         string      `json:"isolate"`
	MaxClients        string      `json:"maxassoc,omitempty"`
	OwnIp             string      `json:"ownip,omitempty"`
	Redirect          string      `json:"wimark_redirect,omitempty"`
	MACAddr           string      `json:"macaddr,omitempty"`
	UbusAcctPeriod    string      `json:"ubus_acct_interval,omitempty"`

	// security
	NasID       string `json:"nasid,omitempty"`
	NasPortID   string `json:"nas_port_id,omitempty"`
	AcctHost    string `json:"acct_server,omitempty"`
	AcctSecret  string `json:"acct_secret,omitempty"`
	AcctPeriod  string `json:"acct_interval,omitempty"`
	AcctPort    string `json:"acct_port,omitempty"`
	AuthHost    string `json:"auth_server,omitempty"`
	AuthSecret  string `json:"auth_secret,omitempty"`
	AuthPort    string `json:"auth_port,omitempty"`
	AcctServers string `json:"acct_server_list,omitempty"`
	AuthServers string `json:"auth_server_list,omitempty"`
	WMNasID     string `json:"wimark_nasid,omitempty"`
	SecType     string `json:"encryption,omitempty"`
	Password    string `json:"key,omitempty"`

	// hotspot 2.0
	HotspotEnabled      string      `json:"hotspot20,omitempty"`
	HSNetType           string      `json:"access_network_type,omitempty"`
	HSInternet          string      `json:"internet,omitempty"`
	HSAsra              string      `json:"asra,omitempty"`
	HSEsr               string      `json:"esr,omitempty"`
	HSUesa              string      `json:"uesa,omitempty"`
	HSVenueGroup        string      `json:"venue_group,omitempty"`
	HSVenueType         string      `json:"venue_type,omitempty"`
	HSVenueNames        interface{} `json:"venue_name,omitempty"`
	HS_HESSID           string      `json:"hessid,omitempty"`
	HSConsortium        interface{} `json:"roaming_consortium,omitempty"`
	HSDomains           interface{} `json:"domain_name,omitempty"`
	HSCellNetworks      interface{} `json:"3gpp_cell_net,omitempty"`
	HSNaiRealms         interface{} `json:"nai_realm,omitempty"`
	HSIpTypeAvail       string      `json:"ipaddr_type_availability,omitempty"`
	HSWanMetrics        string      `json:"wan_metrics,omitempty"`
	HSDgaf              string      `json:"dgaf,omitempty"`
	HSOperFriendlyNames interface{} `json:"oper_friendly_name,omitempty"`
	HSConnCaps          interface{} `json:"conn_capab,omitempty"`
	HSOperClasses       string      `json:"operating_class,omitempty"`

	// WMM
	WMM   string `json:"wmm"`
	Uapsd string `json:"uapsd"`

	// TODO use map[WMMAccessCategory]WMMCategoryConfig
	// and custom marshal/unmarshal
	WmmTxBKAifs  string `json:"wmm_tx_bk_aifs"`
	WmmTxBKCwmin string `json:"wmm_tx_bk_cwmin"`
	WmmTxBKCwmax string `json:"wmm_tx_bk_cwmax"`
	WmmTxBKBurst string `json:"wmm_tx_bk_burst"`

	WmmTxBEAifs  string `json:"wmm_tx_be_aifs"`
	WmmTxBECwmin string `json:"wmm_tx_be_cwmin"`
	WmmTxBECwmax string `json:"wmm_tx_be_cwmax"`
	WmmTxBEBurst string `json:"wmm_tx_be_burst"`

	WmmTxVIAifs  string `json:"wmm_tx_vi_aifs"`
	WmmTxVICwmin string `json:"wmm_tx_vi_cwmin"`
	WmmTxVICwmax string `json:"wmm_tx_vi_cwmax"`
	WmmTxVIBurst string `json:"wmm_tx_vi_burst"`

	WmmTxVOAifs  string `json:"wmm_tx_vo_aifs"`
	WmmTxVOCwmin string `json:"wmm_tx_vo_cwmin"`
	WmmTxVOCwmax string `json:"wmm_tx_vo_cwmax"`
	WmmTxVOBurst string `json:"wmm_tx_vo_burst"`

	WmmAcBKAifs  string `json:"wmm_ac_bk_aifs"`
	WmmAcBKCwmin string `json:"wmm_ac_bk_cwmin"`
	WmmAcBKCwmax string `json:"wmm_ac_bk_cwmax"`
	WmmAcBKTxop  string `json:"wmm_ac_bk_txop"`
	WmmAcBKAcm   string `json:"wmm_ac_bk_acm"`

	WmmAcBEAifs  string `json:"wmm_ac_be_aifs"`
	WmmAcBECwmin string `json:"wmm_ac_be_cwmin"`
	WmmAcBECwmax string `json:"wmm_ac_be_cwmax"`
	WmmAcBETxop  string `json:"wmm_ac_be_txop"`
	WmmAcBEAcm   string `json:"wmm_ac_be_acm"`

	WmmAcVIAifs  string `json:"wmm_ac_vi_aifs"`
	WmmAcVICwmin string `json:"wmm_ac_vi_cwmin"`
	WmmAcVICwmax string `json:"wmm_ac_vi_cwmax"`
	WmmAcVITxop  string `json:"wmm_ac_vi_txop"`
	WmmAcVIAcm   string `json:"wmm_ac_vi_acm"`

	WmmAcVOAifs  string `json:"wmm_ac_vo_aifs"`
	WmmAcVOCwmin string `json:"wmm_ac_vo_cwmin"`
	WmmAcVOCwmax string `json:"wmm_ac_vo_cwmax"`
	WmmAcVOTxop  string `json:"wmm_ac_vo_txop"`
	WmmAcVOAcm   string `json:"wmm_ac_vo_acm"`

	// speed limit
	// limit_fromrf -- FROM client -- UPLOAD
	// limit_torf -- TO client -- DOWNLOAD
	LimitFromRf string `json:"limit_fromrf,omitempty"`
	LimitToRf   string `json:"limit_torf,omitempty"`

	// timeouts
	MaxInactivity string `json:"max_inactivity,omitempty"`

	DisAssocLowAck string `json:"disassoc_low_ack,omitempty"`

	// mesh settings
	MeshID         string `json:"mesh_id"`
	MeshForwarding string `json:"mesh_fwding"`

	BandSteering string `json:"band_steering"`

	FTOverDS string `json:"ft_over_ds"`

	// Config for dae aka Dynamic Authorization Extension client in hostapd
	DaeClient string `json:"dae_client"`
	DaeSecret string `json:"dae_secret"`
	DaePort   string `json:"dae_port"`

	// config for wmwdiscd - rssi based disconnector
	// signal_connect '-50'
	// signal_stay '-60'
	// signal_strikes '3'
	// signal_poll_time '5'
	// signal_drop_reason '3'
	SignalConnect    string `json:"signal_connect"`
	SignalStay       string `json:"signal_stay"`
	SignalStrikes    string `json:"signal_strikes"`
	SignalPollTime   string `json:"signal_poll_time"`
	SignalDropReason string `json:"signal_drop_reason"`

	PMKR1Push          string `json:"pmk_r1_push"`
	FTPskGenerateLocal string `json:"ft_psk_generate_local"`

	// additional MAB attr for special VSA
	RadiusAuthRequestAttr interface{} `json:"radius_auth_req_attr,omitempty"`
}

type UciWifiIface struct {
	Type           string      `json:".type"`
	Addr           string      `json:"macaddr,omitempty"`
	Channel        string      `json:"channel"`
	Power          string      `json:"txpower"`
	BandMode       string      `json:"hwmode"`
	Bandwidth      string      `json:"htmode"`
	Country        string      `json:"country"`
	RequireMode    string      `json:"require_mode"`
	Disabled       string      `json:"disabled"`
	ChanList       interface{} `json:"channels"`
	BasicRate      string      `json:"basic_rate,omitempty"`
	SupportedRates interface{} `json:"supported_rates,omitempty"`
	LegacyRates    string      `json:"legacy_rates,omitempty"`
	LogLevel       string      `json:"log_level,omitempty"`
	MaxInactivity  string      `json:"max_inactivity,omitempty"`
	RSSIThreshold  string      `json:"rssi_reject_assoc_rssi"`
	NoScan         string      `json:"noscan"`
	// Band           string      `json:"band,omitempty"`
	ChanBandwidth string `json:"chanbw,omitempty"`
	CellDensity   string `json:"cell_density,omitempty"`

	HESUBeamformee string `json:"he_su_beamformee,omitempty"`
	HEBSSColor     string `json:"he_bss_color,omitempty"`
}
type UciRadius struct {
	Type       string `json:".type"`
	Addr       string `json:"server_addr"`
	AcctPort   string `json:"acct_port"`
	AuthPort   string `json:"auth_port"`
	AcctSecret string `json:"acct_secret"`
	AuthSecret string `json:"auth_secret"`
}
type innerUciWireless struct {
	Wlans      map[string]UciWifiWlan  `json:"-" inline:"yes,.type:wifi-iface"`
	Interfaces map[string]UciWifiIface `json:"-" inline:"yes,.type:wifi-device"`
	Radiuses   map[string]UciRadius    `json:"-" inline:"yes,.type:radius"`
}
type UciWireless innerUciWireless

//------------- System config ----------------

type UciDotSystem struct {
	Type      string `json:".type"`
	HostName  string `json:"hostname"`
	LogIP     string `json:"log_ip"`
	LogPrefix string `json:"log_prefix"`
	LogProto  string `json:"log_proto"`
	LogRemote string `json:"log_remote"`
	LogPort   string `json:"log_port"`
}
type UciNtp struct {
	Type    string      `json:".type"`
	Servers interface{} `json:"server,omitempty"`
	Enabled string      `json:"enabled"`
}
type innerUciSystem struct {
	Ntp    UciNtp       `json:"ntp"`
	System UciDotSystem `json:"-" inline:"yes,.type:system,unique,.system"`
}
type UciSystem innerUciSystem

//------------- Network config ----------------

type UciNetIface struct {
	Id           string      `json:"id,omitempty"`
	HostId       string      `json:"peerid,omitempty"`
	Tunnel       string      `json:"tunnel,omitempty"`
	Network      string      `json:"network,omitempty"`
	Type         string      `json:"type,omitempty"`
	Proto        string      `json:"proto,omitempty"`
	IfName       interface{} `json:"ifname,omitempty"`
	IpV6         string      `json:"ipv6,omitempty"`
	UciType      string      `json:".type,omitempty"`
	Redirect     string      `json:"wimark_redirect_l2,omitempty"`
	Zone         string      `json:"zone,omitempty"`
	PeerAddr     string      `json:"peeraddr,omitempty"`
	DontFragment string      `json:"df,omitempty"`
}

type UciNetTunnel struct {
	Id         string `json:"id,omitempty"`
	HostId     string `json:"peerid,omitempty"`
	HostIpAddr string `json:"peeraddr,omitempty"`
	Link       string `json:"tunlink,omitempty"`
	Encaps     string `json:"encap,omitempty"`
	UciType    string `json:".type,omitempty"`
}

type UciNetSwitchVlan struct {
	Device    string `json:"device,omitempty"`
	Vlan      string `json:"vlan,omitempty"`
	Vid       string `json:"vid,omitempty"`
	Ports     string `json:"ports,omitempty"`
	UciType   string `json:".type,omitempty"`
	WimarkCfg string `json:"wimark_set,omitempty"`
}

type UciNetSwitchPort struct {
	Number     string `json:"port,omitempty"`
	NativeVlan string `json:"pvid,omitempty"`
	Device     string `json:"device,omitempty"`
	UciType    string `json:".type,omitempty"`
}

type innerUciNetwork struct {
	Tunnels     map[string]UciNetTunnel     `json:"-" inline:"yes,.type:l2tunnel"`
	Interfaces  map[string]UciNetIface      `json:"-" inline:"yes,.type:interface"`
	SwitchVlans map[string]UciNetSwitchVlan `json:"-" inline:"yes,.type:switch_vlan"`
	//SwitchPorts map[string]UciNetSwitchPort `json:"-" inline:"yes,.type:switch_port"`
}
type UciNetwork innerUciNetwork

//------------- All configs ----------------

type UciConfig struct {
	Wimark   UciWimark   `json:"wimark,omitempty"`
	Wireless UciWireless `json:"wireless,omitempty"`
	System   UciSystem   `json:"system,omitempty"`
	Network  UciNetwork  `json:"network,omitempty"`
	Wsnmpd   interface{} `json:"wmsnmpd,omitempty"`
}

//------------- JSON conversion ----------------

func (uci UciWimark) MarshalJSON() (b []byte, e error) {
	return MarshalInline((*innerUciWimark)(&uci))
}
func (uci *UciWimark) UnmarshalJSON(b []byte) error {
	var tmp = map[string]interface{}{
		".type:firmware": &UciFirmwareCfg{},
		".type:scanning": &map[string]UciScanningCfg{},
		".type:redirect": &map[string]UciRedirectCfg{},
		".type:broker":   &map[string]UciBrokerCfg{},
		".type:eth_acct": &UciEthernetAcct{},
		".type:wsmd":     &UciEthernetStat{},
	}
	return UnmarshalInline(b, (*innerUciWimark)(uci), tmp)
}
func (uci UciWireless) MarshalJSON() (b []byte, e error) {
	return MarshalInline((*innerUciWireless)(&uci))
}
func (uci *UciWireless) UnmarshalJSON(b []byte) error {
	var tmp = map[string]interface{}{
		".type:wifi-iface":  &map[string]UciWifiWlan{},
		".type:wifi-device": &map[string]UciWifiIface{},
		".type:radius":      &map[string]UciRadius{},
	}
	return UnmarshalInline(b, (*innerUciWireless)(uci), tmp)
}
func (uci UciNetwork) MarshalJSON() (b []byte, e error) {
	return MarshalInline((*innerUciNetwork)(&uci))
}
func (uci *UciNetwork) UnmarshalJSON(b []byte) error {
	var tmp = map[string]interface{}{
		".type:interface":   &map[string]UciNetIface{},
		".type:l2tunnel":    &map[string]UciNetTunnel{},
		".type:switch_vlan": &map[string]UciNetSwitchVlan{},
		//".type:switch_port": &map[string]UciNetSwitchPort{},
	}
	return UnmarshalInline(b, (*innerUciNetwork)(uci), tmp)
}
func (uci UciSystem) MarshalJSON() (b []byte, e error) {
	return MarshalInline((*innerUciSystem)(&uci))
}
func (uci *UciSystem) UnmarshalJSON(b []byte) error {
	var tmp = map[string]interface{}{
		".type:system": &UciDotSystem{},
	}
	return UnmarshalInline(b, (*innerUciSystem)(uci), tmp)
}

//============= JSONRPC get:network ================

type GANetSubdevice struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	IfName string `json:"ifname"`
	Mac    string `json:"macaddr"`
	Up     bool   `json:"is_up"`
}
type GANetIface struct {
	IfName string   `json:"ifname"`
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Up     bool     `json:"is_up"`
	Mac    string   `json:"macaddr"`
	IPs    []string `json:"ipaddrs"`

	Subdevices map[string]GANetSubdevice `json:"wifi"`
}
type GANetGeneral struct {
	IPAddrs []string `json:"ipaddrs"`
	MACAddr string   `json:"macaddr"`
	GWAddr  string   `json:"gwaddr"`
}
type GAWanInfo struct {
	Interface string `json:"name"`
	Protocol  string `json:"proto"`
}
type GANetwork struct {
	General    GANetGeneral          `json:"general"`
	Interfaces map[string]GANetIface `json:"interfaces"`
	WanInfo    GAWanInfo             `json:"waninfo"`
}

//============= JSONRPC get:wifi ================

type GAWifiNetInfo struct {
	TxPower   int    `json:"txpower"`
	Channel   int    `json:"channel"`
	Frequency string `json:"frequency"`
	IfName    string `json:"ifname"`
	SSID      string `json:"ssid"`
	BSSID     string `json:"bssid"`
}
type GAWifi struct {
	Device string `json:"device"`
	Name   string `json:"name"`
	Up     bool   `json:"up"`

	Networks map[string]GAWifiNetInfo `json:"networks"`
}

//============= JSONRPC get:system ================

type GASysInfo struct{}

//============= JSONRPC get:model ================

type GAModelInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

//============= JSONRPC get:version ================

type GAVersion struct {
	Version
	Statics map[string]Version `json:"static"`
}

//============= JSONRPC wifi:get_radios_state ================

type GAWifiStatus struct {
	Device string            `json:"device"`
	State  CPEInterfaceState `json:"state"`
}

//============= Get all ================

type CPEAllInfo struct {
	Wifi       map[string]GAWifi     `json:"wifi,omitempty"`
	Network    GANetwork             `json:"network,omitempty"`
	Config     UciConfig             `json:"configuration,omitempty"`
	System     GASysInfo             `json:"system,omitempty"`
	Capability Capabilities          `json:"capabilities,omitempty"`
	Model      GAModelInfo           `json:"model,omitempty"`
	Version    GAVersion             `json:"version,omitempty"`
	Status     CPEAgentStatus        `json:"status,omitempty"`
	Packages   map[string]string     `json:"packages,omitempty"`
	Statics    map[string]Version    `json:"statics,omitempty"`
	WifiStatus []GAWifiStatus        `json:"wifistatus,omitempty"`
	Latitude   float64               `json:"latitude,omitempty"`
	Longitude  float64               `json:"longitude,omitempty"`
	Errors     map[int]JSONRPC_Error `json:"-"`
}

type CPEMeta struct {
	ModelName string
	ModelId   string
	Caps      Capabilities
	Uci       UciConfig
	Packages  map[string]string
	Statics   map[string]Version
}

//============= To set.lua script ================

type SLChain struct {
	Id     string         `json:"id"`
	Policy FirewallPolicy `json:"policy,omitempty"`
	Rules  []L2Rule       `json:"rules"`
}
type SLChainObject struct {
	Type      string            `json:"type"`
	Object    string            `json:"object"`
	Direction FirewallDirection `json:"direction"`
	Target    string            `json:"target"`
}
type SLChainConfig struct {
	Chains  []SLChain       `json:"chains"`
	Objects []SLChainObject `json:"objects"`
}

type SLNatAttrs struct {
	Redirect string `json:"wimark_redirect_l3,omitempty"`
	Addr     string `json:"ipaddr,omitempty"`
	NetMask  string `json:"netmask,omitempty"`
	Iface    string `json:"ifname,omitempty"`
}
type SLNatConfig struct {
	Attrs  SLNatAttrs `json:"attrs"`
	Access bool       `json:"access"`
}

//======== CPE model to CPE script with JSON RPC ====

type SLMessageContentsJSONRPC struct {
	General   UciConfig              `json:"config"`
	Wlans     map[string]UciWifiWlan `json:"wlans,omitempty"`
	Nats      map[string]SLNatConfig `json:"nats,omitempty"`
	Chains    SLChainConfig          `json:"l2chains,omitempty"`
	L3Filter  []string               `json:"l3filter,omitempty"`
	ToDel     map[string]interface{} `json:"-"`
	Timeout   time.Duration          `json:"-"`
	WireTCAdd []CPEWireTCConfig      `json:"-"`
	WireTCDel []CPEWireTCConfig      `json:"-"`
}

//======== Params for NAI realm compile JSON RPC ====

type NAIAuthMethod struct {
	Method int            `json:"method"`
	Params map[string]int `json:"params"`
}
type NAIRealmParams struct {
	Name    string          `json:"name"`
	Methods []NAIAuthMethod `json:"methods"`
}
