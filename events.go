package libwimark

import (
	"encoding/json"
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type SystemEvent struct {
	SystemEventObject
	Timestamp   int64            `json:"timestamp"`
	Subject_id  string           `json:"subject_id"`
	Level       SystemEventLevel `json:"level"`
	Description string           `json:"description"`
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
	doc["description"] = self.Description

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

	var descRaw, descExists = doc["description"]
	if descExists {
		var desc string
		var desErr = json.Unmarshal(descRaw, &desc)
		if desErr == nil {
			self.Description = desc
		} else {
			return desErr
		}
	}

	delete(doc, "description")

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
	out["description"] = self.Description

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

	//description
	{
		var v, k_found = in["description"]
		if k_found {
			if err := v.Unmarshal(&self.Description); err != nil {
				return err
			}

			delete(in, "description")
		}

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

type CPEConnectedData struct {
	Version
	Template UUID `json:"template, omitempty"`
}

type MonitorRuleViolationData struct {
	Cpe_id  UUID `json:"cpe_id"`
	Stat_id UUID `json:"stat_id"`
}

type ClientConnectedData struct {
	Session_id string `json:"session_id"`
	Cpe_id     UUID   `json:"cpe_id"`
	Wlan_id    UUID   `json:"wlan_id"`
	Radio_id   string `json:"radio_id"`
	//Freq       string `json:"freq"`
}

type ClientDisconnectedData struct {
	Session_id string `json:"session_id"`
	Cpe_id     UUID   `json:"cpe_id"`
	Wlan_id    UUID   `json:"wlan_id"`
	Radio_id   string `json:"radio_id"`
	//Freq       string `json:"freq"`
}

type CPEInterfaceStateData struct {
	Interface string            `json:"radio_id"`
	State     CPEInterfaceState `json:"state"`
}

type WLANCentrAccChangeData struct {
	IsDeleted bool   `json:"is_deleted"`
	Radiuses  []UUID `json:"radiuses"`
}

type ServiceFatalErrorData struct {
	Sw_version  string `json:"sw_version"`
	Description string `json:"description"`
}

type ServiceConnectedData Version

type RRMStatusData struct {
	Cpes []RRMCpeStatus `json:"cpes"`
}

type RRMCpeStatus struct {
	CpeID   UUID       `json:"cpe_id"`
	Radio   string     `json:"radio"`
	Channel RRMChannel `json:"channel"`
	Power   int        `json:"power"`
	Error   bool       `json:"error"`
}

type RRMChannel struct {
	Channel int `json:"channel"`
	Mode    int `json:"mode"`   // 0-legacy, 1-HT, 2-VHT
	Width   int `json:"width"`  // 20, 40, 80, 160 (no 80+80 support yet)
	Offset  int `json:"offset"` // -1, 0, 1 - for HT only
	Central int `json:"center"`
}

type FirmwareUploadedData struct {
	CpeIDs []UUID             `json:"cpe_ids"`
	Url    string             `json:"url"`
	Md5Sum string             `json:"md5sum"`
	Mode   FirmwareUpdateMode `json:"mode"`
}

type CpeFirmwareData struct {
	AvailableMd5   string `json:"available_md5"`
	CurrentMd5     string `json:"current_md5"`
	NewFirmware    bool   `json:"new_firmware"`
	GoingToUpgrade bool   `json:"going_to_upgrade"`
	Error          string `json:"error"`
}

type RadiusAccountingSendData struct {
	RadiusList []Radius            `json:"radius_list"`
	Message    RadiusMessageObject `json:"message"`
}
