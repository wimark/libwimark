package libwimark

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

const MQTT_ANY_WILDCARD = "+"

const (
	TOPIC_B_FORMAT   = `B/%s/%s`
	TOPIC_B_REGEXP   = `B/(.*)/(.*)`
	TOPIC_REQ_FORMAT = `REQ/%s/%s/%s/%s/%s/%s`
	TOPIC_REQ_REGEXP = `REQ/(.*)/(.*)/(.*)/(.*)/(.*)/(.*)`
	TOPIC_RSP_FORMAT = `RSP/%s/%s/%s/%s/%s`
	TOPIC_RSP_REGEXP = `RSP/(.*)/(.*)/(.*)/(.*)/(.*)`
)

type Topic interface {
	TopicPath() string
}

type BroadcastTopic struct {
	SenderModule Module
	SenderID     string
}

func ParseBroadcastTopic(s string) *BroadcastTopic {
	var r = regexp.MustCompile(TOPIC_B_REGEXP)
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

func (self *BroadcastTopic) TopicPath() string {
	var sm, _ = self.SenderModule.MarshalJSON()
	var sm_s, _ = strconv.Unquote(string(sm))
	return fmt.Sprintf(TOPIC_B_FORMAT, sm_s, self.SenderID)
}

type RequestTopic struct {
	SenderModule   Module
	SenderID       string
	ReceiverModule Module
	ReceiverID     string
	RequestID      string
	Operation      Operation
}

func (self *RequestTopic) TopicPath() string {
	var u = strconv.Unquote

	var sm,   _ = self.SenderModule.MarshalJSON()
	var sm_s, _ = u(string(sm))
	var rm, _   = self.ReceiverModule.MarshalJSON()
	var rm_s, _ = u(string(rm))
	var op, _   = self.Operation.MarshalJSON()
	var op_s, _ = u(string(op))
	return fmt.Sprintf(TOPIC_REQ_FORMAT, sm_s, self.SenderID, rm_s, self.ReceiverID, self.RequestID, op_s)
}

func (self *RequestTopic) ToResponse() ResponseTopic {
	return ResponseTopic{
		SenderModule:   self.SenderModule,
		SenderID:       self.SenderID,
		ReceiverModule: self.ReceiverModule,
		ReceiverID:     self.ReceiverID,
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

func (self *ResponseTopic) TopicPath() string {
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
		var v = ParseBroadcastTopic(s)
		if v != nil {
			return v
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
	Errors []ModelError `json:"errors"`
}

type DBDataObj struct {
	WLANs map[UUID]WLAN `json:"wlan"`
	CPEs  map[UUID]CPE  `json:"cpe"`
}

type DBDataUUID struct {
	WLANs []UUID `json:"wlan"`
	CPEs  []UUID `json:"cpe"`
}

type DBRequestC DBDataObj
type DBRequestU DBDataObj
type DBRequestR DBDataUUID
type DBRequestD DBDataUUID

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
