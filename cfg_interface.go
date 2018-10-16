package libwimark

// JSONRPC functions for configurer
const (
	JSONRPC_CONFIG_REBOOT = "firstboot"
)

// JSONRPC parameters

// for cpeagent:*
type ConfigInterfaceCpeReboot struct {
	Id       UUID   `json:"cpe"`
	Broker   string `json:"broker"`
	NoReboot bool   `json:"noreboot"`
}
