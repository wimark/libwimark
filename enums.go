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

func (en CPEAgentStatusType) GetPtr() *CPEAgentStatusType { var v = en; return &v }

func (en CPEAgentStatusType) String() string {
	switch en {
	case CPEAgentStatusTypeException:
		return "exception"
	case CPEAgentStatusTypeSuccess:
		return "success"
	case CPEAgentStatusTypeSyntaxError:
		return "syntax"
	case CPEAgentStatusTypeUndefined:
		return "undefined"
	}
	if len(en) == 0 {
		return "undefined"
	}
	panic(errors.New("Invalid value of CPEAgentStatusType: " + string(en)))
}

func (en *CPEAgentStatusType) MarshalJSON() ([]byte, error) {
	switch *en {
	case CPEAgentStatusTypeException:
		return json.Marshal("exception")
	case CPEAgentStatusTypeSuccess:
		return json.Marshal("success")
	case CPEAgentStatusTypeSyntaxError:
		return json.Marshal("syntax")
	case CPEAgentStatusTypeUndefined:
		return json.Marshal("undefined")
	}
	if len(*en) == 0 {
		return json.Marshal("undefined")
	}
	return nil, errors.New("Invalid value of CPEAgentStatusType: " + string(*en))
}

func (en *CPEAgentStatusType) GetBSON() (interface{}, error) {
	switch *en {
	case CPEAgentStatusTypeException:
		return "exception", nil
	case CPEAgentStatusTypeSuccess:
		return "success", nil
	case CPEAgentStatusTypeSyntaxError:
		return "syntax", nil
	case CPEAgentStatusTypeUndefined:
		return "undefined", nil
	}
	if len(*en) == 0 {
		return "undefined", nil
	}
	return nil, errors.New("Invalid value of CPEAgentStatusType: " + string(*en))
}

func (en *CPEAgentStatusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "exception":
		*en = CPEAgentStatusTypeException
		return nil
	case "success":
		*en = CPEAgentStatusTypeSuccess
		return nil
	case "syntax":
		*en = CPEAgentStatusTypeSyntaxError
		return nil
	case "undefined":
		*en = CPEAgentStatusTypeUndefined
		return nil
	}
	if len(s) == 0 {
		*en = CPEAgentStatusTypeUndefined
		return nil
	}
	return errors.New("Unknown CPEAgentStatusType: " + s)
}

func (en *CPEAgentStatusType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "exception":
		*en = CPEAgentStatusTypeException
		return nil
	case "success":
		*en = CPEAgentStatusTypeSuccess
		return nil
	case "syntax":
		*en = CPEAgentStatusTypeSyntaxError
		return nil
	case "undefined":
		*en = CPEAgentStatusTypeUndefined
		return nil
	}
	if len(s) == 0 {
		*en = CPEAgentStatusTypeUndefined
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

func (en CPEInterfaceState) GetPtr() *CPEInterfaceState { var v = en; return &v }

func (en CPEInterfaceState) String() string {
	switch en {
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
	panic(errors.New("Invalid value of CPEInterfaceState: " + string(en)))
}

func (en *CPEInterfaceState) MarshalJSON() ([]byte, error) {
	switch *en {
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
	return nil, errors.New("Invalid value of CPEInterfaceState: " + string(*en))
}

func (en *CPEInterfaceState) GetBSON() (interface{}, error) {
	switch *en {
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
	return nil, errors.New("Invalid value of CPEInterfaceState: " + string(*en))
}

func (en *CPEInterfaceState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ACS":
		*en = CPEInterfaceStateACS
		return nil
	case "COUNTRY_UPDATE":
		*en = CPEInterfaceStateCountryUpdate
		return nil
	case "DFS":
		*en = CPEInterfaceStateDFS
		return nil
	case "DISABLED":
		*en = CPEInterfaceStateDisabled
		return nil
	case "ENABLED":
		*en = CPEInterfaceStateEnabled
		return nil
	case "HT_SCAN":
		*en = CPEInterfaceStateHtScan
		return nil
	case "TERMINATED":
		*en = CPEInterfaceStateTerminated
		return nil
	case "UNINITIALIZED":
		*en = CPEInterfaceStateUninitialized
		return nil
	case "UNKNOWN":
		*en = CPEInterfaceStateUnknown
		return nil
	}
	return errors.New("Unknown CPEInterfaceState: " + s)
}

func (en *CPEInterfaceState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ACS":
		*en = CPEInterfaceStateACS
		return nil
	case "COUNTRY_UPDATE":
		*en = CPEInterfaceStateCountryUpdate
		return nil
	case "DFS":
		*en = CPEInterfaceStateDFS
		return nil
	case "DISABLED":
		*en = CPEInterfaceStateDisabled
		return nil
	case "ENABLED":
		*en = CPEInterfaceStateEnabled
		return nil
	case "HT_SCAN":
		*en = CPEInterfaceStateHtScan
		return nil
	case "TERMINATED":
		*en = CPEInterfaceStateTerminated
		return nil
	case "UNINITIALIZED":
		*en = CPEInterfaceStateUninitialized
		return nil
	case "UNKNOWN":
		*en = CPEInterfaceStateUnknown
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

func (en ClientStatPacketType) GetPtr() *ClientStatPacketType { var v = en; return &v }

func (en ClientStatPacketType) String() string {
	switch en {
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
	panic(errors.New("Invalid value of ClientStatPacketType: " + string(en)))
}

func (en *ClientStatPacketType) MarshalJSON() ([]byte, error) {
	switch *en {
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
	return nil, errors.New("Invalid value of ClientStatPacketType: " + string(*en))
}

func (en *ClientStatPacketType) GetBSON() (interface{}, error) {
	switch *en {
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
	return nil, errors.New("Invalid value of ClientStatPacketType: " + string(*en))
}

func (en *ClientStatPacketType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "interim":
		*en = ClientStatPacketTypeInterim
		return nil
	case "Interim-Update":
		*en = ClientStatPacketTypeInterimOld
		return nil
	case "Accounting-Off":
		*en = ClientStatPacketTypeOffOld
		return nil
	case "Accounting-On":
		*en = ClientStatPacketTypeOnOld
		return nil
	case "start":
		*en = ClientStatPacketTypeStart
		return nil
	case "Start":
		*en = ClientStatPacketTypeStartOld
		return nil
	case "stop":
		*en = ClientStatPacketTypeStop
		return nil
	case "Stop":
		*en = ClientStatPacketTypeStopOld
		return nil
	}
	return errors.New("Unknown ClientStatPacketType: " + s)
}

func (en *ClientStatPacketType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "interim":
		*en = ClientStatPacketTypeInterim
		return nil
	case "Interim-Update":
		*en = ClientStatPacketTypeInterimOld
		return nil
	case "Accounting-Off":
		*en = ClientStatPacketTypeOffOld
		return nil
	case "Accounting-On":
		*en = ClientStatPacketTypeOnOld
		return nil
	case "start":
		*en = ClientStatPacketTypeStart
		return nil
	case "Start":
		*en = ClientStatPacketTypeStartOld
		return nil
	case "stop":
		*en = ClientStatPacketTypeStop
		return nil
	case "Stop":
		*en = ClientStatPacketTypeStopOld
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

func (en ConfigurationStatus) GetPtr() *ConfigurationStatus { var v = en; return &v }

func (en ConfigurationStatus) String() string {
	switch en {
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
	if len(en) == 0 {
		return "empty"
	}
	panic(errors.New("Invalid value of ConfigurationStatus: " + string(en)))
}

func (en *ConfigurationStatus) MarshalJSON() ([]byte, error) {
	switch *en {
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
	if len(*en) == 0 {
		return json.Marshal("empty")
	}
	return nil, errors.New("Invalid value of ConfigurationStatus: " + string(*en))
}

func (en *ConfigurationStatus) GetBSON() (interface{}, error) {
	switch *en {
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
	if len(*en) == 0 {
		return "empty", nil
	}
	return nil, errors.New("Invalid value of ConfigurationStatus: " + string(*en))
}

func (en *ConfigurationStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "pending":
		*en = ConfigurationStatusDontUse1
		return nil
	case "error":
		*en = ConfigurationStatusDontUse2
		return nil
	case "empty":
		*en = ConfigurationStatusEmpty
		return nil
	case "ok":
		*en = ConfigurationStatusOK
		return nil
	case "offline":
		*en = ConfigurationStatusOffline
		return nil
	case "rebooting":
		*en = ConfigurationStatusRebooting
		return nil
	case "updating":
		*en = ConfigurationStatusUpdating
		return nil
	case "upgrading":
		*en = ConfigurationStatusUpgrading
		return nil
	}
	if len(s) == 0 {
		*en = ConfigurationStatusEmpty
		return nil
	}
	return errors.New("Unknown ConfigurationStatus: " + s)
}

func (en *ConfigurationStatus) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "pending":
		*en = ConfigurationStatusDontUse1
		return nil
	case "error":
		*en = ConfigurationStatusDontUse2
		return nil
	case "empty":
		*en = ConfigurationStatusEmpty
		return nil
	case "ok":
		*en = ConfigurationStatusOK
		return nil
	case "offline":
		*en = ConfigurationStatusOffline
		return nil
	case "rebooting":
		*en = ConfigurationStatusRebooting
		return nil
	case "updating":
		*en = ConfigurationStatusUpdating
		return nil
	case "upgrading":
		*en = ConfigurationStatusUpgrading
		return nil
	}
	if len(s) == 0 {
		*en = ConfigurationStatusEmpty
		return nil
	}
	return errors.New("Unknown ConfigurationStatus: " + s)
}

type ConnectionModeType string

const ConnectionModeTypeModeAC ConnectionModeType = "ac"
const ConnectionModeTypeModeAX ConnectionModeType = "ax"
const ConnectionModeTypeModeLegacy ConnectionModeType = "legacy"
const ConnectionModeTypeModeN ConnectionModeType = "n"

func (en ConnectionModeType) GetPtr() *ConnectionModeType { var v = en; return &v }

func (en ConnectionModeType) String() string {
	switch en {
	case ConnectionModeTypeModeAC:
		return "ac"
	case ConnectionModeTypeModeAX:
		return "ax"
	case ConnectionModeTypeModeLegacy:
		return "legacy"
	case ConnectionModeTypeModeN:
		return "n"
	}
	if len(en) == 0 {
		return "legacy"
	}
	panic(errors.New("Invalid value of ConnectionModeType: " + string(en)))
}

func (en *ConnectionModeType) MarshalJSON() ([]byte, error) {
	switch *en {
	case ConnectionModeTypeModeAC:
		return json.Marshal("ac")
	case ConnectionModeTypeModeAX:
		return json.Marshal("ax")
	case ConnectionModeTypeModeLegacy:
		return json.Marshal("legacy")
	case ConnectionModeTypeModeN:
		return json.Marshal("n")
	}
	if len(*en) == 0 {
		return json.Marshal("legacy")
	}
	return nil, errors.New("Invalid value of ConnectionModeType: " + string(*en))
}

func (en *ConnectionModeType) GetBSON() (interface{}, error) {
	switch *en {
	case ConnectionModeTypeModeAC:
		return "ac", nil
	case ConnectionModeTypeModeAX:
		return "ax", nil
	case ConnectionModeTypeModeLegacy:
		return "legacy", nil
	case ConnectionModeTypeModeN:
		return "n", nil
	}
	if len(*en) == 0 {
		return "legacy", nil
	}
	return nil, errors.New("Invalid value of ConnectionModeType: " + string(*en))
}

func (en *ConnectionModeType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ac":
		*en = ConnectionModeTypeModeAC
		return nil
	case "ax":
		*en = ConnectionModeTypeModeAX
		return nil
	case "legacy":
		*en = ConnectionModeTypeModeLegacy
		return nil
	case "n":
		*en = ConnectionModeTypeModeN
		return nil
	}
	if len(s) == 0 {
		*en = ConnectionModeTypeModeLegacy
		return nil
	}
	return errors.New("Unknown ConnectionModeType: " + s)
}

func (en *ConnectionModeType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ac":
		*en = ConnectionModeTypeModeAC
		return nil
	case "ax":
		*en = ConnectionModeTypeModeAX
		return nil
	case "legacy":
		*en = ConnectionModeTypeModeLegacy
		return nil
	case "n":
		*en = ConnectionModeTypeModeN
		return nil
	}
	if len(s) == 0 {
		*en = ConnectionModeTypeModeLegacy
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

func (en ControllerStatusType) GetPtr() *ControllerStatusType { var v = en; return &v }

func (en ControllerStatusType) String() string {
	switch en {
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
	if len(en) == 0 {
		return "empty"
	}
	panic(errors.New("Invalid value of ControllerStatusType: " + string(en)))
}

func (en *ControllerStatusType) MarshalJSON() ([]byte, error) {
	switch *en {
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
	if len(*en) == 0 {
		return json.Marshal("empty")
	}
	return nil, errors.New("Invalid value of ControllerStatusType: " + string(*en))
}

func (en *ControllerStatusType) GetBSON() (interface{}, error) {
	switch *en {
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
	if len(*en) == 0 {
		return "empty", nil
	}
	return nil, errors.New("Invalid value of ControllerStatusType: " + string(*en))
}

func (en *ControllerStatusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "connected":
		*en = ControllerStatusTypeConnected
		return nil
	case "disconnected":
		*en = ControllerStatusTypeDisconnected
		return nil
	case "empty":
		*en = ControllerStatusTypeEmpty
		return nil
	case "error":
		*en = ControllerStatusTypeError
		return nil
	case "provision":
		*en = ControllerStatusTypeProvisioning
		return nil
	case "updating":
		*en = ControllerStatusTypeUpdating
		return nil
	}
	if len(s) == 0 {
		*en = ControllerStatusTypeEmpty
		return nil
	}
	return errors.New("Unknown ControllerStatusType: " + s)
}

func (en *ControllerStatusType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "connected":
		*en = ControllerStatusTypeConnected
		return nil
	case "disconnected":
		*en = ControllerStatusTypeDisconnected
		return nil
	case "empty":
		*en = ControllerStatusTypeEmpty
		return nil
	case "error":
		*en = ControllerStatusTypeError
		return nil
	case "provision":
		*en = ControllerStatusTypeProvisioning
		return nil
	case "updating":
		*en = ControllerStatusTypeUpdating
		return nil
	}
	if len(s) == 0 {
		*en = ControllerStatusTypeEmpty
		return nil
	}
	return errors.New("Unknown ControllerStatusType: " + s)
}

type FirewallDirection string

const FirewallDirectionAny FirewallDirection = "ANY"
const FirewallDirectionIn FirewallDirection = "IN"
const FirewallDirectionOut FirewallDirection = "OUT"

func (en FirewallDirection) GetPtr() *FirewallDirection { var v = en; return &v }

func (en FirewallDirection) String() string {
	switch en {
	case FirewallDirectionAny:
		return "ANY"
	case FirewallDirectionIn:
		return "IN"
	case FirewallDirectionOut:
		return "OUT"
	}
	if len(en) == 0 {
		return "ANY"
	}
	panic(errors.New("Invalid value of FirewallDirection: " + string(en)))
}

func (en *FirewallDirection) MarshalJSON() ([]byte, error) {
	switch *en {
	case FirewallDirectionAny:
		return json.Marshal("ANY")
	case FirewallDirectionIn:
		return json.Marshal("IN")
	case FirewallDirectionOut:
		return json.Marshal("OUT")
	}
	if len(*en) == 0 {
		return json.Marshal("ANY")
	}
	return nil, errors.New("Invalid value of FirewallDirection: " + string(*en))
}

func (en *FirewallDirection) GetBSON() (interface{}, error) {
	switch *en {
	case FirewallDirectionAny:
		return "ANY", nil
	case FirewallDirectionIn:
		return "IN", nil
	case FirewallDirectionOut:
		return "OUT", nil
	}
	if len(*en) == 0 {
		return "ANY", nil
	}
	return nil, errors.New("Invalid value of FirewallDirection: " + string(*en))
}

func (en *FirewallDirection) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ANY":
		*en = FirewallDirectionAny
		return nil
	case "IN":
		*en = FirewallDirectionIn
		return nil
	case "OUT":
		*en = FirewallDirectionOut
		return nil
	}
	if len(s) == 0 {
		*en = FirewallDirectionAny
		return nil
	}
	return errors.New("Unknown FirewallDirection: " + s)
}

func (en *FirewallDirection) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ANY":
		*en = FirewallDirectionAny
		return nil
	case "IN":
		*en = FirewallDirectionIn
		return nil
	case "OUT":
		*en = FirewallDirectionOut
		return nil
	}
	if len(s) == 0 {
		*en = FirewallDirectionAny
		return nil
	}
	return errors.New("Unknown FirewallDirection: " + s)
}

type FirewallPolicy string

const FirewallPolicyAccept FirewallPolicy = "ACCEPT"
const FirewallPolicyDrop FirewallPolicy = "DROP"
const FirewallPolicyEmpty FirewallPolicy = ""
const FirewallPolicyReturn FirewallPolicy = "RETURN"

func (en FirewallPolicy) GetPtr() *FirewallPolicy { var v = en; return &v }

func (en FirewallPolicy) String() string {
	switch en {
	case FirewallPolicyAccept:
		return "ACCEPT"
	case FirewallPolicyDrop:
		return "DROP"
	case FirewallPolicyEmpty:
		return ""
	case FirewallPolicyReturn:
		return "RETURN"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of FirewallPolicy: " + string(en)))
}

func (en *FirewallPolicy) MarshalJSON() ([]byte, error) {
	switch *en {
	case FirewallPolicyAccept:
		return json.Marshal("ACCEPT")
	case FirewallPolicyDrop:
		return json.Marshal("DROP")
	case FirewallPolicyEmpty:
		return json.Marshal("")
	case FirewallPolicyReturn:
		return json.Marshal("RETURN")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of FirewallPolicy: " + string(*en))
}

func (en *FirewallPolicy) GetBSON() (interface{}, error) {
	switch *en {
	case FirewallPolicyAccept:
		return "ACCEPT", nil
	case FirewallPolicyDrop:
		return "DROP", nil
	case FirewallPolicyEmpty:
		return "", nil
	case FirewallPolicyReturn:
		return "RETURN", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of FirewallPolicy: " + string(*en))
}

func (en *FirewallPolicy) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ACCEPT":
		*en = FirewallPolicyAccept
		return nil
	case "DROP":
		*en = FirewallPolicyDrop
		return nil
	case "":
		*en = FirewallPolicyEmpty
		return nil
	case "RETURN":
		*en = FirewallPolicyReturn
		return nil
	}
	if len(s) == 0 {
		*en = FirewallPolicyEmpty
		return nil
	}
	return errors.New("Unknown FirewallPolicy: " + s)
}

func (en *FirewallPolicy) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ACCEPT":
		*en = FirewallPolicyAccept
		return nil
	case "DROP":
		*en = FirewallPolicyDrop
		return nil
	case "":
		*en = FirewallPolicyEmpty
		return nil
	case "RETURN":
		*en = FirewallPolicyReturn
		return nil
	}
	if len(s) == 0 {
		*en = FirewallPolicyEmpty
		return nil
	}
	return errors.New("Unknown FirewallPolicy: " + s)
}

type FirmwareUpdateMode string

const FirmwareUpdateModeCheck FirmwareUpdateMode = "check"
const FirmwareUpdateModeOff FirmwareUpdateMode = "off"
const FirmwareUpdateModeOn FirmwareUpdateMode = "on"

func (en FirmwareUpdateMode) GetPtr() *FirmwareUpdateMode { var v = en; return &v }

func (en FirmwareUpdateMode) String() string {
	switch en {
	case FirmwareUpdateModeCheck:
		return "check"
	case FirmwareUpdateModeOff:
		return "off"
	case FirmwareUpdateModeOn:
		return "on"
	}
	if len(en) == 0 {
		return "check"
	}
	panic(errors.New("Invalid value of FirmwareUpdateMode: " + string(en)))
}

func (en *FirmwareUpdateMode) MarshalJSON() ([]byte, error) {
	switch *en {
	case FirmwareUpdateModeCheck:
		return json.Marshal("check")
	case FirmwareUpdateModeOff:
		return json.Marshal("off")
	case FirmwareUpdateModeOn:
		return json.Marshal("on")
	}
	if len(*en) == 0 {
		return json.Marshal("check")
	}
	return nil, errors.New("Invalid value of FirmwareUpdateMode: " + string(*en))
}

func (en *FirmwareUpdateMode) GetBSON() (interface{}, error) {
	switch *en {
	case FirmwareUpdateModeCheck:
		return "check", nil
	case FirmwareUpdateModeOff:
		return "off", nil
	case FirmwareUpdateModeOn:
		return "on", nil
	}
	if len(*en) == 0 {
		return "check", nil
	}
	return nil, errors.New("Invalid value of FirmwareUpdateMode: " + string(*en))
}

func (en *FirmwareUpdateMode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "check":
		*en = FirmwareUpdateModeCheck
		return nil
	case "off":
		*en = FirmwareUpdateModeOff
		return nil
	case "on":
		*en = FirmwareUpdateModeOn
		return nil
	}
	if len(s) == 0 {
		*en = FirmwareUpdateModeCheck
		return nil
	}
	return errors.New("Unknown FirmwareUpdateMode: " + s)
}

func (en *FirmwareUpdateMode) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "check":
		*en = FirmwareUpdateModeCheck
		return nil
	case "off":
		*en = FirmwareUpdateModeOff
		return nil
	case "on":
		*en = FirmwareUpdateModeOn
		return nil
	}
	if len(s) == 0 {
		*en = FirmwareUpdateModeCheck
		return nil
	}
	return errors.New("Unknown FirmwareUpdateMode: " + s)
}

type L3Protocol string

const L3ProtocolEmpty L3Protocol = ""
const L3ProtocolIP L3Protocol = "ip"
const L3ProtocolIPv4 L3Protocol = "ipv4"
const L3ProtocolIPv6 L3Protocol = "ipv6"

func (en L3Protocol) GetPtr() *L3Protocol { var v = en; return &v }

func (en L3Protocol) String() string {
	switch en {
	case L3ProtocolEmpty:
		return ""
	case L3ProtocolIP:
		return "ip"
	case L3ProtocolIPv4:
		return "ipv4"
	case L3ProtocolIPv6:
		return "ipv6"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of L3Protocol: " + string(en)))
}

func (en *L3Protocol) MarshalJSON() ([]byte, error) {
	switch *en {
	case L3ProtocolEmpty:
		return json.Marshal("")
	case L3ProtocolIP:
		return json.Marshal("ip")
	case L3ProtocolIPv4:
		return json.Marshal("ipv4")
	case L3ProtocolIPv6:
		return json.Marshal("ipv6")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of L3Protocol: " + string(*en))
}

func (en *L3Protocol) GetBSON() (interface{}, error) {
	switch *en {
	case L3ProtocolEmpty:
		return "", nil
	case L3ProtocolIP:
		return "ip", nil
	case L3ProtocolIPv4:
		return "ipv4", nil
	case L3ProtocolIPv6:
		return "ipv6", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of L3Protocol: " + string(*en))
}

func (en *L3Protocol) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "":
		*en = L3ProtocolEmpty
		return nil
	case "ip":
		*en = L3ProtocolIP
		return nil
	case "ipv4":
		*en = L3ProtocolIPv4
		return nil
	case "ipv6":
		*en = L3ProtocolIPv6
		return nil
	}
	if len(s) == 0 {
		*en = L3ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L3Protocol: " + s)
}

func (en *L3Protocol) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "":
		*en = L3ProtocolEmpty
		return nil
	case "ip":
		*en = L3ProtocolIP
		return nil
	case "ipv4":
		*en = L3ProtocolIPv4
		return nil
	case "ipv6":
		*en = L3ProtocolIPv6
		return nil
	}
	if len(s) == 0 {
		*en = L3ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L3Protocol: " + s)
}

type L4Protocol string

const L4ProtocolEmpty L4Protocol = ""
const L4ProtocolTCP L4Protocol = "TCP"
const L4ProtocolUDP L4Protocol = "UDP"

func (en L4Protocol) GetPtr() *L4Protocol { var v = en; return &v }

func (en L4Protocol) String() string {
	switch en {
	case L4ProtocolEmpty:
		return ""
	case L4ProtocolTCP:
		return "TCP"
	case L4ProtocolUDP:
		return "UDP"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of L4Protocol: " + string(en)))
}

func (en *L4Protocol) MarshalJSON() ([]byte, error) {
	switch *en {
	case L4ProtocolEmpty:
		return json.Marshal("")
	case L4ProtocolTCP:
		return json.Marshal("TCP")
	case L4ProtocolUDP:
		return json.Marshal("UDP")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of L4Protocol: " + string(*en))
}

func (en *L4Protocol) GetBSON() (interface{}, error) {
	switch *en {
	case L4ProtocolEmpty:
		return "", nil
	case L4ProtocolTCP:
		return "TCP", nil
	case L4ProtocolUDP:
		return "UDP", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of L4Protocol: " + string(*en))
}

func (en *L4Protocol) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "":
		*en = L4ProtocolEmpty
		return nil
	case "TCP":
		*en = L4ProtocolTCP
		return nil
	case "UDP":
		*en = L4ProtocolUDP
		return nil
	}
	if len(s) == 0 {
		*en = L4ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L4Protocol: " + s)
}

func (en *L4Protocol) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "":
		*en = L4ProtocolEmpty
		return nil
	case "TCP":
		*en = L4ProtocolTCP
		return nil
	case "UDP":
		*en = L4ProtocolUDP
		return nil
	}
	if len(s) == 0 {
		*en = L4ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L4Protocol: " + s)
}

type MCSRequire string

const MCSRequireHT MCSRequire = "ht"
const MCSRequireOff MCSRequire = "off"
const MCSRequireVHT MCSRequire = "vht"

func (en MCSRequire) GetPtr() *MCSRequire { var v = en; return &v }

func (en MCSRequire) String() string {
	switch en {
	case MCSRequireHT:
		return "ht"
	case MCSRequireOff:
		return "off"
	case MCSRequireVHT:
		return "vht"
	}
	if len(en) == 0 {
		return "off"
	}
	panic(errors.New("Invalid value of MCSRequire: " + string(en)))
}

func (en *MCSRequire) MarshalJSON() ([]byte, error) {
	switch *en {
	case MCSRequireHT:
		return json.Marshal("ht")
	case MCSRequireOff:
		return json.Marshal("off")
	case MCSRequireVHT:
		return json.Marshal("vht")
	}
	if len(*en) == 0 {
		return json.Marshal("off")
	}
	return nil, errors.New("Invalid value of MCSRequire: " + string(*en))
}

func (en *MCSRequire) GetBSON() (interface{}, error) {
	switch *en {
	case MCSRequireHT:
		return "ht", nil
	case MCSRequireOff:
		return "off", nil
	case MCSRequireVHT:
		return "vht", nil
	}
	if len(*en) == 0 {
		return "off", nil
	}
	return nil, errors.New("Invalid value of MCSRequire: " + string(*en))
}

func (en *MCSRequire) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ht":
		*en = MCSRequireHT
		return nil
	case "off":
		*en = MCSRequireOff
		return nil
	case "vht":
		*en = MCSRequireVHT
		return nil
	}
	if len(s) == 0 {
		*en = MCSRequireOff
		return nil
	}
	return errors.New("Unknown MCSRequire: " + s)
}

func (en *MCSRequire) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ht":
		*en = MCSRequireHT
		return nil
	case "off":
		*en = MCSRequireOff
		return nil
	case "vht":
		*en = MCSRequireVHT
		return nil
	}
	if len(s) == 0 {
		*en = MCSRequireOff
		return nil
	}
	return errors.New("Unknown MCSRequire: " + s)
}

type MacFilterType string

const MacFilterTypeBlackList MacFilterType = "BlackList"
const MacFilterTypeNone MacFilterType = "None"
const MacFilterTypeWhiteList MacFilterType = "WhiteList"

func (en MacFilterType) GetPtr() *MacFilterType { var v = en; return &v }

func (en MacFilterType) String() string {
	switch en {
	case MacFilterTypeBlackList:
		return "BlackList"
	case MacFilterTypeNone:
		return "None"
	case MacFilterTypeWhiteList:
		return "WhiteList"
	}
	if len(en) == 0 {
		return "None"
	}
	panic(errors.New("Invalid value of MacFilterType: " + string(en)))
}

func (en *MacFilterType) MarshalJSON() ([]byte, error) {
	switch *en {
	case MacFilterTypeBlackList:
		return json.Marshal("BlackList")
	case MacFilterTypeNone:
		return json.Marshal("None")
	case MacFilterTypeWhiteList:
		return json.Marshal("WhiteList")
	}
	if len(*en) == 0 {
		return json.Marshal("None")
	}
	return nil, errors.New("Invalid value of MacFilterType: " + string(*en))
}

func (en *MacFilterType) GetBSON() (interface{}, error) {
	switch *en {
	case MacFilterTypeBlackList:
		return "BlackList", nil
	case MacFilterTypeNone:
		return "None", nil
	case MacFilterTypeWhiteList:
		return "WhiteList", nil
	}
	if len(*en) == 0 {
		return "None", nil
	}
	return nil, errors.New("Invalid value of MacFilterType: " + string(*en))
}

func (en *MacFilterType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "BlackList":
		*en = MacFilterTypeBlackList
		return nil
	case "None":
		*en = MacFilterTypeNone
		return nil
	case "WhiteList":
		*en = MacFilterTypeWhiteList
		return nil
	}
	if len(s) == 0 {
		*en = MacFilterTypeNone
		return nil
	}
	return errors.New("Unknown MacFilterType: " + s)
}

func (en *MacFilterType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "BlackList":
		*en = MacFilterTypeBlackList
		return nil
	case "None":
		*en = MacFilterTypeNone
		return nil
	case "WhiteList":
		*en = MacFilterTypeWhiteList
		return nil
	}
	if len(s) == 0 {
		*en = MacFilterTypeNone
		return nil
	}
	return errors.New("Unknown MacFilterType: " + s)
}

type Module string

const ModuleAC Module = "AC"
const ModuleAnalMW Module = "ANAL-MW"
const ModuleAny Module = "+"
const ModuleBackend Module = "BACKEND"
const ModuleCLI Module = "CLI"
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
const ModuleResampling Module = "RESAMPLING"
const ModuleSnmpWalker Module = "SNMP_WALKER"
const ModuleStat Module = "STAT"
const ModuleStatLBS Module = "STAT-LBS"
const ModuleTunManager Module = "TUN_MANAGER"

func (en Module) GetPtr() *Module { var v = en; return &v }

func (en Module) String() string {
	switch en {
	case ModuleAC:
		return "AC"
	case ModuleAnalMW:
		return "ANAL-MW"
	case ModuleAny:
		return "+"
	case ModuleBackend:
		return "BACKEND"
	case ModuleCLI:
		return "CLI"
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
	case ModuleResampling:
		return "RESAMPLING"
	case ModuleSnmpWalker:
		return "SNMP_WALKER"
	case ModuleStat:
		return "STAT"
	case ModuleStatLBS:
		return "STAT-LBS"
	case ModuleTunManager:
		return "TUN_MANAGER"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of Module: " + string(en)))
}

func (en *Module) MarshalJSON() ([]byte, error) {
	switch *en {
	case ModuleAC:
		return json.Marshal("AC")
	case ModuleAnalMW:
		return json.Marshal("ANAL-MW")
	case ModuleAny:
		return json.Marshal("+")
	case ModuleBackend:
		return json.Marshal("BACKEND")
	case ModuleCLI:
		return json.Marshal("CLI")
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
	case ModuleResampling:
		return json.Marshal("RESAMPLING")
	case ModuleSnmpWalker:
		return json.Marshal("SNMP_WALKER")
	case ModuleStat:
		return json.Marshal("STAT")
	case ModuleStatLBS:
		return json.Marshal("STAT-LBS")
	case ModuleTunManager:
		return json.Marshal("TUN_MANAGER")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of Module: " + string(*en))
}

func (en *Module) GetBSON() (interface{}, error) {
	switch *en {
	case ModuleAC:
		return "AC", nil
	case ModuleAnalMW:
		return "ANAL-MW", nil
	case ModuleAny:
		return "+", nil
	case ModuleBackend:
		return "BACKEND", nil
	case ModuleCLI:
		return "CLI", nil
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
	case ModuleResampling:
		return "RESAMPLING", nil
	case ModuleSnmpWalker:
		return "SNMP_WALKER", nil
	case ModuleStat:
		return "STAT", nil
	case ModuleStatLBS:
		return "STAT-LBS", nil
	case ModuleTunManager:
		return "TUN_MANAGER", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of Module: " + string(*en))
}

func (en *Module) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "AC":
		*en = ModuleAC
		return nil
	case "ANAL-MW":
		*en = ModuleAnalMW
		return nil
	case "+":
		*en = ModuleAny
		return nil
	case "BACKEND":
		*en = ModuleBackend
		return nil
	case "CLI":
		*en = ModuleCLI
		return nil
	case "CPE":
		*en = ModuleCPE
		return nil
	case "CPE_STAT":
		*en = ModuleCPEStat
		return nil
	case "CLIENT-DISTANCE":
		*en = ModuleClientDistance
		return nil
	case "CLIENT_STAT":
		*en = ModuleClientStat
		return nil
	case "CONFIG":
		*en = ModuleConfig
		return nil
	case "DB":
		*en = ModuleDB
		return nil
	case "DUMMY":
		*en = ModuleDummy
		return nil
	case "FW":
		*en = ModuleFW
		return nil
	case "LBS":
		*en = ModuleLBS
		return nil
	case "MQTT_LOG":
		*en = ModuleMQTTLog
		return nil
	case "MEDIATOR":
		*en = ModuleMediator
		return nil
	case "MONITOR":
		*en = ModuleMonitor
		return nil
	case "":
		*en = ModuleNone
		return nil
	case "PORTAL_BACKEND":
		*en = ModulePortalBack
		return nil
	case "RRM":
		*en = ModuleRRM
		return nil
	case "RADAR-EXPORT-MW":
		*en = ModuleRadarExportMW
		return nil
	case "RADAR-MW":
		*en = ModuleRadarMW
		return nil
	case "RADIUS_GATEWAY":
		*en = ModuleRadiusGw
		return nil
	case "REDIRECT":
		*en = ModuleRedirect
		return nil
	case "RESAMPLING":
		*en = ModuleResampling
		return nil
	case "SNMP_WALKER":
		*en = ModuleSnmpWalker
		return nil
	case "STAT":
		*en = ModuleStat
		return nil
	case "STAT-LBS":
		*en = ModuleStatLBS
		return nil
	case "TUN_MANAGER":
		*en = ModuleTunManager
		return nil
	}
	if len(s) == 0 {
		*en = ModuleNone
		return nil
	}
	return errors.New("Unknown Module: " + s)
}

func (en *Module) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "AC":
		*en = ModuleAC
		return nil
	case "ANAL-MW":
		*en = ModuleAnalMW
		return nil
	case "+":
		*en = ModuleAny
		return nil
	case "BACKEND":
		*en = ModuleBackend
		return nil
	case "CLI":
		*en = ModuleCLI
		return nil
	case "CPE":
		*en = ModuleCPE
		return nil
	case "CPE_STAT":
		*en = ModuleCPEStat
		return nil
	case "CLIENT-DISTANCE":
		*en = ModuleClientDistance
		return nil
	case "CLIENT_STAT":
		*en = ModuleClientStat
		return nil
	case "CONFIG":
		*en = ModuleConfig
		return nil
	case "DB":
		*en = ModuleDB
		return nil
	case "DUMMY":
		*en = ModuleDummy
		return nil
	case "FW":
		*en = ModuleFW
		return nil
	case "LBS":
		*en = ModuleLBS
		return nil
	case "MQTT_LOG":
		*en = ModuleMQTTLog
		return nil
	case "MEDIATOR":
		*en = ModuleMediator
		return nil
	case "MONITOR":
		*en = ModuleMonitor
		return nil
	case "":
		*en = ModuleNone
		return nil
	case "PORTAL_BACKEND":
		*en = ModulePortalBack
		return nil
	case "RRM":
		*en = ModuleRRM
		return nil
	case "RADAR-EXPORT-MW":
		*en = ModuleRadarExportMW
		return nil
	case "RADAR-MW":
		*en = ModuleRadarMW
		return nil
	case "RADIUS_GATEWAY":
		*en = ModuleRadiusGw
		return nil
	case "REDIRECT":
		*en = ModuleRedirect
		return nil
	case "RESAMPLING":
		*en = ModuleResampling
		return nil
	case "SNMP_WALKER":
		*en = ModuleSnmpWalker
		return nil
	case "STAT":
		*en = ModuleStat
		return nil
	case "STAT-LBS":
		*en = ModuleStatLBS
		return nil
	case "TUN_MANAGER":
		*en = ModuleTunManager
		return nil
	}
	if len(s) == 0 {
		*en = ModuleNone
		return nil
	}
	return errors.New("Unknown Module: " + s)
}

type NotifyType string

const NotifyTypeEmail NotifyType = "email"
const NotifyTypeEmpty NotifyType = ""
const NotifyTypeMqtt NotifyType = "mqtt"
const NotifyTypeTelegram NotifyType = "telegram"

func (en NotifyType) GetPtr() *NotifyType { var v = en; return &v }

func (en NotifyType) String() string {
	switch en {
	case NotifyTypeEmail:
		return "email"
	case NotifyTypeEmpty:
		return ""
	case NotifyTypeMqtt:
		return "mqtt"
	case NotifyTypeTelegram:
		return "telegram"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of NotifyType: " + string(en)))
}

func (en *NotifyType) MarshalJSON() ([]byte, error) {
	switch *en {
	case NotifyTypeEmail:
		return json.Marshal("email")
	case NotifyTypeEmpty:
		return json.Marshal("")
	case NotifyTypeMqtt:
		return json.Marshal("mqtt")
	case NotifyTypeTelegram:
		return json.Marshal("telegram")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of NotifyType: " + string(*en))
}

func (en *NotifyType) GetBSON() (interface{}, error) {
	switch *en {
	case NotifyTypeEmail:
		return "email", nil
	case NotifyTypeEmpty:
		return "", nil
	case NotifyTypeMqtt:
		return "mqtt", nil
	case NotifyTypeTelegram:
		return "telegram", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of NotifyType: " + string(*en))
}

func (en *NotifyType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "email":
		*en = NotifyTypeEmail
		return nil
	case "":
		*en = NotifyTypeEmpty
		return nil
	case "mqtt":
		*en = NotifyTypeMqtt
		return nil
	case "telegram":
		*en = NotifyTypeTelegram
		return nil
	}
	if len(s) == 0 {
		*en = NotifyTypeEmpty
		return nil
	}
	return errors.New("Unknown NotifyType: " + s)
}

func (en *NotifyType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "email":
		*en = NotifyTypeEmail
		return nil
	case "":
		*en = NotifyTypeEmpty
		return nil
	case "mqtt":
		*en = NotifyTypeMqtt
		return nil
	case "telegram":
		*en = NotifyTypeTelegram
		return nil
	}
	if len(s) == 0 {
		*en = NotifyTypeEmpty
		return nil
	}
	return errors.New("Unknown NotifyType: " + s)
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

func (en Operation) GetPtr() *Operation { var v = en; return &v }

func (en Operation) String() string {
	switch en {
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
	panic(errors.New("Invalid value of Operation: " + string(en)))
}

func (en *Operation) MarshalJSON() ([]byte, error) {
	switch *en {
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
	return nil, errors.New("Invalid value of Operation: " + string(*en))
}

func (en *Operation) GetBSON() (interface{}, error) {
	switch *en {
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
	return nil, errors.New("Invalid value of Operation: " + string(*en))
}

func (en *Operation) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		*en = OperationAny
		return nil
	case "STATUS":
		*en = OperationCPEStatus
		return nil
	case "C":
		*en = OperationCreate
		return nil
	case "D":
		*en = OperationDelete
		return nil
	case "JSONRPC":
		*en = OperationJSONRPC
		return nil
	case "LUA":
		*en = OperationLuaScript
		return nil
	case "R":
		*en = OperationRead
		return nil
	case "SH":
		*en = OperationSHScript
		return nil
	case "U":
		*en = OperationUpdate
		return nil
	}
	return errors.New("Unknown Operation: " + s)
}

func (en *Operation) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "+":
		*en = OperationAny
		return nil
	case "STATUS":
		*en = OperationCPEStatus
		return nil
	case "C":
		*en = OperationCreate
		return nil
	case "D":
		*en = OperationDelete
		return nil
	case "JSONRPC":
		*en = OperationJSONRPC
		return nil
	case "LUA":
		*en = OperationLuaScript
		return nil
	case "R":
		*en = OperationRead
		return nil
	case "SH":
		*en = OperationSHScript
		return nil
	case "U":
		*en = OperationUpdate
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

func (en PortalAuthType) GetPtr() *PortalAuthType { var v = en; return &v }

func (en PortalAuthType) String() string {
	switch en {
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
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of PortalAuthType: " + string(en)))
}

func (en *PortalAuthType) MarshalJSON() ([]byte, error) {
	switch *en {
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
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of PortalAuthType: " + string(*en))
}

func (en *PortalAuthType) GetBSON() (interface{}, error) {
	switch *en {
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
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of PortalAuthType: " + string(*en))
}

func (en *PortalAuthType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "external":
		*en = PortalAuthTypeExternal
		return nil
	case "":
		*en = PortalAuthTypeNone
		return nil
	case "oauth2":
		*en = PortalAuthTypeOAuth2
		return nil
	case "radius":
		*en = PortalAuthTypeRADIUS
		return nil
	case "sms":
		*en = PortalAuthTypeSMS
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthType: " + s)
}

func (en *PortalAuthType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "external":
		*en = PortalAuthTypeExternal
		return nil
	case "":
		*en = PortalAuthTypeNone
		return nil
	case "oauth2":
		*en = PortalAuthTypeOAuth2
		return nil
	case "radius":
		*en = PortalAuthTypeRADIUS
		return nil
	case "sms":
		*en = PortalAuthTypeSMS
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthType: " + s)
}

type PortalProfileType string

const PortalProfileTypeFree PortalProfileType = "free"
const PortalProfileTypePremium PortalProfileType = "premium"
const PortalProfileTypeSponsor PortalProfileType = "sponsor"

func (en PortalProfileType) GetPtr() *PortalProfileType { var v = en; return &v }

func (en PortalProfileType) String() string {
	switch en {
	case PortalProfileTypeFree:
		return "free"
	case PortalProfileTypePremium:
		return "premium"
	case PortalProfileTypeSponsor:
		return "sponsor"
	}
	if len(en) == 0 {
		return "free"
	}
	panic(errors.New("Invalid value of PortalProfileType: " + string(en)))
}

func (en *PortalProfileType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalProfileTypeFree:
		return json.Marshal("free")
	case PortalProfileTypePremium:
		return json.Marshal("premium")
	case PortalProfileTypeSponsor:
		return json.Marshal("sponsor")
	}
	if len(*en) == 0 {
		return json.Marshal("free")
	}
	return nil, errors.New("Invalid value of PortalProfileType: " + string(*en))
}

func (en *PortalProfileType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalProfileTypeFree:
		return "free", nil
	case PortalProfileTypePremium:
		return "premium", nil
	case PortalProfileTypeSponsor:
		return "sponsor", nil
	}
	if len(*en) == 0 {
		return "free", nil
	}
	return nil, errors.New("Invalid value of PortalProfileType: " + string(*en))
}

func (en *PortalProfileType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "free":
		*en = PortalProfileTypeFree
		return nil
	case "premium":
		*en = PortalProfileTypePremium
		return nil
	case "sponsor":
		*en = PortalProfileTypeSponsor
		return nil
	}
	if len(s) == 0 {
		*en = PortalProfileTypeFree
		return nil
	}
	return errors.New("Unknown PortalProfileType: " + s)
}

func (en *PortalProfileType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "free":
		*en = PortalProfileTypeFree
		return nil
	case "premium":
		*en = PortalProfileTypePremium
		return nil
	case "sponsor":
		*en = PortalProfileTypeSponsor
		return nil
	}
	if len(s) == 0 {
		*en = PortalProfileTypeFree
		return nil
	}
	return errors.New("Unknown PortalProfileType: " + s)
}

type RRMAlgoType string

const RRMAlgoTypeBlind RRMAlgoType = "Blind"
const RRMAlgoTypeDummy RRMAlgoType = "Dummy"
const RRMAlgoTypeGreed RRMAlgoType = "Greed"

func (en RRMAlgoType) GetPtr() *RRMAlgoType { var v = en; return &v }

func (en RRMAlgoType) String() string {
	switch en {
	case RRMAlgoTypeBlind:
		return "Blind"
	case RRMAlgoTypeDummy:
		return "Dummy"
	case RRMAlgoTypeGreed:
		return "Greed"
	}
	if len(en) == 0 {
		return "Greed"
	}
	panic(errors.New("Invalid value of RRMAlgoType: " + string(en)))
}

func (en *RRMAlgoType) MarshalJSON() ([]byte, error) {
	switch *en {
	case RRMAlgoTypeBlind:
		return json.Marshal("Blind")
	case RRMAlgoTypeDummy:
		return json.Marshal("Dummy")
	case RRMAlgoTypeGreed:
		return json.Marshal("Greed")
	}
	if len(*en) == 0 {
		return json.Marshal("Greed")
	}
	return nil, errors.New("Invalid value of RRMAlgoType: " + string(*en))
}

func (en *RRMAlgoType) GetBSON() (interface{}, error) {
	switch *en {
	case RRMAlgoTypeBlind:
		return "Blind", nil
	case RRMAlgoTypeDummy:
		return "Dummy", nil
	case RRMAlgoTypeGreed:
		return "Greed", nil
	}
	if len(*en) == 0 {
		return "Greed", nil
	}
	return nil, errors.New("Invalid value of RRMAlgoType: " + string(*en))
}

func (en *RRMAlgoType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Blind":
		*en = RRMAlgoTypeBlind
		return nil
	case "Dummy":
		*en = RRMAlgoTypeDummy
		return nil
	case "Greed":
		*en = RRMAlgoTypeGreed
		return nil
	}
	if len(s) == 0 {
		*en = RRMAlgoTypeGreed
		return nil
	}
	return errors.New("Unknown RRMAlgoType: " + s)
}

func (en *RRMAlgoType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "Blind":
		*en = RRMAlgoTypeBlind
		return nil
	case "Dummy":
		*en = RRMAlgoTypeDummy
		return nil
	case "Greed":
		*en = RRMAlgoTypeGreed
		return nil
	}
	if len(s) == 0 {
		*en = RRMAlgoTypeGreed
		return nil
	}
	return errors.New("Unknown RRMAlgoType: " + s)
}

type RadarExportFilter string

const RadarExportFilterAll RadarExportFilter = "all"
const RadarExportFilterNew RadarExportFilter = "new"
const RadarExportFilterReturn RadarExportFilter = "return"

func (en RadarExportFilter) GetPtr() *RadarExportFilter { var v = en; return &v }

func (en RadarExportFilter) String() string {
	switch en {
	case RadarExportFilterAll:
		return "all"
	case RadarExportFilterNew:
		return "new"
	case RadarExportFilterReturn:
		return "return"
	}
	if len(en) == 0 {
		return "all"
	}
	panic(errors.New("Invalid value of RadarExportFilter: " + string(en)))
}

func (en *RadarExportFilter) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadarExportFilterAll:
		return json.Marshal("all")
	case RadarExportFilterNew:
		return json.Marshal("new")
	case RadarExportFilterReturn:
		return json.Marshal("return")
	}
	if len(*en) == 0 {
		return json.Marshal("all")
	}
	return nil, errors.New("Invalid value of RadarExportFilter: " + string(*en))
}

func (en *RadarExportFilter) GetBSON() (interface{}, error) {
	switch *en {
	case RadarExportFilterAll:
		return "all", nil
	case RadarExportFilterNew:
		return "new", nil
	case RadarExportFilterReturn:
		return "return", nil
	}
	if len(*en) == 0 {
		return "all", nil
	}
	return nil, errors.New("Invalid value of RadarExportFilter: " + string(*en))
}

func (en *RadarExportFilter) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "all":
		*en = RadarExportFilterAll
		return nil
	case "new":
		*en = RadarExportFilterNew
		return nil
	case "return":
		*en = RadarExportFilterReturn
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportFilterAll
		return nil
	}
	return errors.New("Unknown RadarExportFilter: " + s)
}

func (en *RadarExportFilter) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "all":
		*en = RadarExportFilterAll
		return nil
	case "new":
		*en = RadarExportFilterNew
		return nil
	case "return":
		*en = RadarExportFilterReturn
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportFilterAll
		return nil
	}
	return errors.New("Unknown RadarExportFilter: " + s)
}

type RadarExportFormat string

const RadarExportFormatCSV RadarExportFormat = "csv"
const RadarExportFormatJson RadarExportFormat = "json"
const RadarExportFormatTxt RadarExportFormat = "txt"

func (en RadarExportFormat) GetPtr() *RadarExportFormat { var v = en; return &v }

func (en RadarExportFormat) String() string {
	switch en {
	case RadarExportFormatCSV:
		return "csv"
	case RadarExportFormatJson:
		return "json"
	case RadarExportFormatTxt:
		return "txt"
	}
	if len(en) == 0 {
		return "csv"
	}
	panic(errors.New("Invalid value of RadarExportFormat: " + string(en)))
}

func (en *RadarExportFormat) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadarExportFormatCSV:
		return json.Marshal("csv")
	case RadarExportFormatJson:
		return json.Marshal("json")
	case RadarExportFormatTxt:
		return json.Marshal("txt")
	}
	if len(*en) == 0 {
		return json.Marshal("csv")
	}
	return nil, errors.New("Invalid value of RadarExportFormat: " + string(*en))
}

func (en *RadarExportFormat) GetBSON() (interface{}, error) {
	switch *en {
	case RadarExportFormatCSV:
		return "csv", nil
	case RadarExportFormatJson:
		return "json", nil
	case RadarExportFormatTxt:
		return "txt", nil
	}
	if len(*en) == 0 {
		return "csv", nil
	}
	return nil, errors.New("Invalid value of RadarExportFormat: " + string(*en))
}

func (en *RadarExportFormat) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*en = RadarExportFormatCSV
		return nil
	case "json":
		*en = RadarExportFormatJson
		return nil
	case "txt":
		*en = RadarExportFormatTxt
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportFormatCSV
		return nil
	}
	return errors.New("Unknown RadarExportFormat: " + s)
}

func (en *RadarExportFormat) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*en = RadarExportFormatCSV
		return nil
	case "json":
		*en = RadarExportFormatJson
		return nil
	case "txt":
		*en = RadarExportFormatTxt
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportFormatCSV
		return nil
	}
	return errors.New("Unknown RadarExportFormat: " + s)
}

type RadarExportMacs string

const RadarExportMacsAll RadarExportMacs = "all"
const RadarExportMacsFake RadarExportMacs = "fake"
const RadarExportMacsReal RadarExportMacs = "real"

func (en RadarExportMacs) GetPtr() *RadarExportMacs { var v = en; return &v }

func (en RadarExportMacs) String() string {
	switch en {
	case RadarExportMacsAll:
		return "all"
	case RadarExportMacsFake:
		return "fake"
	case RadarExportMacsReal:
		return "real"
	}
	if len(en) == 0 {
		return "all"
	}
	panic(errors.New("Invalid value of RadarExportMacs: " + string(en)))
}

func (en *RadarExportMacs) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadarExportMacsAll:
		return json.Marshal("all")
	case RadarExportMacsFake:
		return json.Marshal("fake")
	case RadarExportMacsReal:
		return json.Marshal("real")
	}
	if len(*en) == 0 {
		return json.Marshal("all")
	}
	return nil, errors.New("Invalid value of RadarExportMacs: " + string(*en))
}

func (en *RadarExportMacs) GetBSON() (interface{}, error) {
	switch *en {
	case RadarExportMacsAll:
		return "all", nil
	case RadarExportMacsFake:
		return "fake", nil
	case RadarExportMacsReal:
		return "real", nil
	}
	if len(*en) == 0 {
		return "all", nil
	}
	return nil, errors.New("Invalid value of RadarExportMacs: " + string(*en))
}

func (en *RadarExportMacs) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "all":
		*en = RadarExportMacsAll
		return nil
	case "fake":
		*en = RadarExportMacsFake
		return nil
	case "real":
		*en = RadarExportMacsReal
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportMacsAll
		return nil
	}
	return errors.New("Unknown RadarExportMacs: " + s)
}

func (en *RadarExportMacs) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "all":
		*en = RadarExportMacsAll
		return nil
	case "fake":
		*en = RadarExportMacsFake
		return nil
	case "real":
		*en = RadarExportMacsReal
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportMacsAll
		return nil
	}
	return errors.New("Unknown RadarExportMacs: " + s)
}

type RadarExportStatus string

const RadarExportStatusCreated RadarExportStatus = "created"
const RadarExportStatusFinished RadarExportStatus = "finished"
const RadarExportStatusRunning RadarExportStatus = "running"
const RadarExportStatusUpdated RadarExportStatus = "updated"

func (en RadarExportStatus) GetPtr() *RadarExportStatus { var v = en; return &v }

func (en RadarExportStatus) String() string {
	switch en {
	case RadarExportStatusCreated:
		return "created"
	case RadarExportStatusFinished:
		return "finished"
	case RadarExportStatusRunning:
		return "running"
	case RadarExportStatusUpdated:
		return "updated"
	}
	if len(en) == 0 {
		return "created"
	}
	panic(errors.New("Invalid value of RadarExportStatus: " + string(en)))
}

func (en *RadarExportStatus) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadarExportStatusCreated:
		return json.Marshal("created")
	case RadarExportStatusFinished:
		return json.Marshal("finished")
	case RadarExportStatusRunning:
		return json.Marshal("running")
	case RadarExportStatusUpdated:
		return json.Marshal("updated")
	}
	if len(*en) == 0 {
		return json.Marshal("created")
	}
	return nil, errors.New("Invalid value of RadarExportStatus: " + string(*en))
}

func (en *RadarExportStatus) GetBSON() (interface{}, error) {
	switch *en {
	case RadarExportStatusCreated:
		return "created", nil
	case RadarExportStatusFinished:
		return "finished", nil
	case RadarExportStatusRunning:
		return "running", nil
	case RadarExportStatusUpdated:
		return "updated", nil
	}
	if len(*en) == 0 {
		return "created", nil
	}
	return nil, errors.New("Invalid value of RadarExportStatus: " + string(*en))
}

func (en *RadarExportStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "created":
		*en = RadarExportStatusCreated
		return nil
	case "finished":
		*en = RadarExportStatusFinished
		return nil
	case "running":
		*en = RadarExportStatusRunning
		return nil
	case "updated":
		*en = RadarExportStatusUpdated
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportStatusCreated
		return nil
	}
	return errors.New("Unknown RadarExportStatus: " + s)
}

func (en *RadarExportStatus) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "created":
		*en = RadarExportStatusCreated
		return nil
	case "finished":
		*en = RadarExportStatusFinished
		return nil
	case "running":
		*en = RadarExportStatusRunning
		return nil
	case "updated":
		*en = RadarExportStatusUpdated
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportStatusCreated
		return nil
	}
	return errors.New("Unknown RadarExportStatus: " + s)
}

type RadarExportType string

const RadarExportTypeBeePro RadarExportType = "beepro"
const RadarExportTypeEmail RadarExportType = "email"
const RadarExportTypeExternal RadarExportType = "external"
const RadarExportTypeMytarget RadarExportType = "mytarget"
const RadarExportTypeYandex RadarExportType = "yandex"

func (en RadarExportType) GetPtr() *RadarExportType { var v = en; return &v }

func (en RadarExportType) String() string {
	switch en {
	case RadarExportTypeBeePro:
		return "beepro"
	case RadarExportTypeEmail:
		return "email"
	case RadarExportTypeExternal:
		return "external"
	case RadarExportTypeMytarget:
		return "mytarget"
	case RadarExportTypeYandex:
		return "yandex"
	}
	if len(en) == 0 {
		return "email"
	}
	panic(errors.New("Invalid value of RadarExportType: " + string(en)))
}

func (en *RadarExportType) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadarExportTypeBeePro:
		return json.Marshal("beepro")
	case RadarExportTypeEmail:
		return json.Marshal("email")
	case RadarExportTypeExternal:
		return json.Marshal("external")
	case RadarExportTypeMytarget:
		return json.Marshal("mytarget")
	case RadarExportTypeYandex:
		return json.Marshal("yandex")
	}
	if len(*en) == 0 {
		return json.Marshal("email")
	}
	return nil, errors.New("Invalid value of RadarExportType: " + string(*en))
}

func (en *RadarExportType) GetBSON() (interface{}, error) {
	switch *en {
	case RadarExportTypeBeePro:
		return "beepro", nil
	case RadarExportTypeEmail:
		return "email", nil
	case RadarExportTypeExternal:
		return "external", nil
	case RadarExportTypeMytarget:
		return "mytarget", nil
	case RadarExportTypeYandex:
		return "yandex", nil
	}
	if len(*en) == 0 {
		return "email", nil
	}
	return nil, errors.New("Invalid value of RadarExportType: " + string(*en))
}

func (en *RadarExportType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "beepro":
		*en = RadarExportTypeBeePro
		return nil
	case "email":
		*en = RadarExportTypeEmail
		return nil
	case "external":
		*en = RadarExportTypeExternal
		return nil
	case "mytarget":
		*en = RadarExportTypeMytarget
		return nil
	case "yandex":
		*en = RadarExportTypeYandex
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportTypeEmail
		return nil
	}
	return errors.New("Unknown RadarExportType: " + s)
}

func (en *RadarExportType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "beepro":
		*en = RadarExportTypeBeePro
		return nil
	case "email":
		*en = RadarExportTypeEmail
		return nil
	case "external":
		*en = RadarExportTypeExternal
		return nil
	case "mytarget":
		*en = RadarExportTypeMytarget
		return nil
	case "yandex":
		*en = RadarExportTypeYandex
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportTypeEmail
		return nil
	}
	return errors.New("Unknown RadarExportType: " + s)
}

type RadiusMessageType string

const RadiusMessageTypeAccessAccept RadiusMessageType = "access-accept"
const RadiusMessageTypeAccessReject RadiusMessageType = "access-reject"
const RadiusMessageTypeAccessRequest RadiusMessageType = "access-request"
const RadiusMessageTypeAccountingRequest RadiusMessageType = "accounting"

func (en RadiusMessageType) GetPtr() *RadiusMessageType { var v = en; return &v }

func (en RadiusMessageType) String() string {
	switch en {
	case RadiusMessageTypeAccessAccept:
		return "access-accept"
	case RadiusMessageTypeAccessReject:
		return "access-reject"
	case RadiusMessageTypeAccessRequest:
		return "access-request"
	case RadiusMessageTypeAccountingRequest:
		return "accounting"
	}
	if len(en) == 0 {
		return "accounting"
	}
	panic(errors.New("Invalid value of RadiusMessageType: " + string(en)))
}

func (en *RadiusMessageType) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadiusMessageTypeAccessAccept:
		return json.Marshal("access-accept")
	case RadiusMessageTypeAccessReject:
		return json.Marshal("access-reject")
	case RadiusMessageTypeAccessRequest:
		return json.Marshal("access-request")
	case RadiusMessageTypeAccountingRequest:
		return json.Marshal("accounting")
	}
	if len(*en) == 0 {
		return json.Marshal("accounting")
	}
	return nil, errors.New("Invalid value of RadiusMessageType: " + string(*en))
}

func (en *RadiusMessageType) GetBSON() (interface{}, error) {
	switch *en {
	case RadiusMessageTypeAccessAccept:
		return "access-accept", nil
	case RadiusMessageTypeAccessReject:
		return "access-reject", nil
	case RadiusMessageTypeAccessRequest:
		return "access-request", nil
	case RadiusMessageTypeAccountingRequest:
		return "accounting", nil
	}
	if len(*en) == 0 {
		return "accounting", nil
	}
	return nil, errors.New("Invalid value of RadiusMessageType: " + string(*en))
}

func (en *RadiusMessageType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "access-accept":
		*en = RadiusMessageTypeAccessAccept
		return nil
	case "access-reject":
		*en = RadiusMessageTypeAccessReject
		return nil
	case "access-request":
		*en = RadiusMessageTypeAccessRequest
		return nil
	case "accounting":
		*en = RadiusMessageTypeAccountingRequest
		return nil
	}
	if len(s) == 0 {
		*en = RadiusMessageTypeAccountingRequest
		return nil
	}
	return errors.New("Unknown RadiusMessageType: " + s)
}

func (en *RadiusMessageType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "access-accept":
		*en = RadiusMessageTypeAccessAccept
		return nil
	case "access-reject":
		*en = RadiusMessageTypeAccessReject
		return nil
	case "access-request":
		*en = RadiusMessageTypeAccessRequest
		return nil
	case "accounting":
		*en = RadiusMessageTypeAccountingRequest
		return nil
	}
	if len(s) == 0 {
		*en = RadiusMessageTypeAccountingRequest
		return nil
	}
	return errors.New("Unknown RadiusMessageType: " + s)
}

type ReportActionType string

const ReportActionTypeEmail ReportActionType = "email"
const ReportActionTypeScript ReportActionType = "script"

func (en ReportActionType) GetPtr() *ReportActionType { var v = en; return &v }

func (en ReportActionType) String() string {
	switch en {
	case ReportActionTypeEmail:
		return "email"
	case ReportActionTypeScript:
		return "script"
	}
	if len(en) == 0 {
		return "email"
	}
	panic(errors.New("Invalid value of ReportActionType: " + string(en)))
}

func (en *ReportActionType) MarshalJSON() ([]byte, error) {
	switch *en {
	case ReportActionTypeEmail:
		return json.Marshal("email")
	case ReportActionTypeScript:
		return json.Marshal("script")
	}
	if len(*en) == 0 {
		return json.Marshal("email")
	}
	return nil, errors.New("Invalid value of ReportActionType: " + string(*en))
}

func (en *ReportActionType) GetBSON() (interface{}, error) {
	switch *en {
	case ReportActionTypeEmail:
		return "email", nil
	case ReportActionTypeScript:
		return "script", nil
	}
	if len(*en) == 0 {
		return "email", nil
	}
	return nil, errors.New("Invalid value of ReportActionType: " + string(*en))
}

func (en *ReportActionType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "email":
		*en = ReportActionTypeEmail
		return nil
	case "script":
		*en = ReportActionTypeScript
		return nil
	}
	if len(s) == 0 {
		*en = ReportActionTypeEmail
		return nil
	}
	return errors.New("Unknown ReportActionType: " + s)
}

func (en *ReportActionType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "email":
		*en = ReportActionTypeEmail
		return nil
	case "script":
		*en = ReportActionTypeScript
		return nil
	}
	if len(s) == 0 {
		*en = ReportActionTypeEmail
		return nil
	}
	return errors.New("Unknown ReportActionType: " + s)
}

type ReportFormat string

const ReportFormatCSV ReportFormat = "csv"
const ReportFormatJson ReportFormat = "json"
const ReportFormatTxt ReportFormat = "txt"
const ReportFormatXLSX ReportFormat = "xlsx"

func (en ReportFormat) GetPtr() *ReportFormat { var v = en; return &v }

func (en ReportFormat) String() string {
	switch en {
	case ReportFormatCSV:
		return "csv"
	case ReportFormatJson:
		return "json"
	case ReportFormatTxt:
		return "txt"
	case ReportFormatXLSX:
		return "xlsx"
	}
	if len(en) == 0 {
		return "json"
	}
	panic(errors.New("Invalid value of ReportFormat: " + string(en)))
}

func (en *ReportFormat) MarshalJSON() ([]byte, error) {
	switch *en {
	case ReportFormatCSV:
		return json.Marshal("csv")
	case ReportFormatJson:
		return json.Marshal("json")
	case ReportFormatTxt:
		return json.Marshal("txt")
	case ReportFormatXLSX:
		return json.Marshal("xlsx")
	}
	if len(*en) == 0 {
		return json.Marshal("json")
	}
	return nil, errors.New("Invalid value of ReportFormat: " + string(*en))
}

func (en *ReportFormat) GetBSON() (interface{}, error) {
	switch *en {
	case ReportFormatCSV:
		return "csv", nil
	case ReportFormatJson:
		return "json", nil
	case ReportFormatTxt:
		return "txt", nil
	case ReportFormatXLSX:
		return "xlsx", nil
	}
	if len(*en) == 0 {
		return "json", nil
	}
	return nil, errors.New("Invalid value of ReportFormat: " + string(*en))
}

func (en *ReportFormat) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*en = ReportFormatCSV
		return nil
	case "json":
		*en = ReportFormatJson
		return nil
	case "txt":
		*en = ReportFormatTxt
		return nil
	case "xlsx":
		*en = ReportFormatXLSX
		return nil
	}
	if len(s) == 0 {
		*en = ReportFormatJson
		return nil
	}
	return errors.New("Unknown ReportFormat: " + s)
}

func (en *ReportFormat) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*en = ReportFormatCSV
		return nil
	case "json":
		*en = ReportFormatJson
		return nil
	case "txt":
		*en = ReportFormatTxt
		return nil
	case "xlsx":
		*en = ReportFormatXLSX
		return nil
	}
	if len(s) == 0 {
		*en = ReportFormatJson
		return nil
	}
	return errors.New("Unknown ReportFormat: " + s)
}

type ReportPeriod string

const ReportPeriodDay ReportPeriod = "day"
const ReportPeriodMonth ReportPeriod = "month"
const ReportPeriodNow ReportPeriod = "now"
const ReportPeriodWeek ReportPeriod = "week"

func (en ReportPeriod) GetPtr() *ReportPeriod { var v = en; return &v }

func (en ReportPeriod) String() string {
	switch en {
	case ReportPeriodDay:
		return "day"
	case ReportPeriodMonth:
		return "month"
	case ReportPeriodNow:
		return "now"
	case ReportPeriodWeek:
		return "week"
	}
	if len(en) == 0 {
		return "now"
	}
	panic(errors.New("Invalid value of ReportPeriod: " + string(en)))
}

func (en *ReportPeriod) MarshalJSON() ([]byte, error) {
	switch *en {
	case ReportPeriodDay:
		return json.Marshal("day")
	case ReportPeriodMonth:
		return json.Marshal("month")
	case ReportPeriodNow:
		return json.Marshal("now")
	case ReportPeriodWeek:
		return json.Marshal("week")
	}
	if len(*en) == 0 {
		return json.Marshal("now")
	}
	return nil, errors.New("Invalid value of ReportPeriod: " + string(*en))
}

func (en *ReportPeriod) GetBSON() (interface{}, error) {
	switch *en {
	case ReportPeriodDay:
		return "day", nil
	case ReportPeriodMonth:
		return "month", nil
	case ReportPeriodNow:
		return "now", nil
	case ReportPeriodWeek:
		return "week", nil
	}
	if len(*en) == 0 {
		return "now", nil
	}
	return nil, errors.New("Invalid value of ReportPeriod: " + string(*en))
}

func (en *ReportPeriod) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "day":
		*en = ReportPeriodDay
		return nil
	case "month":
		*en = ReportPeriodMonth
		return nil
	case "now":
		*en = ReportPeriodNow
		return nil
	case "week":
		*en = ReportPeriodWeek
		return nil
	}
	if len(s) == 0 {
		*en = ReportPeriodNow
		return nil
	}
	return errors.New("Unknown ReportPeriod: " + s)
}

func (en *ReportPeriod) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "day":
		*en = ReportPeriodDay
		return nil
	case "month":
		*en = ReportPeriodMonth
		return nil
	case "now":
		*en = ReportPeriodNow
		return nil
	case "week":
		*en = ReportPeriodWeek
		return nil
	}
	if len(s) == 0 {
		*en = ReportPeriodNow
		return nil
	}
	return errors.New("Unknown ReportPeriod: " + s)
}

type ReportSubject string

const ReportSubjectCPECommon ReportSubject = "cpes_common"
const ReportSubjectCPEs ReportSubject = "cpes"
const ReportSubjectClients ReportSubject = "clients"
const ReportSubjectEvents ReportSubject = "events"

func (en ReportSubject) GetPtr() *ReportSubject { var v = en; return &v }

func (en ReportSubject) String() string {
	switch en {
	case ReportSubjectCPECommon:
		return "cpes_common"
	case ReportSubjectCPEs:
		return "cpes"
	case ReportSubjectClients:
		return "clients"
	case ReportSubjectEvents:
		return "events"
	}
	if len(en) == 0 {
		return "clients"
	}
	panic(errors.New("Invalid value of ReportSubject: " + string(en)))
}

func (en *ReportSubject) MarshalJSON() ([]byte, error) {
	switch *en {
	case ReportSubjectCPECommon:
		return json.Marshal("cpes_common")
	case ReportSubjectCPEs:
		return json.Marshal("cpes")
	case ReportSubjectClients:
		return json.Marshal("clients")
	case ReportSubjectEvents:
		return json.Marshal("events")
	}
	if len(*en) == 0 {
		return json.Marshal("clients")
	}
	return nil, errors.New("Invalid value of ReportSubject: " + string(*en))
}

func (en *ReportSubject) GetBSON() (interface{}, error) {
	switch *en {
	case ReportSubjectCPECommon:
		return "cpes_common", nil
	case ReportSubjectCPEs:
		return "cpes", nil
	case ReportSubjectClients:
		return "clients", nil
	case ReportSubjectEvents:
		return "events", nil
	}
	if len(*en) == 0 {
		return "clients", nil
	}
	return nil, errors.New("Invalid value of ReportSubject: " + string(*en))
}

func (en *ReportSubject) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "cpes_common":
		*en = ReportSubjectCPECommon
		return nil
	case "cpes":
		*en = ReportSubjectCPEs
		return nil
	case "clients":
		*en = ReportSubjectClients
		return nil
	case "events":
		*en = ReportSubjectEvents
		return nil
	}
	if len(s) == 0 {
		*en = ReportSubjectClients
		return nil
	}
	return errors.New("Unknown ReportSubject: " + s)
}

func (en *ReportSubject) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "cpes_common":
		*en = ReportSubjectCPECommon
		return nil
	case "cpes":
		*en = ReportSubjectCPEs
		return nil
	case "clients":
		*en = ReportSubjectClients
		return nil
	case "events":
		*en = ReportSubjectEvents
		return nil
	}
	if len(s) == 0 {
		*en = ReportSubjectClients
		return nil
	}
	return errors.New("Unknown ReportSubject: " + s)
}

type ReportType string

const ReportTypeClientsAll ReportType = "clients"
const ReportTypeClientsAuthorized ReportType = "clients_auth"
const ReportTypeCurrent ReportType = "current"
const ReportTypeSummary ReportType = "summary"

func (en ReportType) GetPtr() *ReportType { var v = en; return &v }

func (en ReportType) String() string {
	switch en {
	case ReportTypeClientsAll:
		return "clients"
	case ReportTypeClientsAuthorized:
		return "clients_auth"
	case ReportTypeCurrent:
		return "current"
	case ReportTypeSummary:
		return "summary"
	}
	if len(en) == 0 {
		return "current"
	}
	panic(errors.New("Invalid value of ReportType: " + string(en)))
}

func (en *ReportType) MarshalJSON() ([]byte, error) {
	switch *en {
	case ReportTypeClientsAll:
		return json.Marshal("clients")
	case ReportTypeClientsAuthorized:
		return json.Marshal("clients_auth")
	case ReportTypeCurrent:
		return json.Marshal("current")
	case ReportTypeSummary:
		return json.Marshal("summary")
	}
	if len(*en) == 0 {
		return json.Marshal("current")
	}
	return nil, errors.New("Invalid value of ReportType: " + string(*en))
}

func (en *ReportType) GetBSON() (interface{}, error) {
	switch *en {
	case ReportTypeClientsAll:
		return "clients", nil
	case ReportTypeClientsAuthorized:
		return "clients_auth", nil
	case ReportTypeCurrent:
		return "current", nil
	case ReportTypeSummary:
		return "summary", nil
	}
	if len(*en) == 0 {
		return "current", nil
	}
	return nil, errors.New("Invalid value of ReportType: " + string(*en))
}

func (en *ReportType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "clients":
		*en = ReportTypeClientsAll
		return nil
	case "clients_auth":
		*en = ReportTypeClientsAuthorized
		return nil
	case "current":
		*en = ReportTypeCurrent
		return nil
	case "summary":
		*en = ReportTypeSummary
		return nil
	}
	if len(s) == 0 {
		*en = ReportTypeCurrent
		return nil
	}
	return errors.New("Unknown ReportType: " + s)
}

func (en *ReportType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "clients":
		*en = ReportTypeClientsAll
		return nil
	case "clients_auth":
		*en = ReportTypeClientsAuthorized
		return nil
	case "current":
		*en = ReportTypeCurrent
		return nil
	case "summary":
		*en = ReportTypeSummary
		return nil
	}
	if len(s) == 0 {
		*en = ReportTypeCurrent
		return nil
	}
	return errors.New("Unknown ReportType: " + s)
}

type SecuritySuite string

const SecuritySuiteAES SecuritySuite = "aes"
const SecuritySuiteCCMP SecuritySuite = "ccmp"
const SecuritySuiteTKIP SecuritySuite = "tkip"

func (en SecuritySuite) GetPtr() *SecuritySuite { var v = en; return &v }

func (en SecuritySuite) String() string {
	switch en {
	case SecuritySuiteAES:
		return "aes"
	case SecuritySuiteCCMP:
		return "ccmp"
	case SecuritySuiteTKIP:
		return "tkip"
	}
	panic(errors.New("Invalid value of SecuritySuite: " + string(en)))
}

func (en *SecuritySuite) MarshalJSON() ([]byte, error) {
	switch *en {
	case SecuritySuiteAES:
		return json.Marshal("aes")
	case SecuritySuiteCCMP:
		return json.Marshal("ccmp")
	case SecuritySuiteTKIP:
		return json.Marshal("tkip")
	}
	return nil, errors.New("Invalid value of SecuritySuite: " + string(*en))
}

func (en *SecuritySuite) GetBSON() (interface{}, error) {
	switch *en {
	case SecuritySuiteAES:
		return "aes", nil
	case SecuritySuiteCCMP:
		return "ccmp", nil
	case SecuritySuiteTKIP:
		return "tkip", nil
	}
	return nil, errors.New("Invalid value of SecuritySuite: " + string(*en))
}

func (en *SecuritySuite) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "aes":
		*en = SecuritySuiteAES
		return nil
	case "ccmp":
		*en = SecuritySuiteCCMP
		return nil
	case "tkip":
		*en = SecuritySuiteTKIP
		return nil
	}
	return errors.New("Unknown SecuritySuite: " + s)
}

func (en *SecuritySuite) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "aes":
		*en = SecuritySuiteAES
		return nil
	case "ccmp":
		*en = SecuritySuiteCCMP
		return nil
	case "tkip":
		*en = SecuritySuiteTKIP
		return nil
	}
	return errors.New("Unknown SecuritySuite: " + s)
}

type SecurityType string

const SecurityTypeNone SecurityType = "open"
const SecurityTypeWPA23Enterprise SecurityType = "wpa23enterprise"
const SecurityTypeWPA23Personal SecurityType = "wpa23personal"
const SecurityTypeWPA2Enterprise SecurityType = "wpa2enterprise"
const SecurityTypeWPA2Personal SecurityType = "wpa2personal"
const SecurityTypeWPA3Enterprise SecurityType = "wpa3enterprise"
const SecurityTypeWPA3Personal SecurityType = "wpa3personal"
const SecurityTypeWPAEnterprise SecurityType = "wpaenterprise"
const SecurityTypeWPAPersonal SecurityType = "wpapersonal"

func (en SecurityType) GetPtr() *SecurityType { var v = en; return &v }

func (en SecurityType) String() string {
	switch en {
	case SecurityTypeNone:
		return "open"
	case SecurityTypeWPA23Enterprise:
		return "wpa23enterprise"
	case SecurityTypeWPA23Personal:
		return "wpa23personal"
	case SecurityTypeWPA2Enterprise:
		return "wpa2enterprise"
	case SecurityTypeWPA2Personal:
		return "wpa2personal"
	case SecurityTypeWPA3Enterprise:
		return "wpa3enterprise"
	case SecurityTypeWPA3Personal:
		return "wpa3personal"
	case SecurityTypeWPAEnterprise:
		return "wpaenterprise"
	case SecurityTypeWPAPersonal:
		return "wpapersonal"
	}
	if len(en) == 0 {
		return "open"
	}
	panic(errors.New("Invalid value of SecurityType: " + string(en)))
}

func (en *SecurityType) MarshalJSON() ([]byte, error) {
	switch *en {
	case SecurityTypeNone:
		return json.Marshal("open")
	case SecurityTypeWPA23Enterprise:
		return json.Marshal("wpa23enterprise")
	case SecurityTypeWPA23Personal:
		return json.Marshal("wpa23personal")
	case SecurityTypeWPA2Enterprise:
		return json.Marshal("wpa2enterprise")
	case SecurityTypeWPA2Personal:
		return json.Marshal("wpa2personal")
	case SecurityTypeWPA3Enterprise:
		return json.Marshal("wpa3enterprise")
	case SecurityTypeWPA3Personal:
		return json.Marshal("wpa3personal")
	case SecurityTypeWPAEnterprise:
		return json.Marshal("wpaenterprise")
	case SecurityTypeWPAPersonal:
		return json.Marshal("wpapersonal")
	}
	if len(*en) == 0 {
		return json.Marshal("open")
	}
	return nil, errors.New("Invalid value of SecurityType: " + string(*en))
}

func (en *SecurityType) GetBSON() (interface{}, error) {
	switch *en {
	case SecurityTypeNone:
		return "open", nil
	case SecurityTypeWPA23Enterprise:
		return "wpa23enterprise", nil
	case SecurityTypeWPA23Personal:
		return "wpa23personal", nil
	case SecurityTypeWPA2Enterprise:
		return "wpa2enterprise", nil
	case SecurityTypeWPA2Personal:
		return "wpa2personal", nil
	case SecurityTypeWPA3Enterprise:
		return "wpa3enterprise", nil
	case SecurityTypeWPA3Personal:
		return "wpa3personal", nil
	case SecurityTypeWPAEnterprise:
		return "wpaenterprise", nil
	case SecurityTypeWPAPersonal:
		return "wpapersonal", nil
	}
	if len(*en) == 0 {
		return "open", nil
	}
	return nil, errors.New("Invalid value of SecurityType: " + string(*en))
}

func (en *SecurityType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "open":
		*en = SecurityTypeNone
		return nil
	case "wpa23enterprise":
		*en = SecurityTypeWPA23Enterprise
		return nil
	case "wpa23personal":
		*en = SecurityTypeWPA23Personal
		return nil
	case "wpa2enterprise":
		*en = SecurityTypeWPA2Enterprise
		return nil
	case "wpa2personal":
		*en = SecurityTypeWPA2Personal
		return nil
	case "wpa3enterprise":
		*en = SecurityTypeWPA3Enterprise
		return nil
	case "wpa3personal":
		*en = SecurityTypeWPA3Personal
		return nil
	case "wpaenterprise":
		*en = SecurityTypeWPAEnterprise
		return nil
	case "wpapersonal":
		*en = SecurityTypeWPAPersonal
		return nil
	}
	if len(s) == 0 {
		*en = SecurityTypeNone
		return nil
	}
	return errors.New("Unknown SecurityType: " + s)
}

func (en *SecurityType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "open":
		*en = SecurityTypeNone
		return nil
	case "wpa23enterprise":
		*en = SecurityTypeWPA23Enterprise
		return nil
	case "wpa23personal":
		*en = SecurityTypeWPA23Personal
		return nil
	case "wpa2enterprise":
		*en = SecurityTypeWPA2Enterprise
		return nil
	case "wpa2personal":
		*en = SecurityTypeWPA2Personal
		return nil
	case "wpa3enterprise":
		*en = SecurityTypeWPA3Enterprise
		return nil
	case "wpa3personal":
		*en = SecurityTypeWPA3Personal
		return nil
	case "wpaenterprise":
		*en = SecurityTypeWPAEnterprise
		return nil
	case "wpapersonal":
		*en = SecurityTypeWPAPersonal
		return nil
	}
	if len(s) == 0 {
		*en = SecurityTypeNone
		return nil
	}
	return errors.New("Unknown SecurityType: " + s)
}

type ServiceState string

const ServiceStateConnected ServiceState = "CONNECTED"
const ServiceStateDisconnected ServiceState = "DISCONNECTED"
const ServiceStatePending ServiceState = "PENDING"

func (en ServiceState) GetPtr() *ServiceState { var v = en; return &v }

func (en ServiceState) String() string {
	switch en {
	case ServiceStateConnected:
		return "CONNECTED"
	case ServiceStateDisconnected:
		return "DISCONNECTED"
	case ServiceStatePending:
		return "PENDING"
	}
	panic(errors.New("Invalid value of ServiceState: " + string(en)))
}

func (en *ServiceState) MarshalJSON() ([]byte, error) {
	switch *en {
	case ServiceStateConnected:
		return json.Marshal("CONNECTED")
	case ServiceStateDisconnected:
		return json.Marshal("DISCONNECTED")
	case ServiceStatePending:
		return json.Marshal("PENDING")
	}
	return nil, errors.New("Invalid value of ServiceState: " + string(*en))
}

func (en *ServiceState) GetBSON() (interface{}, error) {
	switch *en {
	case ServiceStateConnected:
		return "CONNECTED", nil
	case ServiceStateDisconnected:
		return "DISCONNECTED", nil
	case ServiceStatePending:
		return "PENDING", nil
	}
	return nil, errors.New("Invalid value of ServiceState: " + string(*en))
}

func (en *ServiceState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*en = ServiceStateConnected
		return nil
	case "DISCONNECTED":
		*en = ServiceStateDisconnected
		return nil
	case "PENDING":
		*en = ServiceStatePending
		return nil
	}
	return errors.New("Unknown ServiceState: " + s)
}

func (en *ServiceState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*en = ServiceStateConnected
		return nil
	case "DISCONNECTED":
		*en = ServiceStateDisconnected
		return nil
	case "PENDING":
		*en = ServiceStatePending
		return nil
	}
	return errors.New("Unknown ServiceState: " + s)
}

type SpeedType string

const SpeedTypeGbit SpeedType = "gbit"
const SpeedTypeKbit SpeedType = "kbit"
const SpeedTypeMbit SpeedType = "mbit"
const SpeedTypeTbit SpeedType = "tbit"

func (en SpeedType) GetPtr() *SpeedType { var v = en; return &v }

func (en SpeedType) String() string {
	switch en {
	case SpeedTypeGbit:
		return "gbit"
	case SpeedTypeKbit:
		return "kbit"
	case SpeedTypeMbit:
		return "mbit"
	case SpeedTypeTbit:
		return "tbit"
	}
	if len(en) == 0 {
		return "kbit"
	}
	panic(errors.New("Invalid value of SpeedType: " + string(en)))
}

func (en *SpeedType) MarshalJSON() ([]byte, error) {
	switch *en {
	case SpeedTypeGbit:
		return json.Marshal("gbit")
	case SpeedTypeKbit:
		return json.Marshal("kbit")
	case SpeedTypeMbit:
		return json.Marshal("mbit")
	case SpeedTypeTbit:
		return json.Marshal("tbit")
	}
	if len(*en) == 0 {
		return json.Marshal("kbit")
	}
	return nil, errors.New("Invalid value of SpeedType: " + string(*en))
}

func (en *SpeedType) GetBSON() (interface{}, error) {
	switch *en {
	case SpeedTypeGbit:
		return "gbit", nil
	case SpeedTypeKbit:
		return "kbit", nil
	case SpeedTypeMbit:
		return "mbit", nil
	case SpeedTypeTbit:
		return "tbit", nil
	}
	if len(*en) == 0 {
		return "kbit", nil
	}
	return nil, errors.New("Invalid value of SpeedType: " + string(*en))
}

func (en *SpeedType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "gbit":
		*en = SpeedTypeGbit
		return nil
	case "kbit":
		*en = SpeedTypeKbit
		return nil
	case "mbit":
		*en = SpeedTypeMbit
		return nil
	case "tbit":
		*en = SpeedTypeTbit
		return nil
	}
	if len(s) == 0 {
		*en = SpeedTypeKbit
		return nil
	}
	return errors.New("Unknown SpeedType: " + s)
}

func (en *SpeedType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "gbit":
		*en = SpeedTypeGbit
		return nil
	case "kbit":
		*en = SpeedTypeKbit
		return nil
	case "mbit":
		*en = SpeedTypeMbit
		return nil
	case "tbit":
		*en = SpeedTypeTbit
		return nil
	}
	if len(s) == 0 {
		*en = SpeedTypeKbit
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
const StatEventRuleTypeCustomerActivity StatEventRuleType = "customer_activity"
const StatEventRuleTypeDisconnected StatEventRuleType = "disconnected"
const StatEventRuleTypeFreeRAM StatEventRuleType = "free_ram"
const StatEventRuleTypeIfaceError StatEventRuleType = "iface_error"

func (en StatEventRuleType) GetPtr() *StatEventRuleType { var v = en; return &v }

func (en StatEventRuleType) String() string {
	switch en {
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
	case StatEventRuleTypeCustomerActivity:
		return "customer_activity"
	case StatEventRuleTypeDisconnected:
		return "disconnected"
	case StatEventRuleTypeFreeRAM:
		return "free_ram"
	case StatEventRuleTypeIfaceError:
		return "iface_error"
	}
	panic(errors.New("Invalid value of StatEventRuleType: " + string(en)))
}

func (en *StatEventRuleType) MarshalJSON() ([]byte, error) {
	switch *en {
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
	case StatEventRuleTypeCustomerActivity:
		return json.Marshal("customer_activity")
	case StatEventRuleTypeDisconnected:
		return json.Marshal("disconnected")
	case StatEventRuleTypeFreeRAM:
		return json.Marshal("free_ram")
	case StatEventRuleTypeIfaceError:
		return json.Marshal("iface_error")
	}
	return nil, errors.New("Invalid value of StatEventRuleType: " + string(*en))
}

func (en *StatEventRuleType) GetBSON() (interface{}, error) {
	switch *en {
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
	case StatEventRuleTypeCustomerActivity:
		return "customer_activity", nil
	case StatEventRuleTypeDisconnected:
		return "disconnected", nil
	case StatEventRuleTypeFreeRAM:
		return "free_ram", nil
	case StatEventRuleTypeIfaceError:
		return "iface_error", nil
	}
	return nil, errors.New("Invalid value of StatEventRuleType: " + string(*en))
}

func (en *StatEventRuleType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "cpu_load":
		*en = StatEventRuleTypeCPUload
		return nil
	case "client_con":
		*en = StatEventRuleTypeClientCon
		return nil
	case "client_dis":
		*en = StatEventRuleTypeClientDis
		return nil
	case "client_far":
		*en = StatEventRuleTypeClientFar
		return nil
	case "config_error":
		*en = StatEventRuleTypeConfigError
		return nil
	case "connected":
		*en = StatEventRuleTypeConnected
		return nil
	case "customer_activity":
		*en = StatEventRuleTypeCustomerActivity
		return nil
	case "disconnected":
		*en = StatEventRuleTypeDisconnected
		return nil
	case "free_ram":
		*en = StatEventRuleTypeFreeRAM
		return nil
	case "iface_error":
		*en = StatEventRuleTypeIfaceError
		return nil
	}
	return errors.New("Unknown StatEventRuleType: " + s)
}

func (en *StatEventRuleType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "cpu_load":
		*en = StatEventRuleTypeCPUload
		return nil
	case "client_con":
		*en = StatEventRuleTypeClientCon
		return nil
	case "client_dis":
		*en = StatEventRuleTypeClientDis
		return nil
	case "client_far":
		*en = StatEventRuleTypeClientFar
		return nil
	case "config_error":
		*en = StatEventRuleTypeConfigError
		return nil
	case "connected":
		*en = StatEventRuleTypeConnected
		return nil
	case "customer_activity":
		*en = StatEventRuleTypeCustomerActivity
		return nil
	case "disconnected":
		*en = StatEventRuleTypeDisconnected
		return nil
	case "free_ram":
		*en = StatEventRuleTypeFreeRAM
		return nil
	case "iface_error":
		*en = StatEventRuleTypeIfaceError
		return nil
	}
	return errors.New("Unknown StatEventRuleType: " + s)
}

type SystemEventLevel string

const SystemEventLevelDEBUG SystemEventLevel = "DEBUG"
const SystemEventLevelERROR SystemEventLevel = "ERROR"
const SystemEventLevelINFO SystemEventLevel = "INFO"
const SystemEventLevelWARNING SystemEventLevel = "WARNING"

func (en SystemEventLevel) GetPtr() *SystemEventLevel { var v = en; return &v }

func (en SystemEventLevel) String() string {
	switch en {
	case SystemEventLevelDEBUG:
		return "DEBUG"
	case SystemEventLevelERROR:
		return "ERROR"
	case SystemEventLevelINFO:
		return "INFO"
	case SystemEventLevelWARNING:
		return "WARNING"
	}
	if len(en) == 0 {
		return "DEBUG"
	}
	panic(errors.New("Invalid value of SystemEventLevel: " + string(en)))
}

func (en *SystemEventLevel) MarshalJSON() ([]byte, error) {
	switch *en {
	case SystemEventLevelDEBUG:
		return json.Marshal("DEBUG")
	case SystemEventLevelERROR:
		return json.Marshal("ERROR")
	case SystemEventLevelINFO:
		return json.Marshal("INFO")
	case SystemEventLevelWARNING:
		return json.Marshal("WARNING")
	}
	if len(*en) == 0 {
		return json.Marshal("DEBUG")
	}
	return nil, errors.New("Invalid value of SystemEventLevel: " + string(*en))
}

func (en *SystemEventLevel) GetBSON() (interface{}, error) {
	switch *en {
	case SystemEventLevelDEBUG:
		return "DEBUG", nil
	case SystemEventLevelERROR:
		return "ERROR", nil
	case SystemEventLevelINFO:
		return "INFO", nil
	case SystemEventLevelWARNING:
		return "WARNING", nil
	}
	if len(*en) == 0 {
		return "DEBUG", nil
	}
	return nil, errors.New("Invalid value of SystemEventLevel: " + string(*en))
}

func (en *SystemEventLevel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "DEBUG":
		*en = SystemEventLevelDEBUG
		return nil
	case "ERROR":
		*en = SystemEventLevelERROR
		return nil
	case "INFO":
		*en = SystemEventLevelINFO
		return nil
	case "WARNING":
		*en = SystemEventLevelWARNING
		return nil
	}
	if len(s) == 0 {
		*en = SystemEventLevelDEBUG
		return nil
	}
	return errors.New("Unknown SystemEventLevel: " + s)
}

func (en *SystemEventLevel) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "DEBUG":
		*en = SystemEventLevelDEBUG
		return nil
	case "ERROR":
		*en = SystemEventLevelERROR
		return nil
	case "INFO":
		*en = SystemEventLevelINFO
		return nil
	case "WARNING":
		*en = SystemEventLevelWARNING
		return nil
	}
	if len(s) == 0 {
		*en = SystemEventLevelDEBUG
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
const SystemEventTypeLocationCacheReload SystemEventType = "LOCATION_CACHE_RELOAD"
const SystemEventTypeLoggedError SystemEventType = "LOGGED_ERROR"
const SystemEventTypeMonitorRuleViolation SystemEventType = "MONITOR_RULE_VIOLATION"
const SystemEventTypeRRMGroupApplyAlgo SystemEventType = "RRM_GROUP_APPLY_ALGO"
const SystemEventTypeRRMStatus SystemEventType = "RRM_STATUS_DATA"
const SystemEventTypeRadarExportUpdate SystemEventType = "RADAR_EXPORT_UPDATE"
const SystemEventTypeRadiusAccountingSend SystemEventType = "RADIUS_ACCOUNTING_SEND"
const SystemEventTypeServiceConnected SystemEventType = "SERVICE_CONNECTED"
const SystemEventTypeServiceDisconnected SystemEventType = "SERVICE_DISCONNECTED"
const SystemEventTypeServiceFatalError SystemEventType = "SERVICE_FATAL_ERROR"
const SystemEventTypeSystemTimeChanged SystemEventType = "SYSTEM_TIME_CHANGE"
const SystemEventTypeUserAuthSuccess SystemEventType = "USER_AUTHORIZATION_SUCCESS"
const SystemEventTypeWLANCentrAccChanged SystemEventType = "WLAN_CENTR_ACC_CHANGE"

func (en SystemEventType) GetPtr() *SystemEventType { var v = en; return &v }

func (en SystemEventType) String() string {
	switch en {
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
	case SystemEventTypeLocationCacheReload:
		return "LOCATION_CACHE_RELOAD"
	case SystemEventTypeLoggedError:
		return "LOGGED_ERROR"
	case SystemEventTypeMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION"
	case SystemEventTypeRRMGroupApplyAlgo:
		return "RRM_GROUP_APPLY_ALGO"
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
	case SystemEventTypeUserAuthSuccess:
		return "USER_AUTHORIZATION_SUCCESS"
	case SystemEventTypeWLANCentrAccChanged:
		return "WLAN_CENTR_ACC_CHANGE"
	}
	panic(errors.New("Invalid value of SystemEventType: " + string(en)))
}

func (en *SystemEventType) MarshalJSON() ([]byte, error) {
	switch *en {
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
	case SystemEventTypeLocationCacheReload:
		return json.Marshal("LOCATION_CACHE_RELOAD")
	case SystemEventTypeLoggedError:
		return json.Marshal("LOGGED_ERROR")
	case SystemEventTypeMonitorRuleViolation:
		return json.Marshal("MONITOR_RULE_VIOLATION")
	case SystemEventTypeRRMGroupApplyAlgo:
		return json.Marshal("RRM_GROUP_APPLY_ALGO")
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
	case SystemEventTypeUserAuthSuccess:
		return json.Marshal("USER_AUTHORIZATION_SUCCESS")
	case SystemEventTypeWLANCentrAccChanged:
		return json.Marshal("WLAN_CENTR_ACC_CHANGE")
	}
	return nil, errors.New("Invalid value of SystemEventType: " + string(*en))
}

func (en *SystemEventType) GetBSON() (interface{}, error) {
	switch *en {
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
	case SystemEventTypeLocationCacheReload:
		return "LOCATION_CACHE_RELOAD", nil
	case SystemEventTypeLoggedError:
		return "LOGGED_ERROR", nil
	case SystemEventTypeMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION", nil
	case SystemEventTypeRRMGroupApplyAlgo:
		return "RRM_GROUP_APPLY_ALGO", nil
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
	case SystemEventTypeUserAuthSuccess:
		return "USER_AUTHORIZATION_SUCCESS", nil
	case SystemEventTypeWLANCentrAccChanged:
		return "WLAN_CENTR_ACC_CHANGE", nil
	}
	return nil, errors.New("Invalid value of SystemEventType: " + string(*en))
}

func (en *SystemEventType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		*en = SystemEventTypeAny
		return nil
	case "CPE_CONNECTED":
		*en = SystemEventTypeCPEConnected
		return nil
	case "CPE_DISCONNECTED":
		*en = SystemEventTypeCPEDisconnected
		return nil
	case "CPE_INTERFACE_STATE":
		*en = SystemEventTypeCPEInterfaceState
		return nil
	case "CLIENT_AUTHORIZATION":
		*en = SystemEventTypeClientAuthorization
		return nil
	case "CLIENT_CONNECTED":
		*en = SystemEventTypeClientConnected
		return nil
	case "CLIENT_DISCONNECTED":
		*en = SystemEventTypeClientDisconnected
		return nil
	case "CPE_FIRMWARE_AVAILABLE":
		*en = SystemEventTypeCpeFirmwareAvailable
		return nil
	case "DHCP_ACK":
		*en = SystemEventTypeDHCPAck
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		*en = SystemEventTypeDaemonSettingsChanged
		return nil
	case "FIRMWARE_UPLOADED":
		*en = SystemEventTypeFirmwareUploaded
		return nil
	case "LOCATION_CACHE_RELOAD":
		*en = SystemEventTypeLocationCacheReload
		return nil
	case "LOGGED_ERROR":
		*en = SystemEventTypeLoggedError
		return nil
	case "MONITOR_RULE_VIOLATION":
		*en = SystemEventTypeMonitorRuleViolation
		return nil
	case "RRM_GROUP_APPLY_ALGO":
		*en = SystemEventTypeRRMGroupApplyAlgo
		return nil
	case "RRM_STATUS_DATA":
		*en = SystemEventTypeRRMStatus
		return nil
	case "RADAR_EXPORT_UPDATE":
		*en = SystemEventTypeRadarExportUpdate
		return nil
	case "RADIUS_ACCOUNTING_SEND":
		*en = SystemEventTypeRadiusAccountingSend
		return nil
	case "SERVICE_CONNECTED":
		*en = SystemEventTypeServiceConnected
		return nil
	case "SERVICE_DISCONNECTED":
		*en = SystemEventTypeServiceDisconnected
		return nil
	case "SERVICE_FATAL_ERROR":
		*en = SystemEventTypeServiceFatalError
		return nil
	case "SYSTEM_TIME_CHANGE":
		*en = SystemEventTypeSystemTimeChanged
		return nil
	case "USER_AUTHORIZATION_SUCCESS":
		*en = SystemEventTypeUserAuthSuccess
		return nil
	case "WLAN_CENTR_ACC_CHANGE":
		*en = SystemEventTypeWLANCentrAccChanged
		return nil
	}
	return errors.New("Unknown SystemEventType: " + s)
}

func (en *SystemEventType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "+":
		*en = SystemEventTypeAny
		return nil
	case "CPE_CONNECTED":
		*en = SystemEventTypeCPEConnected
		return nil
	case "CPE_DISCONNECTED":
		*en = SystemEventTypeCPEDisconnected
		return nil
	case "CPE_INTERFACE_STATE":
		*en = SystemEventTypeCPEInterfaceState
		return nil
	case "CLIENT_AUTHORIZATION":
		*en = SystemEventTypeClientAuthorization
		return nil
	case "CLIENT_CONNECTED":
		*en = SystemEventTypeClientConnected
		return nil
	case "CLIENT_DISCONNECTED":
		*en = SystemEventTypeClientDisconnected
		return nil
	case "CPE_FIRMWARE_AVAILABLE":
		*en = SystemEventTypeCpeFirmwareAvailable
		return nil
	case "DHCP_ACK":
		*en = SystemEventTypeDHCPAck
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		*en = SystemEventTypeDaemonSettingsChanged
		return nil
	case "FIRMWARE_UPLOADED":
		*en = SystemEventTypeFirmwareUploaded
		return nil
	case "LOCATION_CACHE_RELOAD":
		*en = SystemEventTypeLocationCacheReload
		return nil
	case "LOGGED_ERROR":
		*en = SystemEventTypeLoggedError
		return nil
	case "MONITOR_RULE_VIOLATION":
		*en = SystemEventTypeMonitorRuleViolation
		return nil
	case "RRM_GROUP_APPLY_ALGO":
		*en = SystemEventTypeRRMGroupApplyAlgo
		return nil
	case "RRM_STATUS_DATA":
		*en = SystemEventTypeRRMStatus
		return nil
	case "RADAR_EXPORT_UPDATE":
		*en = SystemEventTypeRadarExportUpdate
		return nil
	case "RADIUS_ACCOUNTING_SEND":
		*en = SystemEventTypeRadiusAccountingSend
		return nil
	case "SERVICE_CONNECTED":
		*en = SystemEventTypeServiceConnected
		return nil
	case "SERVICE_DISCONNECTED":
		*en = SystemEventTypeServiceDisconnected
		return nil
	case "SERVICE_FATAL_ERROR":
		*en = SystemEventTypeServiceFatalError
		return nil
	case "SYSTEM_TIME_CHANGE":
		*en = SystemEventTypeSystemTimeChanged
		return nil
	case "USER_AUTHORIZATION_SUCCESS":
		*en = SystemEventTypeUserAuthSuccess
		return nil
	case "WLAN_CENTR_ACC_CHANGE":
		*en = SystemEventTypeWLANCentrAccChanged
		return nil
	}
	return errors.New("Unknown SystemEventType: " + s)
}

type TunManagerRPC string

const TunManagerRPCCreateL2TunnelSession TunManagerRPC = "CreateL2TunnelSession"
const TunManagerRPCDeleteL2TunnelSession TunManagerRPC = "DeleteL2TunnelSession"

func (en TunManagerRPC) GetPtr() *TunManagerRPC { var v = en; return &v }

func (en TunManagerRPC) String() string {
	switch en {
	case TunManagerRPCCreateL2TunnelSession:
		return "CreateL2TunnelSession"
	case TunManagerRPCDeleteL2TunnelSession:
		return "DeleteL2TunnelSession"
	}
	panic(errors.New("Invalid value of TunManagerRPC: " + string(en)))
}

func (en *TunManagerRPC) MarshalJSON() ([]byte, error) {
	switch *en {
	case TunManagerRPCCreateL2TunnelSession:
		return json.Marshal("CreateL2TunnelSession")
	case TunManagerRPCDeleteL2TunnelSession:
		return json.Marshal("DeleteL2TunnelSession")
	}
	return nil, errors.New("Invalid value of TunManagerRPC: " + string(*en))
}

func (en *TunManagerRPC) GetBSON() (interface{}, error) {
	switch *en {
	case TunManagerRPCCreateL2TunnelSession:
		return "CreateL2TunnelSession", nil
	case TunManagerRPCDeleteL2TunnelSession:
		return "DeleteL2TunnelSession", nil
	}
	return nil, errors.New("Invalid value of TunManagerRPC: " + string(*en))
}

func (en *TunManagerRPC) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CreateL2TunnelSession":
		*en = TunManagerRPCCreateL2TunnelSession
		return nil
	case "DeleteL2TunnelSession":
		*en = TunManagerRPCDeleteL2TunnelSession
		return nil
	}
	return errors.New("Unknown TunManagerRPC: " + s)
}

func (en *TunManagerRPC) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "CreateL2TunnelSession":
		*en = TunManagerRPCCreateL2TunnelSession
		return nil
	case "DeleteL2TunnelSession":
		*en = TunManagerRPCDeleteL2TunnelSession
		return nil
	}
	return errors.New("Unknown TunManagerRPC: " + s)
}

type TunnelType string

const TunnelTypeEoGRE TunnelType = "gretap"
const TunnelTypeGRE TunnelType = "gre"
const TunnelTypeL2TP TunnelType = "l2tpv3"

func (en TunnelType) GetPtr() *TunnelType { var v = en; return &v }

func (en TunnelType) String() string {
	switch en {
	case TunnelTypeEoGRE:
		return "gretap"
	case TunnelTypeGRE:
		return "gre"
	case TunnelTypeL2TP:
		return "l2tpv3"
	}
	if len(en) == 0 {
		return "l2tpv3"
	}
	panic(errors.New("Invalid value of TunnelType: " + string(en)))
}

func (en *TunnelType) MarshalJSON() ([]byte, error) {
	switch *en {
	case TunnelTypeEoGRE:
		return json.Marshal("gretap")
	case TunnelTypeGRE:
		return json.Marshal("gre")
	case TunnelTypeL2TP:
		return json.Marshal("l2tpv3")
	}
	if len(*en) == 0 {
		return json.Marshal("l2tpv3")
	}
	return nil, errors.New("Invalid value of TunnelType: " + string(*en))
}

func (en *TunnelType) GetBSON() (interface{}, error) {
	switch *en {
	case TunnelTypeEoGRE:
		return "gretap", nil
	case TunnelTypeGRE:
		return "gre", nil
	case TunnelTypeL2TP:
		return "l2tpv3", nil
	}
	if len(*en) == 0 {
		return "l2tpv3", nil
	}
	return nil, errors.New("Invalid value of TunnelType: " + string(*en))
}

func (en *TunnelType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "gretap":
		*en = TunnelTypeEoGRE
		return nil
	case "gre":
		*en = TunnelTypeGRE
		return nil
	case "l2tpv3":
		*en = TunnelTypeL2TP
		return nil
	}
	if len(s) == 0 {
		*en = TunnelTypeL2TP
		return nil
	}
	return errors.New("Unknown TunnelType: " + s)
}

func (en *TunnelType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "gretap":
		*en = TunnelTypeEoGRE
		return nil
	case "gre":
		*en = TunnelTypeGRE
		return nil
	case "l2tpv3":
		*en = TunnelTypeL2TP
		return nil
	}
	if len(s) == 0 {
		*en = TunnelTypeL2TP
		return nil
	}
	return errors.New("Unknown TunnelType: " + s)
}

type WMMAccessCategory string

const WMMAccessCategoryBackground WMMAccessCategory = "BK"
const WMMAccessCategoryBestEffort WMMAccessCategory = "BE"
const WMMAccessCategoryVideo WMMAccessCategory = "VI"
const WMMAccessCategoryVoice WMMAccessCategory = "VO"

func (en WMMAccessCategory) GetPtr() *WMMAccessCategory { var v = en; return &v }

func (en WMMAccessCategory) String() string {
	switch en {
	case WMMAccessCategoryBackground:
		return "BK"
	case WMMAccessCategoryBestEffort:
		return "BE"
	case WMMAccessCategoryVideo:
		return "VI"
	case WMMAccessCategoryVoice:
		return "VO"
	}
	if len(en) == 0 {
		return "BK"
	}
	panic(errors.New("Invalid value of WMMAccessCategory: " + string(en)))
}

func (en *WMMAccessCategory) MarshalJSON() ([]byte, error) {
	switch *en {
	case WMMAccessCategoryBackground:
		return json.Marshal("BK")
	case WMMAccessCategoryBestEffort:
		return json.Marshal("BE")
	case WMMAccessCategoryVideo:
		return json.Marshal("VI")
	case WMMAccessCategoryVoice:
		return json.Marshal("VO")
	}
	if len(*en) == 0 {
		return json.Marshal("BK")
	}
	return nil, errors.New("Invalid value of WMMAccessCategory: " + string(*en))
}

func (en *WMMAccessCategory) GetBSON() (interface{}, error) {
	switch *en {
	case WMMAccessCategoryBackground:
		return "BK", nil
	case WMMAccessCategoryBestEffort:
		return "BE", nil
	case WMMAccessCategoryVideo:
		return "VI", nil
	case WMMAccessCategoryVoice:
		return "VO", nil
	}
	if len(*en) == 0 {
		return "BK", nil
	}
	return nil, errors.New("Invalid value of WMMAccessCategory: " + string(*en))
}

func (en *WMMAccessCategory) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "BK":
		*en = WMMAccessCategoryBackground
		return nil
	case "BE":
		*en = WMMAccessCategoryBestEffort
		return nil
	case "VI":
		*en = WMMAccessCategoryVideo
		return nil
	case "VO":
		*en = WMMAccessCategoryVoice
		return nil
	}
	if len(s) == 0 {
		*en = WMMAccessCategoryBackground
		return nil
	}
	return errors.New("Unknown WMMAccessCategory: " + s)
}

func (en *WMMAccessCategory) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "BK":
		*en = WMMAccessCategoryBackground
		return nil
	case "BE":
		*en = WMMAccessCategoryBestEffort
		return nil
	case "VI":
		*en = WMMAccessCategoryVideo
		return nil
	case "VO":
		*en = WMMAccessCategoryVoice
		return nil
	}
	if len(s) == 0 {
		*en = WMMAccessCategoryBackground
		return nil
	}
	return errors.New("Unknown WMMAccessCategory: " + s)
}

type WirelessClientState string

const WirelessClientStateConnected WirelessClientState = "CONNECTED"
const WirelessClientStateDisconnected WirelessClientState = "DISCONNECTED"

func (en WirelessClientState) GetPtr() *WirelessClientState { var v = en; return &v }

func (en WirelessClientState) String() string {
	switch en {
	case WirelessClientStateConnected:
		return "CONNECTED"
	case WirelessClientStateDisconnected:
		return "DISCONNECTED"
	}
	panic(errors.New("Invalid value of WirelessClientState: " + string(en)))
}

func (en *WirelessClientState) MarshalJSON() ([]byte, error) {
	switch *en {
	case WirelessClientStateConnected:
		return json.Marshal("CONNECTED")
	case WirelessClientStateDisconnected:
		return json.Marshal("DISCONNECTED")
	}
	return nil, errors.New("Invalid value of WirelessClientState: " + string(*en))
}

func (en *WirelessClientState) GetBSON() (interface{}, error) {
	switch *en {
	case WirelessClientStateConnected:
		return "CONNECTED", nil
	case WirelessClientStateDisconnected:
		return "DISCONNECTED", nil
	}
	return nil, errors.New("Invalid value of WirelessClientState: " + string(*en))
}

func (en *WirelessClientState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*en = WirelessClientStateConnected
		return nil
	case "DISCONNECTED":
		*en = WirelessClientStateDisconnected
		return nil
	}
	return errors.New("Unknown WirelessClientState: " + s)
}

func (en *WirelessClientState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*en = WirelessClientStateConnected
		return nil
	case "DISCONNECTED":
		*en = WirelessClientStateDisconnected
		return nil
	}
	return errors.New("Unknown WirelessClientState: " + s)
}

type WirelessClientType string

const WirelessClientTypeCamera WirelessClientType = "camera"
const WirelessClientTypeOther WirelessClientType = "other"
const WirelessClientTypeWired WirelessClientType = "wired"

func (en WirelessClientType) GetPtr() *WirelessClientType { var v = en; return &v }

func (en WirelessClientType) String() string {
	switch en {
	case WirelessClientTypeCamera:
		return "camera"
	case WirelessClientTypeOther:
		return "other"
	case WirelessClientTypeWired:
		return "wired"
	}
	if len(en) == 0 {
		return "other"
	}
	panic(errors.New("Invalid value of WirelessClientType: " + string(en)))
}

func (en *WirelessClientType) MarshalJSON() ([]byte, error) {
	switch *en {
	case WirelessClientTypeCamera:
		return json.Marshal("camera")
	case WirelessClientTypeOther:
		return json.Marshal("other")
	case WirelessClientTypeWired:
		return json.Marshal("wired")
	}
	if len(*en) == 0 {
		return json.Marshal("other")
	}
	return nil, errors.New("Invalid value of WirelessClientType: " + string(*en))
}

func (en *WirelessClientType) GetBSON() (interface{}, error) {
	switch *en {
	case WirelessClientTypeCamera:
		return "camera", nil
	case WirelessClientTypeOther:
		return "other", nil
	case WirelessClientTypeWired:
		return "wired", nil
	}
	if len(*en) == 0 {
		return "other", nil
	}
	return nil, errors.New("Invalid value of WirelessClientType: " + string(*en))
}

func (en *WirelessClientType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "camera":
		*en = WirelessClientTypeCamera
		return nil
	case "other":
		*en = WirelessClientTypeOther
		return nil
	case "wired":
		*en = WirelessClientTypeWired
		return nil
	}
	if len(s) == 0 {
		*en = WirelessClientTypeOther
		return nil
	}
	return errors.New("Unknown WirelessClientType: " + s)
}

func (en *WirelessClientType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "camera":
		*en = WirelessClientTypeCamera
		return nil
	case "other":
		*en = WirelessClientTypeOther
		return nil
	case "wired":
		*en = WirelessClientTypeWired
		return nil
	}
	if len(s) == 0 {
		*en = WirelessClientTypeOther
		return nil
	}
	return errors.New("Unknown WirelessClientType: " + s)
}
