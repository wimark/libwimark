package libwimark

import (
	"encoding/json"
	"errors"

	"github.com/globalsign/mgo/bson"
)

type CPEAgentStatusType string

const CPEAgentStatusTypeException CPEAgentStatusType = "exception"
const CPEAgentStatusTypeSuccess CPEAgentStatusType = "success"
const CPEAgentStatusTypeSyntaxError CPEAgentStatusType = "syntax"
const CPEAgentStatusTypeUndefined CPEAgentStatusType = "undefined"

func (self CPEAgentStatusType) GetPtr() *CPEAgentStatusType { var v = self; return &v }

func (self CPEAgentStatusType) String() string {
	switch self {
	case CPEAgentStatusTypeException:
		return "exception"
	case CPEAgentStatusTypeSuccess:
		return "success"
	case CPEAgentStatusTypeSyntaxError:
		return "syntax"
	case CPEAgentStatusTypeUndefined:
		return "undefined"
	}
	if len(self) == 0 {
		return "undefined"
	}
	panic(errors.New("Invalid value of CPEAgentStatusType: " + string(self)))
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
	return nil, errors.New("Invalid value of CPEAgentStatusType: " + string(*self))
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
	return nil, errors.New("Invalid value of CPEAgentStatusType: " + string(*self))
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
	return errors.New("Unknown CPEAgentStatusType: " + s)
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
	return errors.New("Unknown CPEAgentStatusType: " + s)
}

type CPEInterfaceState string

const CPEInterfaceStateACS CPEInterfaceState = "ACS"
const CPEInterfaceStateCountryUpdate CPEInterfaceState = "COUNTRY_UPDATE"
const CPEInterfaceStateDFS CPEInterfaceState = "DFS"
const CPEInterfaceStateDisabled CPEInterfaceState = "DISABLED"
const CPEInterfaceStateEnabled CPEInterfaceState = "ENABLED"
const CPEInterfaceStateHtScan CPEInterfaceState = "HT_SCAN"
const CPEInterfaceStateTerminated CPEInterfaceState = "TERMINATED"
const CPEInterfaceStateUninitialized CPEInterfaceState = "UNINITIALIZED"
const CPEInterfaceStateUnknown CPEInterfaceState = "UNKNOWN"

func (self CPEInterfaceState) GetPtr() *CPEInterfaceState { var v = self; return &v }

func (self CPEInterfaceState) String() string {
	switch self {
	case CPEInterfaceStateACS:
		return "ACS"
	case CPEInterfaceStateCountryUpdate:
		return "COUNTRY_UPDATE"
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
	panic(errors.New("Invalid value of CPEInterfaceState: " + string(self)))
}

func (self *CPEInterfaceState) MarshalJSON() ([]byte, error) {
	switch *self {
	case CPEInterfaceStateACS:
		return json.Marshal("ACS")
	case CPEInterfaceStateCountryUpdate:
		return json.Marshal("COUNTRY_UPDATE")
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
	return nil, errors.New("Invalid value of CPEInterfaceState: " + string(*self))
}

func (self *CPEInterfaceState) GetBSON() (interface{}, error) {
	switch *self {
	case CPEInterfaceStateACS:
		return "ACS", nil
	case CPEInterfaceStateCountryUpdate:
		return "COUNTRY_UPDATE", nil
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
	return nil, errors.New("Invalid value of CPEInterfaceState: " + string(*self))
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
	case "COUNTRY_UPDATE":
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
	return errors.New("Unknown CPEInterfaceState: " + s)
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
	case "COUNTRY_UPDATE":
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
	return errors.New("Unknown CPEInterfaceState: " + s)
}

type ClientStatPacketType string

const ClientStatPacketTypeInterim ClientStatPacketType = "interim"
const ClientStatPacketTypeInterimOld ClientStatPacketType = "Interim-Update"
const ClientStatPacketTypeOffOld ClientStatPacketType = "Accounting-Off"
const ClientStatPacketTypeOnOld ClientStatPacketType = "Accounting-On"
const ClientStatPacketTypeStart ClientStatPacketType = "start"
const ClientStatPacketTypeStartOld ClientStatPacketType = "Start"
const ClientStatPacketTypeStop ClientStatPacketType = "stop"
const ClientStatPacketTypeStopOld ClientStatPacketType = "Stop"

func (self ClientStatPacketType) GetPtr() *ClientStatPacketType { var v = self; return &v }

func (self ClientStatPacketType) String() string {
	switch self {
	case ClientStatPacketTypeInterim:
		return "interim"
	case ClientStatPacketTypeInterimOld:
		return "Interim-Update"
	case ClientStatPacketTypeOffOld:
		return "Accounting-Off"
	case ClientStatPacketTypeOnOld:
		return "Accounting-On"
	case ClientStatPacketTypeStart:
		return "start"
	case ClientStatPacketTypeStartOld:
		return "Start"
	case ClientStatPacketTypeStop:
		return "stop"
	case ClientStatPacketTypeStopOld:
		return "Stop"
	}
	panic(errors.New("Invalid value of ClientStatPacketType: " + string(self)))
}

func (self *ClientStatPacketType) MarshalJSON() ([]byte, error) {
	switch *self {
	case ClientStatPacketTypeInterim:
		return json.Marshal("interim")
	case ClientStatPacketTypeInterimOld:
		return json.Marshal("Interim-Update")
	case ClientStatPacketTypeOffOld:
		return json.Marshal("Accounting-Off")
	case ClientStatPacketTypeOnOld:
		return json.Marshal("Accounting-On")
	case ClientStatPacketTypeStart:
		return json.Marshal("start")
	case ClientStatPacketTypeStartOld:
		return json.Marshal("Start")
	case ClientStatPacketTypeStop:
		return json.Marshal("stop")
	case ClientStatPacketTypeStopOld:
		return json.Marshal("Stop")
	}
	return nil, errors.New("Invalid value of ClientStatPacketType: " + string(*self))
}

func (self *ClientStatPacketType) GetBSON() (interface{}, error) {
	switch *self {
	case ClientStatPacketTypeInterim:
		return "interim", nil
	case ClientStatPacketTypeInterimOld:
		return "Interim-Update", nil
	case ClientStatPacketTypeOffOld:
		return "Accounting-Off", nil
	case ClientStatPacketTypeOnOld:
		return "Accounting-On", nil
	case ClientStatPacketTypeStart:
		return "start", nil
	case ClientStatPacketTypeStartOld:
		return "Start", nil
	case ClientStatPacketTypeStop:
		return "stop", nil
	case ClientStatPacketTypeStopOld:
		return "Stop", nil
	}
	return nil, errors.New("Invalid value of ClientStatPacketType: " + string(*self))
}

func (self *ClientStatPacketType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "interim":
		*self = ClientStatPacketTypeInterim
		return nil
	case "Interim-Update":
		*self = ClientStatPacketTypeInterimOld
		return nil
	case "Accounting-Off":
		*self = ClientStatPacketTypeOffOld
		return nil
	case "Accounting-On":
		*self = ClientStatPacketTypeOnOld
		return nil
	case "start":
		*self = ClientStatPacketTypeStart
		return nil
	case "Start":
		*self = ClientStatPacketTypeStartOld
		return nil
	case "stop":
		*self = ClientStatPacketTypeStop
		return nil
	case "Stop":
		*self = ClientStatPacketTypeStopOld
		return nil
	}
	return errors.New("Unknown ClientStatPacketType: " + s)
}

func (self *ClientStatPacketType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "interim":
		*self = ClientStatPacketTypeInterim
		return nil
	case "Interim-Update":
		*self = ClientStatPacketTypeInterimOld
		return nil
	case "Accounting-Off":
		*self = ClientStatPacketTypeOffOld
		return nil
	case "Accounting-On":
		*self = ClientStatPacketTypeOnOld
		return nil
	case "start":
		*self = ClientStatPacketTypeStart
		return nil
	case "Start":
		*self = ClientStatPacketTypeStartOld
		return nil
	case "stop":
		*self = ClientStatPacketTypeStop
		return nil
	case "Stop":
		*self = ClientStatPacketTypeStopOld
		return nil
	}
	return errors.New("Unknown ClientStatPacketType: " + s)
}

type ConfigurationStatus string

const ConfigurationStatusDontUse1 ConfigurationStatus = "pending"
const ConfigurationStatusDontUse2 ConfigurationStatus = "error"
const ConfigurationStatusEmpty ConfigurationStatus = "empty"
const ConfigurationStatusOK ConfigurationStatus = "ok"
const ConfigurationStatusOffline ConfigurationStatus = "offline"
const ConfigurationStatusRebooting ConfigurationStatus = "rebooting"
const ConfigurationStatusUpdating ConfigurationStatus = "updating"
const ConfigurationStatusUpgrading ConfigurationStatus = "upgrading"

func (self ConfigurationStatus) GetPtr() *ConfigurationStatus { var v = self; return &v }

func (self ConfigurationStatus) String() string {
	switch self {
	case ConfigurationStatusDontUse1:
		return "pending"
	case ConfigurationStatusDontUse2:
		return "error"
	case ConfigurationStatusEmpty:
		return "empty"
	case ConfigurationStatusOK:
		return "ok"
	case ConfigurationStatusOffline:
		return "offline"
	case ConfigurationStatusRebooting:
		return "rebooting"
	case ConfigurationStatusUpdating:
		return "updating"
	case ConfigurationStatusUpgrading:
		return "upgrading"
	}
	if len(self) == 0 {
		return "empty"
	}
	panic(errors.New("Invalid value of ConfigurationStatus: " + string(self)))
}

func (self *ConfigurationStatus) MarshalJSON() ([]byte, error) {
	switch *self {
	case ConfigurationStatusDontUse1:
		return json.Marshal("pending")
	case ConfigurationStatusDontUse2:
		return json.Marshal("error")
	case ConfigurationStatusEmpty:
		return json.Marshal("empty")
	case ConfigurationStatusOK:
		return json.Marshal("ok")
	case ConfigurationStatusOffline:
		return json.Marshal("offline")
	case ConfigurationStatusRebooting:
		return json.Marshal("rebooting")
	case ConfigurationStatusUpdating:
		return json.Marshal("updating")
	case ConfigurationStatusUpgrading:
		return json.Marshal("upgrading")
	}
	if len(*self) == 0 {
		return json.Marshal("empty")
	}
	return nil, errors.New("Invalid value of ConfigurationStatus: " + string(*self))
}

func (self *ConfigurationStatus) GetBSON() (interface{}, error) {
	switch *self {
	case ConfigurationStatusDontUse1:
		return "pending", nil
	case ConfigurationStatusDontUse2:
		return "error", nil
	case ConfigurationStatusEmpty:
		return "empty", nil
	case ConfigurationStatusOK:
		return "ok", nil
	case ConfigurationStatusOffline:
		return "offline", nil
	case ConfigurationStatusRebooting:
		return "rebooting", nil
	case ConfigurationStatusUpdating:
		return "updating", nil
	case ConfigurationStatusUpgrading:
		return "upgrading", nil
	}
	if len(*self) == 0 {
		return "empty", nil
	}
	return nil, errors.New("Invalid value of ConfigurationStatus: " + string(*self))
}

func (self *ConfigurationStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "pending":
		*self = ConfigurationStatusDontUse1
		return nil
	case "error":
		*self = ConfigurationStatusDontUse2
		return nil
	case "empty":
		*self = ConfigurationStatusEmpty
		return nil
	case "ok":
		*self = ConfigurationStatusOK
		return nil
	case "offline":
		*self = ConfigurationStatusOffline
		return nil
	case "rebooting":
		*self = ConfigurationStatusRebooting
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
	return errors.New("Unknown ConfigurationStatus: " + s)
}

func (self *ConfigurationStatus) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "pending":
		*self = ConfigurationStatusDontUse1
		return nil
	case "error":
		*self = ConfigurationStatusDontUse2
		return nil
	case "empty":
		*self = ConfigurationStatusEmpty
		return nil
	case "ok":
		*self = ConfigurationStatusOK
		return nil
	case "offline":
		*self = ConfigurationStatusOffline
		return nil
	case "rebooting":
		*self = ConfigurationStatusRebooting
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
	return errors.New("Unknown ConfigurationStatus: " + s)
}

type ConnectionModeType string

const ConnectionModeTypeModeAC ConnectionModeType = "ac"
const ConnectionModeTypeModeLegacy ConnectionModeType = "legacy"
const ConnectionModeTypeModeN ConnectionModeType = "n"

func (self ConnectionModeType) GetPtr() *ConnectionModeType { var v = self; return &v }

func (self ConnectionModeType) String() string {
	switch self {
	case ConnectionModeTypeModeAC:
		return "ac"
	case ConnectionModeTypeModeLegacy:
		return "legacy"
	case ConnectionModeTypeModeN:
		return "n"
	}
	if len(self) == 0 {
		return "legacy"
	}
	panic(errors.New("Invalid value of ConnectionModeType: " + string(self)))
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
	return nil, errors.New("Invalid value of ConnectionModeType: " + string(*self))
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
	return nil, errors.New("Invalid value of ConnectionModeType: " + string(*self))
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
	return errors.New("Unknown ConnectionModeType: " + s)
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
	return errors.New("Unknown ConnectionModeType: " + s)
}

type ControllerStatusType string

const ControllerStatusTypeConnected ControllerStatusType = "connected"
const ControllerStatusTypeDisconnected ControllerStatusType = "disconnected"
const ControllerStatusTypeEmpty ControllerStatusType = "empty"
const ControllerStatusTypeError ControllerStatusType = "error"
const ControllerStatusTypeProvisioning ControllerStatusType = "provision"
const ControllerStatusTypeUpdating ControllerStatusType = "updating"

func (self ControllerStatusType) GetPtr() *ControllerStatusType { var v = self; return &v }

func (self ControllerStatusType) String() string {
	switch self {
	case ControllerStatusTypeConnected:
		return "connected"
	case ControllerStatusTypeDisconnected:
		return "disconnected"
	case ControllerStatusTypeEmpty:
		return "empty"
	case ControllerStatusTypeError:
		return "error"
	case ControllerStatusTypeProvisioning:
		return "provision"
	case ControllerStatusTypeUpdating:
		return "updating"
	}
	if len(self) == 0 {
		return "empty"
	}
	panic(errors.New("Invalid value of ControllerStatusType: " + string(self)))
}

func (self *ControllerStatusType) MarshalJSON() ([]byte, error) {
	switch *self {
	case ControllerStatusTypeConnected:
		return json.Marshal("connected")
	case ControllerStatusTypeDisconnected:
		return json.Marshal("disconnected")
	case ControllerStatusTypeEmpty:
		return json.Marshal("empty")
	case ControllerStatusTypeError:
		return json.Marshal("error")
	case ControllerStatusTypeProvisioning:
		return json.Marshal("provision")
	case ControllerStatusTypeUpdating:
		return json.Marshal("updating")
	}
	if len(*self) == 0 {
		return json.Marshal("empty")
	}
	return nil, errors.New("Invalid value of ControllerStatusType: " + string(*self))
}

func (self *ControllerStatusType) GetBSON() (interface{}, error) {
	switch *self {
	case ControllerStatusTypeConnected:
		return "connected", nil
	case ControllerStatusTypeDisconnected:
		return "disconnected", nil
	case ControllerStatusTypeEmpty:
		return "empty", nil
	case ControllerStatusTypeError:
		return "error", nil
	case ControllerStatusTypeProvisioning:
		return "provision", nil
	case ControllerStatusTypeUpdating:
		return "updating", nil
	}
	if len(*self) == 0 {
		return "empty", nil
	}
	return nil, errors.New("Invalid value of ControllerStatusType: " + string(*self))
}

func (self *ControllerStatusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "connected":
		*self = ControllerStatusTypeConnected
		return nil
	case "disconnected":
		*self = ControllerStatusTypeDisconnected
		return nil
	case "empty":
		*self = ControllerStatusTypeEmpty
		return nil
	case "error":
		*self = ControllerStatusTypeError
		return nil
	case "provision":
		*self = ControllerStatusTypeProvisioning
		return nil
	case "updating":
		*self = ControllerStatusTypeUpdating
		return nil
	}
	if len(s) == 0 {
		*self = ControllerStatusTypeEmpty
		return nil
	}
	return errors.New("Unknown ControllerStatusType: " + s)
}

func (self *ControllerStatusType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "connected":
		*self = ControllerStatusTypeConnected
		return nil
	case "disconnected":
		*self = ControllerStatusTypeDisconnected
		return nil
	case "empty":
		*self = ControllerStatusTypeEmpty
		return nil
	case "error":
		*self = ControllerStatusTypeError
		return nil
	case "provision":
		*self = ControllerStatusTypeProvisioning
		return nil
	case "updating":
		*self = ControllerStatusTypeUpdating
		return nil
	}
	if len(s) == 0 {
		*self = ControllerStatusTypeEmpty
		return nil
	}
	return errors.New("Unknown ControllerStatusType: " + s)
}

type FirewallDirection string

const FirewallDirectionAny FirewallDirection = "ANY"
const FirewallDirectionIn FirewallDirection = "IN"
const FirewallDirectionOut FirewallDirection = "OUT"

func (self FirewallDirection) GetPtr() *FirewallDirection { var v = self; return &v }

func (self FirewallDirection) String() string {
	switch self {
	case FirewallDirectionAny:
		return "ANY"
	case FirewallDirectionIn:
		return "IN"
	case FirewallDirectionOut:
		return "OUT"
	}
	if len(self) == 0 {
		return "ANY"
	}
	panic(errors.New("Invalid value of FirewallDirection: " + string(self)))
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
	return nil, errors.New("Invalid value of FirewallDirection: " + string(*self))
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
	return nil, errors.New("Invalid value of FirewallDirection: " + string(*self))
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
	return errors.New("Unknown FirewallDirection: " + s)
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
	return errors.New("Unknown FirewallDirection: " + s)
}

type FirewallPolicy string

const FirewallPolicyAccept FirewallPolicy = "ACCEPT"
const FirewallPolicyDrop FirewallPolicy = "DROP"
const FirewallPolicyEmpty FirewallPolicy = ""
const FirewallPolicyReturn FirewallPolicy = "RETURN"

func (self FirewallPolicy) GetPtr() *FirewallPolicy { var v = self; return &v }

func (self FirewallPolicy) String() string {
	switch self {
	case FirewallPolicyAccept:
		return "ACCEPT"
	case FirewallPolicyDrop:
		return "DROP"
	case FirewallPolicyEmpty:
		return ""
	case FirewallPolicyReturn:
		return "RETURN"
	}
	if len(self) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of FirewallPolicy: " + string(self)))
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
	return nil, errors.New("Invalid value of FirewallPolicy: " + string(*self))
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
	return nil, errors.New("Invalid value of FirewallPolicy: " + string(*self))
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
	return errors.New("Unknown FirewallPolicy: " + s)
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
	return errors.New("Unknown FirewallPolicy: " + s)
}

type FirmwareUpdateMode string

const FirmwareUpdateModeCheck FirmwareUpdateMode = "check"
const FirmwareUpdateModeOff FirmwareUpdateMode = "off"
const FirmwareUpdateModeOn FirmwareUpdateMode = "on"

func (self FirmwareUpdateMode) GetPtr() *FirmwareUpdateMode { var v = self; return &v }

func (self FirmwareUpdateMode) String() string {
	switch self {
	case FirmwareUpdateModeCheck:
		return "check"
	case FirmwareUpdateModeOff:
		return "off"
	case FirmwareUpdateModeOn:
		return "on"
	}
	if len(self) == 0 {
		return "check"
	}
	panic(errors.New("Invalid value of FirmwareUpdateMode: " + string(self)))
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
	return nil, errors.New("Invalid value of FirmwareUpdateMode: " + string(*self))
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
	return nil, errors.New("Invalid value of FirmwareUpdateMode: " + string(*self))
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
	return errors.New("Unknown FirmwareUpdateMode: " + s)
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
	return errors.New("Unknown FirmwareUpdateMode: " + s)
}

type L3Protocol string

const L3ProtocolEmpty L3Protocol = ""
const L3ProtocolIP L3Protocol = "ip"
const L3ProtocolIPv4 L3Protocol = "ipv4"
const L3ProtocolIPv6 L3Protocol = "ipv6"

func (self L3Protocol) GetPtr() *L3Protocol { var v = self; return &v }

func (self L3Protocol) String() string {
	switch self {
	case L3ProtocolEmpty:
		return ""
	case L3ProtocolIP:
		return "ip"
	case L3ProtocolIPv4:
		return "ipv4"
	case L3ProtocolIPv6:
		return "ipv6"
	}
	if len(self) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of L3Protocol: " + string(self)))
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
	return nil, errors.New("Invalid value of L3Protocol: " + string(*self))
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
	return nil, errors.New("Invalid value of L3Protocol: " + string(*self))
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
	return errors.New("Unknown L3Protocol: " + s)
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
	return errors.New("Unknown L3Protocol: " + s)
}

type L4Protocol string

const L4ProtocolEmpty L4Protocol = ""
const L4ProtocolTCP L4Protocol = "TCP"
const L4ProtocolUDP L4Protocol = "UDP"

func (self L4Protocol) GetPtr() *L4Protocol { var v = self; return &v }

func (self L4Protocol) String() string {
	switch self {
	case L4ProtocolEmpty:
		return ""
	case L4ProtocolTCP:
		return "TCP"
	case L4ProtocolUDP:
		return "UDP"
	}
	if len(self) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of L4Protocol: " + string(self)))
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
	return nil, errors.New("Invalid value of L4Protocol: " + string(*self))
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
	return nil, errors.New("Invalid value of L4Protocol: " + string(*self))
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
	return errors.New("Unknown L4Protocol: " + s)
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
	return errors.New("Unknown L4Protocol: " + s)
}

type MCSRequire string

const MCSRequireHT MCSRequire = "ht"
const MCSRequireOff MCSRequire = "off"
const MCSRequireVHT MCSRequire = "vht"

func (self MCSRequire) GetPtr() *MCSRequire { var v = self; return &v }

func (self MCSRequire) String() string {
	switch self {
	case MCSRequireHT:
		return "ht"
	case MCSRequireOff:
		return "off"
	case MCSRequireVHT:
		return "vht"
	}
	if len(self) == 0 {
		return "off"
	}
	panic(errors.New("Invalid value of MCSRequire: " + string(self)))
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
	return nil, errors.New("Invalid value of MCSRequire: " + string(*self))
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
	return nil, errors.New("Invalid value of MCSRequire: " + string(*self))
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
	return errors.New("Unknown MCSRequire: " + s)
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
	return errors.New("Unknown MCSRequire: " + s)
}

type MacFilterType string

const MacFilterTypeBlackList MacFilterType = "BlackList"
const MacFilterTypeNone MacFilterType = "None"
const MacFilterTypeWhiteList MacFilterType = "WhiteList"

func (self MacFilterType) GetPtr() *MacFilterType { var v = self; return &v }

func (self MacFilterType) String() string {
	switch self {
	case MacFilterTypeBlackList:
		return "BlackList"
	case MacFilterTypeNone:
		return "None"
	case MacFilterTypeWhiteList:
		return "WhiteList"
	}
	if len(self) == 0 {
		return "None"
	}
	panic(errors.New("Invalid value of MacFilterType: " + string(self)))
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
	return nil, errors.New("Invalid value of MacFilterType: " + string(*self))
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
	return nil, errors.New("Invalid value of MacFilterType: " + string(*self))
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
	return errors.New("Unknown MacFilterType: " + s)
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
	return errors.New("Unknown MacFilterType: " + s)
}

type Module string

const ModuleAC Module = "AC"
const ModuleAnalMW Module = "ANAL-MW"
const ModuleAny Module = "+"
const ModuleBackend Module = "BACKEND"
const ModuleCPE Module = "CPE"
const ModuleCPEStat Module = "CPE_STAT"
const ModuleClientDistance Module = "CLIENT-DISTANCE"
const ModuleClientStat Module = "CLIENT_STAT"
const ModuleConfig Module = "CONFIG"
const ModuleDB Module = "DB"
const ModuleDummy Module = "DUMMY"
const ModuleFW Module = "FW"
const ModuleLBS Module = "LBS"
const ModuleMQTTLog Module = "MQTT_LOG"
const ModuleMediator Module = "MEDIATOR"
const ModuleMonitor Module = "MONITOR"
const ModuleNone Module = ""
const ModulePortalBack Module = "PORTAL_BACKEND"
const ModuleRRM Module = "RRM"
const ModuleRadarExportMW Module = "RADAR-EXPORT-MW"
const ModuleRadarMW Module = "RADAR-MW"
const ModuleRadiusGw Module = "RADIUS_GATEWAY"
const ModuleRedirect Module = "REDIRECT"
const ModuleStat Module = "STAT"
const ModuleStatLBS Module = "STAT-LBS"
const ModuleTunManager Module = "TUN_MANAGER"
const ModuleSnmpWalker Module = "SNMP_WALKER"

func (self Module) GetPtr() *Module { var v = self; return &v }

func (self Module) String() string {
	switch self {
	case ModuleAC:
		return "AC"
	case ModuleAnalMW:
		return "ANAL-MW"
	case ModuleAny:
		return "+"
	case ModuleBackend:
		return "BACKEND"
	case ModuleCPE:
		return "CPE"
	case ModuleCPEStat:
		return "CPE_STAT"
	case ModuleClientDistance:
		return "CLIENT-DISTANCE"
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
	case ModuleMediator:
		return "MEDIATOR"
	case ModuleMonitor:
		return "MONITOR"
	case ModuleNone:
		return ""
	case ModulePortalBack:
		return "PORTAL_BACKEND"
	case ModuleRRM:
		return "RRM"
	case ModuleRadarExportMW:
		return "RADAR-EXPORT-MW"
	case ModuleRadarMW:
		return "RADAR-MW"
	case ModuleRadiusGw:
		return "RADIUS_GATEWAY"
	case ModuleRedirect:
		return "REDIRECT"
	case ModuleStat:
		return "STAT"
	case ModuleStatLBS:
		return "STAT-LBS"
	case ModuleTunManager:
		return "TUN_MANAGER"
	case ModuleSnmpWalker:
		return "SNMP_WALKER"

	}
	if len(self) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of Module: " + string(self)))
}

func (self *Module) MarshalJSON() ([]byte, error) {
	switch *self {
	case ModuleAC:
		return json.Marshal("AC")
	case ModuleAnalMW:
		return json.Marshal("ANAL-MW")
	case ModuleAny:
		return json.Marshal("+")
	case ModuleBackend:
		return json.Marshal("BACKEND")
	case ModuleCPE:
		return json.Marshal("CPE")
	case ModuleCPEStat:
		return json.Marshal("CPE_STAT")
	case ModuleClientDistance:
		return json.Marshal("CLIENT-DISTANCE")
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
	case ModuleMediator:
		return json.Marshal("MEDIATOR")
	case ModuleMonitor:
		return json.Marshal("MONITOR")
	case ModuleNone:
		return json.Marshal("")
	case ModulePortalBack:
		return json.Marshal("PORTAL_BACKEND")
	case ModuleRRM:
		return json.Marshal("RRM")
	case ModuleRadarExportMW:
		return json.Marshal("RADAR-EXPORT-MW")
	case ModuleRadarMW:
		return json.Marshal("RADAR-MW")
	case ModuleRadiusGw:
		return json.Marshal("RADIUS_GATEWAY")
	case ModuleRedirect:
		return json.Marshal("REDIRECT")
	case ModuleStat:
		return json.Marshal("STAT")
	case ModuleStatLBS:
		return json.Marshal("STAT-LBS")
	case ModuleTunManager:
		return json.Marshal("TUN_MANAGER")
	case ModuleSnmpWalker:
		return json.Marshal("SNMP_WALKER")
	}
	if len(*self) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of Module: " + string(*self))
}

func (self *Module) GetBSON() (interface{}, error) {
	switch *self {
	case ModuleAC:
		return "AC", nil
	case ModuleAnalMW:
		return "ANAL-MW", nil
	case ModuleAny:
		return "+", nil
	case ModuleBackend:
		return "BACKEND", nil
	case ModuleCPE:
		return "CPE", nil
	case ModuleCPEStat:
		return "CPE_STAT", nil
	case ModuleClientDistance:
		return "CLIENT-DISTANCE", nil
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
	case ModuleMediator:
		return "MEDIATOR", nil
	case ModuleMonitor:
		return "MONITOR", nil
	case ModuleNone:
		return "", nil
	case ModulePortalBack:
		return "PORTAL_BACKEND", nil
	case ModuleRRM:
		return "RRM", nil
	case ModuleRadarExportMW:
		return "RADAR-EXPORT-MW", nil
	case ModuleRadarMW:
		return "RADAR-MW", nil
	case ModuleRadiusGw:
		return "RADIUS_GATEWAY", nil
	case ModuleRedirect:
		return "REDIRECT", nil
	case ModuleStat:
		return "STAT", nil
	case ModuleStatLBS:
		return "STAT-LBS", nil
	case ModuleTunManager:
		return "TUN_MANAGER", nil
	case ModuleSnmpWalker:
		return "SNMP_WALKER",nil
	}
	if len(*self) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of Module: " + string(*self))
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
	case "ANAL-MW":
		*self = ModuleAnalMW
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
	case "CLIENT-DISTANCE":
		*self = ModuleClientDistance
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
	case "MEDIATOR":
		*self = ModuleMediator
		return nil
	case "MONITOR":
		*self = ModuleMonitor
		return nil
	case "":
		*self = ModuleNone
		return nil
	case "PORTAL_BACKEND":
		*self = ModulePortalBack
		return nil
	case "RRM":
		*self = ModuleRRM
		return nil
	case "RADAR-EXPORT-MW":
		*self = ModuleRadarExportMW
		return nil
	case "RADAR-MW":
		*self = ModuleRadarMW
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
	case "STAT-LBS":
		*self = ModuleStatLBS
		return nil
	case "TUN_MANAGER":
		*self = ModuleTunManager
		return nil
	case "SNMP_WALKER":
		*self = ModuleSnmpWalker
		return nil

	}
	if len(s) == 0 {
		*self = ModuleNone
		return nil
	}
	return errors.New("Unknown Module: " + s)
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
	case "ANAL-MW":
		*self = ModuleAnalMW
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
	case "CLIENT-DISTANCE":
		*self = ModuleClientDistance
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
	case "MEDIATOR":
		*self = ModuleMediator
		return nil
	case "MONITOR":
		*self = ModuleMonitor
		return nil
	case "":
		*self = ModuleNone
		return nil
	case "PORTAL_BACKEND":
		*self = ModulePortalBack
		return nil
	case "RRM":
		*self = ModuleRRM
		return nil
	case "RADAR-EXPORT-MW":
		*self = ModuleRadarExportMW
		return nil
	case "RADAR-MW":
		*self = ModuleRadarMW
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
	case "STAT-LBS":
		*self = ModuleStatLBS
		return nil
	case "TUN_MANAGER":
		*self = ModuleTunManager
		return nil
	case "SNMP_WALKER":
		*self = ModuleSnmpWalker
		return nil
	}
	if len(s) == 0 {
		*self = ModuleNone
		return nil
	}
	return errors.New("Unknown Module: " + s)
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

func (self Operation) String() string {
	switch self {
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
	panic(errors.New("Invalid value of Operation: " + string(self)))
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
	return nil, errors.New("Invalid value of Operation: " + string(*self))
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
	return nil, errors.New("Invalid value of Operation: " + string(*self))
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
	return errors.New("Unknown Operation: " + s)
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
	return errors.New("Unknown Operation: " + s)
}

type PortalAuthType string

const PortalAuthTypeExternal PortalAuthType = "external"
const PortalAuthTypeNone PortalAuthType = ""
const PortalAuthTypeOAuth2 PortalAuthType = "oauth2"
const PortalAuthTypeRADIUS PortalAuthType = "radius"
const PortalAuthTypeSMS PortalAuthType = "sms"

func (self PortalAuthType) GetPtr() *PortalAuthType { var v = self; return &v }

func (self PortalAuthType) String() string {
	switch self {
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
	if len(self) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of PortalAuthType: " + string(self)))
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
	return nil, errors.New("Invalid value of PortalAuthType: " + string(*self))
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
	return nil, errors.New("Invalid value of PortalAuthType: " + string(*self))
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
	return errors.New("Unknown PortalAuthType: " + s)
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
	return errors.New("Unknown PortalAuthType: " + s)
}

type PortalProfileType string

const PortalProfileTypeFree PortalProfileType = "free"
const PortalProfileTypePremium PortalProfileType = "premium"
const PortalProfileTypeSponsor PortalProfileType = "sponsor"

func (self PortalProfileType) GetPtr() *PortalProfileType { var v = self; return &v }

func (self PortalProfileType) String() string {
	switch self {
	case PortalProfileTypeFree:
		return "free"
	case PortalProfileTypePremium:
		return "premium"
	case PortalProfileTypeSponsor:
		return "sponsor"
	}
	if len(self) == 0 {
		return "free"
	}
	panic(errors.New("Invalid value of PortalProfileType: " + string(self)))
}

func (self *PortalProfileType) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalProfileTypeFree:
		return json.Marshal("free")
	case PortalProfileTypePremium:
		return json.Marshal("premium")
	case PortalProfileTypeSponsor:
		return json.Marshal("sponsor")
	}
	if len(*self) == 0 {
		return json.Marshal("free")
	}
	return nil, errors.New("Invalid value of PortalProfileType: " + string(*self))
}

func (self *PortalProfileType) GetBSON() (interface{}, error) {
	switch *self {
	case PortalProfileTypeFree:
		return "free", nil
	case PortalProfileTypePremium:
		return "premium", nil
	case PortalProfileTypeSponsor:
		return "sponsor", nil
	}
	if len(*self) == 0 {
		return "free", nil
	}
	return nil, errors.New("Invalid value of PortalProfileType: " + string(*self))
}

func (self *PortalProfileType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "free":
		*self = PortalProfileTypeFree
		return nil
	case "premium":
		*self = PortalProfileTypePremium
		return nil
	case "sponsor":
		*self = PortalProfileTypeSponsor
		return nil
	}
	if len(s) == 0 {
		*self = PortalProfileTypeFree
		return nil
	}
	return errors.New("Unknown PortalProfileType: " + s)
}

func (self *PortalProfileType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "free":
		*self = PortalProfileTypeFree
		return nil
	case "premium":
		*self = PortalProfileTypePremium
		return nil
	case "sponsor":
		*self = PortalProfileTypeSponsor
		return nil
	}
	if len(s) == 0 {
		*self = PortalProfileTypeFree
		return nil
	}
	return errors.New("Unknown PortalProfileType: " + s)
}

type RRMAlgoType string

const RRMAlgoTypeBlind RRMAlgoType = "Blind"
const RRMAlgoTypeDummy RRMAlgoType = "Dummy"
const RRMAlgoTypeGreed RRMAlgoType = "Greed"

func (self RRMAlgoType) GetPtr() *RRMAlgoType { var v = self; return &v }

func (self RRMAlgoType) String() string {
	switch self {
	case RRMAlgoTypeBlind:
		return "Blind"
	case RRMAlgoTypeDummy:
		return "Dummy"
	case RRMAlgoTypeGreed:
		return "Greed"
	}
	if len(self) == 0 {
		return "Greed"
	}
	panic(errors.New("Invalid value of RRMAlgoType: " + string(self)))
}

func (self *RRMAlgoType) MarshalJSON() ([]byte, error) {
	switch *self {
	case RRMAlgoTypeBlind:
		return json.Marshal("Blind")
	case RRMAlgoTypeDummy:
		return json.Marshal("Dummy")
	case RRMAlgoTypeGreed:
		return json.Marshal("Greed")
	}
	if len(*self) == 0 {
		return json.Marshal("Greed")
	}
	return nil, errors.New("Invalid value of RRMAlgoType: " + string(*self))
}

func (self *RRMAlgoType) GetBSON() (interface{}, error) {
	switch *self {
	case RRMAlgoTypeBlind:
		return "Blind", nil
	case RRMAlgoTypeDummy:
		return "Dummy", nil
	case RRMAlgoTypeGreed:
		return "Greed", nil
	}
	if len(*self) == 0 {
		return "Greed", nil
	}
	return nil, errors.New("Invalid value of RRMAlgoType: " + string(*self))
}

func (self *RRMAlgoType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Blind":
		*self = RRMAlgoTypeBlind
		return nil
	case "Dummy":
		*self = RRMAlgoTypeDummy
		return nil
	case "Greed":
		*self = RRMAlgoTypeGreed
		return nil
	}
	if len(s) == 0 {
		*self = RRMAlgoTypeGreed
		return nil
	}
	return errors.New("Unknown RRMAlgoType: " + s)
}

func (self *RRMAlgoType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "Blind":
		*self = RRMAlgoTypeBlind
		return nil
	case "Dummy":
		*self = RRMAlgoTypeDummy
		return nil
	case "Greed":
		*self = RRMAlgoTypeGreed
		return nil
	}
	if len(s) == 0 {
		*self = RRMAlgoTypeGreed
		return nil
	}
	return errors.New("Unknown RRMAlgoType: " + s)
}

type RadarExportFilter string

const RadarExportFilterAll RadarExportFilter = "all"
const RadarExportFilterNew RadarExportFilter = "new"
const RadarExportFilterReturn RadarExportFilter = "return"

func (self RadarExportFilter) GetPtr() *RadarExportFilter { var v = self; return &v }

func (self RadarExportFilter) String() string {
	switch self {
	case RadarExportFilterAll:
		return "all"
	case RadarExportFilterNew:
		return "new"
	case RadarExportFilterReturn:
		return "return"
	}
	if len(self) == 0 {
		return "all"
	}
	panic(errors.New("Invalid value of RadarExportFilter: " + string(self)))
}

func (self *RadarExportFilter) MarshalJSON() ([]byte, error) {
	switch *self {
	case RadarExportFilterAll:
		return json.Marshal("all")
	case RadarExportFilterNew:
		return json.Marshal("new")
	case RadarExportFilterReturn:
		return json.Marshal("return")
	}
	if len(*self) == 0 {
		return json.Marshal("all")
	}
	return nil, errors.New("Invalid value of RadarExportFilter: " + string(*self))
}

func (self *RadarExportFilter) GetBSON() (interface{}, error) {
	switch *self {
	case RadarExportFilterAll:
		return "all", nil
	case RadarExportFilterNew:
		return "new", nil
	case RadarExportFilterReturn:
		return "return", nil
	}
	if len(*self) == 0 {
		return "all", nil
	}
	return nil, errors.New("Invalid value of RadarExportFilter: " + string(*self))
}

func (self *RadarExportFilter) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "all":
		*self = RadarExportFilterAll
		return nil
	case "new":
		*self = RadarExportFilterNew
		return nil
	case "return":
		*self = RadarExportFilterReturn
		return nil
	}
	if len(s) == 0 {
		*self = RadarExportFilterAll
		return nil
	}
	return errors.New("Unknown RadarExportFilter: " + s)
}

func (self *RadarExportFilter) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "all":
		*self = RadarExportFilterAll
		return nil
	case "new":
		*self = RadarExportFilterNew
		return nil
	case "return":
		*self = RadarExportFilterReturn
		return nil
	}
	if len(s) == 0 {
		*self = RadarExportFilterAll
		return nil
	}
	return errors.New("Unknown RadarExportFilter: " + s)
}

type RadarExportFormat string

const RadarExportFormatCSV RadarExportFormat = "csv"
const RadarExportFormatJson RadarExportFormat = "json"
const RadarExportFormatTxt RadarExportFormat = "txt"

func (self RadarExportFormat) GetPtr() *RadarExportFormat { var v = self; return &v }

func (self RadarExportFormat) String() string {
	switch self {
	case RadarExportFormatCSV:
		return "csv"
	case RadarExportFormatJson:
		return "json"
	case RadarExportFormatTxt:
		return "txt"
	}
	if len(self) == 0 {
		return "csv"
	}
	panic(errors.New("Invalid value of RadarExportFormat: " + string(self)))
}

func (self *RadarExportFormat) MarshalJSON() ([]byte, error) {
	switch *self {
	case RadarExportFormatCSV:
		return json.Marshal("csv")
	case RadarExportFormatJson:
		return json.Marshal("json")
	case RadarExportFormatTxt:
		return json.Marshal("txt")
	}
	if len(*self) == 0 {
		return json.Marshal("csv")
	}
	return nil, errors.New("Invalid value of RadarExportFormat: " + string(*self))
}

func (self *RadarExportFormat) GetBSON() (interface{}, error) {
	switch *self {
	case RadarExportFormatCSV:
		return "csv", nil
	case RadarExportFormatJson:
		return "json", nil
	case RadarExportFormatTxt:
		return "txt", nil
	}
	if len(*self) == 0 {
		return "csv", nil
	}
	return nil, errors.New("Invalid value of RadarExportFormat: " + string(*self))
}

func (self *RadarExportFormat) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*self = RadarExportFormatCSV
		return nil
	case "json":
		*self = RadarExportFormatJson
		return nil
	case "txt":
		*self = RadarExportFormatTxt
		return nil
	}
	if len(s) == 0 {
		*self = RadarExportFormatCSV
		return nil
	}
	return errors.New("Unknown RadarExportFormat: " + s)
}

func (self *RadarExportFormat) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*self = RadarExportFormatCSV
		return nil
	case "json":
		*self = RadarExportFormatJson
		return nil
	case "txt":
		*self = RadarExportFormatTxt
		return nil
	}
	if len(s) == 0 {
		*self = RadarExportFormatCSV
		return nil
	}
	return errors.New("Unknown RadarExportFormat: " + s)
}

type RadarExportStatus string

const RadarExportStatusCreated RadarExportStatus = "created"
const RadarExportStatusFinished RadarExportStatus = "finished"
const RadarExportStatusRunning RadarExportStatus = "running"
const RadarExportStatusUpdated RadarExportStatus = "updated"

func (self RadarExportStatus) GetPtr() *RadarExportStatus { var v = self; return &v }

func (self RadarExportStatus) String() string {
	switch self {
	case RadarExportStatusCreated:
		return "created"
	case RadarExportStatusFinished:
		return "finished"
	case RadarExportStatusRunning:
		return "running"
	case RadarExportStatusUpdated:
		return "updated"
	}
	if len(self) == 0 {
		return "created"
	}
	panic(errors.New("Invalid value of RadarExportStatus: " + string(self)))
}

func (self *RadarExportStatus) MarshalJSON() ([]byte, error) {
	switch *self {
	case RadarExportStatusCreated:
		return json.Marshal("created")
	case RadarExportStatusFinished:
		return json.Marshal("finished")
	case RadarExportStatusRunning:
		return json.Marshal("running")
	case RadarExportStatusUpdated:
		return json.Marshal("updated")
	}
	if len(*self) == 0 {
		return json.Marshal("created")
	}
	return nil, errors.New("Invalid value of RadarExportStatus: " + string(*self))
}

func (self *RadarExportStatus) GetBSON() (interface{}, error) {
	switch *self {
	case RadarExportStatusCreated:
		return "created", nil
	case RadarExportStatusFinished:
		return "finished", nil
	case RadarExportStatusRunning:
		return "running", nil
	case RadarExportStatusUpdated:
		return "updated", nil
	}
	if len(*self) == 0 {
		return "created", nil
	}
	return nil, errors.New("Invalid value of RadarExportStatus: " + string(*self))
}

func (self *RadarExportStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "created":
		*self = RadarExportStatusCreated
		return nil
	case "finished":
		*self = RadarExportStatusFinished
		return nil
	case "running":
		*self = RadarExportStatusRunning
		return nil
	case "updated":
		*self = RadarExportStatusUpdated
		return nil
	}
	if len(s) == 0 {
		*self = RadarExportStatusCreated
		return nil
	}
	return errors.New("Unknown RadarExportStatus: " + s)
}

func (self *RadarExportStatus) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "created":
		*self = RadarExportStatusCreated
		return nil
	case "finished":
		*self = RadarExportStatusFinished
		return nil
	case "running":
		*self = RadarExportStatusRunning
		return nil
	case "updated":
		*self = RadarExportStatusUpdated
		return nil
	}
	if len(s) == 0 {
		*self = RadarExportStatusCreated
		return nil
	}
	return errors.New("Unknown RadarExportStatus: " + s)
}

type RadarExportType string

const RadarExportTypeEmail RadarExportType = "email"
const RadarExportTypeExternal RadarExportType = "external"
const RadarExportTypeMytarget RadarExportType = "mytarget"
const RadarExportTypeYandex RadarExportType = "yandex"

func (self RadarExportType) GetPtr() *RadarExportType { var v = self; return &v }

func (self RadarExportType) String() string {
	switch self {
	case RadarExportTypeEmail:
		return "email"
	case RadarExportTypeExternal:
		return "external"
	case RadarExportTypeMytarget:
		return "mytarget"
	case RadarExportTypeYandex:
		return "yandex"
	}
	if len(self) == 0 {
		return "email"
	}
	panic(errors.New("Invalid value of RadarExportType: " + string(self)))
}

func (self *RadarExportType) MarshalJSON() ([]byte, error) {
	switch *self {
	case RadarExportTypeEmail:
		return json.Marshal("email")
	case RadarExportTypeExternal:
		return json.Marshal("external")
	case RadarExportTypeMytarget:
		return json.Marshal("mytarget")
	case RadarExportTypeYandex:
		return json.Marshal("yandex")
	}
	if len(*self) == 0 {
		return json.Marshal("email")
	}
	return nil, errors.New("Invalid value of RadarExportType: " + string(*self))
}

func (self *RadarExportType) GetBSON() (interface{}, error) {
	switch *self {
	case RadarExportTypeEmail:
		return "email", nil
	case RadarExportTypeExternal:
		return "external", nil
	case RadarExportTypeMytarget:
		return "mytarget", nil
	case RadarExportTypeYandex:
		return "yandex", nil
	}
	if len(*self) == 0 {
		return "email", nil
	}
	return nil, errors.New("Invalid value of RadarExportType: " + string(*self))
}

func (self *RadarExportType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "email":
		*self = RadarExportTypeEmail
		return nil
	case "external":
		*self = RadarExportTypeExternal
		return nil
	case "mytarget":
		*self = RadarExportTypeMytarget
		return nil
	case "yandex":
		*self = RadarExportTypeYandex
		return nil
	}
	if len(s) == 0 {
		*self = RadarExportTypeEmail
		return nil
	}
	return errors.New("Unknown RadarExportType: " + s)
}

func (self *RadarExportType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "email":
		*self = RadarExportTypeEmail
		return nil
	case "external":
		*self = RadarExportTypeExternal
		return nil
	case "mytarget":
		*self = RadarExportTypeMytarget
		return nil
	case "yandex":
		*self = RadarExportTypeYandex
		return nil
	}
	if len(s) == 0 {
		*self = RadarExportTypeEmail
		return nil
	}
	return errors.New("Unknown RadarExportType: " + s)
}

type RadiusMessageType string

const RadiusMessageTypeAccessAccept RadiusMessageType = "access-accept"
const RadiusMessageTypeAccessReject RadiusMessageType = "access-reject"
const RadiusMessageTypeAccessRequest RadiusMessageType = "access-request"
const RadiusMessageTypeAccountingRequest RadiusMessageType = "accounting"

func (self RadiusMessageType) GetPtr() *RadiusMessageType { var v = self; return &v }

func (self RadiusMessageType) String() string {
	switch self {
	case RadiusMessageTypeAccessAccept:
		return "access-accept"
	case RadiusMessageTypeAccessReject:
		return "access-reject"
	case RadiusMessageTypeAccessRequest:
		return "access-request"
	case RadiusMessageTypeAccountingRequest:
		return "accounting"
	}
	if len(self) == 0 {
		return "accounting"
	}
	panic(errors.New("Invalid value of RadiusMessageType: " + string(self)))
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
	return nil, errors.New("Invalid value of RadiusMessageType: " + string(*self))
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
	return nil, errors.New("Invalid value of RadiusMessageType: " + string(*self))
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
	return errors.New("Unknown RadiusMessageType: " + s)
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
	return errors.New("Unknown RadiusMessageType: " + s)
}

type ReportActionType string

const ReportActionTypeEmail ReportActionType = "email"
const ReportActionTypeScript ReportActionType = "script"

func (self ReportActionType) GetPtr() *ReportActionType { var v = self; return &v }

func (self ReportActionType) String() string {
	switch self {
	case ReportActionTypeEmail:
		return "email"
	case ReportActionTypeScript:
		return "script"
	}
	if len(self) == 0 {
		return "email"
	}
	panic(errors.New("Invalid value of ReportActionType: " + string(self)))
}

func (self *ReportActionType) MarshalJSON() ([]byte, error) {
	switch *self {
	case ReportActionTypeEmail:
		return json.Marshal("email")
	case ReportActionTypeScript:
		return json.Marshal("script")
	}
	if len(*self) == 0 {
		return json.Marshal("email")
	}
	return nil, errors.New("Invalid value of ReportActionType: " + string(*self))
}

func (self *ReportActionType) GetBSON() (interface{}, error) {
	switch *self {
	case ReportActionTypeEmail:
		return "email", nil
	case ReportActionTypeScript:
		return "script", nil
	}
	if len(*self) == 0 {
		return "email", nil
	}
	return nil, errors.New("Invalid value of ReportActionType: " + string(*self))
}

func (self *ReportActionType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "email":
		*self = ReportActionTypeEmail
		return nil
	case "script":
		*self = ReportActionTypeScript
		return nil
	}
	if len(s) == 0 {
		*self = ReportActionTypeEmail
		return nil
	}
	return errors.New("Unknown ReportActionType: " + s)
}

func (self *ReportActionType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "email":
		*self = ReportActionTypeEmail
		return nil
	case "script":
		*self = ReportActionTypeScript
		return nil
	}
	if len(s) == 0 {
		*self = ReportActionTypeEmail
		return nil
	}
	return errors.New("Unknown ReportActionType: " + s)
}

type ReportFormat string

const ReportFormatCSV ReportFormat = "csv"
const ReportFormatJson ReportFormat = "json"
const ReportFormatTxt ReportFormat = "txt"
const ReportFormatXLSX ReportFormat = "xlsx"

func (self ReportFormat) GetPtr() *ReportFormat { var v = self; return &v }

func (self ReportFormat) String() string {
	switch self {
	case ReportFormatCSV:
		return "csv"
	case ReportFormatJson:
		return "json"
	case ReportFormatTxt:
		return "txt"
	case ReportFormatXLSX:
		return "xlsx"
	}
	if len(self) == 0 {
		return "json"
	}
	panic(errors.New("Invalid value of ReportFormat: " + string(self)))
}

func (self *ReportFormat) MarshalJSON() ([]byte, error) {
	switch *self {
	case ReportFormatCSV:
		return json.Marshal("csv")
	case ReportFormatJson:
		return json.Marshal("json")
	case ReportFormatTxt:
		return json.Marshal("txt")
	case ReportFormatXLSX:
		return json.Marshal("xlsx")
	}
	if len(*self) == 0 {
		return json.Marshal("json")
	}
	return nil, errors.New("Invalid value of ReportFormat: " + string(*self))
}

func (self *ReportFormat) GetBSON() (interface{}, error) {
	switch *self {
	case ReportFormatCSV:
		return "csv", nil
	case ReportFormatJson:
		return "json", nil
	case ReportFormatTxt:
		return "txt", nil
	case ReportFormatXLSX:
		return "xlsx", nil
	}
	if len(*self) == 0 {
		return "json", nil
	}
	return nil, errors.New("Invalid value of ReportFormat: " + string(*self))
}

func (self *ReportFormat) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*self = ReportFormatCSV
		return nil
	case "json":
		*self = ReportFormatJson
		return nil
	case "txt":
		*self = ReportFormatTxt
		return nil
	case "xlsx":
		*self = ReportFormatXLSX
		return nil
	}
	if len(s) == 0 {
		*self = ReportFormatJson
		return nil
	}
	return errors.New("Unknown ReportFormat: " + s)
}

func (self *ReportFormat) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*self = ReportFormatCSV
		return nil
	case "json":
		*self = ReportFormatJson
		return nil
	case "txt":
		*self = ReportFormatTxt
		return nil
	case "xlsx":
		*self = ReportFormatXLSX
		return nil
	}
	if len(s) == 0 {
		*self = ReportFormatJson
		return nil
	}
	return errors.New("Unknown ReportFormat: " + s)
}

type ReportPeriod string

const ReportPeriodDay ReportPeriod = "day"
const ReportPeriodMonth ReportPeriod = "month"
const ReportPeriodNow ReportPeriod = "now"
const ReportPeriodWeek ReportPeriod = "week"

func (self ReportPeriod) GetPtr() *ReportPeriod { var v = self; return &v }

func (self ReportPeriod) String() string {
	switch self {
	case ReportPeriodDay:
		return "day"
	case ReportPeriodMonth:
		return "month"
	case ReportPeriodNow:
		return "now"
	case ReportPeriodWeek:
		return "week"
	}
	if len(self) == 0 {
		return "now"
	}
	panic(errors.New("Invalid value of ReportPeriod: " + string(self)))
}

func (self *ReportPeriod) MarshalJSON() ([]byte, error) {
	switch *self {
	case ReportPeriodDay:
		return json.Marshal("day")
	case ReportPeriodMonth:
		return json.Marshal("month")
	case ReportPeriodNow:
		return json.Marshal("now")
	case ReportPeriodWeek:
		return json.Marshal("week")
	}
	if len(*self) == 0 {
		return json.Marshal("now")
	}
	return nil, errors.New("Invalid value of ReportPeriod: " + string(*self))
}

func (self *ReportPeriod) GetBSON() (interface{}, error) {
	switch *self {
	case ReportPeriodDay:
		return "day", nil
	case ReportPeriodMonth:
		return "month", nil
	case ReportPeriodNow:
		return "now", nil
	case ReportPeriodWeek:
		return "week", nil
	}
	if len(*self) == 0 {
		return "now", nil
	}
	return nil, errors.New("Invalid value of ReportPeriod: " + string(*self))
}

func (self *ReportPeriod) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "day":
		*self = ReportPeriodDay
		return nil
	case "month":
		*self = ReportPeriodMonth
		return nil
	case "now":
		*self = ReportPeriodNow
		return nil
	case "week":
		*self = ReportPeriodWeek
		return nil
	}
	if len(s) == 0 {
		*self = ReportPeriodNow
		return nil
	}
	return errors.New("Unknown ReportPeriod: " + s)
}

func (self *ReportPeriod) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "day":
		*self = ReportPeriodDay
		return nil
	case "month":
		*self = ReportPeriodMonth
		return nil
	case "now":
		*self = ReportPeriodNow
		return nil
	case "week":
		*self = ReportPeriodWeek
		return nil
	}
	if len(s) == 0 {
		*self = ReportPeriodNow
		return nil
	}
	return errors.New("Unknown ReportPeriod: " + s)
}

type ReportSubject string

const ReportSubjectCPECommon ReportSubject = "cpes_common"
const ReportSubjectCPEs ReportSubject = "cpes"
const ReportSubjectClients ReportSubject = "clients"
const ReportSubjectEvents ReportSubject = "events"

func (self ReportSubject) GetPtr() *ReportSubject { var v = self; return &v }

func (self ReportSubject) String() string {
	switch self {
	case ReportSubjectCPECommon:
		return "cpes_common"
	case ReportSubjectCPEs:
		return "cpes"
	case ReportSubjectClients:
		return "clients"
	case ReportSubjectEvents:
		return "events"
	}
	if len(self) == 0 {
		return "clients"
	}
	panic(errors.New("Invalid value of ReportSubject: " + string(self)))
}

func (self *ReportSubject) MarshalJSON() ([]byte, error) {
	switch *self {
	case ReportSubjectCPECommon:
		return json.Marshal("cpes_common")
	case ReportSubjectCPEs:
		return json.Marshal("cpes")
	case ReportSubjectClients:
		return json.Marshal("clients")
	case ReportSubjectEvents:
		return json.Marshal("events")
	}
	if len(*self) == 0 {
		return json.Marshal("clients")
	}
	return nil, errors.New("Invalid value of ReportSubject: " + string(*self))
}

func (self *ReportSubject) GetBSON() (interface{}, error) {
	switch *self {
	case ReportSubjectCPECommon:
		return "cpes_common", nil
	case ReportSubjectCPEs:
		return "cpes", nil
	case ReportSubjectClients:
		return "clients", nil
	case ReportSubjectEvents:
		return "events", nil
	}
	if len(*self) == 0 {
		return "clients", nil
	}
	return nil, errors.New("Invalid value of ReportSubject: " + string(*self))
}

func (self *ReportSubject) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "cpes_common":
		*self = ReportSubjectCPECommon
		return nil
	case "cpes":
		*self = ReportSubjectCPEs
		return nil
	case "clients":
		*self = ReportSubjectClients
		return nil
	case "events":
		*self = ReportSubjectEvents
		return nil
	}
	if len(s) == 0 {
		*self = ReportSubjectClients
		return nil
	}
	return errors.New("Unknown ReportSubject: " + s)
}

func (self *ReportSubject) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "cpes_common":
		*self = ReportSubjectCPECommon
		return nil
	case "cpes":
		*self = ReportSubjectCPEs
		return nil
	case "clients":
		*self = ReportSubjectClients
		return nil
	case "events":
		*self = ReportSubjectEvents
		return nil
	}
	if len(s) == 0 {
		*self = ReportSubjectClients
		return nil
	}
	return errors.New("Unknown ReportSubject: " + s)
}

type ReportType string

const ReportTypeCurrent ReportType = "current"
const ReportTypeSummary ReportType = "summary"

func (self ReportType) GetPtr() *ReportType { var v = self; return &v }

func (self ReportType) String() string {
	switch self {
	case ReportTypeCurrent:
		return "current"
	case ReportTypeSummary:
		return "summary"
	}
	if len(self) == 0 {
		return "current"
	}
	panic(errors.New("Invalid value of ReportType: " + string(self)))
}

func (self *ReportType) MarshalJSON() ([]byte, error) {
	switch *self {
	case ReportTypeCurrent:
		return json.Marshal("current")
	case ReportTypeSummary:
		return json.Marshal("summary")
	}
	if len(*self) == 0 {
		return json.Marshal("current")
	}
	return nil, errors.New("Invalid value of ReportType: " + string(*self))
}

func (self *ReportType) GetBSON() (interface{}, error) {
	switch *self {
	case ReportTypeCurrent:
		return "current", nil
	case ReportTypeSummary:
		return "summary", nil
	}
	if len(*self) == 0 {
		return "current", nil
	}
	return nil, errors.New("Invalid value of ReportType: " + string(*self))
}

func (self *ReportType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "current":
		*self = ReportTypeCurrent
		return nil
	case "summary":
		*self = ReportTypeSummary
		return nil
	}
	if len(s) == 0 {
		*self = ReportTypeCurrent
		return nil
	}
	return errors.New("Unknown ReportType: " + s)
}

func (self *ReportType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "current":
		*self = ReportTypeCurrent
		return nil
	case "summary":
		*self = ReportTypeSummary
		return nil
	}
	if len(s) == 0 {
		*self = ReportTypeCurrent
		return nil
	}
	return errors.New("Unknown ReportType: " + s)
}

type SecuritySuite string

const SecuritySuiteAES SecuritySuite = "aes"
const SecuritySuiteTKIP SecuritySuite = "tkip"

func (self SecuritySuite) GetPtr() *SecuritySuite { var v = self; return &v }

func (self SecuritySuite) String() string {
	switch self {
	case SecuritySuiteAES:
		return "aes"
	case SecuritySuiteTKIP:
		return "tkip"
	}
	panic(errors.New("Invalid value of SecuritySuite: " + string(self)))
}

func (self *SecuritySuite) MarshalJSON() ([]byte, error) {
	switch *self {
	case SecuritySuiteAES:
		return json.Marshal("aes")
	case SecuritySuiteTKIP:
		return json.Marshal("tkip")
	}
	return nil, errors.New("Invalid value of SecuritySuite: " + string(*self))
}

func (self *SecuritySuite) GetBSON() (interface{}, error) {
	switch *self {
	case SecuritySuiteAES:
		return "aes", nil
	case SecuritySuiteTKIP:
		return "tkip", nil
	}
	return nil, errors.New("Invalid value of SecuritySuite: " + string(*self))
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
	return errors.New("Unknown SecuritySuite: " + s)
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
	return errors.New("Unknown SecuritySuite: " + s)
}

type SecurityType string

const SecurityTypeNone SecurityType = "open"
const SecurityTypeWPA2Enterprise SecurityType = "wpa2enterprise"
const SecurityTypeWPA2Personal SecurityType = "wpa2personal"
const SecurityTypeWPAEnterprise SecurityType = "wpaenterprise"
const SecurityTypeWPAPersonal SecurityType = "wpapersonal"

func (self SecurityType) GetPtr() *SecurityType { var v = self; return &v }

func (self SecurityType) String() string {
	switch self {
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
	if len(self) == 0 {
		return "open"
	}
	panic(errors.New("Invalid value of SecurityType: " + string(self)))
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
	return nil, errors.New("Invalid value of SecurityType: " + string(*self))
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
	return nil, errors.New("Invalid value of SecurityType: " + string(*self))
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
	return errors.New("Unknown SecurityType: " + s)
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
	return errors.New("Unknown SecurityType: " + s)
}

type ServiceState string

const ServiceStateConnected ServiceState = "CONNECTED"
const ServiceStateDisconnected ServiceState = "DISCONNECTED"
const ServiceStatePending ServiceState = "PENDING"

func (self ServiceState) GetPtr() *ServiceState { var v = self; return &v }

func (self ServiceState) String() string {
	switch self {
	case ServiceStateConnected:
		return "CONNECTED"
	case ServiceStateDisconnected:
		return "DISCONNECTED"
	case ServiceStatePending:
		return "PENDING"
	}
	panic(errors.New("Invalid value of ServiceState: " + string(self)))
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
	return nil, errors.New("Invalid value of ServiceState: " + string(*self))
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
	return nil, errors.New("Invalid value of ServiceState: " + string(*self))
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
	return errors.New("Unknown ServiceState: " + s)
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
	return errors.New("Unknown ServiceState: " + s)
}

type SpeedType string

const SpeedTypeGbit SpeedType = "gbit"
const SpeedTypeKbit SpeedType = "kbit"
const SpeedTypeMbit SpeedType = "mbit"
const SpeedTypeTbit SpeedType = "tbit"

func (self SpeedType) GetPtr() *SpeedType { var v = self; return &v }

func (self SpeedType) String() string {
	switch self {
	case SpeedTypeGbit:
		return "gbit"
	case SpeedTypeKbit:
		return "kbit"
	case SpeedTypeMbit:
		return "mbit"
	case SpeedTypeTbit:
		return "tbit"
	}
	if len(self) == 0 {
		return "kbit"
	}
	panic(errors.New("Invalid value of SpeedType: " + string(self)))
}

func (self *SpeedType) MarshalJSON() ([]byte, error) {
	switch *self {
	case SpeedTypeGbit:
		return json.Marshal("gbit")
	case SpeedTypeKbit:
		return json.Marshal("kbit")
	case SpeedTypeMbit:
		return json.Marshal("mbit")
	case SpeedTypeTbit:
		return json.Marshal("tbit")
	}
	if len(*self) == 0 {
		return json.Marshal("kbit")
	}
	return nil, errors.New("Invalid value of SpeedType: " + string(*self))
}

func (self *SpeedType) GetBSON() (interface{}, error) {
	switch *self {
	case SpeedTypeGbit:
		return "gbit", nil
	case SpeedTypeKbit:
		return "kbit", nil
	case SpeedTypeMbit:
		return "mbit", nil
	case SpeedTypeTbit:
		return "tbit", nil
	}
	if len(*self) == 0 {
		return "kbit", nil
	}
	return nil, errors.New("Invalid value of SpeedType: " + string(*self))
}

func (self *SpeedType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "gbit":
		*self = SpeedTypeGbit
		return nil
	case "kbit":
		*self = SpeedTypeKbit
		return nil
	case "mbit":
		*self = SpeedTypeMbit
		return nil
	case "tbit":
		*self = SpeedTypeTbit
		return nil
	}
	if len(s) == 0 {
		*self = SpeedTypeKbit
		return nil
	}
	return errors.New("Unknown SpeedType: " + s)
}

func (self *SpeedType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "gbit":
		*self = SpeedTypeGbit
		return nil
	case "kbit":
		*self = SpeedTypeKbit
		return nil
	case "mbit":
		*self = SpeedTypeMbit
		return nil
	case "tbit":
		*self = SpeedTypeTbit
		return nil
	}
	if len(s) == 0 {
		*self = SpeedTypeKbit
		return nil
	}
	return errors.New("Unknown SpeedType: " + s)
}

type StatEventRuleType string

const StatEventRuleTypeCPUload StatEventRuleType = "cpu_load"
const StatEventRuleTypeClientCon StatEventRuleType = "client_con"
const StatEventRuleTypeClientDis StatEventRuleType = "client_dis"
const StatEventRuleTypeClientFar StatEventRuleType = "client_far"
const StatEventRuleTypeConfigError StatEventRuleType = "config_error"
const StatEventRuleTypeConnected StatEventRuleType = "connected"
const StatEventRuleTypeDisconnected StatEventRuleType = "disconnected"
const StatEventRuleTypeFreeRAM StatEventRuleType = "free_ram"
const StatEventRuleTypeIfaceError StatEventRuleType = "iface_error"

func (self StatEventRuleType) GetPtr() *StatEventRuleType { var v = self; return &v }

func (self StatEventRuleType) String() string {
	switch self {
	case StatEventRuleTypeCPUload:
		return "cpu_load"
	case StatEventRuleTypeClientCon:
		return "client_con"
	case StatEventRuleTypeClientDis:
		return "client_dis"
	case StatEventRuleTypeClientFar:
		return "client_far"
	case StatEventRuleTypeConfigError:
		return "config_error"
	case StatEventRuleTypeConnected:
		return "connected"
	case StatEventRuleTypeDisconnected:
		return "disconnected"
	case StatEventRuleTypeFreeRAM:
		return "free_ram"
	case StatEventRuleTypeIfaceError:
		return "iface_error"
	}
	panic(errors.New("Invalid value of StatEventRuleType: " + string(self)))
}

func (self *StatEventRuleType) MarshalJSON() ([]byte, error) {
	switch *self {
	case StatEventRuleTypeCPUload:
		return json.Marshal("cpu_load")
	case StatEventRuleTypeClientCon:
		return json.Marshal("client_con")
	case StatEventRuleTypeClientDis:
		return json.Marshal("client_dis")
	case StatEventRuleTypeClientFar:
		return json.Marshal("client_far")
	case StatEventRuleTypeConfigError:
		return json.Marshal("config_error")
	case StatEventRuleTypeConnected:
		return json.Marshal("connected")
	case StatEventRuleTypeDisconnected:
		return json.Marshal("disconnected")
	case StatEventRuleTypeFreeRAM:
		return json.Marshal("free_ram")
	case StatEventRuleTypeIfaceError:
		return json.Marshal("iface_error")
	}
	return nil, errors.New("Invalid value of StatEventRuleType: " + string(*self))
}

func (self *StatEventRuleType) GetBSON() (interface{}, error) {
	switch *self {
	case StatEventRuleTypeCPUload:
		return "cpu_load", nil
	case StatEventRuleTypeClientCon:
		return "client_con", nil
	case StatEventRuleTypeClientDis:
		return "client_dis", nil
	case StatEventRuleTypeClientFar:
		return "client_far", nil
	case StatEventRuleTypeConfigError:
		return "config_error", nil
	case StatEventRuleTypeConnected:
		return "connected", nil
	case StatEventRuleTypeDisconnected:
		return "disconnected", nil
	case StatEventRuleTypeFreeRAM:
		return "free_ram", nil
	case StatEventRuleTypeIfaceError:
		return "iface_error", nil
	}
	return nil, errors.New("Invalid value of StatEventRuleType: " + string(*self))
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
	case "client_con":
		*self = StatEventRuleTypeClientCon
		return nil
	case "client_dis":
		*self = StatEventRuleTypeClientDis
		return nil
	case "client_far":
		*self = StatEventRuleTypeClientFar
		return nil
	case "config_error":
		*self = StatEventRuleTypeConfigError
		return nil
	case "connected":
		*self = StatEventRuleTypeConnected
		return nil
	case "disconnected":
		*self = StatEventRuleTypeDisconnected
		return nil
	case "free_ram":
		*self = StatEventRuleTypeFreeRAM
		return nil
	case "iface_error":
		*self = StatEventRuleTypeIfaceError
		return nil
	}
	return errors.New("Unknown StatEventRuleType: " + s)
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
	case "client_con":
		*self = StatEventRuleTypeClientCon
		return nil
	case "client_dis":
		*self = StatEventRuleTypeClientDis
		return nil
	case "client_far":
		*self = StatEventRuleTypeClientFar
		return nil
	case "config_error":
		*self = StatEventRuleTypeConfigError
		return nil
	case "connected":
		*self = StatEventRuleTypeConnected
		return nil
	case "disconnected":
		*self = StatEventRuleTypeDisconnected
		return nil
	case "free_ram":
		*self = StatEventRuleTypeFreeRAM
		return nil
	case "iface_error":
		*self = StatEventRuleTypeIfaceError
		return nil
	}
	return errors.New("Unknown StatEventRuleType: " + s)
}

type SystemEventLevel string

const SystemEventLevelDEBUG SystemEventLevel = "DEBUG"
const SystemEventLevelERROR SystemEventLevel = "ERROR"
const SystemEventLevelINFO SystemEventLevel = "INFO"
const SystemEventLevelWARNING SystemEventLevel = "WARNING"

func (self SystemEventLevel) GetPtr() *SystemEventLevel { var v = self; return &v }

func (self SystemEventLevel) String() string {
	switch self {
	case SystemEventLevelDEBUG:
		return "DEBUG"
	case SystemEventLevelERROR:
		return "ERROR"
	case SystemEventLevelINFO:
		return "INFO"
	case SystemEventLevelWARNING:
		return "WARNING"
	}
	if len(self) == 0 {
		return "DEBUG"
	}
	panic(errors.New("Invalid value of SystemEventLevel: " + string(self)))
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
	return nil, errors.New("Invalid value of SystemEventLevel: " + string(*self))
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
	return nil, errors.New("Invalid value of SystemEventLevel: " + string(*self))
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
	return errors.New("Unknown SystemEventLevel: " + s)
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
	return errors.New("Unknown SystemEventLevel: " + s)
}

type SystemEventType string

const SystemEventTypeAny SystemEventType = "+"
const SystemEventTypeCPEConnected SystemEventType = "CPE_CONNECTED"
const SystemEventTypeCPEDisconnected SystemEventType = "CPE_DISCONNECTED"
const SystemEventTypeCPEInterfaceState SystemEventType = "CPE_INTERFACE_STATE"
const SystemEventTypeClientAuthorization SystemEventType = "CLIENT_AUTHORIZATION"
const SystemEventTypeClientConnected SystemEventType = "CLIENT_CONNECTED"
const SystemEventTypeClientDisconnected SystemEventType = "CLIENT_DISCONNECTED"
const SystemEventTypeCpeFirmwareAvailable SystemEventType = "CPE_FIRMWARE_AVAILABLE"
const SystemEventTypeDHCPAck SystemEventType = "DHCP_ACK"
const SystemEventTypeDaemonSettingsChanged SystemEventType = "DAEMON_SETTINGS_CHANGE"
const SystemEventTypeFirmwareUploaded SystemEventType = "FIRMWARE_UPLOADED"
const SystemEventTypeLoggedError SystemEventType = "LOGGED_ERROR"
const SystemEventTypeMonitorRuleViolation SystemEventType = "MONITOR_RULE_VIOLATION"
const SystemEventTypeRRMStatus SystemEventType = "RRM_STATUS_DATA"
const SystemEventTypeRadarExportUpdate SystemEventType = "RADAR_EXPORT_UPDATE"
const SystemEventTypeRadiusAccountingSend SystemEventType = "RADIUS_ACCOUNTING_SEND"
const SystemEventTypeServiceConnected SystemEventType = "SERVICE_CONNECTED"
const SystemEventTypeServiceDisconnected SystemEventType = "SERVICE_DISCONNECTED"
const SystemEventTypeServiceFatalError SystemEventType = "SERVICE_FATAL_ERROR"
const SystemEventTypeSystemTimeChanged SystemEventType = "SYSTEM_TIME_CHANGE"
const SystemEventTypeWLANCentrAccChanged SystemEventType = "WLAN_CENTR_ACC_CHANGE"

func (self SystemEventType) GetPtr() *SystemEventType { var v = self; return &v }

func (self SystemEventType) String() string {
	switch self {
	case SystemEventTypeAny:
		return "+"
	case SystemEventTypeCPEConnected:
		return "CPE_CONNECTED"
	case SystemEventTypeCPEDisconnected:
		return "CPE_DISCONNECTED"
	case SystemEventTypeCPEInterfaceState:
		return "CPE_INTERFACE_STATE"
	case SystemEventTypeClientAuthorization:
		return "CLIENT_AUTHORIZATION"
	case SystemEventTypeClientConnected:
		return "CLIENT_CONNECTED"
	case SystemEventTypeClientDisconnected:
		return "CLIENT_DISCONNECTED"
	case SystemEventTypeCpeFirmwareAvailable:
		return "CPE_FIRMWARE_AVAILABLE"
	case SystemEventTypeDHCPAck:
		return "DHCP_ACK"
	case SystemEventTypeDaemonSettingsChanged:
		return "DAEMON_SETTINGS_CHANGE"
	case SystemEventTypeFirmwareUploaded:
		return "FIRMWARE_UPLOADED"
	case SystemEventTypeLoggedError:
		return "LOGGED_ERROR"
	case SystemEventTypeMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION"
	case SystemEventTypeRRMStatus:
		return "RRM_STATUS_DATA"
	case SystemEventTypeRadarExportUpdate:
		return "RADAR_EXPORT_UPDATE"
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
	panic(errors.New("Invalid value of SystemEventType: " + string(self)))
}

func (self *SystemEventType) MarshalJSON() ([]byte, error) {
	switch *self {
	case SystemEventTypeAny:
		return json.Marshal("+")
	case SystemEventTypeCPEConnected:
		return json.Marshal("CPE_CONNECTED")
	case SystemEventTypeCPEDisconnected:
		return json.Marshal("CPE_DISCONNECTED")
	case SystemEventTypeCPEInterfaceState:
		return json.Marshal("CPE_INTERFACE_STATE")
	case SystemEventTypeClientAuthorization:
		return json.Marshal("CLIENT_AUTHORIZATION")
	case SystemEventTypeClientConnected:
		return json.Marshal("CLIENT_CONNECTED")
	case SystemEventTypeClientDisconnected:
		return json.Marshal("CLIENT_DISCONNECTED")
	case SystemEventTypeCpeFirmwareAvailable:
		return json.Marshal("CPE_FIRMWARE_AVAILABLE")
	case SystemEventTypeDHCPAck:
		return json.Marshal("DHCP_ACK")
	case SystemEventTypeDaemonSettingsChanged:
		return json.Marshal("DAEMON_SETTINGS_CHANGE")
	case SystemEventTypeFirmwareUploaded:
		return json.Marshal("FIRMWARE_UPLOADED")
	case SystemEventTypeLoggedError:
		return json.Marshal("LOGGED_ERROR")
	case SystemEventTypeMonitorRuleViolation:
		return json.Marshal("MONITOR_RULE_VIOLATION")
	case SystemEventTypeRRMStatus:
		return json.Marshal("RRM_STATUS_DATA")
	case SystemEventTypeRadarExportUpdate:
		return json.Marshal("RADAR_EXPORT_UPDATE")
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
	return nil, errors.New("Invalid value of SystemEventType: " + string(*self))
}

func (self *SystemEventType) GetBSON() (interface{}, error) {
	switch *self {
	case SystemEventTypeAny:
		return "+", nil
	case SystemEventTypeCPEConnected:
		return "CPE_CONNECTED", nil
	case SystemEventTypeCPEDisconnected:
		return "CPE_DISCONNECTED", nil
	case SystemEventTypeCPEInterfaceState:
		return "CPE_INTERFACE_STATE", nil
	case SystemEventTypeClientAuthorization:
		return "CLIENT_AUTHORIZATION", nil
	case SystemEventTypeClientConnected:
		return "CLIENT_CONNECTED", nil
	case SystemEventTypeClientDisconnected:
		return "CLIENT_DISCONNECTED", nil
	case SystemEventTypeCpeFirmwareAvailable:
		return "CPE_FIRMWARE_AVAILABLE", nil
	case SystemEventTypeDHCPAck:
		return "DHCP_ACK", nil
	case SystemEventTypeDaemonSettingsChanged:
		return "DAEMON_SETTINGS_CHANGE", nil
	case SystemEventTypeFirmwareUploaded:
		return "FIRMWARE_UPLOADED", nil
	case SystemEventTypeLoggedError:
		return "LOGGED_ERROR", nil
	case SystemEventTypeMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION", nil
	case SystemEventTypeRRMStatus:
		return "RRM_STATUS_DATA", nil
	case SystemEventTypeRadarExportUpdate:
		return "RADAR_EXPORT_UPDATE", nil
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
	return nil, errors.New("Invalid value of SystemEventType: " + string(*self))
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
	case "CPE_CONNECTED":
		*self = SystemEventTypeCPEConnected
		return nil
	case "CPE_DISCONNECTED":
		*self = SystemEventTypeCPEDisconnected
		return nil
	case "CPE_INTERFACE_STATE":
		*self = SystemEventTypeCPEInterfaceState
		return nil
	case "CLIENT_AUTHORIZATION":
		*self = SystemEventTypeClientAuthorization
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
	case "DHCP_ACK":
		*self = SystemEventTypeDHCPAck
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		*self = SystemEventTypeDaemonSettingsChanged
		return nil
	case "FIRMWARE_UPLOADED":
		*self = SystemEventTypeFirmwareUploaded
		return nil
	case "LOGGED_ERROR":
		*self = SystemEventTypeLoggedError
		return nil
	case "MONITOR_RULE_VIOLATION":
		*self = SystemEventTypeMonitorRuleViolation
		return nil
	case "RRM_STATUS_DATA":
		*self = SystemEventTypeRRMStatus
		return nil
	case "RADAR_EXPORT_UPDATE":
		*self = SystemEventTypeRadarExportUpdate
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
	return errors.New("Unknown SystemEventType: " + s)
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
	case "CPE_CONNECTED":
		*self = SystemEventTypeCPEConnected
		return nil
	case "CPE_DISCONNECTED":
		*self = SystemEventTypeCPEDisconnected
		return nil
	case "CPE_INTERFACE_STATE":
		*self = SystemEventTypeCPEInterfaceState
		return nil
	case "CLIENT_AUTHORIZATION":
		*self = SystemEventTypeClientAuthorization
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
	case "DHCP_ACK":
		*self = SystemEventTypeDHCPAck
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		*self = SystemEventTypeDaemonSettingsChanged
		return nil
	case "FIRMWARE_UPLOADED":
		*self = SystemEventTypeFirmwareUploaded
		return nil
	case "LOGGED_ERROR":
		*self = SystemEventTypeLoggedError
		return nil
	case "MONITOR_RULE_VIOLATION":
		*self = SystemEventTypeMonitorRuleViolation
		return nil
	case "RRM_STATUS_DATA":
		*self = SystemEventTypeRRMStatus
		return nil
	case "RADAR_EXPORT_UPDATE":
		*self = SystemEventTypeRadarExportUpdate
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
	return errors.New("Unknown SystemEventType: " + s)
}

type TunManagerRPC string

const TunManagerRPCCreateL2TunnelSession TunManagerRPC = "CreateL2TunnelSession"
const TunManagerRPCDeleteL2TunnelSession TunManagerRPC = "DeleteL2TunnelSession"

func (self TunManagerRPC) GetPtr() *TunManagerRPC { var v = self; return &v }

func (self TunManagerRPC) String() string {
	switch self {
	case TunManagerRPCCreateL2TunnelSession:
		return "CreateL2TunnelSession"
	case TunManagerRPCDeleteL2TunnelSession:
		return "DeleteL2TunnelSession"
	}
	panic(errors.New("Invalid value of TunManagerRPC: " + string(self)))
}

func (self *TunManagerRPC) MarshalJSON() ([]byte, error) {
	switch *self {
	case TunManagerRPCCreateL2TunnelSession:
		return json.Marshal("CreateL2TunnelSession")
	case TunManagerRPCDeleteL2TunnelSession:
		return json.Marshal("DeleteL2TunnelSession")
	}
	return nil, errors.New("Invalid value of TunManagerRPC: " + string(*self))
}

func (self *TunManagerRPC) GetBSON() (interface{}, error) {
	switch *self {
	case TunManagerRPCCreateL2TunnelSession:
		return "CreateL2TunnelSession", nil
	case TunManagerRPCDeleteL2TunnelSession:
		return "DeleteL2TunnelSession", nil
	}
	return nil, errors.New("Invalid value of TunManagerRPC: " + string(*self))
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
	return errors.New("Unknown TunManagerRPC: " + s)
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
	return errors.New("Unknown TunManagerRPC: " + s)
}

type WMMAccessCategory string

const WMMAccessCategoryBackground WMMAccessCategory = "BK"
const WMMAccessCategoryBestEffort WMMAccessCategory = "BE"
const WMMAccessCategoryVideo WMMAccessCategory = "VI"
const WMMAccessCategoryVoice WMMAccessCategory = "VO"

func (self WMMAccessCategory) GetPtr() *WMMAccessCategory { var v = self; return &v }

func (self WMMAccessCategory) String() string {
	switch self {
	case WMMAccessCategoryBackground:
		return "BK"
	case WMMAccessCategoryBestEffort:
		return "BE"
	case WMMAccessCategoryVideo:
		return "VI"
	case WMMAccessCategoryVoice:
		return "VO"
	}
	if len(self) == 0 {
		return "BK"
	}
	panic(errors.New("Invalid value of WMMAccessCategory: " + string(self)))
}

func (self *WMMAccessCategory) MarshalJSON() ([]byte, error) {
	switch *self {
	case WMMAccessCategoryBackground:
		return json.Marshal("BK")
	case WMMAccessCategoryBestEffort:
		return json.Marshal("BE")
	case WMMAccessCategoryVideo:
		return json.Marshal("VI")
	case WMMAccessCategoryVoice:
		return json.Marshal("VO")
	}
	if len(*self) == 0 {
		return json.Marshal("BK")
	}
	return nil, errors.New("Invalid value of WMMAccessCategory: " + string(*self))
}

func (self *WMMAccessCategory) GetBSON() (interface{}, error) {
	switch *self {
	case WMMAccessCategoryBackground:
		return "BK", nil
	case WMMAccessCategoryBestEffort:
		return "BE", nil
	case WMMAccessCategoryVideo:
		return "VI", nil
	case WMMAccessCategoryVoice:
		return "VO", nil
	}
	if len(*self) == 0 {
		return "BK", nil
	}
	return nil, errors.New("Invalid value of WMMAccessCategory: " + string(*self))
}

func (self *WMMAccessCategory) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "BK":
		*self = WMMAccessCategoryBackground
		return nil
	case "BE":
		*self = WMMAccessCategoryBestEffort
		return nil
	case "VI":
		*self = WMMAccessCategoryVideo
		return nil
	case "VO":
		*self = WMMAccessCategoryVoice
		return nil
	}
	if len(s) == 0 {
		*self = WMMAccessCategoryBackground
		return nil
	}
	return errors.New("Unknown WMMAccessCategory: " + s)
}

func (self *WMMAccessCategory) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "BK":
		*self = WMMAccessCategoryBackground
		return nil
	case "BE":
		*self = WMMAccessCategoryBestEffort
		return nil
	case "VI":
		*self = WMMAccessCategoryVideo
		return nil
	case "VO":
		*self = WMMAccessCategoryVoice
		return nil
	}
	if len(s) == 0 {
		*self = WMMAccessCategoryBackground
		return nil
	}
	return errors.New("Unknown WMMAccessCategory: " + s)
}

type WirelessClientState string

const WirelessClientStateConnected WirelessClientState = "CONNECTED"
const WirelessClientStateDisconnected WirelessClientState = "DISCONNECTED"

func (self WirelessClientState) GetPtr() *WirelessClientState { var v = self; return &v }

func (self WirelessClientState) String() string {
	switch self {
	case WirelessClientStateConnected:
		return "CONNECTED"
	case WirelessClientStateDisconnected:
		return "DISCONNECTED"
	}
	panic(errors.New("Invalid value of WirelessClientState: " + string(self)))
}

func (self *WirelessClientState) MarshalJSON() ([]byte, error) {
	switch *self {
	case WirelessClientStateConnected:
		return json.Marshal("CONNECTED")
	case WirelessClientStateDisconnected:
		return json.Marshal("DISCONNECTED")
	}
	return nil, errors.New("Invalid value of WirelessClientState: " + string(*self))
}

func (self *WirelessClientState) GetBSON() (interface{}, error) {
	switch *self {
	case WirelessClientStateConnected:
		return "CONNECTED", nil
	case WirelessClientStateDisconnected:
		return "DISCONNECTED", nil
	}
	return nil, errors.New("Invalid value of WirelessClientState: " + string(*self))
}

func (self *WirelessClientState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*self = WirelessClientStateConnected
		return nil
	case "DISCONNECTED":
		*self = WirelessClientStateDisconnected
		return nil
	}
	return errors.New("Unknown WirelessClientState: " + s)
}

func (self *WirelessClientState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*self = WirelessClientStateConnected
		return nil
	case "DISCONNECTED":
		*self = WirelessClientStateDisconnected
		return nil
	}
	return errors.New("Unknown WirelessClientState: " + s)
}

type WirelessClientType string

const WirelessClientTypeCamera WirelessClientType = "camera"
const WirelessClientTypeOther WirelessClientType = "other"
const WirelessClientTypeWired WirelessClientType = "wired"

func (self WirelessClientType) GetPtr() *WirelessClientType { var v = self; return &v }

func (self WirelessClientType) String() string {
	switch self {
	case WirelessClientTypeCamera:
		return "camera"
	case WirelessClientTypeOther:
		return "other"
	case WirelessClientTypeWired:
		return "wired"
	}
	if len(self) == 0 {
		return "other"
	}
	panic(errors.New("Invalid value of WirelessClientType: " + string(self)))
}

func (self *WirelessClientType) MarshalJSON() ([]byte, error) {
	switch *self {
	case WirelessClientTypeCamera:
		return json.Marshal("camera")
	case WirelessClientTypeOther:
		return json.Marshal("other")
	case WirelessClientTypeWired:
		return json.Marshal("wired")
	}
	if len(*self) == 0 {
		return json.Marshal("other")
	}
	return nil, errors.New("Invalid value of WirelessClientType: " + string(*self))
}

func (self *WirelessClientType) GetBSON() (interface{}, error) {
	switch *self {
	case WirelessClientTypeCamera:
		return "camera", nil
	case WirelessClientTypeOther:
		return "other", nil
	case WirelessClientTypeWired:
		return "wired", nil
	}
	if len(*self) == 0 {
		return "other", nil
	}
	return nil, errors.New("Invalid value of WirelessClientType: " + string(*self))
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
	case "wired":
		*self = WirelessClientTypeWired
		return nil
	}
	if len(s) == 0 {
		*self = WirelessClientTypeOther
		return nil
	}
	return errors.New("Unknown WirelessClientType: " + s)
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
	case "wired":
		*self = WirelessClientTypeWired
		return nil
	}
	if len(s) == 0 {
		*self = WirelessClientTypeOther
		return nil
	}
	return errors.New("Unknown WirelessClientType: " + s)
}
