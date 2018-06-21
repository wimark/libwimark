package libwimark

// JSONRPC functions of RRM daemon
const (
	JSONRPC_RRM_UPDATE_GROUP = "UpdateRRMGroup"
)

// JSONRPC parameters

type RRMTimerParams struct {
	Timeout int `json:"timeout"`
}
