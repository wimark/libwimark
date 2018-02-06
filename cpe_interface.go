package libwimark

// JSONRPC functions from CPE for CPEAGENT
const (
	JSONRPC_CPE_CPEAGENT_PROCESS_REQ = "cpeagent:process_request"
	JSONRPC_CPE_CPEAGENT_GET_STATICS = "cpeagent:get_statics"
	JSONRPC_CPE_CPEAGENT_GET_METHODS = "cpeagent:get_methods"
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
)

// JSONRPC functions from CPE for NETWORK
const (
	JSONRPC_CPE_NETWORK_GET_BOARD  = "network:get_boardinfo"
	JSONRPC_CPE_NETWORK_RELOAD     = "network:reload"
	JSONRPC_CPE_NETWORK_L2AP       = "network:l2ap"
	JSONRPC_CPE_NETWORK_VLAN       = "network:vlan"
	JSONRPC_CPE_NETWORK_CLEAN_L2TP = "network:clean_l2tp"
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

// JSONRPC parameters

// for firmware:*
type CPEFirmwareUpgradeParams struct {
	Config       CPEFirmwareConfig `json:"config"`
	FastResponce bool              `json:"fast_responce"`
}
type CPEFirmwareConfig struct {
	FileUrl      string             `json:"file"`
	StorageUrl   string             `json:"storage"`
	Md5SumsUrl   string             `json:"md5sums"`
	CheckTimeout int                `json:"timeout"`
	Mode         FirmwareUpdateMode `json:"mode"`
	AvailableMd5 string             `json:"available_md5"`
	ForceUpdate  bool               `json:"force_update"`
}
type CPEFirmwareConfigResponse struct {
	Action string `json:"action"`
	CpeFirmwareData
}
