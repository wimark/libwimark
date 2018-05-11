package libwimark

import (
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type WimarkErrorCode string

const WimarkErrorCodeBadEnvironment WimarkErrorCode = "ERROR_BAD_ENV"
const WimarkErrorCodeCPE WimarkErrorCode = "ERROR_CPE"
const WimarkErrorCodeDB WimarkErrorCode = "ERROR_DB"
const WimarkErrorCodeJson WimarkErrorCode = "ERROR_JSON"
const WimarkErrorCodeJsonRpc WimarkErrorCode = "ERROR_JSONRPC"
const WimarkErrorCodeModuleOffline WimarkErrorCode = "ERROR_MODULE_OFFLINE"
const WimarkErrorCodeMqtt WimarkErrorCode = "ERROR_MQTT"
const WimarkErrorCodeNoError WimarkErrorCode = ""
const WimarkErrorCodeObjectNotExist WimarkErrorCode = "ERROR_OBJECT_NOT_EXIST"
const WimarkErrorCodeOther WimarkErrorCode = "ERROR_OTHER"
const WimarkErrorCodeProtocol WimarkErrorCode = "ERROR_PROTOCOL"
const WimarkErrorCodeRequestCheck WimarkErrorCode = "ERROR_REQUEST_CHECK"
const WimarkErrorCodeRspTimeout WimarkErrorCode = "ERROR_RSP_TIMEOUT"

func (self WimarkErrorCode) GetPtr() *WimarkErrorCode { var v = self; return &v }

func (self *WimarkErrorCode) String() string {
	switch *self {
	case WimarkErrorCodeBadEnvironment:
		return "ERROR_BAD_ENV"
	case WimarkErrorCodeCPE:
		return "ERROR_CPE"
	case WimarkErrorCodeDB:
		return "ERROR_DB"
	case WimarkErrorCodeJson:
		return "ERROR_JSON"
	case WimarkErrorCodeJsonRpc:
		return "ERROR_JSONRPC"
	case WimarkErrorCodeModuleOffline:
		return "ERROR_MODULE_OFFLINE"
	case WimarkErrorCodeMqtt:
		return "ERROR_MQTT"
	case WimarkErrorCodeNoError:
		return ""
	case WimarkErrorCodeObjectNotExist:
		return "ERROR_OBJECT_NOT_EXIST"
	case WimarkErrorCodeOther:
		return "ERROR_OTHER"
	case WimarkErrorCodeProtocol:
		return "ERROR_PROTOCOL"
	case WimarkErrorCodeRequestCheck:
		return "ERROR_REQUEST_CHECK"
	case WimarkErrorCodeRspTimeout:
		return "ERROR_RSP_TIMEOUT"
	}
	if len(*self) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of WimarkErrorCode"))
}

func (self *WimarkErrorCode) MarshalJSON() ([]byte, error) {
	switch *self {
	case WimarkErrorCodeBadEnvironment:
		return json.Marshal("ERROR_BAD_ENV")
	case WimarkErrorCodeCPE:
		return json.Marshal("ERROR_CPE")
	case WimarkErrorCodeDB:
		return json.Marshal("ERROR_DB")
	case WimarkErrorCodeJson:
		return json.Marshal("ERROR_JSON")
	case WimarkErrorCodeJsonRpc:
		return json.Marshal("ERROR_JSONRPC")
	case WimarkErrorCodeModuleOffline:
		return json.Marshal("ERROR_MODULE_OFFLINE")
	case WimarkErrorCodeMqtt:
		return json.Marshal("ERROR_MQTT")
	case WimarkErrorCodeNoError:
		return json.Marshal("")
	case WimarkErrorCodeObjectNotExist:
		return json.Marshal("ERROR_OBJECT_NOT_EXIST")
	case WimarkErrorCodeOther:
		return json.Marshal("ERROR_OTHER")
	case WimarkErrorCodeProtocol:
		return json.Marshal("ERROR_PROTOCOL")
	case WimarkErrorCodeRequestCheck:
		return json.Marshal("ERROR_REQUEST_CHECK")
	case WimarkErrorCodeRspTimeout:
		return json.Marshal("ERROR_RSP_TIMEOUT")
	}
	if len(*self) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of WimarkErrorCode")
}

func (self *WimarkErrorCode) GetBSON() (interface{}, error) {
	switch *self {
	case WimarkErrorCodeBadEnvironment:
		return "ERROR_BAD_ENV", nil
	case WimarkErrorCodeCPE:
		return "ERROR_CPE", nil
	case WimarkErrorCodeDB:
		return "ERROR_DB", nil
	case WimarkErrorCodeJson:
		return "ERROR_JSON", nil
	case WimarkErrorCodeJsonRpc:
		return "ERROR_JSONRPC", nil
	case WimarkErrorCodeModuleOffline:
		return "ERROR_MODULE_OFFLINE", nil
	case WimarkErrorCodeMqtt:
		return "ERROR_MQTT", nil
	case WimarkErrorCodeNoError:
		return "", nil
	case WimarkErrorCodeObjectNotExist:
		return "ERROR_OBJECT_NOT_EXIST", nil
	case WimarkErrorCodeOther:
		return "ERROR_OTHER", nil
	case WimarkErrorCodeProtocol:
		return "ERROR_PROTOCOL", nil
	case WimarkErrorCodeRequestCheck:
		return "ERROR_REQUEST_CHECK", nil
	case WimarkErrorCodeRspTimeout:
		return "ERROR_RSP_TIMEOUT", nil
	}
	if len(*self) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of WimarkErrorCode")
}

func (self *WimarkErrorCode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ERROR_BAD_ENV":
		*self = WimarkErrorCodeBadEnvironment
		return nil
	case "ERROR_CPE":
		*self = WimarkErrorCodeCPE
		return nil
	case "ERROR_DB":
		*self = WimarkErrorCodeDB
		return nil
	case "ERROR_JSON":
		*self = WimarkErrorCodeJson
		return nil
	case "ERROR_JSONRPC":
		*self = WimarkErrorCodeJsonRpc
		return nil
	case "ERROR_MODULE_OFFLINE":
		*self = WimarkErrorCodeModuleOffline
		return nil
	case "ERROR_MQTT":
		*self = WimarkErrorCodeMqtt
		return nil
	case "":
		*self = WimarkErrorCodeNoError
		return nil
	case "ERROR_OBJECT_NOT_EXIST":
		*self = WimarkErrorCodeObjectNotExist
		return nil
	case "ERROR_OTHER":
		*self = WimarkErrorCodeOther
		return nil
	case "ERROR_PROTOCOL":
		*self = WimarkErrorCodeProtocol
		return nil
	case "ERROR_REQUEST_CHECK":
		*self = WimarkErrorCodeRequestCheck
		return nil
	case "ERROR_RSP_TIMEOUT":
		*self = WimarkErrorCodeRspTimeout
		return nil
	}
	if len(s) == 0 {
		*self = WimarkErrorCodeNoError
		return nil
	}
	return errors.New("Unknown WimarkErrorCode")
}

func (self *WimarkErrorCode) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ERROR_BAD_ENV":
		*self = WimarkErrorCodeBadEnvironment
		return nil
	case "ERROR_CPE":
		*self = WimarkErrorCodeCPE
		return nil
	case "ERROR_DB":
		*self = WimarkErrorCodeDB
		return nil
	case "ERROR_JSON":
		*self = WimarkErrorCodeJson
		return nil
	case "ERROR_JSONRPC":
		*self = WimarkErrorCodeJsonRpc
		return nil
	case "ERROR_MODULE_OFFLINE":
		*self = WimarkErrorCodeModuleOffline
		return nil
	case "ERROR_MQTT":
		*self = WimarkErrorCodeMqtt
		return nil
	case "":
		*self = WimarkErrorCodeNoError
		return nil
	case "ERROR_OBJECT_NOT_EXIST":
		*self = WimarkErrorCodeObjectNotExist
		return nil
	case "ERROR_OTHER":
		*self = WimarkErrorCodeOther
		return nil
	case "ERROR_PROTOCOL":
		*self = WimarkErrorCodeProtocol
		return nil
	case "ERROR_REQUEST_CHECK":
		*self = WimarkErrorCodeRequestCheck
		return nil
	case "ERROR_RSP_TIMEOUT":
		*self = WimarkErrorCodeRspTimeout
		return nil
	}
	if len(s) == 0 {
		*self = WimarkErrorCodeNoError
		return nil
	}
	return errors.New("Unknown WimarkErrorCode")
}
