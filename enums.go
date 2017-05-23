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

type EventTypeIface interface {
	EventTypeIfaceFunc()
}
type EventType struct{ EventTypeIface }

func (self *EventType) Value() EventTypeIface { return self.EventTypeIface }

type EventStatRuleViolation struct{}

func (EventStatRuleViolation) EventTypeIfaceFunc() {}

type EventStatRulesChanged struct{}

func (EventStatRulesChanged) EventTypeIfaceFunc() {}
func (self *EventType) String() string {
	switch self.EventTypeIface.(type) {
	case EventStatRuleViolation:
		return "STAT_RULE_VIOLATION"
	case EventStatRulesChanged:
		return "STAT_RULES_CHANGED"

	}
	panic(errors.New("Not implemented"))

}
func (self EventType) MarshalJSON() ([]byte, error) {
	switch self.Value().(type) {
	case EventStatRuleViolation:
		return json.Marshal("STAT_RULE_VIOLATION")
	case EventStatRulesChanged:
		return json.Marshal("STAT_RULES_CHANGED")

	}
	return nil, errors.New("Not implemented")

}
func (self EventType) GetBSON() (interface{}, error) {
	var v = self.Value()
	if v == nil {
		return nil, errors.New("EventType cannot be nil")
	}
	switch v.(type) {
	case EventStatRuleViolation:
		return "STAT_RULE_VIOLATION", nil
	case EventStatRulesChanged:
		return "STAT_RULES_CHANGED", nil

	}
	return nil, errors.New("Not implemented")

}
func (self *EventType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "STAT_RULE_VIOLATION":
		self.EventTypeIface = EventStatRuleViolation{}
		return nil
	case "STAT_RULES_CHANGED":
		self.EventTypeIface = EventStatRulesChanged{}
		return nil

	}
	return errors.New("Unknown EventType")

}

func (self *EventType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "STAT_RULE_VIOLATION":
		self.EventTypeIface = EventStatRuleViolation{}
		return nil
	case "STAT_RULES_CHANGED":
		self.EventTypeIface = EventStatRulesChanged{}
		return nil

	}
	return errors.New("Unknown EventType")

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

type ModuleConfig struct{}

func (ModuleConfig) ModuleIfaceFunc() {}

type ModuleDB struct{}

func (ModuleDB) ModuleIfaceFunc() {}

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
	case ModuleConfig:
		return "CONFIG"
	case ModuleDB:
		return "DB"
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
	case ModuleConfig:
		return json.Marshal("CONFIG")
	case ModuleDB:
		return json.Marshal("DB")
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
	case ModuleConfig:
		return "CONFIG", nil
	case ModuleDB:
		return "DB", nil
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
	case "CONFIG":
		self.ModuleIface = ModuleConfig{}
		return nil
	case "DB":
		self.ModuleIface = ModuleDB{}
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
	case "CONFIG":
		self.ModuleIface = ModuleConfig{}
		return nil
	case "DB":
		self.ModuleIface = ModuleDB{}
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
