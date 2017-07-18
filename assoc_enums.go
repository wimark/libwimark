package libwimark

import "encoding/json"
import "errors"

type CPEAgentError struct {
	Type CPEAgentStatusType `json:"type"`
	Data interface{}        `json:"data"`
}

func (self *CPEAgentError) UnmarshalJSON(b []byte) error {
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
	_ = data_found
	var t CPEAgentStatusType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t.Value().(type) {
	case CPEAgentStatusException:
		if !data_found {
			return errors.New("No associated data found for enum CPEAgentError")
		}
		var d string
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case CPEAgentStatusSuccess:
		break
	case CPEAgentStatusSyntaxError:
		if !data_found {
			return errors.New("No associated data found for enum CPEAgentError")
		}
		var d string
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case CPEAgentStatusUndefined:
		if !data_found {
			return errors.New("No associated data found for enum CPEAgentError")
		}
		var d string
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d

	}
	self.Type = t
	return nil

}

type CPEInterfaceInfo struct {
	Type CPEInterfaceType `json:"type"`
	Data interface{}      `json:"data"`
}

func (self *CPEInterfaceInfo) UnmarshalJSON(b []byte) error {
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
	_ = data_found
	var t CPEInterfaceType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t.Value().(type) {
	case InterfaceWiFi:
		if !data_found {
			return errors.New("No associated data found for enum CPEInterfaceInfo")
		}
		var d WiFiData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case InterfaceWired:
		if !data_found {
			return errors.New("No associated data found for enum CPEInterfaceInfo")
		}
		var d WiredData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d

	}
	self.Type = t
	return nil

}

type EnumSecurity struct {
	Type SecurityType `json:"type"`
	Data interface{}  `json:"data"`
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
	_ = data_found
	var t SecurityType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t.Value().(type) {
	case WPA2Enterprise:
		if !data_found {
			return errors.New("No associated data found for enum EnumSecurity")
		}
		var d WPA2EnterpriseData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case WPA2Personal:
		if !data_found {
			return errors.New("No associated data found for enum EnumSecurity")
		}
		var d WPA2PersonalData
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d

	}
	self.Type = t
	return nil

}

type StatEventRuleObject struct {
	Type StatEventRuleType `json:"type"`
	Data interface{}       `json:"data"`
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
	_ = data_found
	var t StatEventRuleType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t.Value().(type) {
	case StatEventCPUload:
		if !data_found {
			return errors.New("No associated data found for enum StatEventRuleObject")
		}
		var d LimitBetween
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case StatEventFreeRAM:
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

type WirelessClientObject struct {
	Type WirelessClientType `json:"type"`
	Data interface{}        `json:"data"`
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
	_ = data_found
	var t WirelessClientType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	switch t.Value().(type) {
	case WirelessClientCamera:
		if !data_found {
			return errors.New("No associated data found for enum WirelessClientObject")
		}
		var d CameraClient
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d
	case WirelessClientOther:
		if !data_found {
			return errors.New("No associated data found for enum WirelessClientObject")
		}
		var d OtherClient
		var data_err = json.Unmarshal(data_raw, &d)
		if data_err != nil {
			return data_err
		}
		self.Data = &d

	}
	self.Type = t
	return nil

}
