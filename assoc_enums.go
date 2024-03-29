package libwimark

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/globalsign/mgo/bson"
)

type EnumSecurity struct {
	Type SecurityType "json:\"type\""
	Data interface{}  "json:\"data\""
}

func (en *EnumSecurity) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(b, &doc); err != nil {
		return err
	}
	if doc == nil {
		return nil
	}
	var t_raw, t_found = doc["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = doc["data"]
	if bytes.Equal(data_raw, []byte("null")) {
		data_found = false
	}
	var t SecurityType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t {
	case SecurityTypeNone:
		en.Data = nil
	case SecurityTypeWPA23Enterprise:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2EnterpriseData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPA23Personal:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2PersonalData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPA2Enterprise:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2EnterpriseData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPA2Personal:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2PersonalData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPA3Enterprise:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2EnterpriseData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPA3Personal:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2PersonalData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPAEnterprise:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPAEnterpriseData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPAPersonal:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPAPersonalData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	}
	en.Type = t
	return nil
}

func (en *EnumSecurity) SetBSON(v bson.Raw) error {
	var in = map[string]bson.Raw{}
	if err := v.Unmarshal(&in); err != nil {
		return err
	}
	if in == nil {
		return nil
	}
	var t_raw, t_found = in["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = in["data"]
	if bytes.Equal(data_raw.Data, []byte("null")) {
		data_found = false
	}
	var t SecurityType
	if t_err := t_raw.Unmarshal(&t); t_err != nil {
		return t_err
	}
	switch t {
	case SecurityTypeNone:
		en.Data = nil
	case SecurityTypeWPA23Enterprise:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2EnterpriseData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPA23Personal:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2PersonalData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPA2Enterprise:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2EnterpriseData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPA2Personal:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2PersonalData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPA3Enterprise:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2EnterpriseData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPA3Personal:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPA2PersonalData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPAEnterprise:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPAEnterpriseData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SecurityTypeWPAPersonal:
		if !data_found {
			return errors.New("no associated data found for enum EnumSecurity")
		}
		var d WPAPersonalData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	}
	en.Type = t
	return nil
}

type RRMAlgoObject struct {
	Type RRMAlgoType "json:\"type\""
	Data interface{} "json:\"data\""
}

func (en *RRMAlgoObject) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(b, &doc); err != nil {
		return err
	}
	if doc == nil {
		return nil
	}
	var t_raw, t_found = doc["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = doc["data"]
	if bytes.Equal(data_raw, []byte("null")) {
		data_found = false
	}
	var t RRMAlgoType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t {
	case RRMAlgoTypeBlind:
		if !data_found {
			return errors.New("no associated data found for enum RRMAlgoObject")
		}
		var d RRMTimerParams
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case RRMAlgoTypeDummy:
		if !data_found {
			return errors.New("no associated data found for enum RRMAlgoObject")
		}
		var d RRMTimerParams
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case RRMAlgoTypeGreed:
		if !data_found {
			return errors.New("no associated data found for enum RRMAlgoObject")
		}
		var d RRMGreedParams
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	}
	en.Type = t
	return nil
}

func (en *RRMAlgoObject) SetBSON(v bson.Raw) error {
	var in = map[string]bson.Raw{}
	if err := v.Unmarshal(&in); err != nil {
		return err
	}
	if in == nil {
		return nil
	}
	var t_raw, t_found = in["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = in["data"]
	if bytes.Equal(data_raw.Data, []byte("null")) {
		data_found = false
	}
	var t RRMAlgoType
	if t_err := t_raw.Unmarshal(&t); t_err != nil {
		return t_err
	}
	switch t {
	case RRMAlgoTypeBlind:
		if !data_found {
			return errors.New("no associated data found for enum RRMAlgoObject")
		}
		var d RRMTimerParams
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case RRMAlgoTypeDummy:
		if !data_found {
			return errors.New("no associated data found for enum RRMAlgoObject")
		}
		var d RRMTimerParams
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case RRMAlgoTypeGreed:
		if !data_found {
			return errors.New("no associated data found for enum RRMAlgoObject")
		}
		var d RRMGreedParams
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	}
	en.Type = t
	return nil
}

type RadiusMessageObject struct {
	Type RadiusMessageType "json:\"type\""
	Data interface{}       "json:\"data\""
}

func (en *RadiusMessageObject) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(b, &doc); err != nil {
		return err
	}
	if doc == nil {
		return nil
	}
	var t_raw, t_found = doc["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = doc["data"]
	if bytes.Equal(data_raw, []byte("null")) {
		data_found = false
	}
	var t RadiusMessageType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t {
	case RadiusMessageTypeAccessAccept:
		if !data_found {
			return errors.New("no associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessAccept
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case RadiusMessageTypeAccessReject:
		if !data_found {
			return errors.New("no associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessReject
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case RadiusMessageTypeAccessRequest:
		if !data_found {
			return errors.New("no associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessRequest
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case RadiusMessageTypeAccountingRequest:
		if !data_found {
			return errors.New("no associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccountingRequest
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	}
	en.Type = t
	return nil
}

func (en *RadiusMessageObject) SetBSON(v bson.Raw) error {
	var in = map[string]bson.Raw{}
	if err := v.Unmarshal(&in); err != nil {
		return err
	}
	if in == nil {
		return nil
	}
	var t_raw, t_found = in["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = in["data"]
	if bytes.Equal(data_raw.Data, []byte("null")) {
		data_found = false
	}
	var t RadiusMessageType
	if t_err := t_raw.Unmarshal(&t); t_err != nil {
		return t_err
	}
	switch t {
	case RadiusMessageTypeAccessAccept:
		if !data_found {
			return errors.New("no associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessAccept
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case RadiusMessageTypeAccessReject:
		if !data_found {
			return errors.New("no associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessReject
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case RadiusMessageTypeAccessRequest:
		if !data_found {
			return errors.New("no associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessRequest
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case RadiusMessageTypeAccountingRequest:
		if !data_found {
			return errors.New("no associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccountingRequest
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	}
	en.Type = t
	return nil
}

type StatEventRuleObject struct {
	Type StatEventRuleType "json:\"type\""
	Data interface{}       "json:\"data\""
}

func (en *StatEventRuleObject) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(b, &doc); err != nil {
		return err
	}
	if doc == nil {
		return nil
	}
	var t_raw, t_found = doc["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = doc["data"]
	if bytes.Equal(data_raw, []byte("null")) {
		data_found = false
	}
	var t StatEventRuleType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t {
	case StatEventRuleTypeCPUload:
		if !data_found {
			return errors.New("no associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case StatEventRuleTypeClientCon:
		en.Data = nil
	case StatEventRuleTypeClientDis:
		en.Data = nil
	case StatEventRuleTypeClientFar:
		if !data_found {
			return errors.New("no associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case StatEventRuleTypeConfigError:
		en.Data = nil
	case StatEventRuleTypeConnected:
		en.Data = nil
	case StatEventRuleTypeCustomerActivity:
		if !data_found {
			return errors.New("no associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case StatEventRuleTypeDisconnected:
		en.Data = nil
	case StatEventRuleTypeFreeRAM:
		if !data_found {
			return errors.New("no associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case StatEventRuleTypeIfaceError:
		en.Data = nil
	}
	en.Type = t
	return nil
}

func (en *StatEventRuleObject) SetBSON(v bson.Raw) error {
	var in = map[string]bson.Raw{}
	if err := v.Unmarshal(&in); err != nil {
		return err
	}
	if in == nil {
		return nil
	}
	var t_raw, t_found = in["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = in["data"]
	if bytes.Equal(data_raw.Data, []byte("null")) {
		data_found = false
	}
	var t StatEventRuleType
	if t_err := t_raw.Unmarshal(&t); t_err != nil {
		return t_err
	}
	switch t {
	case StatEventRuleTypeCPUload:
		if !data_found {
			return errors.New("no associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case StatEventRuleTypeClientCon:
		en.Data = nil
	case StatEventRuleTypeClientDis:
		en.Data = nil
	case StatEventRuleTypeClientFar:
		if !data_found {
			return errors.New("no associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case StatEventRuleTypeConfigError:
		en.Data = nil
	case StatEventRuleTypeConnected:
		en.Data = nil
	case StatEventRuleTypeCustomerActivity:
		if !data_found {
			return errors.New("no associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case StatEventRuleTypeDisconnected:
		en.Data = nil
	case StatEventRuleTypeFreeRAM:
		if !data_found {
			return errors.New("no associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case StatEventRuleTypeIfaceError:
		en.Data = nil
	}
	en.Type = t
	return nil
}

type SystemEventObject struct {
	Type SystemEventType "json:\"type\""
	Data interface{}     "json:\"data\""
}

func (en *SystemEventObject) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(b, &doc); err != nil {
		return err
	}
	if doc == nil {
		return nil
	}
	var t_raw, t_found = doc["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = doc["data"]
	if bytes.Equal(data_raw, []byte("null")) {
		data_found = false
	}
	var t SystemEventType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t {
	case SystemEventTypeCPEConnected:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d CPEConnectedData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeCPEDisconnected:
		en.Data = nil
	case SystemEventTypeCPEInterfaceState:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d CPEInterfaceStateData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeClientAuthorization:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ClientAuthorizationData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeClientConnected:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ClientConnectedData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeClientDisconnected:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ClientDisconnectedData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeCpeFirmwareAvailable:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d CpeFirmwareData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeDHCPAck:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d DHCPAckData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeDaemonSettingsChanged:
		en.Data = nil
	case SystemEventTypeFirmwareUploaded:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d FirmwareUploadedData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeLocationCacheReload:
		en.Data = nil
	case SystemEventTypeLoggedError:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ModelError
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeMonitorRuleViolation:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d MonitorRuleViolationData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeRRMStatus:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d RRMStatusData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeRadarExportUpdate:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d RadarExportUpdate
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeRadiusAccountingSend:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d RadiusAccountingSendData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeServiceConnected:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ServiceConnectedData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeServiceDisconnected:
		en.Data = nil
	case SystemEventTypeServiceFatalError:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ServiceFatalErrorData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeSystemTimeChanged:
		en.Data = nil
	case SystemEventTypeUserAuthSuccess:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d UserAuthorizationSuccessData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeWLANCentrAccChanged:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d WLANCentrAccChangeData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	}
	en.Type = t
	return nil
}

func (en *SystemEventObject) SetBSON(v bson.Raw) error {
	var in = map[string]bson.Raw{}
	if err := v.Unmarshal(&in); err != nil {
		return err
	}
	if in == nil {
		return nil
	}
	var t_raw, t_found = in["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = in["data"]
	if bytes.Equal(data_raw.Data, []byte("null")) {
		data_found = false
	}
	var t SystemEventType
	if t_err := t_raw.Unmarshal(&t); t_err != nil {
		return t_err
	}
	switch t {
	case SystemEventTypeCPEConnected:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d CPEConnectedData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeCPEDisconnected:
		en.Data = nil
	case SystemEventTypeCPEInterfaceState:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d CPEInterfaceStateData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeClientAuthorization:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ClientAuthorizationData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeClientConnected:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ClientConnectedData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeClientDisconnected:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ClientDisconnectedData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeCpeFirmwareAvailable:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d CpeFirmwareData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeDHCPAck:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d DHCPAckData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeDaemonSettingsChanged:
		en.Data = nil
	case SystemEventTypeFirmwareUploaded:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d FirmwareUploadedData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeLocationCacheReload:
		en.Data = nil
	case SystemEventTypeLoggedError:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ModelError
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeMonitorRuleViolation:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d MonitorRuleViolationData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeRRMStatus:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d RRMStatusData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeRadarExportUpdate:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d RadarExportUpdate
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeRadiusAccountingSend:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d RadiusAccountingSendData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeServiceConnected:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ServiceConnectedData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeServiceDisconnected:
		en.Data = nil
	case SystemEventTypeServiceFatalError:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d ServiceFatalErrorData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeSystemTimeChanged:
		en.Data = nil
	case SystemEventTypeUserAuthSuccess:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d UserAuthorizationSuccessData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case SystemEventTypeWLANCentrAccChanged:
		if !data_found {
			return errors.New("no associated data found for enum SystemEventObject")
		}
		var d WLANCentrAccChangeData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	}
	en.Type = t
	return nil
}

type WirelessClientObject struct {
	Type WirelessClientType "json:\"type\""
	Data interface{}        "json:\"data\""
}

func (en *WirelessClientObject) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(b, &doc); err != nil {
		return err
	}
	if doc == nil {
		return nil
	}
	var t_raw, t_found = doc["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = doc["data"]
	if bytes.Equal(data_raw, []byte("null")) {
		data_found = false
	}
	var t WirelessClientType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t {
	case WirelessClientTypeCamera:
		if !data_found {
			return errors.New("no associated data found for enum WirelessClientObject")
		}
		var d CameraClientData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case WirelessClientTypeOther:
		if !data_found {
			return errors.New("no associated data found for enum WirelessClientObject")
		}
		var d OtherClientData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case WirelessClientTypeWired:
		if !data_found {
			return errors.New("no associated data found for enum WirelessClientObject")
		}
		var d OtherClientData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	}
	en.Type = t
	return nil
}

func (en *WirelessClientObject) SetBSON(v bson.Raw) error {
	var in = map[string]bson.Raw{}
	if err := v.Unmarshal(&in); err != nil {
		return err
	}
	if in == nil {
		return nil
	}
	var t_raw, t_found = in["type"]
	if !t_found {
		return nil
	}
	var data_raw, data_found = in["data"]
	if bytes.Equal(data_raw.Data, []byte("null")) {
		data_found = false
	}
	var t WirelessClientType
	if t_err := t_raw.Unmarshal(&t); t_err != nil {
		return t_err
	}
	switch t {
	case WirelessClientTypeCamera:
		if !data_found {
			return errors.New("no associated data found for enum WirelessClientObject")
		}
		var d CameraClientData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case WirelessClientTypeOther:
		if !data_found {
			return errors.New("no associated data found for enum WirelessClientObject")
		}
		var d OtherClientData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	case WirelessClientTypeWired:
		if !data_found {
			return errors.New("no associated data found for enum WirelessClientObject")
		}
		var d OtherClientData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		en.Data = &d
	}
	en.Type = t
	return nil
}
