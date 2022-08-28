package libwimark

// JSONRPC functions of RRM daemon
const (
	JSONRPC_RRM_UPDATE_GROUP = "UpdateRRMGroup"
	JSONRPC_RRM_FORCE_GROUP  = "ForceRRMGroup"
)

// JSONRPC parameters

type RRMParamUpdateGroup struct {
	Id        UUID         `json:"id"`
	Group     RRMGroup     `json:"group"`
	Operation RRMOperation `json:"operation"`
}

// Algorithms parameters

type RRMTimerParams struct {
	Timeout int `json:"timeout"`
}

type RRMGreedParams struct {
	RRMTimerParams `bson:",inline"`
	ManagePower    bool `json:"manage_power" bson:"manage_power"`
}

//go:generate enumer -type=RRMOperation -json -sql -transform=snake -output rrm_inteface_opeartion_string.go

// RRMOperation -тип rrm операции
type RRMOperation uint8

const (
	CreateRRMOperation RRMOperation = iota
	UpdateRRMOperation
	DeleteRRMOperation
)
