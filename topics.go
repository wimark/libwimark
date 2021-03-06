package libwimark

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const MQTT_ANY_WILDCARD = "+"
const MQTT_MULTILEVEL_WILDCARD = "#"

const (
	TOPIC_STATUS_FORMAT = `B/%s/%s`
	TOPIC_STATUS_REGEXP = `B/(.*)/(.*)`
	TOPIC_LOG_FORMAT    = `LOG/%s/%s`
	TOPIC_LOG_REGEXP    = `LOG/(.*)/(.*)`
	TOPIC_EVENT_FORMAT  = `EVENT/%s/%s/%s`
	TOPIC_EVENT_REGEXP  = `EVENT/(.*)/(.*)/(.*)`
	TOPIC_REQ_FORMAT    = `REQ/%s/%s/%s/%s/%s/%s/%s`
	TOPIC_REQ_REGEXP    = `REQ/(.*)/(.*)/(.*)/(.*)/(.*)/(.*)/(.*)`
	TOPIC_RSP_FORMAT    = `RSP/%s/%s/%s/%s/%s`
	TOPIC_RSP_REGEXP    = `RSP/(.*)/(.*)/(.*)/(.*)/(.*)`
)

var errTopicParse = errors.New("failed to parse topic")

type Topic interface {
	TopicPath() string
}

type BroadcastTopic struct {
	SenderModule Module
	SenderID     string
}

func ParseBroadcastTopic(s string, format string) (BroadcastTopic, error) {
	var v BroadcastTopic
	var r = regexp.MustCompile(format)
	var ds = r.FindAllStringSubmatch(s, -1)
	if len(ds) == 1 {
		var data = ds[0]
		if data != nil && len(data) == 2+1 {
			var smodule Module
			var smodule_err = json.Unmarshal([]byte(strconv.Quote(data[1])), &smodule)
			if smodule_err == nil {

				v.SenderModule = smodule
				v.SenderID = data[2]

				return v, nil
			}
		}
	}

	return v, errTopicParse
}

func (topic BroadcastTopic) TopicPathGeneric(format string) string {
	var sm, _ = topic.SenderModule.MarshalJSON()
	var sm_s, _ = strconv.Unquote(string(sm))
	return fmt.Sprintf(format, sm_s, topic.SenderID)
}

type StatusTopic BroadcastTopic

func ParseStatusTopic(s string) (StatusTopic, error) {
	var v, err = ParseBroadcastTopic(s, TOPIC_STATUS_REGEXP)
	return StatusTopic(v), err
}

func (topic StatusTopic) TopicPath() string {
	return (BroadcastTopic)(topic).TopicPathGeneric(TOPIC_STATUS_FORMAT)
}

type LogTopic BroadcastTopic

func ParseLogTopic(s string) (LogTopic, error) {
	var v, err = ParseBroadcastTopic(s, TOPIC_LOG_REGEXP)
	return LogTopic(v), err
}

func (topic LogTopic) TopicPath() string {
	return (BroadcastTopic)(topic).TopicPathGeneric(TOPIC_LOG_FORMAT)
}

type EventTopic struct {
	SenderModule Module
	SenderID     string
	Type         SystemEventType
}

func ParseEventTopic(s string) (EventTopic, error) {
	var v EventTopic
	var r = regexp.MustCompile(TOPIC_EVENT_REGEXP)
	var ds = r.FindAllStringSubmatch(s, -1)
	if len(ds) == 1 {
		var data = ds[0]
		if data != nil && len(data) == 3+1 {
			var smodule Module
			var event_type SystemEventType
			var smodule_err = json.Unmarshal([]byte(strconv.Quote(data[1])), &smodule)
			var event_type_err = json.Unmarshal([]byte(strconv.Quote(data[3])), &event_type)
			if smodule_err == nil && event_type_err == nil {
				v.SenderModule = smodule
				v.SenderID = data[2]
				v.Type = event_type

				return v, nil
			}
		}
	}

	return v, errTopicParse
}

func (topic EventTopic) TopicPath() string {
	var sm, _ = topic.SenderModule.MarshalJSON()
	var sm_s, _ = strconv.Unquote(string(sm))
	var t, _ = topic.Type.MarshalJSON()
	var t_s, _ = strconv.Unquote(string(t))
	return fmt.Sprintf(TOPIC_EVENT_FORMAT, sm_s, topic.SenderID, t_s)
}

type RequestTopic struct {
	SenderModule   Module
	SenderID       string
	ReceiverModule Module
	ReceiverID     string
	RequestID      string
	Operation      Operation
	Tag            string
}

func (topic RequestTopic) TopicPath() string {
	var u = strconv.Unquote

	var sm, _ = topic.SenderModule.MarshalJSON()
	var sm_s, _ = u(string(sm))
	var rm, _ = topic.ReceiverModule.MarshalJSON()
	var rm_s, _ = u(string(rm))
	var op, _ = topic.Operation.MarshalJSON()
	var op_s, _ = u(string(op))
	return fmt.Sprintf(TOPIC_REQ_FORMAT, sm_s, topic.SenderID, rm_s, topic.ReceiverID, topic.RequestID, op_s, topic.Tag)
}

func (topic RequestTopic) ToResponse() ResponseTopic {
	return ResponseTopic{
		SenderModule:   topic.ReceiverModule,
		SenderID:       topic.ReceiverID,
		ReceiverModule: topic.SenderModule,
		ReceiverID:     topic.SenderID,
		RequestID:      topic.RequestID,
	}
}

func ParseRequestTopic(s string) (RequestTopic, error) {
	var v RequestTopic
	var r = regexp.MustCompile(TOPIC_REQ_REGEXP)
	var ds = r.FindAllStringSubmatch(s, -1)
	if len(ds) == 1 {
		var data = ds[0]
		if data != nil && len(data) == 7+1 {
			var smodule Module
			var rmodule Module
			var op Operation

			var smodule_err = json.Unmarshal([]byte(strconv.Quote(data[1])), &smodule)
			var rmodule_err = json.Unmarshal([]byte(strconv.Quote(data[3])), &rmodule)
			var op_err = json.Unmarshal([]byte(strconv.Quote(data[6])), &op)

			if smodule_err == nil && rmodule_err == nil && op_err == nil {
				v.SenderModule = smodule
				v.SenderID = data[2]
				v.ReceiverModule = rmodule
				v.ReceiverID = data[4]
				v.RequestID = data[5]
				v.Operation = op
				v.Tag = data[7]

				return v, nil
			}
		}
	}

	return v, errTopicParse
}

type ResponseTopic struct {
	SenderModule   Module
	SenderID       string
	ReceiverModule Module
	ReceiverID     string
	RequestID      string
}

func (topic ResponseTopic) TopicPath() string {
	var u = strconv.Unquote

	var sm, _ = topic.SenderModule.MarshalJSON()
	var rm, _ = topic.ReceiverModule.MarshalJSON()

	var sm_s, _ = u(string(sm))
	var rm_s, _ = u(string(rm))
	return fmt.Sprintf(TOPIC_RSP_FORMAT, sm_s, topic.SenderID, rm_s, topic.ReceiverID, topic.RequestID)
}

func ParseResponseTopic(s string) (ResponseTopic, error) {
	var v ResponseTopic
	var r = regexp.MustCompile(TOPIC_RSP_REGEXP)
	var ds = r.FindAllStringSubmatch(s, -1)
	if len(ds) == 1 {
		var data = ds[0]
		if data != nil && len(data) == 5+1 {
			var smodule Module
			var rmodule Module

			var smodule_err = json.Unmarshal([]byte(strconv.Quote(data[1])), &smodule)
			var rmodule_err = json.Unmarshal([]byte(strconv.Quote(data[3])), &rmodule)

			if smodule_err == nil && rmodule_err == nil {
				v.SenderModule = smodule
				v.SenderID = data[2]
				v.ReceiverModule = rmodule
				v.ReceiverID = data[4]
				v.RequestID = data[5]

				return v, nil
			}
		}
	}

	return v, errTopicParse
}

func ParseTopicPath(s string) Topic {
	{
		var v, err = ParseStatusTopic(s)
		if err == nil {
			return v
		}
	}
	{
		var v, err = ParseLogTopic(s)
		if err == nil {
			return v
		}
	}
	{
		var v, err = ParseEventTopic(s)
		if err == nil {
			return v
		}
	}
	{
		var v, err = ParseRequestTopic(s)
		if err == nil {
			return v
		}
	}
	{
		var v, err = ParseResponseTopic(s)
		if err == nil {
			return v
		}
	}

	return nil
}
