package libwimark

import (
	"encoding/json"
)
import (
	"errors"
)
import (
	"gopkg.in/mgo.v2/bson"
)

type BandwidthType string

const BandwidthTypeHT20 BandwidthType = "HT20"
const BandwidthTypeHT40 BandwidthType = "HT40"
const BandwidthTypeHT40Minus BandwidthType = "HT40-"
const BandwidthTypeHT40Plus BandwidthType = "HT40+"
const BandwidthTypeVHT160 BandwidthType = "VHT160"
const BandwidthTypeVHT20 BandwidthType = "VHT20"
const BandwidthTypeVHT40 BandwidthType = "VHT40"
const BandwidthTypeVHT80 BandwidthType = "VHT80"

func (self BandwidthType) GetPtr() *BandwidthType { var v = self; return &v }

func (self *BandwidthType) String() string {
	switch *self {
	case BandwidthTypeHT20:
		return "HT20"
	case BandwidthTypeHT40:
		return "HT40"
	case BandwidthTypeHT40Minus:
		return "HT40-"
	case BandwidthTypeHT40Plus:
		return "HT40+"
	case BandwidthTypeVHT160:
		return "VHT160"
	case BandwidthTypeVHT20:
		return "VHT20"
	case BandwidthTypeVHT40:
		return "VHT40"
	case BandwidthTypeVHT80:
		return "VHT80"

	}
	panic(errors.New("Invalid value of BandwidthType"))

}

func (self *BandwidthType) MarshalJSON() ([]byte, error) {
	switch *self {
	case BandwidthTypeHT20:
		return json.Marshal("HT20")
	case BandwidthTypeHT40:
		return json.Marshal("HT40")
	case BandwidthTypeHT40Minus:
		return json.Marshal("HT40-")
	case BandwidthTypeHT40Plus:
		return json.Marshal("HT40+")
	case BandwidthTypeVHT160:
		return json.Marshal("VHT160")
	case BandwidthTypeVHT20:
		return json.Marshal("VHT20")
	case BandwidthTypeVHT40:
		return json.Marshal("VHT40")
	case BandwidthTypeVHT80:
		return json.Marshal("VHT80")

	}
	return nil, errors.New("Invalid value of BandwidthType")

}

func (self *BandwidthType) GetBSON() (interface{}, error) {
	switch *self {
	case BandwidthTypeHT20:
		return "HT20", nil
	case BandwidthTypeHT40:
		return "HT40", nil
	case BandwidthTypeHT40Minus:
		return "HT40-", nil
	case BandwidthTypeHT40Plus:
		return "HT40+", nil
	case BandwidthTypeVHT160:
		return "VHT160", nil
	case BandwidthTypeVHT20:
		return "VHT20", nil
	case BandwidthTypeVHT40:
		return "VHT40", nil
	case BandwidthTypeVHT80:
		return "VHT80", nil

	}
	return nil, errors.New("Invalid value of BandwidthType")

}
func (self *BandwidthType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "HT20":
		*self = BandwidthTypeHT20
		return nil
	case "HT40":
		*self = BandwidthTypeHT40
		return nil
	case "HT40-":
		*self = BandwidthTypeHT40Minus
		return nil
	case "HT40+":
		*self = BandwidthTypeHT40Plus
		return nil
	case "VHT160":
		*self = BandwidthTypeVHT160
		return nil
	case "VHT20":
		*self = BandwidthTypeVHT20
		return nil
	case "VHT40":
		*self = BandwidthTypeVHT40
		return nil
	case "VHT80":
		*self = BandwidthTypeVHT80
		return nil

	}
	return errors.New("Unknown BandwidthType")

}

func (self *BandwidthType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "HT20":
		*self = BandwidthTypeHT20
		return nil
	case "HT40":
		*self = BandwidthTypeHT40
		return nil
	case "HT40-":
		*self = BandwidthTypeHT40Minus
		return nil
	case "HT40+":
		*self = BandwidthTypeHT40Plus
		return nil
	case "VHT160":
		*self = BandwidthTypeVHT160
		return nil
	case "VHT20":
		*self = BandwidthTypeVHT20
		return nil
	case "VHT40":
		*self = BandwidthTypeVHT40
		return nil
	case "VHT80":
		*self = BandwidthTypeVHT80
		return nil

	}
	return errors.New("Unknown BandwidthType")

}

type CPEAgentStatusType string

const CPEAgentStatusTypeException CPEAgentStatusType = "exception"
const CPEAgentStatusTypeSuccess CPEAgentStatusType = "success"
const CPEAgentStatusTypeSyntaxError CPEAgentStatusType = "syntax"
const CPEAgentStatusTypeUndefined CPEAgentStatusType = "undefined"

func (self CPEAgentStatusType) GetPtr() *CPEAgentStatusType { var v = self; return &v }

func (self *CPEAgentStatusType) String() string {
	switch *self {
	case CPEAgentStatusTypeException:
		return "exception"
	case CPEAgentStatusTypeSuccess:
		return "success"
	case CPEAgentStatusTypeSyntaxError:
		return "syntax"
	case CPEAgentStatusTypeUndefined:
		return "undefined"

	}
	panic(errors.New("Invalid value of CPEAgentStatusType"))

}

func (self *CPEAgentStatusType) MarshalJSON() ([]byte, error) {
	switch *self {
	case CPEAgentStatusTypeException:
		return json.Marshal("exception")
	case CPEAgentStatusTypeSuccess:
		return json.Marshal("success")
	case CPEAgentStatusTypeSyntaxError:
		return json.Marshal("syntax")
	case CPEAgentStatusTypeUndefined:
		return json.Marshal("undefined")

	}
	return nil, errors.New("Invalid value of CPEAgentStatusType")

}

func (self *CPEAgentStatusType) GetBSON() (interface{}, error) {
	switch *self {
	case CPEAgentStatusTypeException:
		return "exception", nil
	case CPEAgentStatusTypeSuccess:
		return "success", nil
	case CPEAgentStatusTypeSyntaxError:
		return "syntax", nil
	case CPEAgentStatusTypeUndefined:
		return "undefined", nil

	}
	return nil, errors.New("Invalid value of CPEAgentStatusType")

}
func (self *CPEAgentStatusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "exception":
		*self = CPEAgentStatusTypeException
		return nil
	case "success":
		*self = CPEAgentStatusTypeSuccess
		return nil
	case "syntax":
		*self = CPEAgentStatusTypeSyntaxError
		return nil
	case "undefined":
		*self = CPEAgentStatusTypeUndefined
		return nil

	}
	return errors.New("Unknown CPEAgentStatusType")

}

func (self *CPEAgentStatusType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "exception":
		*self = CPEAgentStatusTypeException
		return nil
	case "success":
		*self = CPEAgentStatusTypeSuccess
		return nil
	case "syntax":
		*self = CPEAgentStatusTypeSyntaxError
		return nil
	case "undefined":
		*self = CPEAgentStatusTypeUndefined
		return nil

	}
	return errors.New("Unknown CPEAgentStatusType")

}

type CPEInterfaceType string

const CPEInterfaceTypeWiFi CPEInterfaceType = "wifi"
const CPEInterfaceTypeWired CPEInterfaceType = "wired"

func (self CPEInterfaceType) GetPtr() *CPEInterfaceType { var v = self; return &v }

func (self *CPEInterfaceType) String() string {
	switch *self {
	case CPEInterfaceTypeWiFi:
		return "wifi"
	case CPEInterfaceTypeWired:
		return "wired"

	}
	panic(errors.New("Invalid value of CPEInterfaceType"))

}

func (self *CPEInterfaceType) MarshalJSON() ([]byte, error) {
	switch *self {
	case CPEInterfaceTypeWiFi:
		return json.Marshal("wifi")
	case CPEInterfaceTypeWired:
		return json.Marshal("wired")

	}
	return nil, errors.New("Invalid value of CPEInterfaceType")

}

func (self *CPEInterfaceType) GetBSON() (interface{}, error) {
	switch *self {
	case CPEInterfaceTypeWiFi:
		return "wifi", nil
	case CPEInterfaceTypeWired:
		return "wired", nil

	}
	return nil, errors.New("Invalid value of CPEInterfaceType")

}
func (self *CPEInterfaceType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "wifi":
		*self = CPEInterfaceTypeWiFi
		return nil
	case "wired":
		*self = CPEInterfaceTypeWired
		return nil

	}
	return errors.New("Unknown CPEInterfaceType")

}

func (self *CPEInterfaceType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "wifi":
		*self = CPEInterfaceTypeWiFi
		return nil
	case "wired":
		*self = CPEInterfaceTypeWired
		return nil

	}
	return errors.New("Unknown CPEInterfaceType")

}

type ClientStatPacketType string

const ClientStatPacketTypeInterim ClientStatPacketType = "Interim-Update"
const ClientStatPacketTypeOff ClientStatPacketType = "Accounting-Off"
const ClientStatPacketTypeOn ClientStatPacketType = "Accounting-On"
const ClientStatPacketTypeStart ClientStatPacketType = "Start"
const ClientStatPacketTypeStop ClientStatPacketType = "Stop"

func (self ClientStatPacketType) GetPtr() *ClientStatPacketType { var v = self; return &v }

func (self *ClientStatPacketType) String() string {
	switch *self {
	case ClientStatPacketTypeInterim:
		return "Interim-Update"
	case ClientStatPacketTypeOff:
		return "Accounting-Off"
	case ClientStatPacketTypeOn:
		return "Accounting-On"
	case ClientStatPacketTypeStart:
		return "Start"
	case ClientStatPacketTypeStop:
		return "Stop"

	}
	panic(errors.New("Invalid value of ClientStatPacketType"))

}

func (self *ClientStatPacketType) MarshalJSON() ([]byte, error) {
	switch *self {
	case ClientStatPacketTypeInterim:
		return json.Marshal("Interim-Update")
	case ClientStatPacketTypeOff:
		return json.Marshal("Accounting-Off")
	case ClientStatPacketTypeOn:
		return json.Marshal("Accounting-On")
	case ClientStatPacketTypeStart:
		return json.Marshal("Start")
	case ClientStatPacketTypeStop:
		return json.Marshal("Stop")

	}
	return nil, errors.New("Invalid value of ClientStatPacketType")

}

func (self *ClientStatPacketType) GetBSON() (interface{}, error) {
	switch *self {
	case ClientStatPacketTypeInterim:
		return "Interim-Update", nil
	case ClientStatPacketTypeOff:
		return "Accounting-Off", nil
	case ClientStatPacketTypeOn:
		return "Accounting-On", nil
	case ClientStatPacketTypeStart:
		return "Start", nil
	case ClientStatPacketTypeStop:
		return "Stop", nil

	}
	return nil, errors.New("Invalid value of ClientStatPacketType")

}
func (self *ClientStatPacketType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Interim-Update":
		*self = ClientStatPacketTypeInterim
		return nil
	case "Accounting-Off":
		*self = ClientStatPacketTypeOff
		return nil
	case "Accounting-On":
		*self = ClientStatPacketTypeOn
		return nil
	case "Start":
		*self = ClientStatPacketTypeStart
		return nil
	case "Stop":
		*self = ClientStatPacketTypeStop
		return nil

	}
	return errors.New("Unknown ClientStatPacketType")

}

func (self *ClientStatPacketType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "Interim-Update":
		*self = ClientStatPacketTypeInterim
		return nil
	case "Accounting-Off":
		*self = ClientStatPacketTypeOff
		return nil
	case "Accounting-On":
		*self = ClientStatPacketTypeOn
		return nil
	case "Start":
		*self = ClientStatPacketTypeStart
		return nil
	case "Stop":
		*self = ClientStatPacketTypeStop
		return nil

	}
	return errors.New("Unknown ClientStatPacketType")

}

type ConfigurationStatus string

const ConfigurationStatusEmpty ConfigurationStatus = "empty"
const ConfigurationStatusError ConfigurationStatus = "error"
const ConfigurationStatusOK ConfigurationStatus = "ok"
const ConfigurationStatusPending ConfigurationStatus = "pending"
const ConfigurationStatusUpdating ConfigurationStatus = "updating"

func (self ConfigurationStatus) GetPtr() *ConfigurationStatus { var v = self; return &v }

func (self *ConfigurationStatus) String() string {
	switch *self {
	case ConfigurationStatusEmpty:
		return "empty"
	case ConfigurationStatusError:
		return "error"
	case ConfigurationStatusOK:
		return "ok"
	case ConfigurationStatusPending:
		return "pending"
	case ConfigurationStatusUpdating:
		return "updating"

	}
	panic(errors.New("Invalid value of ConfigurationStatus"))

}

func (self *ConfigurationStatus) MarshalJSON() ([]byte, error) {
	switch *self {
	case ConfigurationStatusEmpty:
		return json.Marshal("empty")
	case ConfigurationStatusError:
		return json.Marshal("error")
	case ConfigurationStatusOK:
		return json.Marshal("ok")
	case ConfigurationStatusPending:
		return json.Marshal("pending")
	case ConfigurationStatusUpdating:
		return json.Marshal("updating")

	}
	return nil, errors.New("Invalid value of ConfigurationStatus")

}

func (self *ConfigurationStatus) GetBSON() (interface{}, error) {
	switch *self {
	case ConfigurationStatusEmpty:
		return "empty", nil
	case ConfigurationStatusError:
		return "error", nil
	case ConfigurationStatusOK:
		return "ok", nil
	case ConfigurationStatusPending:
		return "pending", nil
	case ConfigurationStatusUpdating:
		return "updating", nil

	}
	return nil, errors.New("Invalid value of ConfigurationStatus")

}
func (self *ConfigurationStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "empty":
		*self = ConfigurationStatusEmpty
		return nil
	case "error":
		*self = ConfigurationStatusError
		return nil
	case "ok":
		*self = ConfigurationStatusOK
		return nil
	case "pending":
		*self = ConfigurationStatusPending
		return nil
	case "updating":
		*self = ConfigurationStatusUpdating
		return nil

	}
	return errors.New("Unknown ConfigurationStatus")

}

func (self *ConfigurationStatus) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "empty":
		*self = ConfigurationStatusEmpty
		return nil
	case "error":
		*self = ConfigurationStatusError
		return nil
	case "ok":
		*self = ConfigurationStatusOK
		return nil
	case "pending":
		*self = ConfigurationStatusPending
		return nil
	case "updating":
		*self = ConfigurationStatusUpdating
		return nil

	}
	return errors.New("Unknown ConfigurationStatus")

}

type MacFilterType string

const MacFilterTypeBlackList MacFilterType = "BlackList"
const MacFilterTypeNone MacFilterType = "None"
const MacFilterTypeWhiteList MacFilterType = "WhiteList"

func (self MacFilterType) GetPtr() *MacFilterType { var v = self; return &v }

func (self *MacFilterType) String() string {
	switch *self {
	case MacFilterTypeBlackList:
		return "BlackList"
	case MacFilterTypeNone:
		return "None"
	case MacFilterTypeWhiteList:
		return "WhiteList"

	}
	panic(errors.New("Invalid value of MacFilterType"))

}

func (self *MacFilterType) MarshalJSON() ([]byte, error) {
	switch *self {
	case MacFilterTypeBlackList:
		return json.Marshal("BlackList")
	case MacFilterTypeNone:
		return json.Marshal("None")
	case MacFilterTypeWhiteList:
		return json.Marshal("WhiteList")

	}
	return nil, errors.New("Invalid value of MacFilterType")

}

func (self *MacFilterType) GetBSON() (interface{}, error) {
	switch *self {
	case MacFilterTypeBlackList:
		return "BlackList", nil
	case MacFilterTypeNone:
		return "None", nil
	case MacFilterTypeWhiteList:
		return "WhiteList", nil

	}
	return nil, errors.New("Invalid value of MacFilterType")

}
func (self *MacFilterType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "BlackList":
		*self = MacFilterTypeBlackList
		return nil
	case "None":
		*self = MacFilterTypeNone
		return nil
	case "WhiteList":
		*self = MacFilterTypeWhiteList
		return nil

	}
	return errors.New("Unknown MacFilterType")

}

func (self *MacFilterType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "BlackList":
		*self = MacFilterTypeBlackList
		return nil
	case "None":
		*self = MacFilterTypeNone
		return nil
	case "WhiteList":
		*self = MacFilterTypeWhiteList
		return nil

	}
	return errors.New("Unknown MacFilterType")

}

type Module string

const ModuleAC Module = "AC"
const ModuleAny Module = "+"
const ModuleBackend Module = "BACKEND"
const ModuleCPE Module = "CPE"
const ModuleCPEStat Module = "CPE_STAT"
const ModuleClientStat Module = "CLIENT_STAT"
const ModuleConfig Module = "CONFIG"
const ModuleDB Module = "DB"
const ModuleDummy Module = "DUMMY"
const ModuleLBS Module = "LBS"
const ModuleMQTTLog Module = "MQTT_LOG"
const ModuleMonitor Module = "MONITOR"
const ModuleStat Module = "STAT"

func (self Module) GetPtr() *Module { var v = self; return &v }

func (self *Module) String() string {
	switch *self {
	case ModuleAC:
		return "AC"
	case ModuleAny:
		return "+"
	case ModuleBackend:
		return "BACKEND"
	case ModuleCPE:
		return "CPE"
	case ModuleCPEStat:
		return "CPE_STAT"
	case ModuleClientStat:
		return "CLIENT_STAT"
	case ModuleConfig:
		return "CONFIG"
	case ModuleDB:
		return "DB"
	case ModuleDummy:
		return "DUMMY"
	case ModuleLBS:
		return "LBS"
	case ModuleMQTTLog:
		return "MQTT_LOG"
	case ModuleMonitor:
		return "MONITOR"
	case ModuleStat:
		return "STAT"

	}
	panic(errors.New("Invalid value of Module"))

}

func (self *Module) MarshalJSON() ([]byte, error) {
	switch *self {
	case ModuleAC:
		return json.Marshal("AC")
	case ModuleAny:
		return json.Marshal("+")
	case ModuleBackend:
		return json.Marshal("BACKEND")
	case ModuleCPE:
		return json.Marshal("CPE")
	case ModuleCPEStat:
		return json.Marshal("CPE_STAT")
	case ModuleClientStat:
		return json.Marshal("CLIENT_STAT")
	case ModuleConfig:
		return json.Marshal("CONFIG")
	case ModuleDB:
		return json.Marshal("DB")
	case ModuleDummy:
		return json.Marshal("DUMMY")
	case ModuleLBS:
		return json.Marshal("LBS")
	case ModuleMQTTLog:
		return json.Marshal("MQTT_LOG")
	case ModuleMonitor:
		return json.Marshal("MONITOR")
	case ModuleStat:
		return json.Marshal("STAT")

	}
	return nil, errors.New("Invalid value of Module")

}

func (self *Module) GetBSON() (interface{}, error) {
	switch *self {
	case ModuleAC:
		return "AC", nil
	case ModuleAny:
		return "+", nil
	case ModuleBackend:
		return "BACKEND", nil
	case ModuleCPE:
		return "CPE", nil
	case ModuleCPEStat:
		return "CPE_STAT", nil
	case ModuleClientStat:
		return "CLIENT_STAT", nil
	case ModuleConfig:
		return "CONFIG", nil
	case ModuleDB:
		return "DB", nil
	case ModuleDummy:
		return "DUMMY", nil
	case ModuleLBS:
		return "LBS", nil
	case ModuleMQTTLog:
		return "MQTT_LOG", nil
	case ModuleMonitor:
		return "MONITOR", nil
	case ModuleStat:
		return "STAT", nil

	}
	return nil, errors.New("Invalid value of Module")

}
func (self *Module) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "AC":
		*self = ModuleAC
		return nil
	case "+":
		*self = ModuleAny
		return nil
	case "BACKEND":
		*self = ModuleBackend
		return nil
	case "CPE":
		*self = ModuleCPE
		return nil
	case "CPE_STAT":
		*self = ModuleCPEStat
		return nil
	case "CLIENT_STAT":
		*self = ModuleClientStat
		return nil
	case "CONFIG":
		*self = ModuleConfig
		return nil
	case "DB":
		*self = ModuleDB
		return nil
	case "DUMMY":
		*self = ModuleDummy
		return nil
	case "LBS":
		*self = ModuleLBS
		return nil
	case "MQTT_LOG":
		*self = ModuleMQTTLog
		return nil
	case "MONITOR":
		*self = ModuleMonitor
		return nil
	case "STAT":
		*self = ModuleStat
		return nil

	}
	return errors.New("Unknown Module")

}

func (self *Module) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "AC":
		*self = ModuleAC
		return nil
	case "+":
		*self = ModuleAny
		return nil
	case "BACKEND":
		*self = ModuleBackend
		return nil
	case "CPE":
		*self = ModuleCPE
		return nil
	case "CPE_STAT":
		*self = ModuleCPEStat
		return nil
	case "CLIENT_STAT":
		*self = ModuleClientStat
		return nil
	case "CONFIG":
		*self = ModuleConfig
		return nil
	case "DB":
		*self = ModuleDB
		return nil
	case "DUMMY":
		*self = ModuleDummy
		return nil
	case "LBS":
		*self = ModuleLBS
		return nil
	case "MQTT_LOG":
		*self = ModuleMQTTLog
		return nil
	case "MONITOR":
		*self = ModuleMonitor
		return nil
	case "STAT":
		*self = ModuleStat
		return nil

	}
	return errors.New("Unknown Module")

}

type Operation string

const OperationAny Operation = "+"
const OperationCPEStatus Operation = "STATUS"
const OperationCreate Operation = "C"
const OperationDelete Operation = "D"
const OperationLuaScript Operation = "LUA"
const OperationRead Operation = "R"
const OperationSHScript Operation = "SH"
const OperationUpdate Operation = "U"

func (self Operation) GetPtr() *Operation { var v = self; return &v }

func (self *Operation) String() string {
	switch *self {
	case OperationAny:
		return "+"
	case OperationCPEStatus:
		return "STATUS"
	case OperationCreate:
		return "C"
	case OperationDelete:
		return "D"
	case OperationLuaScript:
		return "LUA"
	case OperationRead:
		return "R"
	case OperationSHScript:
		return "SH"
	case OperationUpdate:
		return "U"

	}
	panic(errors.New("Invalid value of Operation"))

}

func (self *Operation) MarshalJSON() ([]byte, error) {
	switch *self {
	case OperationAny:
		return json.Marshal("+")
	case OperationCPEStatus:
		return json.Marshal("STATUS")
	case OperationCreate:
		return json.Marshal("C")
	case OperationDelete:
		return json.Marshal("D")
	case OperationLuaScript:
		return json.Marshal("LUA")
	case OperationRead:
		return json.Marshal("R")
	case OperationSHScript:
		return json.Marshal("SH")
	case OperationUpdate:
		return json.Marshal("U")

	}
	return nil, errors.New("Invalid value of Operation")

}

func (self *Operation) GetBSON() (interface{}, error) {
	switch *self {
	case OperationAny:
		return "+", nil
	case OperationCPEStatus:
		return "STATUS", nil
	case OperationCreate:
		return "C", nil
	case OperationDelete:
		return "D", nil
	case OperationLuaScript:
		return "LUA", nil
	case OperationRead:
		return "R", nil
	case OperationSHScript:
		return "SH", nil
	case OperationUpdate:
		return "U", nil

	}
	return nil, errors.New("Invalid value of Operation")

}
func (self *Operation) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		*self = OperationAny
		return nil
	case "STATUS":
		*self = OperationCPEStatus
		return nil
	case "C":
		*self = OperationCreate
		return nil
	case "D":
		*self = OperationDelete
		return nil
	case "LUA":
		*self = OperationLuaScript
		return nil
	case "R":
		*self = OperationRead
		return nil
	case "SH":
		*self = OperationSHScript
		return nil
	case "U":
		*self = OperationUpdate
		return nil

	}
	return errors.New("Unknown Operation")

}

func (self *Operation) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "+":
		*self = OperationAny
		return nil
	case "STATUS":
		*self = OperationCPEStatus
		return nil
	case "C":
		*self = OperationCreate
		return nil
	case "D":
		*self = OperationDelete
		return nil
	case "LUA":
		*self = OperationLuaScript
		return nil
	case "R":
		*self = OperationRead
		return nil
	case "SH":
		*self = OperationSHScript
		return nil
	case "U":
		*self = OperationUpdate
		return nil

	}
	return errors.New("Unknown Operation")

}

type RadiusType string

const RadiusTypeAccounting RadiusType = "acct"
const RadiusTypeAuthentication RadiusType = "auth"
const RadiusTypeBoth RadiusType = "both"

func (self RadiusType) GetPtr() *RadiusType { var v = self; return &v }

func (self *RadiusType) String() string {
	switch *self {
	case RadiusTypeAccounting:
		return "acct"
	case RadiusTypeAuthentication:
		return "auth"
	case RadiusTypeBoth:
		return "both"

	}
	panic(errors.New("Invalid value of RadiusType"))

}

func (self *RadiusType) MarshalJSON() ([]byte, error) {
	switch *self {
	case RadiusTypeAccounting:
		return json.Marshal("acct")
	case RadiusTypeAuthentication:
		return json.Marshal("auth")
	case RadiusTypeBoth:
		return json.Marshal("both")

	}
	return nil, errors.New("Invalid value of RadiusType")

}

func (self *RadiusType) GetBSON() (interface{}, error) {
	switch *self {
	case RadiusTypeAccounting:
		return "acct", nil
	case RadiusTypeAuthentication:
		return "auth", nil
	case RadiusTypeBoth:
		return "both", nil

	}
	return nil, errors.New("Invalid value of RadiusType")

}
func (self *RadiusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "acct":
		*self = RadiusTypeAccounting
		return nil
	case "auth":
		*self = RadiusTypeAuthentication
		return nil
	case "both":
		*self = RadiusTypeBoth
		return nil

	}
	return errors.New("Unknown RadiusType")

}

func (self *RadiusType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "acct":
		*self = RadiusTypeAccounting
		return nil
	case "auth":
		*self = RadiusTypeAuthentication
		return nil
	case "both":
		*self = RadiusTypeBoth
		return nil

	}
	return errors.New("Unknown RadiusType")

}

type SecuritySuite string

const SecuritySuiteAES SecuritySuite = "aes"
const SecuritySuiteTKIP SecuritySuite = "tkip"

func (self SecuritySuite) GetPtr() *SecuritySuite { var v = self; return &v }

func (self *SecuritySuite) String() string {
	switch *self {
	case SecuritySuiteAES:
		return "aes"
	case SecuritySuiteTKIP:
		return "tkip"

	}
	panic(errors.New("Invalid value of SecuritySuite"))

}

func (self *SecuritySuite) MarshalJSON() ([]byte, error) {
	switch *self {
	case SecuritySuiteAES:
		return json.Marshal("aes")
	case SecuritySuiteTKIP:
		return json.Marshal("tkip")

	}
	return nil, errors.New("Invalid value of SecuritySuite")

}

func (self *SecuritySuite) GetBSON() (interface{}, error) {
	switch *self {
	case SecuritySuiteAES:
		return "aes", nil
	case SecuritySuiteTKIP:
		return "tkip", nil

	}
	return nil, errors.New("Invalid value of SecuritySuite")

}
func (self *SecuritySuite) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "aes":
		*self = SecuritySuiteAES
		return nil
	case "tkip":
		*self = SecuritySuiteTKIP
		return nil

	}
	return errors.New("Unknown SecuritySuite")

}

func (self *SecuritySuite) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "aes":
		*self = SecuritySuiteAES
		return nil
	case "tkip":
		*self = SecuritySuiteTKIP
		return nil

	}
	return errors.New("Unknown SecuritySuite")

}

type SecurityType string

const SecurityTypeWPA2Enterprise SecurityType = "wpa2enterprise"
const SecurityTypeWPA2Personal SecurityType = "wpa2personal"
const SecurityTypeWPAEnterprise SecurityType = "wpaenterprise"
const SecurityTypeWPAPersonal SecurityType = "wpapersonal"

func (self SecurityType) GetPtr() *SecurityType { var v = self; return &v }

func (self *SecurityType) String() string {
	switch *self {
	case SecurityTypeWPA2Enterprise:
		return "wpa2enterprise"
	case SecurityTypeWPA2Personal:
		return "wpa2personal"
	case SecurityTypeWPAEnterprise:
		return "wpaenterprise"
	case SecurityTypeWPAPersonal:
		return "wpapersonal"

	}
	panic(errors.New("Invalid value of SecurityType"))

}

func (self *SecurityType) MarshalJSON() ([]byte, error) {
	switch *self {
	case SecurityTypeWPA2Enterprise:
		return json.Marshal("wpa2enterprise")
	case SecurityTypeWPA2Personal:
		return json.Marshal("wpa2personal")
	case SecurityTypeWPAEnterprise:
		return json.Marshal("wpaenterprise")
	case SecurityTypeWPAPersonal:
		return json.Marshal("wpapersonal")

	}
	return nil, errors.New("Invalid value of SecurityType")

}

func (self *SecurityType) GetBSON() (interface{}, error) {
	switch *self {
	case SecurityTypeWPA2Enterprise:
		return "wpa2enterprise", nil
	case SecurityTypeWPA2Personal:
		return "wpa2personal", nil
	case SecurityTypeWPAEnterprise:
		return "wpaenterprise", nil
	case SecurityTypeWPAPersonal:
		return "wpapersonal", nil

	}
	return nil, errors.New("Invalid value of SecurityType")

}
func (self *SecurityType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "wpa2enterprise":
		*self = SecurityTypeWPA2Enterprise
		return nil
	case "wpa2personal":
		*self = SecurityTypeWPA2Personal
		return nil
	case "wpaenterprise":
		*self = SecurityTypeWPAEnterprise
		return nil
	case "wpapersonal":
		*self = SecurityTypeWPAPersonal
		return nil

	}
	return errors.New("Unknown SecurityType")

}

func (self *SecurityType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "wpa2enterprise":
		*self = SecurityTypeWPA2Enterprise
		return nil
	case "wpa2personal":
		*self = SecurityTypeWPA2Personal
		return nil
	case "wpaenterprise":
		*self = SecurityTypeWPAEnterprise
		return nil
	case "wpapersonal":
		*self = SecurityTypeWPAPersonal
		return nil

	}
	return errors.New("Unknown SecurityType")

}

type ServiceState string

const ServiceStateConnected ServiceState = "CONNECTED"
const ServiceStateDisconnected ServiceState = "DISCONNECTED"
const ServiceStatePending ServiceState = "PENDING"

func (self ServiceState) GetPtr() *ServiceState { var v = self; return &v }

func (self *ServiceState) String() string {
	switch *self {
	case ServiceStateConnected:
		return "CONNECTED"
	case ServiceStateDisconnected:
		return "DISCONNECTED"
	case ServiceStatePending:
		return "PENDING"

	}
	panic(errors.New("Invalid value of ServiceState"))

}

func (self *ServiceState) MarshalJSON() ([]byte, error) {
	switch *self {
	case ServiceStateConnected:
		return json.Marshal("CONNECTED")
	case ServiceStateDisconnected:
		return json.Marshal("DISCONNECTED")
	case ServiceStatePending:
		return json.Marshal("PENDING")

	}
	return nil, errors.New("Invalid value of ServiceState")

}

func (self *ServiceState) GetBSON() (interface{}, error) {
	switch *self {
	case ServiceStateConnected:
		return "CONNECTED", nil
	case ServiceStateDisconnected:
		return "DISCONNECTED", nil
	case ServiceStatePending:
		return "PENDING", nil

	}
	return nil, errors.New("Invalid value of ServiceState")

}
func (self *ServiceState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*self = ServiceStateConnected
		return nil
	case "DISCONNECTED":
		*self = ServiceStateDisconnected
		return nil
	case "PENDING":
		*self = ServiceStatePending
		return nil

	}
	return errors.New("Unknown ServiceState")

}

func (self *ServiceState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*self = ServiceStateConnected
		return nil
	case "DISCONNECTED":
		*self = ServiceStateDisconnected
		return nil
	case "PENDING":
		*self = ServiceStatePending
		return nil

	}
	return errors.New("Unknown ServiceState")

}

type StatEventRuleType string

const StatEventRuleTypeCPUload StatEventRuleType = "cpu_load"
const StatEventRuleTypeFreeRAM StatEventRuleType = "free_ram"

func (self StatEventRuleType) GetPtr() *StatEventRuleType { var v = self; return &v }

func (self *StatEventRuleType) String() string {
	switch *self {
	case StatEventRuleTypeCPUload:
		return "cpu_load"
	case StatEventRuleTypeFreeRAM:
		return "free_ram"

	}
	panic(errors.New("Invalid value of StatEventRuleType"))

}

func (self *StatEventRuleType) MarshalJSON() ([]byte, error) {
	switch *self {
	case StatEventRuleTypeCPUload:
		return json.Marshal("cpu_load")
	case StatEventRuleTypeFreeRAM:
		return json.Marshal("free_ram")

	}
	return nil, errors.New("Invalid value of StatEventRuleType")

}

func (self *StatEventRuleType) GetBSON() (interface{}, error) {
	switch *self {
	case StatEventRuleTypeCPUload:
		return "cpu_load", nil
	case StatEventRuleTypeFreeRAM:
		return "free_ram", nil

	}
	return nil, errors.New("Invalid value of StatEventRuleType")

}
func (self *StatEventRuleType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "cpu_load":
		*self = StatEventRuleTypeCPUload
		return nil
	case "free_ram":
		*self = StatEventRuleTypeFreeRAM
		return nil

	}
	return errors.New("Unknown StatEventRuleType")

}

func (self *StatEventRuleType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "cpu_load":
		*self = StatEventRuleTypeCPUload
		return nil
	case "free_ram":
		*self = StatEventRuleTypeFreeRAM
		return nil

	}
	return errors.New("Unknown StatEventRuleType")

}

type SystemEventLevel string

const SystemEventLevelDEBUG SystemEventLevel = "DEBUG"
const SystemEventLevelERROR SystemEventLevel = "ERROR"
const SystemEventLevelINFO SystemEventLevel = "INFO"
const SystemEventLevelWARNING SystemEventLevel = "WARNING"

func (self SystemEventLevel) GetPtr() *SystemEventLevel { var v = self; return &v }

func (self *SystemEventLevel) String() string {
	switch *self {
	case SystemEventLevelDEBUG:
		return "DEBUG"
	case SystemEventLevelERROR:
		return "ERROR"
	case SystemEventLevelINFO:
		return "INFO"
	case SystemEventLevelWARNING:
		return "WARNING"

	}
	panic(errors.New("Invalid value of SystemEventLevel"))

}

func (self *SystemEventLevel) MarshalJSON() ([]byte, error) {
	switch *self {
	case SystemEventLevelDEBUG:
		return json.Marshal("DEBUG")
	case SystemEventLevelERROR:
		return json.Marshal("ERROR")
	case SystemEventLevelINFO:
		return json.Marshal("INFO")
	case SystemEventLevelWARNING:
		return json.Marshal("WARNING")

	}
	return nil, errors.New("Invalid value of SystemEventLevel")

}

func (self *SystemEventLevel) GetBSON() (interface{}, error) {
	switch *self {
	case SystemEventLevelDEBUG:
		return "DEBUG", nil
	case SystemEventLevelERROR:
		return "ERROR", nil
	case SystemEventLevelINFO:
		return "INFO", nil
	case SystemEventLevelWARNING:
		return "WARNING", nil

	}
	return nil, errors.New("Invalid value of SystemEventLevel")

}
func (self *SystemEventLevel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "DEBUG":
		*self = SystemEventLevelDEBUG
		return nil
	case "ERROR":
		*self = SystemEventLevelERROR
		return nil
	case "INFO":
		*self = SystemEventLevelINFO
		return nil
	case "WARNING":
		*self = SystemEventLevelWARNING
		return nil

	}
	return errors.New("Unknown SystemEventLevel")

}

func (self *SystemEventLevel) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "DEBUG":
		*self = SystemEventLevelDEBUG
		return nil
	case "ERROR":
		*self = SystemEventLevelERROR
		return nil
	case "INFO":
		*self = SystemEventLevelINFO
		return nil
	case "WARNING":
		*self = SystemEventLevelWARNING
		return nil

	}
	return errors.New("Unknown SystemEventLevel")

}

type SystemEventType string

const SystemEventTypeAny SystemEventType = "+"
const SystemEventTypeCPEConfigurationError SystemEventType = "CPE_CONFIGURATION_ERROR"
const SystemEventTypeCPEConfigurationSuccess SystemEventType = "CPE_CONFIGURATION_SUCCESS"
const SystemEventTypeCPEConnected SystemEventType = "CPE_CONNECTED"
const SystemEventTypeCPEDisconnected SystemEventType = "CPE_DISCONNECTED"
const SystemEventTypeClientConnected SystemEventType = "CLIENT_CONNECTED"
const SystemEventTypeClientDisconnected SystemEventType = "CLIENT_DISCONNECTED"
const SystemEventTypeDaemonSettingsChanged SystemEventType = "DAEMON_SETTINGS_CHANGE"
const SystemEventTypeMonitorRuleViolation SystemEventType = "MONITOR_RULE_VIOLATION"
const SystemEventTypeServiceConnected SystemEventType = "SERVICE_CONNECTED"
const SystemEventTypeServiceDisconnected SystemEventType = "SERVICE_DISCONNECTED"
const SystemEventTypeServiceFatalError SystemEventType = "SERVICE_FATAL_ERROR"
const SystemEventTypeSystemTimeChanged SystemEventType = "SYSTEM_TIME_CHANGE"
const SystemEventTypeWLANCentrAccChanged SystemEventType = "WLAN_CENTR_ACC_CHANGE"

func (self SystemEventType) GetPtr() *SystemEventType { var v = self; return &v }

func (self *SystemEventType) String() string {
	switch *self {
	case SystemEventTypeAny:
		return "+"
	case SystemEventTypeCPEConfigurationError:
		return "CPE_CONFIGURATION_ERROR"
	case SystemEventTypeCPEConfigurationSuccess:
		return "CPE_CONFIGURATION_SUCCESS"
	case SystemEventTypeCPEConnected:
		return "CPE_CONNECTED"
	case SystemEventTypeCPEDisconnected:
		return "CPE_DISCONNECTED"
	case SystemEventTypeClientConnected:
		return "CLIENT_CONNECTED"
	case SystemEventTypeClientDisconnected:
		return "CLIENT_DISCONNECTED"
	case SystemEventTypeDaemonSettingsChanged:
		return "DAEMON_SETTINGS_CHANGE"
	case SystemEventTypeMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION"
	case SystemEventTypeServiceConnected:
		return "SERVICE_CONNECTED"
	case SystemEventTypeServiceDisconnected:
		return "SERVICE_DISCONNECTED"
	case SystemEventTypeServiceFatalError:
		return "SERVICE_FATAL_ERROR"
	case SystemEventTypeSystemTimeChanged:
		return "SYSTEM_TIME_CHANGE"
	case SystemEventTypeWLANCentrAccChanged:
		return "WLAN_CENTR_ACC_CHANGE"

	}
	panic(errors.New("Invalid value of SystemEventType"))

}

func (self *SystemEventType) MarshalJSON() ([]byte, error) {
	switch *self {
	case SystemEventTypeAny:
		return json.Marshal("+")
	case SystemEventTypeCPEConfigurationError:
		return json.Marshal("CPE_CONFIGURATION_ERROR")
	case SystemEventTypeCPEConfigurationSuccess:
		return json.Marshal("CPE_CONFIGURATION_SUCCESS")
	case SystemEventTypeCPEConnected:
		return json.Marshal("CPE_CONNECTED")
	case SystemEventTypeCPEDisconnected:
		return json.Marshal("CPE_DISCONNECTED")
	case SystemEventTypeClientConnected:
		return json.Marshal("CLIENT_CONNECTED")
	case SystemEventTypeClientDisconnected:
		return json.Marshal("CLIENT_DISCONNECTED")
	case SystemEventTypeDaemonSettingsChanged:
		return json.Marshal("DAEMON_SETTINGS_CHANGE")
	case SystemEventTypeMonitorRuleViolation:
		return json.Marshal("MONITOR_RULE_VIOLATION")
	case SystemEventTypeServiceConnected:
		return json.Marshal("SERVICE_CONNECTED")
	case SystemEventTypeServiceDisconnected:
		return json.Marshal("SERVICE_DISCONNECTED")
	case SystemEventTypeServiceFatalError:
		return json.Marshal("SERVICE_FATAL_ERROR")
	case SystemEventTypeSystemTimeChanged:
		return json.Marshal("SYSTEM_TIME_CHANGE")
	case SystemEventTypeWLANCentrAccChanged:
		return json.Marshal("WLAN_CENTR_ACC_CHANGE")

	}
	return nil, errors.New("Invalid value of SystemEventType")

}

func (self *SystemEventType) GetBSON() (interface{}, error) {
	switch *self {
	case SystemEventTypeAny:
		return "+", nil
	case SystemEventTypeCPEConfigurationError:
		return "CPE_CONFIGURATION_ERROR", nil
	case SystemEventTypeCPEConfigurationSuccess:
		return "CPE_CONFIGURATION_SUCCESS", nil
	case SystemEventTypeCPEConnected:
		return "CPE_CONNECTED", nil
	case SystemEventTypeCPEDisconnected:
		return "CPE_DISCONNECTED", nil
	case SystemEventTypeClientConnected:
		return "CLIENT_CONNECTED", nil
	case SystemEventTypeClientDisconnected:
		return "CLIENT_DISCONNECTED", nil
	case SystemEventTypeDaemonSettingsChanged:
		return "DAEMON_SETTINGS_CHANGE", nil
	case SystemEventTypeMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION", nil
	case SystemEventTypeServiceConnected:
		return "SERVICE_CONNECTED", nil
	case SystemEventTypeServiceDisconnected:
		return "SERVICE_DISCONNECTED", nil
	case SystemEventTypeServiceFatalError:
		return "SERVICE_FATAL_ERROR", nil
	case SystemEventTypeSystemTimeChanged:
		return "SYSTEM_TIME_CHANGE", nil
	case SystemEventTypeWLANCentrAccChanged:
		return "WLAN_CENTR_ACC_CHANGE", nil

	}
	return nil, errors.New("Invalid value of SystemEventType")

}
func (self *SystemEventType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		*self = SystemEventTypeAny
		return nil
	case "CPE_CONFIGURATION_ERROR":
		*self = SystemEventTypeCPEConfigurationError
		return nil
	case "CPE_CONFIGURATION_SUCCESS":
		*self = SystemEventTypeCPEConfigurationSuccess
		return nil
	case "CPE_CONNECTED":
		*self = SystemEventTypeCPEConnected
		return nil
	case "CPE_DISCONNECTED":
		*self = SystemEventTypeCPEDisconnected
		return nil
	case "CLIENT_CONNECTED":
		*self = SystemEventTypeClientConnected
		return nil
	case "CLIENT_DISCONNECTED":
		*self = SystemEventTypeClientDisconnected
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		*self = SystemEventTypeDaemonSettingsChanged
		return nil
	case "MONITOR_RULE_VIOLATION":
		*self = SystemEventTypeMonitorRuleViolation
		return nil
	case "SERVICE_CONNECTED":
		*self = SystemEventTypeServiceConnected
		return nil
	case "SERVICE_DISCONNECTED":
		*self = SystemEventTypeServiceDisconnected
		return nil
	case "SERVICE_FATAL_ERROR":
		*self = SystemEventTypeServiceFatalError
		return nil
	case "SYSTEM_TIME_CHANGE":
		*self = SystemEventTypeSystemTimeChanged
		return nil
	case "WLAN_CENTR_ACC_CHANGE":
		*self = SystemEventTypeWLANCentrAccChanged
		return nil

	}
	return errors.New("Unknown SystemEventType")

}

func (self *SystemEventType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "+":
		*self = SystemEventTypeAny
		return nil
	case "CPE_CONFIGURATION_ERROR":
		*self = SystemEventTypeCPEConfigurationError
		return nil
	case "CPE_CONFIGURATION_SUCCESS":
		*self = SystemEventTypeCPEConfigurationSuccess
		return nil
	case "CPE_CONNECTED":
		*self = SystemEventTypeCPEConnected
		return nil
	case "CPE_DISCONNECTED":
		*self = SystemEventTypeCPEDisconnected
		return nil
	case "CLIENT_CONNECTED":
		*self = SystemEventTypeClientConnected
		return nil
	case "CLIENT_DISCONNECTED":
		*self = SystemEventTypeClientDisconnected
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		*self = SystemEventTypeDaemonSettingsChanged
		return nil
	case "MONITOR_RULE_VIOLATION":
		*self = SystemEventTypeMonitorRuleViolation
		return nil
	case "SERVICE_CONNECTED":
		*self = SystemEventTypeServiceConnected
		return nil
	case "SERVICE_DISCONNECTED":
		*self = SystemEventTypeServiceDisconnected
		return nil
	case "SERVICE_FATAL_ERROR":
		*self = SystemEventTypeServiceFatalError
		return nil
	case "SYSTEM_TIME_CHANGE":
		*self = SystemEventTypeSystemTimeChanged
		return nil
	case "WLAN_CENTR_ACC_CHANGE":
		*self = SystemEventTypeWLANCentrAccChanged
		return nil

	}
	return errors.New("Unknown SystemEventType")

}

type WirelessClientState string

const WirelessClientStateCONNECTED WirelessClientState = "CONNECTED"
const WirelessClientStateDISCONNECTED WirelessClientState = "DISCONNECTED"

func (self WirelessClientState) GetPtr() *WirelessClientState { var v = self; return &v }

func (self *WirelessClientState) String() string {
	switch *self {
	case WirelessClientStateCONNECTED:
		return "CONNECTED"
	case WirelessClientStateDISCONNECTED:
		return "DISCONNECTED"

	}
	panic(errors.New("Invalid value of WirelessClientState"))

}

func (self *WirelessClientState) MarshalJSON() ([]byte, error) {
	switch *self {
	case WirelessClientStateCONNECTED:
		return json.Marshal("CONNECTED")
	case WirelessClientStateDISCONNECTED:
		return json.Marshal("DISCONNECTED")

	}
	return nil, errors.New("Invalid value of WirelessClientState")

}

func (self *WirelessClientState) GetBSON() (interface{}, error) {
	switch *self {
	case WirelessClientStateCONNECTED:
		return "CONNECTED", nil
	case WirelessClientStateDISCONNECTED:
		return "DISCONNECTED", nil

	}
	return nil, errors.New("Invalid value of WirelessClientState")

}
func (self *WirelessClientState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*self = WirelessClientStateCONNECTED
		return nil
	case "DISCONNECTED":
		*self = WirelessClientStateDISCONNECTED
		return nil

	}
	return errors.New("Unknown WirelessClientState")

}

func (self *WirelessClientState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*self = WirelessClientStateCONNECTED
		return nil
	case "DISCONNECTED":
		*self = WirelessClientStateDISCONNECTED
		return nil

	}
	return errors.New("Unknown WirelessClientState")

}

type WirelessClientType string

const WirelessClientTypeCamera WirelessClientType = "camera"
const WirelessClientTypeOther WirelessClientType = "other"

func (self WirelessClientType) GetPtr() *WirelessClientType { var v = self; return &v }

func (self *WirelessClientType) String() string {
	switch *self {
	case WirelessClientTypeCamera:
		return "camera"
	case WirelessClientTypeOther:
		return "other"

	}
	panic(errors.New("Invalid value of WirelessClientType"))

}

func (self *WirelessClientType) MarshalJSON() ([]byte, error) {
	switch *self {
	case WirelessClientTypeCamera:
		return json.Marshal("camera")
	case WirelessClientTypeOther:
		return json.Marshal("other")

	}
	return nil, errors.New("Invalid value of WirelessClientType")

}

func (self *WirelessClientType) GetBSON() (interface{}, error) {
	switch *self {
	case WirelessClientTypeCamera:
		return "camera", nil
	case WirelessClientTypeOther:
		return "other", nil

	}
	return nil, errors.New("Invalid value of WirelessClientType")

}
func (self *WirelessClientType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "camera":
		*self = WirelessClientTypeCamera
		return nil
	case "other":
		*self = WirelessClientTypeOther
		return nil

	}
	return errors.New("Unknown WirelessClientType")

}

func (self *WirelessClientType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "camera":
		*self = WirelessClientTypeCamera
		return nil
	case "other":
		*self = WirelessClientTypeOther
		return nil

	}
	return errors.New("Unknown WirelessClientType")

}
