package libwimark

import (
	"encoding/json"
	"math/rand"
)

const JSON_RPC_VERSION = "2.0"

// clientRequest represents a JSON-RPC request sent by a client.
type JSONRPCClientRequest struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`
	// A String containing the name of the method to be invoked.
	Method string `json:"method"`
	// Object to pass as request parameter to the method.
	Params interface{} `json:"params,omitempty"`
	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	Id int `json:"id,omitempty"`
}
type JSONRPCClientRequestList []JSONRPCClientRequest

// clientResponse represents a JSON-RPC response returned to a client.
type JSONRPCClientResponse struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`

	Result interface{}    `json:"result,omitempty"`
	Error  *JSONRPC_Error `json:"error,omitempty"`
	Id     int            `json:"id"`
}
type JSONRPCClientResponseList []JSONRPCClientResponse

type JSONRPC_ErrorCode int

type JSONRPC_Error struct {
	// code of the error
	Code JSONRPC_ErrorCode `json:"code"`
	// description of the error
	Message string `json:"messsage"`
	// optional data to attach
	Data interface{} `json:"data,omitempty"`
}

// RFC error codes
const (
	E_PARSE       JSONRPC_ErrorCode = -32700
	E_INVALID_REQ JSONRPC_ErrorCode = -32600
	E_NO_METHOD   JSONRPC_ErrorCode = -32601
	E_BAD_PARAMS  JSONRPC_ErrorCode = -32602
	E_INTERNAL    JSONRPC_ErrorCode = -32603
	E_SERVER      JSONRPC_ErrorCode = -32000
)

// EncodeClientRequest encodes parameters for a JSON-RPC client request.
func EncodeRPCRequest(method string, args interface{}) ([]byte, error) {
	c := &JSONRPCClientRequest{
		Version: JSON_RPC_VERSION,
		Method:  method,
		Params:  args,
		Id:      rand.Int(),
	}
	return json.Marshal(c)
}

func NewJSONRPCRequest(method string, args interface{}, id int) JSONRPCClientRequest {
	return JSONRPCClientRequest{
		Version: JSON_RPC_VERSION,
		Method:  method,
		Params:  args,
		Id:      id,
	}
}
