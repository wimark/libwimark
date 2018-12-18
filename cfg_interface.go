package libwimark

// JSONRPC functions for configurer
const (
	JSONRPC_CONFIG_REBOOT = "reboot"
)

// JSONRPC parameters

// for cpeagent:*
type ConfigInterfaceCpeReboot struct {
	Ids     []UUID `json:"cpes"`
	Broker  string `json:"broker"`
	NoReset bool   `json:"noreset"`
}
