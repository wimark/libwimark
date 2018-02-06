package libwimark

// JSONRPC functions of RRM daemon
const (
	JSONRPC_RRM_UPDATE_GROUP = "UpdateRRMGroup"

	JSONRPC_RRM_UPDATE_GRUOP_RSP_STATUS_SUCCESS = "SUCCESS"
	JSONRPC_RRM_UPDATE_GRUOP_RSP_STATUS_ERROR   = "ERROR"
)

// RRM Algorithms
const (
	RRM_ALGO_DUMMY = "Dummy"
)

func MakeRRMJSONRPCRequest(method string, args interface{}) *JSONRPCClientRequest {
	return &JSONRPCClientRequest{
		Version: JSON_RPC_VERSION,
		Method:  method,
		Params:  args,
	}
}

type UpdateRRMGroupParamsRSP struct {
	Status      string
	Description string
}
