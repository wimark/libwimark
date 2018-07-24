package libwimark

// JSONRPC functions of RRM daemon
const (
	JSONRPC_RRM_UPDATE_GROUP = "UpdateRRMGroup"
)

// JSONRPC parameters

type RRMParamUpdateGroup struct {
	Id    UUID     `json:"id"`
	Group RRMGroup `json:"group"`
}

// Algorithms parameters

type RRMTimerParams struct {
	Timeout int `json:"timeout"`
}

type RRMGreedParams struct {
	RRMTimerParams
	ManagePower bool `json:"manage_power"`
}
