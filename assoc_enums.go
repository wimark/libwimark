package libwimark

import (
	"github.com/vorot93/goutil"
)
import (
	"encoding/json"
)

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
	var data_m = goutil.Document{}
	var s_err = json.Unmarshal(data_raw, &data_m)
	if s_err != nil {
		return s_err
	}
	var data interface{}
	var data_err error
	switch t.Value().(type) {
	case InterfaceWiFi:
		data, data_err = data_m.ToValue(func() interface{} { return &WiFiData{} })
	case InterfaceWired:
		data, data_err = data_m.ToValue(func() interface{} { return &WiredData{} })

	}
	if data_err != nil {
		return data_err
	}
	self.T = t
	self.D = data.(CPEInterfaceType)
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
	var data_m = goutil.Document{}
	var s_err = json.Unmarshal(data_raw, &data_m)
	if s_err != nil {
		return s_err
	}
	var data interface{}
	var data_err error
	switch t.Value().(type) {
	case WPA2Enterprise:
		data, data_err = data_m.ToValue(func() interface{} { return &WPA2EnterpriseData{} })
	case WPA2Personal:
		data, data_err = data_m.ToValue(func() interface{} { return &WPA2PersonalData{} })

	}
	if data_err != nil {
		return data_err
	}
	self.T = t
	self.D = data.(SecurityType)
	return nil

}
