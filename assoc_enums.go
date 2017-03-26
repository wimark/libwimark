package libwimark

import (
	"github.com/vorot93/goutil"
)
import (
	"encoding/json"
)

type CPEAgentError struct {
	T CPEAgentStatusType
	D interface{}
}

func (self *CPEAgentError) MarshalJSON() ([]byte, error) {
	var t, terr = json.Marshal(self.T)
	if terr != nil {
		return nil, terr
	}
	var d, derr = json.Marshal(self.D)
	if derr != nil {
		return nil, derr
	}
	return json.Marshal(goutil.Document{"type": t, "data": d})

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
	if !data_found {
		return nil
	}
	var t CPEAgentStatusType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	var data interface{} = nil
	var data_err error
	switch t.Value().(type) {
	case CPEAgentStatusException:
		var d string
		data_err = json.Unmarshal(data_raw, &d)
		data = &d
	case CPEAgentStatusSuccess:
		var d goutil.Document
		data_err = json.Unmarshal(data_raw, &d)
		data = &d
	case CPEAgentStatusSyntaxError:
		var d string
		data_err = json.Unmarshal(data_raw, &d)
		data = &d
	case CPEAgentStatusUndefined:
		var d string
		data_err = json.Unmarshal(data_raw, &d)
		data = &d

	}
	if data_err != nil {
		return data_err
	}
	self.T = t
	self.D = data
	return nil

}

type CPEInterfaceInfo struct {
	T CPEInterfaceType
	D interface{}
}

func (self *CPEInterfaceInfo) MarshalJSON() ([]byte, error) {
	var t, terr = json.Marshal(self.T)
	if terr != nil {
		return nil, terr
	}
	var d, derr = json.Marshal(self.D)
	if derr != nil {
		return nil, derr
	}
	return json.Marshal(goutil.Document{"type": t, "data": d})

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
	if !data_found {
		return nil
	}
	var t CPEInterfaceType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	var data interface{} = nil
	var data_err error
	switch t.Value().(type) {
	case InterfaceWiFi:
		var d WiFiData
		data_err = json.Unmarshal(data_raw, &d)
		data = &d
	case InterfaceWired:
		var d WiredData
		data_err = json.Unmarshal(data_raw, &d)
		data = &d

	}
	if data_err != nil {
		return data_err
	}
	self.T = t
	self.D = data
	return nil

}

type EnumSecurity struct {
	T SecurityType
	D interface{}
}

func (self *EnumSecurity) MarshalJSON() ([]byte, error) {
	var t, terr = json.Marshal(self.T)
	if terr != nil {
		return nil, terr
	}
	var d, derr = json.Marshal(self.D)
	if derr != nil {
		return nil, derr
	}
	return json.Marshal(goutil.Document{"type": t, "data": d})

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
	if !data_found {
		return nil
	}
	var t SecurityType
	if t_err := json.Unmarshal(t_raw, &t); t_err != nil {
		return t_err
	}
	var data interface{} = nil
	var data_err error
	switch t.Value().(type) {
	case WPA2Enterprise:
		var d WPA2EnterpriseData
		data_err = json.Unmarshal(data_raw, &d)
		data = &d
	case WPA2Personal:
		var d WPA2PersonalData
		data_err = json.Unmarshal(data_raw, &d)
		data = &d

	}
	if data_err != nil {
		return data_err
	}
	self.T = t
	self.D = data
	return nil

}
