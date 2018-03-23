package libwimark

import (
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

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
	if len(*self) == 0 {
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
	if len(*self) == 0 {
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
	if len(*self) == 0 {
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
	if len(s) == 0 {
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
	if len(s) == 0 {
		*self = CPEAgentStatusTypeUndefined
		return nil
	}
	return errors.New("Unknown CPEAgentStatusType")
}

type CPEInterfaceState string

const CPEInterfaceStateACS CPEInterfaceState = "ACS"
const CPEInterfaceStateCountryUpdate CPEInterfaceState = "CONTRY_UPDATE"
const CPEInterfaceStateDFS CPEInterfaceState = "DFS"
const CPEInterfaceStateDisabled CPEInterfaceState = "DISABLED"
const CPEInterfaceStateEnabled CPEInterfaceState = "ENABLED"
const CPEInterfaceStateHtScan CPEInterfaceState = "HT_SCAN"
const CPEInterfaceStateTerminated CPEInterfaceState = "TERMINATED"
const CPEInterfaceStateUninitialized CPEInterfaceState = "UNINITIALIZED"
const CPEInterfaceStateUnknown CPEInterfaceState = "UNKNOWN"

func (self CPEInterfaceState) GetPtr() *CPEInterfaceState { var v = self; return &v }

func (self *CPEInterfaceState) String() string {
	switch *self {
	case CPEInterfaceStateACS:
		return "ACS"
	case CPEInterfaceStateCountryUpdate:
		return "CONTRY_UPDATE"
	case CPEInterfaceStateDFS:
		return "DFS"
	case CPEInterfaceStateDisabled:
		return "DISABLED"
	case CPEInterfaceStateEnabled:
		return "ENABLED"
	case CPEInterfaceStateHtScan:
		return "HT_SCAN"
	case CPEInterfaceStateTerminated:
		return "TERMINATED"
	case CPEInterfaceStateUninitialized:
		return "UNINITIALIZED"
	case CPEInterfaceStateUnknown:
		return "UNKNOWN"
	}
	panic(errors.New("Invalid value of CPEInterfaceState"))
}

func (self *CPEInterfaceState) MarshalJSON() ([]byte, error) {
	switch *self {
	case CPEInterfaceStateACS:
		return json.Marshal("ACS")
	case CPEInterfaceStateCountryUpdate:
		return json.Marshal("CONTRY_UPDATE")
	case CPEInterfaceStateDFS:
		return json.Marshal("DFS")
	case CPEInterfaceStateDisabled:
		return json.Marshal("DISABLED")
	case CPEInterfaceStateEnabled:
		return json.Marshal("ENABLED")
	case CPEInterfaceStateHtScan:
		return json.Marshal("HT_SCAN")
	case CPEInterfaceStateTerminated:
		return json.Marshal("TERMINATED")
	case CPEInterfaceStateUninitialized:
		return json.Marshal("UNINITIALIZED")
	case CPEInterfaceStateUnknown:
		return json.Marshal("UNKNOWN")
	}
	return nil, errors.New("Invalid value of CPEInterfaceState")
}

func (self *CPEInterfaceState) GetBSON() (interface{}, error) {
	switch *self {
	case CPEInterfaceStateACS:
		return "ACS", nil
	case CPEInterfaceStateCountryUpdate:
		return "CONTRY_UPDATE", nil
	case CPEInterfaceStateDFS:
		return "DFS", nil
	case CPEInterfaceStateDisabled:
		return "DISABLED", nil
	case CPEInterfaceStateEnabled:
		return "ENABLED", nil
	case CPEInterfaceStateHtScan:
		return "HT_SCAN", nil
	case CPEInterfaceStateTerminated:
		return "TERMINATED", nil
	case CPEInterfaceStateUninitialized:
		return "UNINITIALIZED", nil
	case CPEInterfaceStateUnknown:
		return "UNKNOWN", nil
	}
	return nil, errors.New("Invalid value of CPEInterfaceState")
}

func (self *CPEInterfaceState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ACS":
		*self = CPEInterfaceStateACS
		return nil
	case "CONTRY_UPDATE":
		*self = CPEInterfaceStateCountryUpdate
		return nil
	case "DFS":
		*self = CPEInterfaceStateDFS
		return nil
	case "DISABLED":
		*self = CPEInterfaceStateDisabled
		return nil
	case "ENABLED":
		*self = CPEInterfaceStateEnabled
		return nil
	case "HT_SCAN":
		*self = CPEInterfaceStateHtScan
		return nil
	case "TERMINATED":
		*self = CPEInterfaceStateTerminated
		return nil
	case "UNINITIALIZED":
		*self = CPEInterfaceStateUninitialized
		return nil
	case "UNKNOWN":
		*self = CPEInterfaceStateUnknown
		return nil
	}
	return errors.New("Unknown CPEInterfaceState")
}

func (self *CPEInterfaceState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ACS":
		*self = CPEInterfaceStateACS
		return nil
	case "CONTRY_UPDATE":
		*self = CPEInterfaceStateCountryUpdate
		return nil
	case "DFS":
		*self = CPEInterfaceStateDFS
		return nil
	case "DISABLED":
		*self = CPEInterfaceStateDisabled
		return nil
	case "ENABLED":
		*self = CPEInterfaceStateEnabled
		return nil
	case "HT_SCAN":
		*self = CPEInterfaceStateHtScan
		return nil
	case "TERMINATED":
		*self = CPEInterfaceStateTerminated
		return nil
	case "UNINITIALIZED":
		*self = CPEInterfaceStateUninitialized
		return nil
	case "UNKNOWN":
		*self = CPEInterfaceStateUnknown
		return nil
	}
	return errors.New("Unknown CPEInterfaceState")
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
const ConfigurationStatusUpgrading ConfigurationStatus = "upgrading"

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
	case ConfigurationStatusUpgrading:
		return "upgrading"
	}
	if len(*self) == 0 {
		return "empty"
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
	case ConfigurationStatusUpgrading:
		return json.Marshal("upgrading")
	}
	if len(*self) == 0 {
		return json.Marshal("empty")
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
	case ConfigurationStatusUpgrading:
		return "upgrading", nil
	}
	if len(*self) == 0 {
		return "empty", nil
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
	case "upgrading":
		*self = ConfigurationStatusUpgrading
		return nil
	}
	if len(s) == 0 {
		*self = ConfigurationStatusEmpty
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
	case "upgrading":
		*self = ConfigurationStatusUpgrading
		return nil
	}
	if len(s) == 0 {
		*self = ConfigurationStatusEmpty
		return nil
	}
	return errors.New("Unknown ConfigurationStatus")
}

type ConnectionModeType string

const ConnectionModeTypeModeAC ConnectionModeType = "ac"
const ConnectionModeTypeModeLegacy ConnectionModeType = "legacy"
const ConnectionModeTypeModeN ConnectionModeType = "n"

func (self ConnectionModeType) GetPtr() *ConnectionModeType { var v = self; return &v }

func (self *ConnectionModeType) String() string {
	switch *self {
	case ConnectionModeTypeModeAC:
		return "ac"
	case ConnectionModeTypeModeLegacy:
		return "legacy"
	case ConnectionModeTypeModeN:
		return "n"
	}
	if len(*self) == 0 {
		return "legacy"
	}
	panic(errors.New("Invalid value of ConnectionModeType"))
}

func (self *ConnectionModeType) MarshalJSON() ([]byte, error) {
	switch *self {
	case ConnectionModeTypeModeAC:
		return json.Marshal("ac")
	case ConnectionModeTypeModeLegacy:
		return json.Marshal("legacy")
	case ConnectionModeTypeModeN:
		return json.Marshal("n")
	}
	if len(*self) == 0 {
		return json.Marshal("legacy")
	}
	return nil, errors.New("Invalid value of ConnectionModeType")
}

func (self *ConnectionModeType) GetBSON() (interface{}, error) {
	switch *self {
	case ConnectionModeTypeModeAC:
		return "ac", nil
	case ConnectionModeTypeModeLegacy:
		return "legacy", nil
	case ConnectionModeTypeModeN:
		return "n", nil
	}
	if len(*self) == 0 {
		return "legacy", nil
	}
	return nil, errors.New("Invalid value of ConnectionModeType")
}

func (self *ConnectionModeType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ac":
		*self = ConnectionModeTypeModeAC
		return nil
	case "legacy":
		*self = ConnectionModeTypeModeLegacy
		return nil
	case "n":
		*self = ConnectionModeTypeModeN
		return nil
	}
	if len(s) == 0 {
		*self = ConnectionModeTypeModeLegacy
		return nil
	}
	return errors.New("Unknown ConnectionModeType")
}

func (self *ConnectionModeType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ac":
		*self = ConnectionModeTypeModeAC
		return nil
	case "legacy":
		*self = ConnectionModeTypeModeLegacy
		return nil
	case "n":
		*self = ConnectionModeTypeModeN
		return nil
	}
	if len(s) == 0 {
		*self = ConnectionModeTypeModeLegacy
		return nil
	}
	return errors.New("Unknown ConnectionModeType")
}

type FirewallDirection string

const FirewallDirectionAny FirewallDirection = "ANY"
const FirewallDirectionIn FirewallDirection = "IN"
const FirewallDirectionOut FirewallDirection = "OUT"

func (self FirewallDirection) GetPtr() *FirewallDirection { var v = self; return &v }

func (self *FirewallDirection) String() string {
	switch *self {
	case FirewallDirectionAny:
		return "ANY"
	case FirewallDirectionIn:
		return "IN"
	case FirewallDirectionOut:
		return "OUT"
	}
	if len(*self) == 0 {
		return "ANY"
	}
	panic(errors.New("Invalid value of FirewallDirection"))
}

func (self *FirewallDirection) MarshalJSON() ([]byte, error) {
	switch *self {
	case FirewallDirectionAny:
		return json.Marshal("ANY")
	case FirewallDirectionIn:
		return json.Marshal("IN")
	case FirewallDirectionOut:
		return json.Marshal("OUT")
	}
	if len(*self) == 0 {
		return json.Marshal("ANY")
	}
	return nil, errors.New("Invalid value of FirewallDirection")
}

func (self *FirewallDirection) GetBSON() (interface{}, error) {
	switch *self {
	case FirewallDirectionAny:
		return "ANY", nil
	case FirewallDirectionIn:
		return "IN", nil
	case FirewallDirectionOut:
		return "OUT", nil
	}
	if len(*self) == 0 {
		return "ANY", nil
	}
	return nil, errors.New("Invalid value of FirewallDirection")
}

func (self *FirewallDirection) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ANY":
		*self = FirewallDirectionAny
		return nil
	case "IN":
		*self = FirewallDirectionIn
		return nil
	case "OUT":
		*self = FirewallDirectionOut
		return nil
	}
	if len(s) == 0 {
		*self = FirewallDirectionAny
		return nil
	}
	return errors.New("Unknown FirewallDirection")
}

func (self *FirewallDirection) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ANY":
		*self = FirewallDirectionAny
		return nil
	case "IN":
		*self = FirewallDirectionIn
		return nil
	case "OUT":
		*self = FirewallDirectionOut
		return nil
	}
	if len(s) == 0 {
		*self = FirewallDirectionAny
		return nil
	}
	return errors.New("Unknown FirewallDirection")
}

type FirewallPolicy string

const FirewallPolicyAccept FirewallPolicy = "ACCEPT"
const FirewallPolicyDrop FirewallPolicy = "DROP"
const FirewallPolicyEmpty FirewallPolicy = ""
const FirewallPolicyReturn FirewallPolicy = "RETURN"

func (self FirewallPolicy) GetPtr() *FirewallPolicy { var v = self; return &v }

func (self *FirewallPolicy) String() string {
	switch *self {
	case FirewallPolicyAccept:
		return "ACCEPT"
	case FirewallPolicyDrop:
		return "DROP"
	case FirewallPolicyEmpty:
		return ""
	case FirewallPolicyReturn:
		return "RETURN"
	}
	if len(*self) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of FirewallPolicy"))
}

func (self *FirewallPolicy) MarshalJSON() ([]byte, error) {
	switch *self {
	case FirewallPolicyAccept:
		return json.Marshal("ACCEPT")
	case FirewallPolicyDrop:
		return json.Marshal("DROP")
	case FirewallPolicyEmpty:
		return json.Marshal("")
	case FirewallPolicyReturn:
		return json.Marshal("RETURN")
	}
	if len(*self) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of FirewallPolicy")
}

func (self *FirewallPolicy) GetBSON() (interface{}, error) {
	switch *self {
	case FirewallPolicyAccept:
		return "ACCEPT", nil
	case FirewallPolicyDrop:
		return "DROP", nil
	case FirewallPolicyEmpty:
		return "", nil
	case FirewallPolicyReturn:
		return "RETURN", nil
	}
	if len(*self) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of FirewallPolicy")
}

func (self *FirewallPolicy) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ACCEPT":
		*self = FirewallPolicyAccept
		return nil
	case "DROP":
		*self = FirewallPolicyDrop
		return nil
	case "":
		*self = FirewallPolicyEmpty
		return nil
	case "RETURN":
		*self = FirewallPolicyReturn
		return nil
	}
	if len(s) == 0 {
		*self = FirewallPolicyEmpty
		return nil
	}
	return errors.New("Unknown FirewallPolicy")
}

func (self *FirewallPolicy) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ACCEPT":
		*self = FirewallPolicyAccept
		return nil
	case "DROP":
		*self = FirewallPolicyDrop
		return nil
	case "":
		*self = FirewallPolicyEmpty
		return nil
	case "RETURN":
		*self = FirewallPolicyReturn
		return nil
	}
	if len(s) == 0 {
		*self = FirewallPolicyEmpty
		return nil
	}
	return errors.New("Unknown FirewallPolicy")
}

type FirmwareUpdateMode string

const FirmwareUpdateModeCheck FirmwareUpdateMode = "check"
const FirmwareUpdateModeOff FirmwareUpdateMode = "off"
const FirmwareUpdateModeOn FirmwareUpdateMode = "on"

func (self FirmwareUpdateMode) GetPtr() *FirmwareUpdateMode { var v = self; return &v }

func (self *FirmwareUpdateMode) String() string {
	switch *self {
	case FirmwareUpdateModeCheck:
		return "check"
	case FirmwareUpdateModeOff:
		return "off"
	case FirmwareUpdateModeOn:
		return "on"
	}
	if len(*self) == 0 {
		return "check"
	}
	panic(errors.New("Invalid value of FirmwareUpdateMode"))
}

func (self *FirmwareUpdateMode) MarshalJSON() ([]byte, error) {
	switch *self {
	case FirmwareUpdateModeCheck:
		return json.Marshal("check")
	case FirmwareUpdateModeOff:
		return json.Marshal("off")
	case FirmwareUpdateModeOn:
		return json.Marshal("on")
	}
	if len(*self) == 0 {
		return json.Marshal("check")
	}
	return nil, errors.New("Invalid value of FirmwareUpdateMode")
}

func (self *FirmwareUpdateMode) GetBSON() (interface{}, error) {
	switch *self {
	case FirmwareUpdateModeCheck:
		return "check", nil
	case FirmwareUpdateModeOff:
		return "off", nil
	case FirmwareUpdateModeOn:
		return "on", nil
	}
	if len(*self) == 0 {
		return "check", nil
	}
	return nil, errors.New("Invalid value of FirmwareUpdateMode")
}

func (self *FirmwareUpdateMode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "check":
		*self = FirmwareUpdateModeCheck
		return nil
	case "off":
		*self = FirmwareUpdateModeOff
		return nil
	case "on":
		*self = FirmwareUpdateModeOn
		return nil
	}
	if len(s) == 0 {
		*self = FirmwareUpdateModeCheck
		return nil
	}
	return errors.New("Unknown FirmwareUpdateMode")
}

func (self *FirmwareUpdateMode) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "check":
		*self = FirmwareUpdateModeCheck
		return nil
	case "off":
		*self = FirmwareUpdateModeOff
		return nil
	case "on":
		*self = FirmwareUpdateModeOn
		return nil
	}
	if len(s) == 0 {
		*self = FirmwareUpdateModeCheck
		return nil
	}
	return errors.New("Unknown FirmwareUpdateMode")
}

type L3Protocol string

const L3ProtocolEmpty L3Protocol = ""
const L3ProtocolIP L3Protocol = "ip"
const L3ProtocolIPv4 L3Protocol = "ipv4"
const L3ProtocolIPv6 L3Protocol = "ipv6"

func (self L3Protocol) GetPtr() *L3Protocol { var v = self; return &v }

func (self *L3Protocol) String() string {
	switch *self {
	case L3ProtocolEmpty:
		return ""
	case L3ProtocolIP:
		return "ip"
	case L3ProtocolIPv4:
		return "ipv4"
	case L3ProtocolIPv6:
		return "ipv6"
	}
	if len(*self) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of L3Protocol"))
}

func (self *L3Protocol) MarshalJSON() ([]byte, error) {
	switch *self {
	case L3ProtocolEmpty:
		return json.Marshal("")
	case L3ProtocolIP:
		return json.Marshal("ip")
	case L3ProtocolIPv4:
		return json.Marshal("ipv4")
	case L3ProtocolIPv6:
		return json.Marshal("ipv6")
	}
	if len(*self) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of L3Protocol")
}

func (self *L3Protocol) GetBSON() (interface{}, error) {
	switch *self {
	case L3ProtocolEmpty:
		return "", nil
	case L3ProtocolIP:
		return "ip", nil
	case L3ProtocolIPv4:
		return "ipv4", nil
	case L3ProtocolIPv6:
		return "ipv6", nil
	}
	if len(*self) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of L3Protocol")
}

func (self *L3Protocol) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "":
		*self = L3ProtocolEmpty
		return nil
	case "ip":
		*self = L3ProtocolIP
		return nil
	case "ipv4":
		*self = L3ProtocolIPv4
		return nil
	case "ipv6":
		*self = L3ProtocolIPv6
		return nil
	}
	if len(s) == 0 {
		*self = L3ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L3Protocol")
}

func (self *L3Protocol) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "":
		*self = L3ProtocolEmpty
		return nil
	case "ip":
		*self = L3ProtocolIP
		return nil
	case "ipv4":
		*self = L3ProtocolIPv4
		return nil
	case "ipv6":
		*self = L3ProtocolIPv6
		return nil
	}
	if len(s) == 0 {
		*self = L3ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L3Protocol")
}

type L4Protocol string

const L4ProtocolEmpty L4Protocol = ""
const L4ProtocolTCP L4Protocol = "TCP"
const L4ProtocolUDP L4Protocol = "UDP"

func (self L4Protocol) GetPtr() *L4Protocol { var v = self; return &v }

func (self *L4Protocol) String() string {
	switch *self {
	case L4ProtocolEmpty:
		return ""
	case L4ProtocolTCP:
		return "TCP"
	case L4ProtocolUDP:
		return "UDP"
	}
	if len(*self) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of L4Protocol"))
}

func (self *L4Protocol) MarshalJSON() ([]byte, error) {
	switch *self {
	case L4ProtocolEmpty:
		return json.Marshal("")
	case L4ProtocolTCP:
		return json.Marshal("TCP")
	case L4ProtocolUDP:
		return json.Marshal("UDP")
	}
	if len(*self) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of L4Protocol")
}

func (self *L4Protocol) GetBSON() (interface{}, error) {
	switch *self {
	case L4ProtocolEmpty:
		return "", nil
	case L4ProtocolTCP:
		return "TCP", nil
	case L4ProtocolUDP:
		return "UDP", nil
	}
	if len(*self) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of L4Protocol")
}

func (self *L4Protocol) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "":
		*self = L4ProtocolEmpty
		return nil
	case "TCP":
		*self = L4ProtocolTCP
		return nil
	case "UDP":
		*self = L4ProtocolUDP
		return nil
	}
	if len(s) == 0 {
		*self = L4ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L4Protocol")
}

func (self *L4Protocol) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "":
		*self = L4ProtocolEmpty
		return nil
	case "TCP":
		*self = L4ProtocolTCP
		return nil
	case "UDP":
		*self = L4ProtocolUDP
		return nil
	}
	if len(s) == 0 {
		*self = L4ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L4Protocol")
}

type MCSRequire string

const MCSRequireHT MCSRequire = "ht"
const MCSRequireOff MCSRequire = "off"
const MCSRequireVHT MCSRequire = "vht"

func (self MCSRequire) GetPtr() *MCSRequire { var v = self; return &v }

func (self *MCSRequire) String() string {
	switch *self {
	case MCSRequireHT:
		return "ht"
	case MCSRequireOff:
		return "off"
	case MCSRequireVHT:
		return "vht"
	}
	if len(*self) == 0 {
		return "off"
	}
	panic(errors.New("Invalid value of MCSRequire"))
}

func (self *MCSRequire) MarshalJSON() ([]byte, error) {
	switch *self {
	case MCSRequireHT:
		return json.Marshal("ht")
	case MCSRequireOff:
		return json.Marshal("off")
	case MCSRequireVHT:
		return json.Marshal("vht")
	}
	if len(*self) == 0 {
		return json.Marshal("off")
	}
	return nil, errors.New("Invalid value of MCSRequire")
}

func (self *MCSRequire) GetBSON() (interface{}, error) {
	switch *self {
	case MCSRequireHT:
		return "ht", nil
	case MCSRequireOff:
		return "off", nil
	case MCSRequireVHT:
		return "vht", nil
	}
	if len(*self) == 0 {
		return "off", nil
	}
	return nil, errors.New("Invalid value of MCSRequire")
}

func (self *MCSRequire) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ht":
		*self = MCSRequireHT
		return nil
	case "off":
		*self = MCSRequireOff
		return nil
	case "vht":
		*self = MCSRequireVHT
		return nil
	}
	if len(s) == 0 {
		*self = MCSRequireOff
		return nil
	}
	return errors.New("Unknown MCSRequire")
}

func (self *MCSRequire) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ht":
		*self = MCSRequireHT
		return nil
	case "off":
		*self = MCSRequireOff
		return nil
	case "vht":
		*self = MCSRequireVHT
		return nil
	}
	if len(s) == 0 {
		*self = MCSRequireOff
		return nil
	}
	return errors.New("Unknown MCSRequire")
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
	if len(*self) == 0 {
		return "None"
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
	if len(*self) == 0 {
		return json.Marshal("None")
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
	if len(*self) == 0 {
		return "None", nil
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
	if len(s) == 0 {
		*self = MacFilterTypeNone
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
	if len(s) == 0 {
		*self = MacFilterTypeNone
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
const ModuleFW Module = "FW"
const ModuleLBS Module = "LBS"
const ModuleMQTTLog Module = "MQTT_LOG"
const ModuleMonitor Module = "MONITOR"
const ModulePortalBack Module = "PORTAL_BACKEND"
const ModuleRRM Module = "RRM"
const ModuleRadiusGw Module = "RADIUS_GATEWAY"
const ModuleRedirect Module = "REDIRECT"
const ModuleStat Module = "STAT"
const ModuleTunManager Module = "TUN_MANAGER"

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
	case ModuleFW:
		return "FW"
	case ModuleLBS:
		return "LBS"
	case ModuleMQTTLog:
		return "MQTT_LOG"
	case ModuleMonitor:
		return "MONITOR"
	case ModulePortalBack:
		return "PORTAL_BACKEND"
	case ModuleRRM:
		return "RRM"
	case ModuleRadiusGw:
		return "RADIUS_GATEWAY"
	case ModuleRedirect:
		return "REDIRECT"
	case ModuleStat:
		return "STAT"
	case ModuleTunManager:
		return "TUN_MANAGER"
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
	case ModuleFW:
		return json.Marshal("FW")
	case ModuleLBS:
		return json.Marshal("LBS")
	case ModuleMQTTLog:
		return json.Marshal("MQTT_LOG")
	case ModuleMonitor:
		return json.Marshal("MONITOR")
	case ModulePortalBack:
		return json.Marshal("PORTAL_BACKEND")
	case ModuleRRM:
		return json.Marshal("RRM")
	case ModuleRadiusGw:
		return json.Marshal("RADIUS_GATEWAY")
	case ModuleRedirect:
		return json.Marshal("REDIRECT")
	case ModuleStat:
		return json.Marshal("STAT")
	case ModuleTunManager:
		return json.Marshal("TUN_MANAGER")
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
	case ModuleFW:
		return "FW", nil
	case ModuleLBS:
		return "LBS", nil
	case ModuleMQTTLog:
		return "MQTT_LOG", nil
	case ModuleMonitor:
		return "MONITOR", nil
	case ModulePortalBack:
		return "PORTAL_BACKEND", nil
	case ModuleRRM:
		return "RRM", nil
	case ModuleRadiusGw:
		return "RADIUS_GATEWAY", nil
	case ModuleRedirect:
		return "REDIRECT", nil
	case ModuleStat:
		return "STAT", nil
	case ModuleTunManager:
		return "TUN_MANAGER", nil
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
	case "FW":
		*self = ModuleFW
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
	case "PORTAL_BACKEND":
		*self = ModulePortalBack
		return nil
	case "RRM":
		*self = ModuleRRM
		return nil
	case "RADIUS_GATEWAY":
		*self = ModuleRadiusGw
		return nil
	case "REDIRECT":
		*self = ModuleRedirect
		return nil
	case "STAT":
		*self = ModuleStat
		return nil
	case "TUN_MANAGER":
		*self = ModuleTunManager
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
	case "FW":
		*self = ModuleFW
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
	case "PORTAL_BACKEND":
		*self = ModulePortalBack
		return nil
	case "RRM":
		*self = ModuleRRM
		return nil
	case "RADIUS_GATEWAY":
		*self = ModuleRadiusGw
		return nil
	case "REDIRECT":
		*self = ModuleRedirect
		return nil
	case "STAT":
		*self = ModuleStat
		return nil
	case "TUN_MANAGER":
		*self = ModuleTunManager
		return nil
	}
	return errors.New("Unknown Module")
}

type Operation string

const OperationAny Operation = "+"
const OperationCPEStatus Operation = "STATUS"
const OperationCreate Operation = "C"
const OperationDelete Operation = "D"
const OperationJSONRPC Operation = "JSONRPC"
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
	case OperationJSONRPC:
		return "JSONRPC"
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
	case OperationJSONRPC:
		return json.Marshal("JSONRPC")
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
	case OperationJSONRPC:
		return "JSONRPC", nil
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
	case "JSONRPC":
		*self = OperationJSONRPC
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
	case "JSONRPC":
		*self = OperationJSONRPC
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

type PortalAuthType string

const PortalAuthTypeExternal PortalAuthType = "external"
const PortalAuthTypeNone PortalAuthType = ""
const PortalAuthTypeOAuth2 PortalAuthType = "oauth2"
const PortalAuthTypeRADIUS PortalAuthType = "radius"
const PortalAuthTypeSMS PortalAuthType = "sms"

func (self PortalAuthType) GetPtr() *PortalAuthType { var v = self; return &v }

func (self *PortalAuthType) String() string {
	switch *self {
	case PortalAuthTypeExternal:
		return "external"
	case PortalAuthTypeNone:
		return ""
	case PortalAuthTypeOAuth2:
		return "oauth2"
	case PortalAuthTypeRADIUS:
		return "radius"
	case PortalAuthTypeSMS:
		return "sms"
	}
	if len(*self) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of PortalAuthType"))
}

func (self *PortalAuthType) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalAuthTypeExternal:
		return json.Marshal("external")
	case PortalAuthTypeNone:
		return json.Marshal("")
	case PortalAuthTypeOAuth2:
		return json.Marshal("oauth2")
	case PortalAuthTypeRADIUS:
		return json.Marshal("radius")
	case PortalAuthTypeSMS:
		return json.Marshal("sms")
	}
	if len(*self) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of PortalAuthType")
}

func (self *PortalAuthType) GetBSON() (interface{}, error) {
	switch *self {
	case PortalAuthTypeExternal:
		return "external", nil
	case PortalAuthTypeNone:
		return "", nil
	case PortalAuthTypeOAuth2:
		return "oauth2", nil
	case PortalAuthTypeRADIUS:
		return "radius", nil
	case PortalAuthTypeSMS:
		return "sms", nil
	}
	if len(*self) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of PortalAuthType")
}

func (self *PortalAuthType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "external":
		*self = PortalAuthTypeExternal
		return nil
	case "":
		*self = PortalAuthTypeNone
		return nil
	case "oauth2":
		*self = PortalAuthTypeOAuth2
		return nil
	case "radius":
		*self = PortalAuthTypeRADIUS
		return nil
	case "sms":
		*self = PortalAuthTypeSMS
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthType")
}

func (self *PortalAuthType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "external":
		*self = PortalAuthTypeExternal
		return nil
	case "":
		*self = PortalAuthTypeNone
		return nil
	case "oauth2":
		*self = PortalAuthTypeOAuth2
		return nil
	case "radius":
		*self = PortalAuthTypeRADIUS
		return nil
	case "sms":
		*self = PortalAuthTypeSMS
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthType")
}

type RadiusMessageType string

const RadiusMessageTypeAccessAccept RadiusMessageType = "access-accept"
const RadiusMessageTypeAccessReject RadiusMessageType = "access-reject"
const RadiusMessageTypeAccessRequest RadiusMessageType = "access-request"
const RadiusMessageTypeAccountingRequest RadiusMessageType = "accounting"

func (self RadiusMessageType) GetPtr() *RadiusMessageType { var v = self; return &v }

func (self *RadiusMessageType) String() string {
	switch *self {
	case RadiusMessageTypeAccessAccept:
		return "access-accept"
	case RadiusMessageTypeAccessReject:
		return "access-reject"
	case RadiusMessageTypeAccessRequest:
		return "access-request"
	case RadiusMessageTypeAccountingRequest:
		return "accounting"
	}
	if len(*self) == 0 {
		return "accounting"
	}
	panic(errors.New("Invalid value of RadiusMessageType"))
}

func (self *RadiusMessageType) MarshalJSON() ([]byte, error) {
	switch *self {
	case RadiusMessageTypeAccessAccept:
		return json.Marshal("access-accept")
	case RadiusMessageTypeAccessReject:
		return json.Marshal("access-reject")
	case RadiusMessageTypeAccessRequest:
		return json.Marshal("access-request")
	case RadiusMessageTypeAccountingRequest:
		return json.Marshal("accounting")
	}
	if len(*self) == 0 {
		return json.Marshal("accounting")
	}
	return nil, errors.New("Invalid value of RadiusMessageType")
}

func (self *RadiusMessageType) GetBSON() (interface{}, error) {
	switch *self {
	case RadiusMessageTypeAccessAccept:
		return "access-accept", nil
	case RadiusMessageTypeAccessReject:
		return "access-reject", nil
	case RadiusMessageTypeAccessRequest:
		return "access-request", nil
	case RadiusMessageTypeAccountingRequest:
		return "accounting", nil
	}
	if len(*self) == 0 {
		return "accounting", nil
	}
	return nil, errors.New("Invalid value of RadiusMessageType")
}

func (self *RadiusMessageType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "access-accept":
		*self = RadiusMessageTypeAccessAccept
		return nil
	case "access-reject":
		*self = RadiusMessageTypeAccessReject
		return nil
	case "access-request":
		*self = RadiusMessageTypeAccessRequest
		return nil
	case "accounting":
		*self = RadiusMessageTypeAccountingRequest
		return nil
	}
	if len(s) == 0 {
		*self = RadiusMessageTypeAccountingRequest
		return nil
	}
	return errors.New("Unknown RadiusMessageType")
}

func (self *RadiusMessageType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "access-accept":
		*self = RadiusMessageTypeAccessAccept
		return nil
	case "access-reject":
		*self = RadiusMessageTypeAccessReject
		return nil
	case "access-request":
		*self = RadiusMessageTypeAccessRequest
		return nil
	case "accounting":
		*self = RadiusMessageTypeAccountingRequest
		return nil
	}
	if len(s) == 0 {
		*self = RadiusMessageTypeAccountingRequest
		return nil
	}
	return errors.New("Unknown RadiusMessageType")
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

const SecurityTypeNone SecurityType = "open"
const SecurityTypeWPA2Enterprise SecurityType = "wpa2enterprise"
const SecurityTypeWPA2Personal SecurityType = "wpa2personal"
const SecurityTypeWPAEnterprise SecurityType = "wpaenterprise"
const SecurityTypeWPAPersonal SecurityType = "wpapersonal"

func (self SecurityType) GetPtr() *SecurityType { var v = self; return &v }

func (self *SecurityType) String() string {
	switch *self {
	case SecurityTypeNone:
		return "open"
	case SecurityTypeWPA2Enterprise:
		return "wpa2enterprise"
	case SecurityTypeWPA2Personal:
		return "wpa2personal"
	case SecurityTypeWPAEnterprise:
		return "wpaenterprise"
	case SecurityTypeWPAPersonal:
		return "wpapersonal"
	}
	if len(*self) == 0 {
		return "open"
	}
	panic(errors.New("Invalid value of SecurityType"))
}

func (self *SecurityType) MarshalJSON() ([]byte, error) {
	switch *self {
	case SecurityTypeNone:
		return json.Marshal("open")
	case SecurityTypeWPA2Enterprise:
		return json.Marshal("wpa2enterprise")
	case SecurityTypeWPA2Personal:
		return json.Marshal("wpa2personal")
	case SecurityTypeWPAEnterprise:
		return json.Marshal("wpaenterprise")
	case SecurityTypeWPAPersonal:
		return json.Marshal("wpapersonal")
	}
	if len(*self) == 0 {
		return json.Marshal("open")
	}
	return nil, errors.New("Invalid value of SecurityType")
}

func (self *SecurityType) GetBSON() (interface{}, error) {
	switch *self {
	case SecurityTypeNone:
		return "open", nil
	case SecurityTypeWPA2Enterprise:
		return "wpa2enterprise", nil
	case SecurityTypeWPA2Personal:
		return "wpa2personal", nil
	case SecurityTypeWPAEnterprise:
		return "wpaenterprise", nil
	case SecurityTypeWPAPersonal:
		return "wpapersonal", nil
	}
	if len(*self) == 0 {
		return "open", nil
	}
	return nil, errors.New("Invalid value of SecurityType")
}

func (self *SecurityType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "open":
		*self = SecurityTypeNone
		return nil
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
	if len(s) == 0 {
		*self = SecurityTypeNone
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
	case "open":
		*self = SecurityTypeNone
		return nil
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
	if len(s) == 0 {
		*self = SecurityTypeNone
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
	if len(*self) == 0 {
		return "DEBUG"
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
	if len(*self) == 0 {
		return json.Marshal("DEBUG")
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
	if len(*self) == 0 {
		return "DEBUG", nil
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
	if len(s) == 0 {
		*self = SystemEventLevelDEBUG
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
	if len(s) == 0 {
		*self = SystemEventLevelDEBUG
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
const SystemEventTypeCPEInterfaceState SystemEventType = "CPE_INTERFACE_STATE"
const SystemEventTypeClientConnected SystemEventType = "CLIENT_CONNECTED"
const SystemEventTypeClientDisconnected SystemEventType = "CLIENT_DISCONNECTED"
const SystemEventTypeCpeFirmwareAvailable SystemEventType = "CPE_FIRMWARE_AVAILABLE"
const SystemEventTypeDaemonSettingsChanged SystemEventType = "DAEMON_SETTINGS_CHANGE"
const SystemEventTypeFirmwareUploaded SystemEventType = "FIRMWARE_UPLOADED"
const SystemEventTypeMonitorRuleViolation SystemEventType = "MONITOR_RULE_VIOLATION"
const SystemEventTypeRRMStatus SystemEventType = "RRM_STATUS_DATA"
const SystemEventTypeRadiusAccountingSend SystemEventType = "RADIUS_ACCOUNTING_SEND"
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
	case SystemEventTypeCPEInterfaceState:
		return "CPE_INTERFACE_STATE"
	case SystemEventTypeClientConnected:
		return "CLIENT_CONNECTED"
	case SystemEventTypeClientDisconnected:
		return "CLIENT_DISCONNECTED"
	case SystemEventTypeCpeFirmwareAvailable:
		return "CPE_FIRMWARE_AVAILABLE"
	case SystemEventTypeDaemonSettingsChanged:
		return "DAEMON_SETTINGS_CHANGE"
	case SystemEventTypeFirmwareUploaded:
		return "FIRMWARE_UPLOADED"
	case SystemEventTypeMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION"
	case SystemEventTypeRRMStatus:
		return "RRM_STATUS_DATA"
	case SystemEventTypeRadiusAccountingSend:
		return "RADIUS_ACCOUNTING_SEND"
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
	case SystemEventTypeCPEInterfaceState:
		return json.Marshal("CPE_INTERFACE_STATE")
	case SystemEventTypeClientConnected:
		return json.Marshal("CLIENT_CONNECTED")
	case SystemEventTypeClientDisconnected:
		return json.Marshal("CLIENT_DISCONNECTED")
	case SystemEventTypeCpeFirmwareAvailable:
		return json.Marshal("CPE_FIRMWARE_AVAILABLE")
	case SystemEventTypeDaemonSettingsChanged:
		return json.Marshal("DAEMON_SETTINGS_CHANGE")
	case SystemEventTypeFirmwareUploaded:
		return json.Marshal("FIRMWARE_UPLOADED")
	case SystemEventTypeMonitorRuleViolation:
		return json.Marshal("MONITOR_RULE_VIOLATION")
	case SystemEventTypeRRMStatus:
		return json.Marshal("RRM_STATUS_DATA")
	case SystemEventTypeRadiusAccountingSend:
		return json.Marshal("RADIUS_ACCOUNTING_SEND")
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
	case SystemEventTypeCPEInterfaceState:
		return "CPE_INTERFACE_STATE", nil
	case SystemEventTypeClientConnected:
		return "CLIENT_CONNECTED", nil
	case SystemEventTypeClientDisconnected:
		return "CLIENT_DISCONNECTED", nil
	case SystemEventTypeCpeFirmwareAvailable:
		return "CPE_FIRMWARE_AVAILABLE", nil
	case SystemEventTypeDaemonSettingsChanged:
		return "DAEMON_SETTINGS_CHANGE", nil
	case SystemEventTypeFirmwareUploaded:
		return "FIRMWARE_UPLOADED", nil
	case SystemEventTypeMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION", nil
	case SystemEventTypeRRMStatus:
		return "RRM_STATUS_DATA", nil
	case SystemEventTypeRadiusAccountingSend:
		return "RADIUS_ACCOUNTING_SEND", nil
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
	case "CPE_INTERFACE_STATE":
		*self = SystemEventTypeCPEInterfaceState
		return nil
	case "CLIENT_CONNECTED":
		*self = SystemEventTypeClientConnected
		return nil
	case "CLIENT_DISCONNECTED":
		*self = SystemEventTypeClientDisconnected
		return nil
	case "CPE_FIRMWARE_AVAILABLE":
		*self = SystemEventTypeCpeFirmwareAvailable
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		*self = SystemEventTypeDaemonSettingsChanged
		return nil
	case "FIRMWARE_UPLOADED":
		*self = SystemEventTypeFirmwareUploaded
		return nil
	case "MONITOR_RULE_VIOLATION":
		*self = SystemEventTypeMonitorRuleViolation
		return nil
	case "RRM_STATUS_DATA":
		*self = SystemEventTypeRRMStatus
		return nil
	case "RADIUS_ACCOUNTING_SEND":
		*self = SystemEventTypeRadiusAccountingSend
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
	case "CPE_INTERFACE_STATE":
		*self = SystemEventTypeCPEInterfaceState
		return nil
	case "CLIENT_CONNECTED":
		*self = SystemEventTypeClientConnected
		return nil
	case "CLIENT_DISCONNECTED":
		*self = SystemEventTypeClientDisconnected
		return nil
	case "CPE_FIRMWARE_AVAILABLE":
		*self = SystemEventTypeCpeFirmwareAvailable
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		*self = SystemEventTypeDaemonSettingsChanged
		return nil
	case "FIRMWARE_UPLOADED":
		*self = SystemEventTypeFirmwareUploaded
		return nil
	case "MONITOR_RULE_VIOLATION":
		*self = SystemEventTypeMonitorRuleViolation
		return nil
	case "RRM_STATUS_DATA":
		*self = SystemEventTypeRRMStatus
		return nil
	case "RADIUS_ACCOUNTING_SEND":
		*self = SystemEventTypeRadiusAccountingSend
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

type TunManagerRPC string

const TunManagerRPCCreateL2TunnelSession TunManagerRPC = "CreateL2TunnelSession"
const TunManagerRPCDeleteL2TunnelSession TunManagerRPC = "DeleteL2TunnelSession"

func (self TunManagerRPC) GetPtr() *TunManagerRPC { var v = self; return &v }

func (self *TunManagerRPC) String() string {
	switch *self {
	case TunManagerRPCCreateL2TunnelSession:
		return "CreateL2TunnelSession"
	case TunManagerRPCDeleteL2TunnelSession:
		return "DeleteL2TunnelSession"
	}
	panic(errors.New("Invalid value of TunManagerRPC"))
}

func (self *TunManagerRPC) MarshalJSON() ([]byte, error) {
	switch *self {
	case TunManagerRPCCreateL2TunnelSession:
		return json.Marshal("CreateL2TunnelSession")
	case TunManagerRPCDeleteL2TunnelSession:
		return json.Marshal("DeleteL2TunnelSession")
	}
	return nil, errors.New("Invalid value of TunManagerRPC")
}

func (self *TunManagerRPC) GetBSON() (interface{}, error) {
	switch *self {
	case TunManagerRPCCreateL2TunnelSession:
		return "CreateL2TunnelSession", nil
	case TunManagerRPCDeleteL2TunnelSession:
		return "DeleteL2TunnelSession", nil
	}
	return nil, errors.New("Invalid value of TunManagerRPC")
}

func (self *TunManagerRPC) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CreateL2TunnelSession":
		*self = TunManagerRPCCreateL2TunnelSession
		return nil
	case "DeleteL2TunnelSession":
		*self = TunManagerRPCDeleteL2TunnelSession
		return nil
	}
	return errors.New("Unknown TunManagerRPC")
}

func (self *TunManagerRPC) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "CreateL2TunnelSession":
		*self = TunManagerRPCCreateL2TunnelSession
		return nil
	case "DeleteL2TunnelSession":
		*self = TunManagerRPCDeleteL2TunnelSession
		return nil
	}
	return errors.New("Unknown TunManagerRPC")
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
	if len(*self) == 0 {
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
	if len(*self) == 0 {
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
	if len(*self) == 0 {
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
	if len(s) == 0 {
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
	if len(s) == 0 {
		*self = WirelessClientTypeOther
		return nil
	}
	return errors.New("Unknown WirelessClientType")
}
