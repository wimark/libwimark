package libwimark

// JSONRPC functions from CPE for CPEAGENT
const (
	JSONRPC_CPE_CPEAGENT_PROCESS_REQ = "cpeagent:process_request"
)

// JSONRPC functions from CPE for UCI
const (
	JSONRPC_CPE_UCI_SET     = "uci:set"
	JSONRPC_CPE_UCI_GET     = "uci:get"
	JSONRPC_CPE_UCI_APPLY   = "uci:apply"
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
	JSONRPC_CPE_NETWORK_GET_BOARD = "network:get_boardinfo"
	JSONRPC_CPE_NETWORK_RELOAD    = "network:reload"
	JSONRPC_CPE_NETWORK_L2AP      = "network:l2ap"
	JSONRPC_CPE_NETWORK_VLAN      = "network:vlan"
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
