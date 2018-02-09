package libwimark

import (
	"encoding/json"
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type SystemEvent struct {
	SystemEventObject
	Timestamp  int64            `json:"timestamp"`
	Subject_id string           `json:"subject_id"`
	Level      SystemEventLevel `json:"level"`
}

func (self *SystemEvent) MarshalJSON() ([]byte, error) {
	var b []byte
	{
		var err error
		b, err = json.Marshal(self.SystemEventObject)
		if err != nil {
			return nil, err
		}
	}

	var doc Document
	{
		var err error
		err = json.Unmarshal(b, &doc)
		if err != nil {
			return nil, err
		}
	}

	doc["timestamp"] = self.Timestamp
	doc["subject_id"] = self.Subject_id
	doc["level"] = self.Level

	return json.Marshal(doc)
}

func (self *SystemEvent) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	var err = json.Unmarshal(b, &doc)
	if err != nil {
		return err
	}

	if doc == nil {
		return nil
	}

	// may be there is a way how to make code more reusable
	// don't have time to figure out

	var tsRaw, tsExists = doc["timestamp"]
	if tsExists {
		var ts int64
		var tsErr = json.Unmarshal(tsRaw, &ts)
		if tsErr == nil {
			self.Timestamp = ts
		} else {
			return tsErr
		}
	}

	delete(doc, "timestamp")

	var subRaw, subExists = doc["subject_id"]
	if subExists {
		var subject_id string
		var subErr = json.Unmarshal(subRaw, &subject_id)
		if subErr == nil {
			self.Subject_id = subject_id
		} else {
			return subErr
		}
	}

	delete(doc, "subject_id")

	var levelRaw, levelExists = doc["level"]
	if levelExists {
		var level SystemEventLevel
		var levErr = json.Unmarshal(levelRaw, &level)
		if levErr == nil {
			self.Level = level
		} else {
			return levErr
		}
	}

	delete(doc, "level")

	var v, _ = json.Marshal(doc)

	return self.SystemEventObject.UnmarshalJSON(v)
}

func (self *SystemEvent) GetBSON() (interface{}, error) {
	var out = bson.M{}

	levelBson, err := self.Level.GetBSON()
	if err != nil {
		return nil, err
	}

	var obj_b []byte
	{
		var err error
		obj_b, err = bson.Marshal(self.SystemEventObject)
		if err != nil {
			return nil, err
		}
	}

	var obj bson.M
	{
		var err error
		err = bson.Unmarshal(obj_b, &obj)
		if err != nil {
			return nil, err
		}
	}
	out = obj
	out["timestamp"] = self.Timestamp
	out["subject_id"] = self.Subject_id
	out["level"] = levelBson

	return out, nil
}

func (self *SystemEvent) SetBSON(raw bson.Raw) error {
	var in = map[string]bson.Raw{}
	{
		if err := raw.Unmarshal(&in); err != nil {
			return err
		}
	}

	//timestamp
	{
		var v, k_found = in["timestamp"]
		if !k_found {
			return errors.New("No timestamp found")
		}
		if err := v.Unmarshal(&self.Timestamp); err != nil {
			return err
		}

		delete(in, "timestamp")
	}

	//subject_id
	{
		var v, k_found = in["subject_id"]
		if !k_found {
			return errors.New("No subject_id found")
		}
		if err := v.Unmarshal(&self.Subject_id); err != nil {
			return err
		}

		delete(in, "subject_id")
	}

	//subject_id
	{
		var v, k_found = in["level"]
		if !k_found {
			return errors.New("No subject_id found")
		}
		if err := v.Unmarshal(&self.Level); err != nil {
			return err
		}

		delete(in, "level")
	}

	var obj_b, mErr = bson.Marshal(in)
	if mErr != nil {
		return mErr
	}

	if err := bson.Unmarshal(obj_b, &self.SystemEventObject); err != nil {
		return err
	}
	return nil
}

type CPEConnectedData Version

type MonitorRuleViolationData struct {
	Cpe_id  UUID `json:"cpe_id"`
	Stat_id UUID `json:"stat_id"`
}

type ClientConnectedData struct {
	Session_id string `json:"session_id"`
	Freq       string `json:"freq"`
	Cpe_id     UUID   `json:"cpe_id"`
	Wlan_id    UUID   `json:"wlan_id"`
	Radio_id   string `json:"radio_id"`
}

type ClientDisconnectedData struct {
	Session_id string `json:"session_id"`
	Freq       string `json:"freq"`
	Cpe_id     UUID   `json:"cpe_id"`
	Wlan_id    UUID   `json:"wlan_id"`
	Radio_id   string `json:"radio_id"`
}

type CPEConfigurationErrorData struct {
	Wifi   map[string]WifiConfigurationError `json:"wifi,omitempty"`
	Radius string                            `json:"radius,omitempty"`
}

type CPEInterfaceStateData struct {
	Interface string            `json:"radio_id"`
	WLAN      string            `json:"wlan_id"`
	EventName string            `json:"event_name"`
	EventType string            `json:"event_type"`
	State     CPEInterfaceState `json:"state"`
}

type WLANCentrAccChangeData struct {
	IsDeleted bool   `json:"is_deleted"`
	Radiuses  []UUID `json:"radiuses"`
}

type WifiConfigurationError struct {
	Config interface{}            `json:"config"`
	Wlans  map[string]interface{} `json:"wlans"`
}

type ServiceFatalErrorData struct {
	Sw_version  string `json:"sw_version"`
	Description string `json:"description"`
}

type ServiceConnectedData Version

type LBSClientData struct {
	Timestamp int64  `json:"timestamp"`
	CPE       UUID   `json:"cpe"`
	Radio     string `json:"radio"`
	ClientMac string `json:"client_mac"`
	RSSI      int    `json:"rssi"`
}

type RRMStatusData struct {
	Cpes []RRMCpeStatus `json:"cpes"`
}

type RRMCpeStatus struct {
	CpeID   UUID   `json:"cpe_id"`
	Radio   string `json:"radio"`
	Channel int    `json:"channel"`
	Dbm     int    `json:"dbm"`
}

type FirmwareUploadedData struct {
	CpeIDs []UUID `json:"cpe_ids"`
	Url    string `json:"url"`
	Md5Sum string `json:"md5sum"`
}

type CpeFirmwareData struct {
	AvailableMd5  string `json:"available_md5"`
	CurrentMd5    string `json:"current_md5"`
	NewFirmware   bool   `json:"new_firmware"`
	GoingToUpdate bool   `json:"going_to_update"`
	Error         string `json:"error"`
}
