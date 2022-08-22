package libwimark

// JSONRPC functions of RRM daemon
const (
	JSONRPC_RRM_UPDATE_GROUP = "UpdateRRMGroup"
	JSONRPC_RRM_FORCE_GROUP  = "ForceRRMGroup"
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
	RRMTimerParams `bson:",inline"`
	ManagePower    bool `json:"manage_power" bson:"manage_power"`
}
