package libwimark

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

const MQTT_ANY_WILDCARD = "+"

const (
	TOPIC_STATUS_FORMAT = `B/%s/%s`
	TOPIC_STATUS_REGEXP = `B/(.*)/(.*)`
	TOPIC_LOG_FORMAT    = `LOG/%s/%s`
	TOPIC_LOG_REGEXP    = `LOG/(.*)/(.*)`
	TOPIC_EVENT_FORMAT  = `EVENT/%s/%s`
	TOPIC_EVENT_REGEXP  = `EVENT/(.*)/(.*)`
	TOPIC_REQ_FORMAT    = `REQ/%s/%s/%s/%s/%s/%s`
	TOPIC_REQ_REGEXP    = `REQ/(.*)/(.*)/(.*)/(.*)/(.*)/(.*)`
	TOPIC_RSP_FORMAT    = `RSP/%s/%s/%s/%s/%s`
	TOPIC_RSP_REGEXP    = `RSP/(.*)/(.*)/(.*)/(.*)/(.*)`
)

type Topic interface {
	TopicPath() string
}

type BroadcastTopic struct {
	SenderModule Module
	SenderID     string
}

func ParseBroadcastTopic(s string, format string) *BroadcastTopic {
	var r = regexp.MustCompile(format)
	var ds = r.FindAllStringSubmatch(s, -1)
	if ds != nil && len(ds) == 1 {
		var data = ds[0]
		if data != nil && len(data) == 2+1 {
			var smodule Module
			var smodule_err = json.Unmarshal([]byte(strconv.Quote(data[1])), &smodule)
			if smodule_err == nil {
				var v BroadcastTopic

				v.SenderModule = smodule
				v.SenderID = data[2]

				return &v
			}
		}
	}

	return nil
}

func (self BroadcastTopic) TopicPathGeneric(format string) string {
	var sm, _ = self.SenderModule.MarshalJSON()
	var sm_s, _ = strconv.Unquote(string(sm))
	return fmt.Sprintf(format, sm_s, self.SenderID)
}

type StatusTopic BroadcastTopic

func ParseStatusTopic(s string) *StatusTopic {
	return (*StatusTopic)(ParseBroadcastTopic(s, TOPIC_STATUS_REGEXP))
}

func (self StatusTopic) TopicPath() string {
	return (BroadcastTopic)(self).TopicPathGeneric(TOPIC_STATUS_FORMAT)
}

type LogTopic BroadcastTopic

func ParseLogTopic(s string) *LogTopic {
	return (*LogTopic)(ParseBroadcastTopic(s, TOPIC_LOG_REGEXP))
}

func (self LogTopic) TopicPath() string {
	return (BroadcastTopic)(self).TopicPathGeneric(TOPIC_LOG_FORMAT)
}

type EventTopic BroadcastTopic

func ParseEventTopic(s string) *EventTopic {
	return (*EventTopic)(ParseBroadcastTopic(s, TOPIC_EVENT_REGEXP))
}

func (self EventTopic) TopicPath() string {
	return (BroadcastTopic)(self).TopicPathGeneric(TOPIC_EVENT_FORMAT)
}

type RequestTopic struct {
	SenderModule   Module
	SenderID       string
	ReceiverModule Module
	ReceiverID     string
	RequestID      string
	Operation      Operation
}

func (self RequestTopic) TopicPath() string {
	var u = strconv.Unquote

	var sm, _ = self.SenderModule.MarshalJSON()
	var sm_s, _ = u(string(sm))
	var rm, _ = self.ReceiverModule.MarshalJSON()
	var rm_s, _ = u(string(rm))
	var op, _ = self.Operation.MarshalJSON()
	var op_s, _ = u(string(op))
	return fmt.Sprintf(TOPIC_REQ_FORMAT, sm_s, self.SenderID, rm_s, self.ReceiverID, self.RequestID, op_s)
}

func (self RequestTopic) ToResponse() ResponseTopic {
	return ResponseTopic{
		SenderModule:   self.ReceiverModule,
		SenderID:       self.ReceiverID,
		ReceiverModule: self.SenderModule,
		ReceiverID:     self.SenderID,
		RequestID:      self.RequestID,
	}
}

func ParseRequestTopic(s string) *RequestTopic {
	var r = regexp.MustCompile(TOPIC_REQ_REGEXP)
	var ds = r.FindAllStringSubmatch(s, -1)
	if ds != nil && len(ds) == 1 {
		var data = ds[0]
		if data != nil && len(data) == 6+1 {
			var smodule Module
			var rmodule Module
			var op Operation

			var smodule_err = json.Unmarshal([]byte(strconv.Quote(data[1])), &smodule)
			var rmodule_err = json.Unmarshal([]byte(strconv.Quote(data[3])), &rmodule)
			var op_err = json.Unmarshal([]byte(strconv.Quote(data[6])), &op)

			if smodule_err == nil && rmodule_err == nil && op_err == nil {
				var v RequestTopic

				v.SenderModule = smodule
				v.SenderID = data[2]
				v.ReceiverModule = rmodule
				v.ReceiverID = data[4]
				v.RequestID = data[5]
				v.Operation = op

				return &v
			}
		}
	}

	return nil
}

type ResponseTopic struct {
	SenderModule   Module
	SenderID       string
	ReceiverModule Module
	ReceiverID     string
	RequestID      string
}

func (self ResponseTopic) TopicPath() string {
	var u = strconv.Unquote

	var sm, _ = self.SenderModule.MarshalJSON()
	var rm, _ = self.ReceiverModule.MarshalJSON()

	var sm_s, _ = u(string(sm))
	var rm_s, _ = u(string(rm))
	return fmt.Sprintf(TOPIC_RSP_FORMAT, sm_s, self.SenderID, rm_s, self.ReceiverID, self.RequestID)
}

func ParseResponseTopic(s string) *ResponseTopic {
	var r = regexp.MustCompile(TOPIC_RSP_REGEXP)
	var ds = r.FindAllStringSubmatch(s, -1)
	if ds != nil && len(ds) == 1 {
		var data = ds[0]
		if data != nil && len(data) == 5+1 {
			var smodule Module
			var rmodule Module

			var smodule_err = json.Unmarshal([]byte(strconv.Quote(data[1])), &smodule)
			var rmodule_err = json.Unmarshal([]byte(strconv.Quote(data[3])), &rmodule)

			if smodule_err == nil && rmodule_err == nil {
				var v ResponseTopic

				v.SenderModule = smodule
				v.SenderID = data[2]
				v.ReceiverModule = rmodule
				v.ReceiverID = data[4]
				v.RequestID = data[5]

				return &v
			}
		}
	}

	return nil
}

func ParseTopicPath(s string) Topic {
	{
		var v = ParseStatusTopic(s)
		if v != nil {
			return v
		}
	}
	{
		var v = ParseLogTopic(s)
		if v != nil {
			return nil
		}
	}
	{
		var v = ParseRequestTopic(s)
		if v != nil {
			return v
		}
	}
	{
		var v = ParseResponseTopic(s)
		if v != nil {
			return v
		}
	}

	return nil
}

type DBResponseBase struct {
	Errors []ModelError `json:"errors,omitempty"`
}

type DBDataObj struct {
	WLANs  map[UUID]WLAN               `json:"wlan,omitempty"`
	CPEs   map[UUID]CPE                `json:"cpe,omitempty"`
	Stats  map[UUID]Stat               `json:"stat,omitempty"`
	SDS    map[UUID]StatDaemonSettings `json:"stat-daemon-settings,omitempty"`
	Events map[UUID]Event              `json:"event,omitempty"`
}

type DBDataUUID struct {
	WLANs  []UUID `json:"wlan,omitempty"`
	CPEs   []UUID `json:"cpe,omitempty"`
	Stats  []UUID `json:"stat,omitempty"`
	SDS    []UUID `json:"stat-daemon-settings,omitempty"`
	Events []UUID `json:"event,omitempty"`
}

type DBDataMasks struct {
	WLANs  *SimpleMask `json:"wlan,omitempty"`
	CPEs   *CPEMask    `json:"cpe,omitempty"`
	Stats  *SimpleMask `json:"stat,omitempty"`
	SDS    *SimpleMask `json:"stat-daemon-settings,omitempty"`
	Events *SimpleMask `json:"event,omitempty"`
}

type DBRequestC DBDataObj
type DBRequestU DBDataObj
type DBRequestR DBDataMasks
type DBRequestD DBDataMasks

func (self *DBRequestC) UnmarshalJSON(b []byte) error {
	for k, _ := range self.WLANs {
		delete(self.WLANs, k)
	}
	for k, _ := range self.CPEs {
		delete(self.CPEs, k)
	}
	return json.Unmarshal(b, (*DBDataObj)(self))
}

func (self *DBRequestU) UnmarshalJSON(b []byte) error {
	for k, _ := range self.WLANs {
		delete(self.WLANs, k)
	}
	for k, _ := range self.CPEs {
		delete(self.CPEs, k)
	}
	return json.Unmarshal(b, (*DBDataObj)(self))
}

type DBResponseObj struct {
	DBResponseBase `json:",inline"`
	DBDataObj      `json:"data",inline`
}

type DBResponseUUID struct {
	DBResponseBase `json:",inline"`
	DBDataUUID     `json:"data",inline`
}

type DBResponseC DBResponseUUID
type DBResponseR DBResponseObj
type DBResponseU DBResponseUUID
type DBResponseD DBResponseUUID

type ConnectorInfo struct {
	Version   float64  `json:"version"`
	DbType    string   `json:"db_type"`
	DbServers []string `json:"db_servers"`
	Models    []string `json:"models"`
}
