package libwimark

// JSONRPC functions from CPE for CPEAGENT
const (
	JSONRPC_CPE_CPEAGENT_PROCESS_REQ = "cpeagent:process_request"
	JSONRPC_CPE_CPEAGENT_GET_STATICS = "cpeagent:get_statics"
	JSONRPC_CPE_CPEAGENT_GET_METHODS = "cpeagent:get_methods"
	JSONRPC_CPE_CPEAGENT_STATUS      = "cpeagent:status"
	JSONRPC_CPE_CPEAGENT_OPKG        = "cpeagent:opkg"
)

// JSONRPC functions from CPE for UCI
const (
	JSONRPC_CPE_UCI_SET     = "uci:set"
	JSONRPC_CPE_UCI_GET     = "uci:get"
	JSONRPC_CPE_UCI_OPEN    = "uci:open_config"
	JSONRPC_CPE_UCI_APPLY   = "uci:apply_config"
	JSONRPC_CPE_UCI_CHANGES = "uci:changes"
	JSONRPC_CPE_UCI_COMMIT  = "uci:commit"
	JSONRPC_CPE_UCI_FINISH  = "uci:finish"
)

// JSONRPC functions from CPE for WIFI
const (
	JSONRPC_CPE_WIFI_ADD            = "wifi:add"
	JSONRPC_CPE_WIFI_CLEAR          = "wifi:clear"
	JSONRPC_CPE_WIFI_UPDATE         = "wifi:update"
	JSONRPC_CPE_WIFI_ENABLE_RADIO   = "wifi:enable_radio"
	JSONRPC_CPE_WIFI_CONFIGURE_ACCT = "wifi:configure_acct"
	JSONRPC_CPE_WIFI_NAI_COMPILE    = "wifi:compile_nai_realm"
	JSONRPC_CPE_WIFI_RADIO_STATE    = "wifi:get_radios_state"
)

// JSONRPC functions from CPE for NETWORK
const (
	JSONRPC_CPE_NETWORK_GET_BOARD   = "network:get_boardinfo"
	JSONRPC_CPE_NETWORK_RELOAD      = "network:reload"
	JSONRPC_CPE_NETWORK_L2AP        = "network:l2ap"
	JSONRPC_CPE_NETWORK_VLAN        = "network:vlan"
	JSONRPC_CPE_NETWORK_CLEAN_L2TP  = "network:clean_l2tp"
	JSONRPC_CPE_NETWORK_ADD_NAT     = "network:add_nat"
	JSONRPC_CPE_NETWORK_WAN_ACCESS  = "network:wan_access"
	JSONRPC_CPE_NETWORK_WIRE_TC_ADD = "network:wire_tc_add"
	JSONRPC_CPE_NETWORK_WIRE_TC_DEL = "network:wire_tc_del"
)

// JSONRPC functions from CPE for LBS
const (
	JSONRPC_CPE_LBS_CONFIGURE = "lbs:configure"
	JSONRPC_CPE_LBS_CHECK     = "lbs:check"
)

// JSONRPC functions from CPE for GET
const (
	JSONRPC_CPE_GET_SYSTEM       = "get:system"
	JSONRPC_CPE_GET_NETWORK      = "get:network"
	JSONRPC_CPE_GET_WIFI         = "get:wifi"
	JSONRPC_CPE_GET_CAPABILITIES = "get:capabilities"
	JSONRPC_CPE_GET_MODEL        = "get:model"
	JSONRPC_CPE_GET_VERSION      = "get:version"
)

// JSONRPC functions from CPE for LOGGING
const (
	JSONRPC_CPE_LOGGING_CONFIGURE = "logging:configure"
)

// JSONRPC functions from CPE for NOTIFIER
const (
	JSONRPC_CPE_NOTIFIER_OPEN  = "notifier:open"
	JSONRPC_CPE_NOTIFIER_CLOSE = "notifier:close"
	JSONRPC_CPE_NOTIFIER_ERROR = "notifier:error"
)

// JSONRPC functions from CPE for RADIUS
const (
	JSONRPC_CPE_RADIUS_PREPARE = "radius:prepare"
)

// JSONRPC functions from CPE for STATISTIC
const (
	JSONRPC_CPE_STATISTIC_CONFIGURE = "statistic:configure"
)

// JSONRPC functions from CPE for ACL
const (
	JSONRPC_CPE_ACL_L2           = "acl:l2"
	JSONRPC_CPE_ACL_L2_CLEAR     = "acl:l2_clear"
	JSONRPC_CPE_ACL_L2_FROM_UCI  = "acl:l2_from_uci"
	JSONRPC_CPE_ACL_L2_TO_UCI    = "acl:l2_to_uci"
	JSONRPC_CPE_ACL_L2_CLEAR_UCI = "acl:l2_clear_uci"
	JSONRPC_CPE_ACL_L3           = "acl:l3"
	JSONRPC_CPE_ACL_L3_RELOAD    = "acl:l3_reload"
)

// JSONRPC functions from CPE for FW upgrade
const (
	JSONRPC_CPE_FW_CONFIG  = "firmware:config"
	JSONRPC_CPE_FW_CHECK   = "firmware:check"
	JSONRPC_CPE_FW_UPGRADE = "firmware:upgrade"
)

// JSONRPC functions from CPE for l2portal
const (
	JSONRPC_CPE_L2PORTAL_WHITELIST_ADD    = "l2portal:whitelist_add"
	JSONRPC_CPE_L2PORTAL_WHITELIST_REMOVE = "l2portal:whitelist_remove"
)

// JSONRPC functions from CPE for RRM
const (
	JSONRPC_CPE_WIFI_SET_CHANNEL = "wifi:switch_chan"
	JSONRPC_CPE_WIFI_SET_TXPOWER = "wifi:set_txpower"
)

// JSONRPC functions for interaction with CISCO mediator
const (
	JSONRPC_CISCO_GET = "cisco:get"
	JSONRPC_CISCO_SET = "cisco:set"
)

// JSONRPC system functions from CPE
const (
	JSONRPC_CPE_SYSTEM_FIRSTBOOT = "system:firstboot"
	JSONRPC_CPE_SYSTEM_REBOOT    = "system:reboot"
)

// JSONRPC logupload functions for CPE
const (
	JSONRPC_CPE_LOGUPLOAD_START_LOGGING = "logupload:start_logging"
	JSONRPC_CPE_LOGUPLOAD_STOP_LOGGING  = "logupload:stop_logging"

	JSONRPC_CPE_LOGUPLOAD_ENABLE_DEBUG_HOSTAPD   = "logupload:enable_debug_hostapd"
	JSONRPC_CPE_LOGUPLOAD_ENABLE_DEBUG_2_HOSTAPD = "logupload:enable_debug2_hostapd"
	JSONRPC_CPE_LOGUPLOAD_DISABLE_DEBUG_HOSTAPD  = "logupload:disable_debug_hostapd"

	JSONRPC_CPE_LOGUPLOAD_ENABLE_DEBUG_NTPD  = "logupload:enable_debug_ntpd"
	JSONRPC_CPE_LOGUPLOAD_DISABLE_DEBUG_NTPD = "logupload:disable_debug_ntpd"

	JSONRPC_CPE_LOGUPLOAD_ENABLE_DEBUG_DHCPOPTINJ  = "logupload:enable_debug_dhcpoptinj"
	JSONRPC_CPE_LOGUPLOAD_DISABLE_DEBUG_DHCPOPTINJ = "logupload:disable_debug_dhcpoptinj"
)

const (
	CPEClientDisAssocTemplate = "for i in $(cat /proc/net/wireless | grep wlan | awk '{print $1;}' | sed 's/://g'); do echo $i; ubus call hostapd.$i del_client \"{'addr':'%s', 'reason':5, 'deauth':true, 'ban_time':0}\" ; done"
)

// JSONRPC parameters

// for cpeagent:*
type CPEAgentStatusBroker struct {
	Host string `json:"host"`
}
type CPEAgentStatusConnHost struct {
	Host string `json:"address"`
}
type CPEAgentStatusConnInfo struct {
	Remote CPEAgentStatusConnHost `json:"remote"`
	Local  CPEAgentStatusConnHost `json:"local"`
}
type CPEAgentStatus struct {
	State        string                 `json:"state"`
	Connection   CPEAgentStatusConnInfo `json:"conninfo"`
	Broker       CPEAgentStatusBroker   `json:"broker"`
	TunnelBroker CPEAgentStatusBroker   `json:"tunnel_broker"`
	TunnelType   string                 `json:"tunnel"`
}

// for firmware:*
type CPEFirmwareUpgradeParams struct {
	Config       CPEFirmwareConfig `json:"config"`
	FastResponce bool              `json:"fast_response"`
}
type CPEFirmwareConfig struct {
	FileUrl      string             `json:"file,omitempty"`
	StorageUrl   string             `json:"storage,omitempty"`
	Md5SumsUrl   string             `json:"md5sums,omitempty"`
	CheckTimeout int                `json:"timeout,omitempty"`
	Mode         FirmwareUpdateMode `json:"mode"`
	AvailableMd5 string             `json:"available_md5,omitempty"`
	ForceUpgrade bool               `json:"force_upgrade"`
}
type CPEFirmwareConfigResponse struct {
	Action string `json:"action"`
	CpeFirmwareData
}

// for RRM
type CPERRMChannelParams struct {
	Count       int    `json:"bcn_count"`
	Freq        int    `json:"freq"`
	Block       string `json:"block_tx,omitempty"`
	Offset      int    `json:"sec_channel_offset,omitempty"`
	CenterFreq1 int    `json:"center_freq1,omitempty"`
	CenterFreq2 int    `json:"center_freq2,omitempty"`
	Bandwidth   int    `json:"bandwidth,omitempty"`
	HT          string `json:"ht,omitempty"`
	VHT         string `json:"vht,omitempty"`
}

// for tc config on wire nat iface
type CPEWireTCConfig struct {
	Iface string `json:"iface"`
	Max   int    `json:"max,omitempty"`
	Min   int    `json:"min,omitempty"`
	Type  string `json:"type,omitempty"`
}
