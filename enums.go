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

type BandwidthTypeIface interface {
	BandwidthTypeIfaceFunc()
}
type BandwidthType struct{ BandwidthTypeIface }

func (self *BandwidthType) Value() BandwidthTypeIface { return self.BandwidthTypeIface }

type BandwidthHT20 struct{}

func (BandwidthHT20) BandwidthTypeIfaceFunc() {}

type BandwidthHT40 struct{}

func (BandwidthHT40) BandwidthTypeIfaceFunc() {}

type BandwidthHT40Minus struct{}

func (BandwidthHT40Minus) BandwidthTypeIfaceFunc() {}

type BandwidthHT40Plus struct{}

func (BandwidthHT40Plus) BandwidthTypeIfaceFunc() {}

type BandwidthVHT160 struct{}

func (BandwidthVHT160) BandwidthTypeIfaceFunc() {}

type BandwidthVHT20 struct{}

func (BandwidthVHT20) BandwidthTypeIfaceFunc() {}

type BandwidthVHT40 struct{}

func (BandwidthVHT40) BandwidthTypeIfaceFunc() {}

type BandwidthVHT80 struct{}

func (BandwidthVHT80) BandwidthTypeIfaceFunc() {}
func (self *BandwidthType) String() string {
	switch self.BandwidthTypeIface.(type) {
	case BandwidthHT20:
		return "HT20"
	case BandwidthHT40:
		return "HT40"
	case BandwidthHT40Minus:
		return "HT40-"
	case BandwidthHT40Plus:
		return "HT40+"
	case BandwidthVHT160:
		return "VHT160"
	case BandwidthVHT20:
		return "VHT20"
	case BandwidthVHT40:
		return "VHT40"
	case BandwidthVHT80:
		return "VHT80"

	}
	panic(errors.New("Not implemented"))

}
func (self BandwidthType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case BandwidthHT20:
		return json.Marshal("HT20")
	case BandwidthHT40:
		return json.Marshal("HT40")
	case BandwidthHT40Minus:
		return json.Marshal("HT40-")
	case BandwidthHT40Plus:
		return json.Marshal("HT40+")
	case BandwidthVHT160:
		return json.Marshal("VHT160")
	case BandwidthVHT20:
		return json.Marshal("VHT20")
	case BandwidthVHT40:
		return json.Marshal("VHT40")
	case BandwidthVHT80:
		return json.Marshal("VHT80")

	}
	return nil, errors.New("Not implemented")

}
func (self BandwidthType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("BandwidthType cannot be nil")
	}
	switch v.(type) {
	case BandwidthHT20:
		return "HT20", nil
	case BandwidthHT40:
		return "HT40", nil
	case BandwidthHT40Minus:
		return "HT40-", nil
	case BandwidthHT40Plus:
		return "HT40+", nil
	case BandwidthVHT160:
		return "VHT160", nil
	case BandwidthVHT20:
		return "VHT20", nil
	case BandwidthVHT40:
		return "VHT40", nil
	case BandwidthVHT80:
		return "VHT80", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *BandwidthType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "HT20":
		self.BandwidthTypeIface = BandwidthHT20{}
		return nil
	case "HT40":
		self.BandwidthTypeIface = BandwidthHT40{}
		return nil
	case "HT40-":
		self.BandwidthTypeIface = BandwidthHT40Minus{}
		return nil
	case "HT40+":
		self.BandwidthTypeIface = BandwidthHT40Plus{}
		return nil
	case "VHT160":
		self.BandwidthTypeIface = BandwidthVHT160{}
		return nil
	case "VHT20":
		self.BandwidthTypeIface = BandwidthVHT20{}
		return nil
	case "VHT40":
		self.BandwidthTypeIface = BandwidthVHT40{}
		return nil
	case "VHT80":
		self.BandwidthTypeIface = BandwidthVHT80{}
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
		self.BandwidthTypeIface = BandwidthHT20{}
		return nil
	case "HT40":
		self.BandwidthTypeIface = BandwidthHT40{}
		return nil
	case "HT40-":
		self.BandwidthTypeIface = BandwidthHT40Minus{}
		return nil
	case "HT40+":
		self.BandwidthTypeIface = BandwidthHT40Plus{}
		return nil
	case "VHT160":
		self.BandwidthTypeIface = BandwidthVHT160{}
		return nil
	case "VHT20":
		self.BandwidthTypeIface = BandwidthVHT20{}
		return nil
	case "VHT40":
		self.BandwidthTypeIface = BandwidthVHT40{}
		return nil
	case "VHT80":
		self.BandwidthTypeIface = BandwidthVHT80{}
		return nil

	}
	return errors.New("Unknown BandwidthType")

}

type CPEAgentStatusTypeIface interface {
	CPEAgentStatusTypeIfaceFunc()
}
type CPEAgentStatusType struct{ CPEAgentStatusTypeIface }

func (self *CPEAgentStatusType) Value() CPEAgentStatusTypeIface { return self.CPEAgentStatusTypeIface }

type CPEAgentStatusException struct{}

func (CPEAgentStatusException) CPEAgentStatusTypeIfaceFunc() {}

type CPEAgentStatusSuccess struct{}

func (CPEAgentStatusSuccess) CPEAgentStatusTypeIfaceFunc() {}

type CPEAgentStatusSyntaxError struct{}

func (CPEAgentStatusSyntaxError) CPEAgentStatusTypeIfaceFunc() {}

type CPEAgentStatusUndefined struct{}

func (CPEAgentStatusUndefined) CPEAgentStatusTypeIfaceFunc() {}
func (self *CPEAgentStatusType) String() string {
	switch self.CPEAgentStatusTypeIface.(type) {
	case CPEAgentStatusException:
		return "exception"
	case CPEAgentStatusSuccess:
		return "success"
	case CPEAgentStatusSyntaxError:
		return "syntax"
	case CPEAgentStatusUndefined:
		return "undefined"

	}
	panic(errors.New("Not implemented"))

}
func (self CPEAgentStatusType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case CPEAgentStatusException:
		return json.Marshal("exception")
	case CPEAgentStatusSuccess:
		return json.Marshal("success")
	case CPEAgentStatusSyntaxError:
		return json.Marshal("syntax")
	case CPEAgentStatusUndefined:
		return json.Marshal("undefined")

	}
	return nil, errors.New("Not implemented")

}
func (self CPEAgentStatusType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("CPEAgentStatusType cannot be nil")
	}
	switch v.(type) {
	case CPEAgentStatusException:
		return "exception", nil
	case CPEAgentStatusSuccess:
		return "success", nil
	case CPEAgentStatusSyntaxError:
		return "syntax", nil
	case CPEAgentStatusUndefined:
		return "undefined", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *CPEAgentStatusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "exception":
		self.CPEAgentStatusTypeIface = CPEAgentStatusException{}
		return nil
	case "success":
		self.CPEAgentStatusTypeIface = CPEAgentStatusSuccess{}
		return nil
	case "syntax":
		self.CPEAgentStatusTypeIface = CPEAgentStatusSyntaxError{}
		return nil
	case "undefined":
		self.CPEAgentStatusTypeIface = CPEAgentStatusUndefined{}
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
		self.CPEAgentStatusTypeIface = CPEAgentStatusException{}
		return nil
	case "success":
		self.CPEAgentStatusTypeIface = CPEAgentStatusSuccess{}
		return nil
	case "syntax":
		self.CPEAgentStatusTypeIface = CPEAgentStatusSyntaxError{}
		return nil
	case "undefined":
		self.CPEAgentStatusTypeIface = CPEAgentStatusUndefined{}
		return nil

	}
	return errors.New("Unknown CPEAgentStatusType")

}

type CPEInterfaceTypeIface interface {
	CPEInterfaceTypeIfaceFunc()
}
type CPEInterfaceType struct{ CPEInterfaceTypeIface }

func (self *CPEInterfaceType) Value() CPEInterfaceTypeIface { return self.CPEInterfaceTypeIface }

type InterfaceWiFi struct{}

func (InterfaceWiFi) CPEInterfaceTypeIfaceFunc() {}

type InterfaceWired struct{}

func (InterfaceWired) CPEInterfaceTypeIfaceFunc() {}
func (self *CPEInterfaceType) String() string {
	switch self.CPEInterfaceTypeIface.(type) {
	case InterfaceWiFi:
		return "wifi"
	case InterfaceWired:
		return "wired"

	}
	panic(errors.New("Not implemented"))

}
func (self CPEInterfaceType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case InterfaceWiFi:
		return json.Marshal("wifi")
	case InterfaceWired:
		return json.Marshal("wired")

	}
	return nil, errors.New("Not implemented")

}
func (self CPEInterfaceType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("CPEInterfaceType cannot be nil")
	}
	switch v.(type) {
	case InterfaceWiFi:
		return "wifi", nil
	case InterfaceWired:
		return "wired", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *CPEInterfaceType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "wifi":
		self.CPEInterfaceTypeIface = InterfaceWiFi{}
		return nil
	case "wired":
		self.CPEInterfaceTypeIface = InterfaceWired{}
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
		self.CPEInterfaceTypeIface = InterfaceWiFi{}
		return nil
	case "wired":
		self.CPEInterfaceTypeIface = InterfaceWired{}
		return nil

	}
	return errors.New("Unknown CPEInterfaceType")

}

type ClientStatPacketTypeIface interface {
	ClientStatPacketTypeIfaceFunc()
}
type ClientStatPacketType struct{ ClientStatPacketTypeIface }

func (self *ClientStatPacketType) Value() ClientStatPacketTypeIface {
	return self.ClientStatPacketTypeIface
}

type ClientStatInterim struct{}

func (ClientStatInterim) ClientStatPacketTypeIfaceFunc() {}

type ClientStatOff struct{}

func (ClientStatOff) ClientStatPacketTypeIfaceFunc() {}

type ClientStatOn struct{}

func (ClientStatOn) ClientStatPacketTypeIfaceFunc() {}

type ClientStatStart struct{}

func (ClientStatStart) ClientStatPacketTypeIfaceFunc() {}

type ClientStatStop struct{}

func (ClientStatStop) ClientStatPacketTypeIfaceFunc() {}
func (self *ClientStatPacketType) String() string {
	switch self.ClientStatPacketTypeIface.(type) {
	case ClientStatInterim:
		return "Interim-Update"
	case ClientStatOff:
		return "Accounting-Off"
	case ClientStatOn:
		return "Accounting-On"
	case ClientStatStart:
		return "Start"
	case ClientStatStop:
		return "Stop"

	}
	panic(errors.New("Not implemented"))

}
func (self ClientStatPacketType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case ClientStatInterim:
		return json.Marshal("Interim-Update")
	case ClientStatOff:
		return json.Marshal("Accounting-Off")
	case ClientStatOn:
		return json.Marshal("Accounting-On")
	case ClientStatStart:
		return json.Marshal("Start")
	case ClientStatStop:
		return json.Marshal("Stop")

	}
	return nil, errors.New("Not implemented")

}
func (self ClientStatPacketType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("ClientStatPacketType cannot be nil")
	}
	switch v.(type) {
	case ClientStatInterim:
		return "Interim-Update", nil
	case ClientStatOff:
		return "Accounting-Off", nil
	case ClientStatOn:
		return "Accounting-On", nil
	case ClientStatStart:
		return "Start", nil
	case ClientStatStop:
		return "Stop", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *ClientStatPacketType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Interim-Update":
		self.ClientStatPacketTypeIface = ClientStatInterim{}
		return nil
	case "Accounting-Off":
		self.ClientStatPacketTypeIface = ClientStatOff{}
		return nil
	case "Accounting-On":
		self.ClientStatPacketTypeIface = ClientStatOn{}
		return nil
	case "Start":
		self.ClientStatPacketTypeIface = ClientStatStart{}
		return nil
	case "Stop":
		self.ClientStatPacketTypeIface = ClientStatStop{}
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
		self.ClientStatPacketTypeIface = ClientStatInterim{}
		return nil
	case "Accounting-Off":
		self.ClientStatPacketTypeIface = ClientStatOff{}
		return nil
	case "Accounting-On":
		self.ClientStatPacketTypeIface = ClientStatOn{}
		return nil
	case "Start":
		self.ClientStatPacketTypeIface = ClientStatStart{}
		return nil
	case "Stop":
		self.ClientStatPacketTypeIface = ClientStatStop{}
		return nil

	}
	return errors.New("Unknown ClientStatPacketType")

}

type ConfigurationStatusIface interface {
	ConfigurationStatusIfaceFunc()
}
type ConfigurationStatus struct{ ConfigurationStatusIface }

func (self *ConfigurationStatus) Value() ConfigurationStatusIface {
	return self.ConfigurationStatusIface
}

type StatusEmpty struct{}

func (StatusEmpty) ConfigurationStatusIfaceFunc() {}

type StatusError struct{}

func (StatusError) ConfigurationStatusIfaceFunc() {}

type StatusFuture struct{}

func (StatusFuture) ConfigurationStatusIfaceFunc() {}

type StatusOK struct{}

func (StatusOK) ConfigurationStatusIfaceFunc() {}

type StatusPending struct{}

func (StatusPending) ConfigurationStatusIfaceFunc() {}
func (self *ConfigurationStatus) String() string {
	switch self.ConfigurationStatusIface.(type) {
	case StatusEmpty:
		return "empty"
	case StatusError:
		return "error"
	case StatusFuture:
		return "future"
	case StatusOK:
		return "ok"
	case StatusPending:
		return "pending"

	}
	panic(errors.New("Not implemented"))

}
func (self ConfigurationStatus) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case StatusEmpty:
		return json.Marshal("empty")
	case StatusError:
		return json.Marshal("error")
	case StatusFuture:
		return json.Marshal("future")
	case StatusOK:
		return json.Marshal("ok")
	case StatusPending:
		return json.Marshal("pending")

	}
	return nil, errors.New("Not implemented")

}
func (self ConfigurationStatus) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("ConfigurationStatus cannot be nil")
	}
	switch v.(type) {
	case StatusEmpty:
		return "empty", nil
	case StatusError:
		return "error", nil
	case StatusFuture:
		return "future", nil
	case StatusOK:
		return "ok", nil
	case StatusPending:
		return "pending", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *ConfigurationStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "empty":
		self.ConfigurationStatusIface = StatusEmpty{}
		return nil
	case "error":
		self.ConfigurationStatusIface = StatusError{}
		return nil
	case "future":
		self.ConfigurationStatusIface = StatusFuture{}
		return nil
	case "ok":
		self.ConfigurationStatusIface = StatusOK{}
		return nil
	case "pending":
		self.ConfigurationStatusIface = StatusPending{}
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
		self.ConfigurationStatusIface = StatusEmpty{}
		return nil
	case "error":
		self.ConfigurationStatusIface = StatusError{}
		return nil
	case "future":
		self.ConfigurationStatusIface = StatusFuture{}
		return nil
	case "ok":
		self.ConfigurationStatusIface = StatusOK{}
		return nil
	case "pending":
		self.ConfigurationStatusIface = StatusPending{}
		return nil

	}
	return errors.New("Unknown ConfigurationStatus")

}

type MacFilterTypeIface interface {
	MacFilterTypeIfaceFunc()
}
type MacFilterType struct{ MacFilterTypeIface }

func (self *MacFilterType) Value() MacFilterTypeIface { return self.MacFilterTypeIface }

type MacFilterBlackList struct{}

func (MacFilterBlackList) MacFilterTypeIfaceFunc() {}

type MacFilterNone struct{}

func (MacFilterNone) MacFilterTypeIfaceFunc() {}

type MacFilterWhiteList struct{}

func (MacFilterWhiteList) MacFilterTypeIfaceFunc() {}
func (self *MacFilterType) String() string {
	switch self.MacFilterTypeIface.(type) {
	case MacFilterBlackList:
		return "BlackList"
	case MacFilterNone:
		return "None"
	case MacFilterWhiteList:
		return "WhiteList"

	}
	panic(errors.New("Not implemented"))

}
func (self MacFilterType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case MacFilterBlackList:
		return json.Marshal("BlackList")
	case MacFilterNone:
		return json.Marshal("None")
	case MacFilterWhiteList:
		return json.Marshal("WhiteList")

	}
	return nil, errors.New("Not implemented")

}
func (self MacFilterType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("MacFilterType cannot be nil")
	}
	switch v.(type) {
	case MacFilterBlackList:
		return "BlackList", nil
	case MacFilterNone:
		return "None", nil
	case MacFilterWhiteList:
		return "WhiteList", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *MacFilterType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "BlackList":
		self.MacFilterTypeIface = MacFilterBlackList{}
		return nil
	case "None":
		self.MacFilterTypeIface = MacFilterNone{}
		return nil
	case "WhiteList":
		self.MacFilterTypeIface = MacFilterWhiteList{}
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
		self.MacFilterTypeIface = MacFilterBlackList{}
		return nil
	case "None":
		self.MacFilterTypeIface = MacFilterNone{}
		return nil
	case "WhiteList":
		self.MacFilterTypeIface = MacFilterWhiteList{}
		return nil

	}
	return errors.New("Unknown MacFilterType")

}

type ModuleIface interface {
	ModuleIfaceFunc()
}
type Module struct{ ModuleIface }

func (self *Module) Value() ModuleIface { return self.ModuleIface }

type ModuleAny struct{}

func (ModuleAny) ModuleIfaceFunc() {}

type ModuleBackend struct{}

func (ModuleBackend) ModuleIfaceFunc() {}

type ModuleCPE struct{}

func (ModuleCPE) ModuleIfaceFunc() {}

type ModuleClientStat struct{}

func (ModuleClientStat) ModuleIfaceFunc() {}

type ModuleConfig struct{}

func (ModuleConfig) ModuleIfaceFunc() {}

type ModuleDB struct{}

func (ModuleDB) ModuleIfaceFunc() {}

type ModuleDummy struct{}

func (ModuleDummy) ModuleIfaceFunc() {}

type ModuleLBS struct{}

func (ModuleLBS) ModuleIfaceFunc() {}

type ModuleMQTTLog struct{}

func (ModuleMQTTLog) ModuleIfaceFunc() {}

type ModuleMonitor struct{}

func (ModuleMonitor) ModuleIfaceFunc() {}

type ModuleStat struct{}

func (ModuleStat) ModuleIfaceFunc() {}
func (self *Module) String() string {
	switch self.ModuleIface.(type) {
	case ModuleAny:
		return "+"
	case ModuleBackend:
		return "BACKEND"
	case ModuleCPE:
		return "CPE"
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
	panic(errors.New("Not implemented"))

}
func (self Module) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case ModuleAny:
		return json.Marshal("+")
	case ModuleBackend:
		return json.Marshal("BACKEND")
	case ModuleCPE:
		return json.Marshal("CPE")
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
	return nil, errors.New("Not implemented")

}
func (self Module) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("Module cannot be nil")
	}
	switch v.(type) {
	case ModuleAny:
		return "+", nil
	case ModuleBackend:
		return "BACKEND", nil
	case ModuleCPE:
		return "CPE", nil
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
	return nil, errors.New("Not implemented")

}
func (self *Module) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		self.ModuleIface = ModuleAny{}
		return nil
	case "BACKEND":
		self.ModuleIface = ModuleBackend{}
		return nil
	case "CPE":
		self.ModuleIface = ModuleCPE{}
		return nil
	case "CLIENT_STAT":
		self.ModuleIface = ModuleClientStat{}
		return nil
	case "CONFIG":
		self.ModuleIface = ModuleConfig{}
		return nil
	case "DB":
		self.ModuleIface = ModuleDB{}
		return nil
	case "DUMMY":
		self.ModuleIface = ModuleDummy{}
		return nil
	case "LBS":
		self.ModuleIface = ModuleLBS{}
		return nil
	case "MQTT_LOG":
		self.ModuleIface = ModuleMQTTLog{}
		return nil
	case "MONITOR":
		self.ModuleIface = ModuleMonitor{}
		return nil
	case "STAT":
		self.ModuleIface = ModuleStat{}
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
	case "+":
		self.ModuleIface = ModuleAny{}
		return nil
	case "BACKEND":
		self.ModuleIface = ModuleBackend{}
		return nil
	case "CPE":
		self.ModuleIface = ModuleCPE{}
		return nil
	case "CLIENT_STAT":
		self.ModuleIface = ModuleClientStat{}
		return nil
	case "CONFIG":
		self.ModuleIface = ModuleConfig{}
		return nil
	case "DB":
		self.ModuleIface = ModuleDB{}
		return nil
	case "DUMMY":
		self.ModuleIface = ModuleDummy{}
		return nil
	case "LBS":
		self.ModuleIface = ModuleLBS{}
		return nil
	case "MQTT_LOG":
		self.ModuleIface = ModuleMQTTLog{}
		return nil
	case "MONITOR":
		self.ModuleIface = ModuleMonitor{}
		return nil
	case "STAT":
		self.ModuleIface = ModuleStat{}
		return nil

	}
	return errors.New("Unknown Module")

}

type OperationIface interface {
	OperationIfaceFunc()
}
type Operation struct{ OperationIface }

func (self *Operation) Value() OperationIface { return self.OperationIface }

type OperationAny struct{}

func (OperationAny) OperationIfaceFunc() {}

type OperationCPEStatus struct{}

func (OperationCPEStatus) OperationIfaceFunc() {}

type OperationCreate struct{}

func (OperationCreate) OperationIfaceFunc() {}

type OperationDelete struct{}

func (OperationDelete) OperationIfaceFunc() {}

type OperationLuaScript struct{}

func (OperationLuaScript) OperationIfaceFunc() {}

type OperationRead struct{}

func (OperationRead) OperationIfaceFunc() {}

type OperationSHScript struct{}

func (OperationSHScript) OperationIfaceFunc() {}

type OperationUpdate struct{}

func (OperationUpdate) OperationIfaceFunc() {}
func (self *Operation) String() string {
	switch self.OperationIface.(type) {
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
	panic(errors.New("Not implemented"))

}
func (self Operation) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
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
	return nil, errors.New("Not implemented")

}
func (self Operation) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("Operation cannot be nil")
	}
	switch v.(type) {
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
	return nil, errors.New("Not implemented")

}
func (self *Operation) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		self.OperationIface = OperationAny{}
		return nil
	case "STATUS":
		self.OperationIface = OperationCPEStatus{}
		return nil
	case "C":
		self.OperationIface = OperationCreate{}
		return nil
	case "D":
		self.OperationIface = OperationDelete{}
		return nil
	case "LUA":
		self.OperationIface = OperationLuaScript{}
		return nil
	case "R":
		self.OperationIface = OperationRead{}
		return nil
	case "SH":
		self.OperationIface = OperationSHScript{}
		return nil
	case "U":
		self.OperationIface = OperationUpdate{}
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
		self.OperationIface = OperationAny{}
		return nil
	case "STATUS":
		self.OperationIface = OperationCPEStatus{}
		return nil
	case "C":
		self.OperationIface = OperationCreate{}
		return nil
	case "D":
		self.OperationIface = OperationDelete{}
		return nil
	case "LUA":
		self.OperationIface = OperationLuaScript{}
		return nil
	case "R":
		self.OperationIface = OperationRead{}
		return nil
	case "SH":
		self.OperationIface = OperationSHScript{}
		return nil
	case "U":
		self.OperationIface = OperationUpdate{}
		return nil

	}
	return errors.New("Unknown Operation")

}

type RadiusTypeIface interface {
	RadiusTypeIfaceFunc()
}
type RadiusType struct{ RadiusTypeIface }

func (self *RadiusType) Value() RadiusTypeIface { return self.RadiusTypeIface }

type RadiusAccounting struct{}

func (RadiusAccounting) RadiusTypeIfaceFunc() {}

type RadiusAuthentication struct{}

func (RadiusAuthentication) RadiusTypeIfaceFunc() {}

type RadiusBoth struct{}

func (RadiusBoth) RadiusTypeIfaceFunc() {}
func (self *RadiusType) String() string {
	switch self.RadiusTypeIface.(type) {
	case RadiusAccounting:
		return "acct"
	case RadiusAuthentication:
		return "auth"
	case RadiusBoth:
		return "both"

	}
	panic(errors.New("Not implemented"))

}
func (self RadiusType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case RadiusAccounting:
		return json.Marshal("acct")
	case RadiusAuthentication:
		return json.Marshal("auth")
	case RadiusBoth:
		return json.Marshal("both")

	}
	return nil, errors.New("Not implemented")

}
func (self RadiusType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("RadiusType cannot be nil")
	}
	switch v.(type) {
	case RadiusAccounting:
		return "acct", nil
	case RadiusAuthentication:
		return "auth", nil
	case RadiusBoth:
		return "both", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *RadiusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "acct":
		self.RadiusTypeIface = RadiusAccounting{}
		return nil
	case "auth":
		self.RadiusTypeIface = RadiusAuthentication{}
		return nil
	case "both":
		self.RadiusTypeIface = RadiusBoth{}
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
		self.RadiusTypeIface = RadiusAccounting{}
		return nil
	case "auth":
		self.RadiusTypeIface = RadiusAuthentication{}
		return nil
	case "both":
		self.RadiusTypeIface = RadiusBoth{}
		return nil

	}
	return errors.New("Unknown RadiusType")

}

type SecuritySuiteIface interface {
	SecuritySuiteIfaceFunc()
}
type SecuritySuite struct{ SecuritySuiteIface }

func (self *SecuritySuite) Value() SecuritySuiteIface { return self.SecuritySuiteIface }

type AES struct{}

func (AES) SecuritySuiteIfaceFunc() {}

type TKIP struct{}

func (TKIP) SecuritySuiteIfaceFunc() {}
func (self *SecuritySuite) String() string {
	switch self.SecuritySuiteIface.(type) {
	case AES:
		return "aes"
	case TKIP:
		return "tkip"

	}
	panic(errors.New("Not implemented"))

}
func (self SecuritySuite) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case AES:
		return json.Marshal("aes")
	case TKIP:
		return json.Marshal("tkip")

	}
	return nil, errors.New("Not implemented")

}
func (self SecuritySuite) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("SecuritySuite cannot be nil")
	}
	switch v.(type) {
	case AES:
		return "aes", nil
	case TKIP:
		return "tkip", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *SecuritySuite) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "aes":
		self.SecuritySuiteIface = AES{}
		return nil
	case "tkip":
		self.SecuritySuiteIface = TKIP{}
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
		self.SecuritySuiteIface = AES{}
		return nil
	case "tkip":
		self.SecuritySuiteIface = TKIP{}
		return nil

	}
	return errors.New("Unknown SecuritySuite")

}

type SecurityTypeIface interface {
	SecurityTypeIfaceFunc()
}
type SecurityType struct{ SecurityTypeIface }

func (self *SecurityType) Value() SecurityTypeIface { return self.SecurityTypeIface }

type WPA2Enterprise struct{}

func (WPA2Enterprise) SecurityTypeIfaceFunc() {}

type WPA2Personal struct{}

func (WPA2Personal) SecurityTypeIfaceFunc() {}
func (self *SecurityType) String() string {
	switch self.SecurityTypeIface.(type) {
	case WPA2Enterprise:
		return "wpa2enterprise"
	case WPA2Personal:
		return "wpa2personal"

	}
	panic(errors.New("Not implemented"))

}
func (self SecurityType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case WPA2Enterprise:
		return json.Marshal("wpa2enterprise")
	case WPA2Personal:
		return json.Marshal("wpa2personal")

	}
	return nil, errors.New("Not implemented")

}
func (self SecurityType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("SecurityType cannot be nil")
	}
	switch v.(type) {
	case WPA2Enterprise:
		return "wpa2enterprise", nil
	case WPA2Personal:
		return "wpa2personal", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *SecurityType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "wpa2enterprise":
		self.SecurityTypeIface = WPA2Enterprise{}
		return nil
	case "wpa2personal":
		self.SecurityTypeIface = WPA2Personal{}
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
		self.SecurityTypeIface = WPA2Enterprise{}
		return nil
	case "wpa2personal":
		self.SecurityTypeIface = WPA2Personal{}
		return nil

	}
	return errors.New("Unknown SecurityType")

}

type ServiceStateIface interface {
	ServiceStateIfaceFunc()
}
type ServiceState struct{ ServiceStateIface }

func (self *ServiceState) Value() ServiceStateIface { return self.ServiceStateIface }

type ServiceStateConnected struct{}

func (ServiceStateConnected) ServiceStateIfaceFunc() {}

type ServiceStateDisconnected struct{}

func (ServiceStateDisconnected) ServiceStateIfaceFunc() {}

type ServiceStatePending struct{}

func (ServiceStatePending) ServiceStateIfaceFunc() {}
func (self *ServiceState) String() string {
	switch self.ServiceStateIface.(type) {
	case ServiceStateConnected:
		return "CONNECTED"
	case ServiceStateDisconnected:
		return "DISCONNECTED"
	case ServiceStatePending:
		return "PENDING"

	}
	panic(errors.New("Not implemented"))

}
func (self ServiceState) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case ServiceStateConnected:
		return json.Marshal("CONNECTED")
	case ServiceStateDisconnected:
		return json.Marshal("DISCONNECTED")
	case ServiceStatePending:
		return json.Marshal("PENDING")

	}
	return nil, errors.New("Not implemented")

}
func (self ServiceState) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("ServiceState cannot be nil")
	}
	switch v.(type) {
	case ServiceStateConnected:
		return "CONNECTED", nil
	case ServiceStateDisconnected:
		return "DISCONNECTED", nil
	case ServiceStatePending:
		return "PENDING", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *ServiceState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		self.ServiceStateIface = ServiceStateConnected{}
		return nil
	case "DISCONNECTED":
		self.ServiceStateIface = ServiceStateDisconnected{}
		return nil
	case "PENDING":
		self.ServiceStateIface = ServiceStatePending{}
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
		self.ServiceStateIface = ServiceStateConnected{}
		return nil
	case "DISCONNECTED":
		self.ServiceStateIface = ServiceStateDisconnected{}
		return nil
	case "PENDING":
		self.ServiceStateIface = ServiceStatePending{}
		return nil

	}
	return errors.New("Unknown ServiceState")

}

type StatEventRuleTypeIface interface {
	StatEventRuleTypeIfaceFunc()
}
type StatEventRuleType struct{ StatEventRuleTypeIface }

func (self *StatEventRuleType) Value() StatEventRuleTypeIface { return self.StatEventRuleTypeIface }

type StatEventCPUload struct{}

func (StatEventCPUload) StatEventRuleTypeIfaceFunc() {}

type StatEventFreeRAM struct{}

func (StatEventFreeRAM) StatEventRuleTypeIfaceFunc() {}
func (self *StatEventRuleType) String() string {
	switch self.StatEventRuleTypeIface.(type) {
	case StatEventCPUload:
		return "cpu_load"
	case StatEventFreeRAM:
		return "free_ram"

	}
	panic(errors.New("Not implemented"))

}
func (self StatEventRuleType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case StatEventCPUload:
		return json.Marshal("cpu_load")
	case StatEventFreeRAM:
		return json.Marshal("free_ram")

	}
	return nil, errors.New("Not implemented")

}
func (self StatEventRuleType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("StatEventRuleType cannot be nil")
	}
	switch v.(type) {
	case StatEventCPUload:
		return "cpu_load", nil
	case StatEventFreeRAM:
		return "free_ram", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *StatEventRuleType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "cpu_load":
		self.StatEventRuleTypeIface = StatEventCPUload{}
		return nil
	case "free_ram":
		self.StatEventRuleTypeIface = StatEventFreeRAM{}
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
		self.StatEventRuleTypeIface = StatEventCPUload{}
		return nil
	case "free_ram":
		self.StatEventRuleTypeIface = StatEventFreeRAM{}
		return nil

	}
	return errors.New("Unknown StatEventRuleType")

}

type SystemEventLevelIface interface {
	SystemEventLevelIfaceFunc()
}
type SystemEventLevel struct{ SystemEventLevelIface }

func (self *SystemEventLevel) Value() SystemEventLevelIface { return self.SystemEventLevelIface }

type SystemEventLevelDEBUG struct{}

func (SystemEventLevelDEBUG) SystemEventLevelIfaceFunc() {}

type SystemEventLevelERROR struct{}

func (SystemEventLevelERROR) SystemEventLevelIfaceFunc() {}

type SystemEventLevelINFO struct{}

func (SystemEventLevelINFO) SystemEventLevelIfaceFunc() {}

type SystemEventLevelWARNING struct{}

func (SystemEventLevelWARNING) SystemEventLevelIfaceFunc() {}
func (self *SystemEventLevel) String() string {
	switch self.SystemEventLevelIface.(type) {
	case SystemEventLevelDEBUG:
		return "DEBUG"
	case SystemEventLevelERROR:
		return "ERROR"
	case SystemEventLevelINFO:
		return "INFO"
	case SystemEventLevelWARNING:
		return "WARNING"

	}
	panic(errors.New("Not implemented"))

}
func (self SystemEventLevel) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case SystemEventLevelDEBUG:
		return json.Marshal("DEBUG")
	case SystemEventLevelERROR:
		return json.Marshal("ERROR")
	case SystemEventLevelINFO:
		return json.Marshal("INFO")
	case SystemEventLevelWARNING:
		return json.Marshal("WARNING")

	}
	return nil, errors.New("Not implemented")

}
func (self SystemEventLevel) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("SystemEventLevel cannot be nil")
	}
	switch v.(type) {
	case SystemEventLevelDEBUG:
		return "DEBUG", nil
	case SystemEventLevelERROR:
		return "ERROR", nil
	case SystemEventLevelINFO:
		return "INFO", nil
	case SystemEventLevelWARNING:
		return "WARNING", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *SystemEventLevel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "DEBUG":
		self.SystemEventLevelIface = SystemEventLevelDEBUG{}
		return nil
	case "ERROR":
		self.SystemEventLevelIface = SystemEventLevelERROR{}
		return nil
	case "INFO":
		self.SystemEventLevelIface = SystemEventLevelINFO{}
		return nil
	case "WARNING":
		self.SystemEventLevelIface = SystemEventLevelWARNING{}
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
		self.SystemEventLevelIface = SystemEventLevelDEBUG{}
		return nil
	case "ERROR":
		self.SystemEventLevelIface = SystemEventLevelERROR{}
		return nil
	case "INFO":
		self.SystemEventLevelIface = SystemEventLevelINFO{}
		return nil
	case "WARNING":
		self.SystemEventLevelIface = SystemEventLevelWARNING{}
		return nil

	}
	return errors.New("Unknown SystemEventLevel")

}

type SystemEventObjectTypeIface interface {
	SystemEventObjectTypeIfaceFunc()
}
type SystemEventObjectType struct{ SystemEventObjectTypeIface }

func (self *SystemEventObjectType) Value() SystemEventObjectTypeIface {
	return self.SystemEventObjectTypeIface
}

type SystemEventAny struct{}

func (SystemEventAny) SystemEventObjectTypeIfaceFunc() {}

type SystemEventCPEConfigurationError struct{}

func (SystemEventCPEConfigurationError) SystemEventObjectTypeIfaceFunc() {}

type SystemEventCPEConnected struct{}

func (SystemEventCPEConnected) SystemEventObjectTypeIfaceFunc() {}

type SystemEventCPEDisconnected struct{}

func (SystemEventCPEDisconnected) SystemEventObjectTypeIfaceFunc() {}

type SystemEventClientConnected struct{}

func (SystemEventClientConnected) SystemEventObjectTypeIfaceFunc() {}

type SystemEventClientDisconnected struct{}

func (SystemEventClientDisconnected) SystemEventObjectTypeIfaceFunc() {}

type SystemEventDaemonSettingsChanged struct{}

func (SystemEventDaemonSettingsChanged) SystemEventObjectTypeIfaceFunc() {}

type SystemEventMonitorRuleViolation struct{}

func (SystemEventMonitorRuleViolation) SystemEventObjectTypeIfaceFunc() {}

type SystemEventServiceConnected struct{}

func (SystemEventServiceConnected) SystemEventObjectTypeIfaceFunc() {}

type SystemEventServiceDisconnected struct{}

func (SystemEventServiceDisconnected) SystemEventObjectTypeIfaceFunc() {}

type SystemEventServiceFatalError struct{}

func (SystemEventServiceFatalError) SystemEventObjectTypeIfaceFunc() {}
func (self *SystemEventObjectType) String() string {
	switch self.SystemEventObjectTypeIface.(type) {
	case SystemEventAny:
		return "+"
	case SystemEventCPEConfigurationError:
		return "CPE_CONFIGURATION_ERROR"
	case SystemEventCPEConnected:
		return "CPE_CONNECTED"
	case SystemEventCPEDisconnected:
		return "CPE_DISCONNECTED"
	case SystemEventClientConnected:
		return "CLIENT_CONNECTED"
	case SystemEventClientDisconnected:
		return "CLIENT_DISCONNECTED"
	case SystemEventDaemonSettingsChanged:
		return "DAEMON_SETTINGS_CHANGE"
	case SystemEventMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION"
	case SystemEventServiceConnected:
		return "SERVICE_CONNECTED"
	case SystemEventServiceDisconnected:
		return "SERVICE_DISCONNECTED"
	case SystemEventServiceFatalError:
		return "SERVICE_FATAL_ERROR"

	}
	panic(errors.New("Not implemented"))

}
func (self SystemEventObjectType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case SystemEventAny:
		return json.Marshal("+")
	case SystemEventCPEConfigurationError:
		return json.Marshal("CPE_CONFIGURATION_ERROR")
	case SystemEventCPEConnected:
		return json.Marshal("CPE_CONNECTED")
	case SystemEventCPEDisconnected:
		return json.Marshal("CPE_DISCONNECTED")
	case SystemEventClientConnected:
		return json.Marshal("CLIENT_CONNECTED")
	case SystemEventClientDisconnected:
		return json.Marshal("CLIENT_DISCONNECTED")
	case SystemEventDaemonSettingsChanged:
		return json.Marshal("DAEMON_SETTINGS_CHANGE")
	case SystemEventMonitorRuleViolation:
		return json.Marshal("MONITOR_RULE_VIOLATION")
	case SystemEventServiceConnected:
		return json.Marshal("SERVICE_CONNECTED")
	case SystemEventServiceDisconnected:
		return json.Marshal("SERVICE_DISCONNECTED")
	case SystemEventServiceFatalError:
		return json.Marshal("SERVICE_FATAL_ERROR")

	}
	return nil, errors.New("Not implemented")

}
func (self SystemEventObjectType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("SystemEventObjectType cannot be nil")
	}
	switch v.(type) {
	case SystemEventAny:
		return "+", nil
	case SystemEventCPEConfigurationError:
		return "CPE_CONFIGURATION_ERROR", nil
	case SystemEventCPEConnected:
		return "CPE_CONNECTED", nil
	case SystemEventCPEDisconnected:
		return "CPE_DISCONNECTED", nil
	case SystemEventClientConnected:
		return "CLIENT_CONNECTED", nil
	case SystemEventClientDisconnected:
		return "CLIENT_DISCONNECTED", nil
	case SystemEventDaemonSettingsChanged:
		return "DAEMON_SETTINGS_CHANGE", nil
	case SystemEventMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION", nil
	case SystemEventServiceConnected:
		return "SERVICE_CONNECTED", nil
	case SystemEventServiceDisconnected:
		return "SERVICE_DISCONNECTED", nil
	case SystemEventServiceFatalError:
		return "SERVICE_FATAL_ERROR", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *SystemEventObjectType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		self.SystemEventObjectTypeIface = SystemEventAny{}
		return nil
	case "CPE_CONFIGURATION_ERROR":
		self.SystemEventObjectTypeIface = SystemEventCPEConfigurationError{}
		return nil
	case "CPE_CONNECTED":
		self.SystemEventObjectTypeIface = SystemEventCPEConnected{}
		return nil
	case "CPE_DISCONNECTED":
		self.SystemEventObjectTypeIface = SystemEventCPEDisconnected{}
		return nil
	case "CLIENT_CONNECTED":
		self.SystemEventObjectTypeIface = SystemEventClientConnected{}
		return nil
	case "CLIENT_DISCONNECTED":
		self.SystemEventObjectTypeIface = SystemEventClientDisconnected{}
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		self.SystemEventObjectTypeIface = SystemEventDaemonSettingsChanged{}
		return nil
	case "MONITOR_RULE_VIOLATION":
		self.SystemEventObjectTypeIface = SystemEventMonitorRuleViolation{}
		return nil
	case "SERVICE_CONNECTED":
		self.SystemEventObjectTypeIface = SystemEventServiceConnected{}
		return nil
	case "SERVICE_DISCONNECTED":
		self.SystemEventObjectTypeIface = SystemEventServiceDisconnected{}
		return nil
	case "SERVICE_FATAL_ERROR":
		self.SystemEventObjectTypeIface = SystemEventServiceFatalError{}
		return nil

	}
	return errors.New("Unknown SystemEventObjectType")

}

func (self *SystemEventObjectType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "+":
		self.SystemEventObjectTypeIface = SystemEventAny{}
		return nil
	case "CPE_CONFIGURATION_ERROR":
		self.SystemEventObjectTypeIface = SystemEventCPEConfigurationError{}
		return nil
	case "CPE_CONNECTED":
		self.SystemEventObjectTypeIface = SystemEventCPEConnected{}
		return nil
	case "CPE_DISCONNECTED":
		self.SystemEventObjectTypeIface = SystemEventCPEDisconnected{}
		return nil
	case "CLIENT_CONNECTED":
		self.SystemEventObjectTypeIface = SystemEventClientConnected{}
		return nil
	case "CLIENT_DISCONNECTED":
		self.SystemEventObjectTypeIface = SystemEventClientDisconnected{}
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		self.SystemEventObjectTypeIface = SystemEventDaemonSettingsChanged{}
		return nil
	case "MONITOR_RULE_VIOLATION":
		self.SystemEventObjectTypeIface = SystemEventMonitorRuleViolation{}
		return nil
	case "SERVICE_CONNECTED":
		self.SystemEventObjectTypeIface = SystemEventServiceConnected{}
		return nil
	case "SERVICE_DISCONNECTED":
		self.SystemEventObjectTypeIface = SystemEventServiceDisconnected{}
		return nil
	case "SERVICE_FATAL_ERROR":
		self.SystemEventObjectTypeIface = SystemEventServiceFatalError{}
		return nil

	}
	return errors.New("Unknown SystemEventObjectType")

}

type WirelessClientStateIface interface {
	WirelessClientStateIfaceFunc()
}
type WirelessClientState struct{ WirelessClientStateIface }

func (self *WirelessClientState) Value() WirelessClientStateIface {
	return self.WirelessClientStateIface
}

type WirelessClientStateCONNECTED struct{}

func (WirelessClientStateCONNECTED) WirelessClientStateIfaceFunc() {}

type WirelessClientStateDISCONNECTED struct{}

func (WirelessClientStateDISCONNECTED) WirelessClientStateIfaceFunc() {}
func (self *WirelessClientState) String() string {
	switch self.WirelessClientStateIface.(type) {
	case WirelessClientStateCONNECTED:
		return "CONNECTED"
	case WirelessClientStateDISCONNECTED:
		return "DISCONNECTED"

	}
	panic(errors.New("Not implemented"))

}
func (self WirelessClientState) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case WirelessClientStateCONNECTED:
		return json.Marshal("CONNECTED")
	case WirelessClientStateDISCONNECTED:
		return json.Marshal("DISCONNECTED")

	}
	return nil, errors.New("Not implemented")

}
func (self WirelessClientState) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("WirelessClientState cannot be nil")
	}
	switch v.(type) {
	case WirelessClientStateCONNECTED:
		return "CONNECTED", nil
	case WirelessClientStateDISCONNECTED:
		return "DISCONNECTED", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *WirelessClientState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		self.WirelessClientStateIface = WirelessClientStateCONNECTED{}
		return nil
	case "DISCONNECTED":
		self.WirelessClientStateIface = WirelessClientStateDISCONNECTED{}
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
		self.WirelessClientStateIface = WirelessClientStateCONNECTED{}
		return nil
	case "DISCONNECTED":
		self.WirelessClientStateIface = WirelessClientStateDISCONNECTED{}
		return nil

	}
	return errors.New("Unknown WirelessClientState")

}

type WirelessClientTypeIface interface {
	WirelessClientTypeIfaceFunc()
}
type WirelessClientType struct{ WirelessClientTypeIface }

func (self *WirelessClientType) Value() WirelessClientTypeIface { return self.WirelessClientTypeIface }

type WirelessClientCamera struct{}

func (WirelessClientCamera) WirelessClientTypeIfaceFunc() {}

type WirelessClientOther struct{}

func (WirelessClientOther) WirelessClientTypeIfaceFunc() {}
func (self *WirelessClientType) String() string {
	switch self.WirelessClientTypeIface.(type) {
	case WirelessClientCamera:
		return "camera"
	case WirelessClientOther:
		return "other"

	}
	panic(errors.New("Not implemented"))

}
func (self WirelessClientType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case WirelessClientCamera:
		return json.Marshal("camera")
	case WirelessClientOther:
		return json.Marshal("other")

	}
	return nil, errors.New("Not implemented")

}
func (self WirelessClientType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("WirelessClientType cannot be nil")
	}
	switch v.(type) {
	case WirelessClientCamera:
		return "camera", nil
	case WirelessClientOther:
		return "other", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *WirelessClientType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "camera":
		self.WirelessClientTypeIface = WirelessClientCamera{}
		return nil
	case "other":
		self.WirelessClientTypeIface = WirelessClientOther{}
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
		self.WirelessClientTypeIface = WirelessClientCamera{}
		return nil
	case "other":
		self.WirelessClientTypeIface = WirelessClientOther{}
		return nil

	}
	return errors.New("Unknown WirelessClientType")

}
