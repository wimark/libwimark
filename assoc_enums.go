package libwimark

import (
	"bytes"
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type EnumSecurity struct {
	Type SecurityType "json:\"type\""
	Data interface{}  "json:\"data\""
}

func (self *EnumSecurity) UnmarshalJSON(b []byte) error {
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
		break
	case SecurityTypeWPA2Enterprise:
		if !data_found {
			return errors.New("No associated data found for enum EnumSecurity")
		}
		var d WPA2EnterpriseData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SecurityTypeWPA2Personal:
		if !data_found {
			return errors.New("No associated data found for enum EnumSecurity")
		}
		var d WPA2PersonalData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SecurityTypeWPAEnterprise:
		if !data_found {
			return errors.New("No associated data found for enum EnumSecurity")
		}
		var d WPAEnterpriseData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SecurityTypeWPAPersonal:
		if !data_found {
			return errors.New("No associated data found for enum EnumSecurity")
		}
		var d WPAPersonalData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	}
	self.Type = t
	return nil
}

func (self *EnumSecurity) SetBSON(v bson.Raw) error {
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
		break
	case SecurityTypeWPA2Enterprise:
		if !data_found {
			return errors.New("No associated data found for enum EnumSecurity")
		}
		var d WPA2EnterpriseData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SecurityTypeWPA2Personal:
		if !data_found {
			return errors.New("No associated data found for enum EnumSecurity")
		}
		var d WPA2PersonalData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SecurityTypeWPAEnterprise:
		if !data_found {
			return errors.New("No associated data found for enum EnumSecurity")
		}
		var d WPAEnterpriseData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SecurityTypeWPAPersonal:
		if !data_found {
			return errors.New("No associated data found for enum EnumSecurity")
		}
		var d WPAPersonalData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	}
	self.Type = t
	return nil
}

type RadiusMessageObject struct {
	Type RadiusMessageType "json:\"type\""
	Data interface{}       "json:\"data\""
}

func (self *RadiusMessageObject) UnmarshalJSON(b []byte) error {
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
			return errors.New("No associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessAccept
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case RadiusMessageTypeAccessReject:
		if !data_found {
			return errors.New("No associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessReject
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case RadiusMessageTypeAccessRequest:
		if !data_found {
			return errors.New("No associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessRequest
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case RadiusMessageTypeAccountingRequest:
		if !data_found {
			return errors.New("No associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccountingRequest
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	}
	self.Type = t
	return nil
}

func (self *RadiusMessageObject) SetBSON(v bson.Raw) error {
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
			return errors.New("No associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessAccept
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case RadiusMessageTypeAccessReject:
		if !data_found {
			return errors.New("No associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessReject
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case RadiusMessageTypeAccessRequest:
		if !data_found {
			return errors.New("No associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccessRequest
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case RadiusMessageTypeAccountingRequest:
		if !data_found {
			return errors.New("No associated data found for enum RadiusMessageObject")
		}
		var d RadiusAccountingRequest
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	}
	self.Type = t
	return nil
}

type StatEventRuleObject struct {
	Type StatEventRuleType "json:\"type\""
	Data interface{}       "json:\"data\""
}

func (self *StatEventRuleObject) UnmarshalJSON(b []byte) error {
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
			return errors.New("No associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case StatEventRuleTypeFreeRAM:
		if !data_found {
			return errors.New("No associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	}
	self.Type = t
	return nil
}

func (self *StatEventRuleObject) SetBSON(v bson.Raw) error {
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
			return errors.New("No associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case StatEventRuleTypeFreeRAM:
		if !data_found {
			return errors.New("No associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	}
	self.Type = t
	return nil
}

type SystemEventObject struct {
	Type SystemEventType "json:\"type\""
	Data interface{}     "json:\"data\""
}

func (self *SystemEventObject) UnmarshalJSON(b []byte) error {
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
	case SystemEventTypeCPEConfigurationError:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d CPEConfigurationErrorData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeCPEConfigurationSuccess:
		break
	case SystemEventTypeCPEConnected:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d CPEConnectedData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeCPEDisconnected:
		break
	case SystemEventTypeCPEInterfaceState:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d CPEInterfaceStateData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeClientConnected:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d ClientConnectedData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeClientDisconnected:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d ClientDisconnectedData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeCpeFirmwareAvailable:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d CpeFirmwareData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeDaemonSettingsChanged:
		break
	case SystemEventTypeFirmwareUploaded:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d FirmwareUploadedData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeMonitorRuleViolation:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d MonitorRuleViolationData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeRRMStatus:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d RRMStatusData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeRadiusAccountingSend:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d RadiusAccountingSendData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeServiceConnected:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d ServiceConnectedData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeServiceDisconnected:
		break
	case SystemEventTypeServiceFatalError:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d ServiceFatalErrorData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeSystemTimeChanged:
		break
	case SystemEventTypeWLANCentrAccChanged:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d WLANCentrAccChangeData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	}
	self.Type = t
	return nil
}

func (self *SystemEventObject) SetBSON(v bson.Raw) error {
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
	case SystemEventTypeCPEConfigurationError:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d CPEConfigurationErrorData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeCPEConfigurationSuccess:
		break
	case SystemEventTypeCPEConnected:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d CPEConnectedData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeCPEDisconnected:
		break
	case SystemEventTypeCPEInterfaceState:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d CPEInterfaceStateData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeClientConnected:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d ClientConnectedData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeClientDisconnected:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d ClientDisconnectedData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeCpeFirmwareAvailable:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d CpeFirmwareData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeDaemonSettingsChanged:
		break
	case SystemEventTypeFirmwareUploaded:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d FirmwareUploadedData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeMonitorRuleViolation:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d MonitorRuleViolationData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeRRMStatus:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d RRMStatusData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeRadiusAccountingSend:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d RadiusAccountingSendData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeServiceConnected:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d ServiceConnectedData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeServiceDisconnected:
		break
	case SystemEventTypeServiceFatalError:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d ServiceFatalErrorData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case SystemEventTypeSystemTimeChanged:
		break
	case SystemEventTypeWLANCentrAccChanged:
		if !data_found {
			return errors.New("No associated data found for enum SystemEventObject")
		}
		var d WLANCentrAccChangeData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	}
	self.Type = t
	return nil
}

type WirelessClientObject struct {
	Type WirelessClientType "json:\"type\""
	Data interface{}        "json:\"data\""
}

func (self *WirelessClientObject) UnmarshalJSON(b []byte) error {
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
			return errors.New("No associated data found for enum WirelessClientObject")
		}
		var d CameraClientData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case WirelessClientTypeOther:
		if !data_found {
			return errors.New("No associated data found for enum WirelessClientObject")
		}
		var d OtherClientData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	}
	self.Type = t
	return nil
}

func (self *WirelessClientObject) SetBSON(v bson.Raw) error {
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
			return errors.New("No associated data found for enum WirelessClientObject")
		}
		var d CameraClientData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case WirelessClientTypeOther:
		if !data_found {
			return errors.New("No associated data found for enum WirelessClientObject")
		}
		var d OtherClientData
		var data_err = data_raw.Unmarshal(&d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	}
	self.Type = t
	return nil
}
