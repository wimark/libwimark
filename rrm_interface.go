package libwimark

// JSONRPC functions of RRM daemon
const (
	JSONRPC_RRM_UPDATE_GROUP = "UpdateRRMGroup"
)

// RRM Algorithms
const (
	RRM_ALGO_DUMMY = "Dummy"
)

func MakeRRMRequest(args RRMGroup) *JSONRPCClientRequest {
	return &JSONRPCClientRequest{
		Version: JSON_RPC_VERSION,
		Method:  JSONRPC_RRM_UPDATE_GROUP,
		Params:  args,
	}
}
