package libwimark

import (
	"encoding/json"
	"errors"
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
	TOPIC_EVENT_FORMAT  = `EVENT/%s/%s/%s`
	TOPIC_EVENT_REGEXP  = `EVENT/(.*)/(.*)/(.*)`
	TOPIC_REQ_FORMAT    = `REQ/%s/%s/%s/%s/%s/%s`
	TOPIC_REQ_REGEXP    = `REQ/(.*)/(.*)/(.*)/(.*)/(.*)/(.*)`
	TOPIC_RSP_FORMAT    = `RSP/%s/%s/%s/%s/%s`
	TOPIC_RSP_REGEXP    = `RSP/(.*)/(.*)/(.*)/(.*)/(.*)`
)

var parseErr = errors.New("Failed to parse topic")

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
	if ds != nil && len(ds) == 1 {
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

	return v, parseErr
}

func (self BroadcastTopic) TopicPathGeneric(format string) string {
	var sm, _ = self.SenderModule.MarshalJSON()
	var sm_s, _ = strconv.Unquote(string(sm))
	return fmt.Sprintf(format, sm_s, self.SenderID)
}

type StatusTopic BroadcastTopic

func ParseStatusTopic(s string) (StatusTopic, error) {
	var v, err = ParseBroadcastTopic(s, TOPIC_STATUS_REGEXP)
	return StatusTopic(v), err
}

func (self StatusTopic) TopicPath() string {
	return (BroadcastTopic)(self).TopicPathGeneric(TOPIC_STATUS_FORMAT)
}

type LogTopic BroadcastTopic

func ParseLogTopic(s string) (LogTopic, error) {
	var v, err = ParseBroadcastTopic(s, TOPIC_LOG_REGEXP)
	return LogTopic(v), err
}

func (self LogTopic) TopicPath() string {
	return (BroadcastTopic)(self).TopicPathGeneric(TOPIC_LOG_FORMAT)
}

type EventTopic struct {
	SenderModule Module
	SenderID     string
	Type         SystemEventObjectType
}

func ParseEventTopic(s string) (EventTopic, error) {
	var v EventTopic
	var r = regexp.MustCompile(TOPIC_EVENT_REGEXP)
	var ds = r.FindAllStringSubmatch(s, -1)
	if ds != nil && len(ds) == 1 {
		var data = ds[0]
		if data != nil && len(data) == 3+1 {
			var smodule Module
			var event_type SystemEventObjectType
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

	return v, parseErr
}

func (self EventTopic) TopicPath() string {
	var sm, _ = self.SenderModule.MarshalJSON()
	var sm_s, _ = strconv.Unquote(string(sm))
	var t, _ = self.Type.MarshalJSON()
	var t_s, _ = strconv.Unquote(string(t))
	return fmt.Sprintf(TOPIC_EVENT_FORMAT, sm_s, self.SenderID, t_s)
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

func ParseRequestTopic(s string) (RequestTopic, error) {
	var v RequestTopic
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

				v.SenderModule = smodule
				v.SenderID = data[2]
				v.ReceiverModule = rmodule
				v.ReceiverID = data[4]
				v.RequestID = data[5]
				v.Operation = op

				return v, nil
			}
		}
	}

	return v, parseErr
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

func ParseResponseTopic(s string) (ResponseTopic, error) {
	var v ResponseTopic
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
				v.SenderModule = smodule
				v.SenderID = data[2]
				v.ReceiverModule = rmodule
				v.ReceiverID = data[4]
				v.RequestID = data[5]

				return v, nil
			}
		}
	}

	return v, parseErr
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

type DBResponseBase struct {
	Errors []ModelError `json:"errors,omitempty"`
}

type DBDataObj struct {
	WLANs           map[UUID]WLAN            `json:"wlan,omitempty"`
	CPEs            map[UUID]CPE             `json:"cpe,omitempty"`
	Stats           map[UUID]Stat            `json:"stat,omitempty"`
	ClientStats     map[UUID]ClientStat      `json:"client-stat,omitempty"`
	Events          map[UUID]SystemEvent     `json:"event,omitempty"`
	StatEventRules  map[UUID]StatEventRule   `json:"stat-event-rule,omitempty"`
	PollCPE         map[UUID]CPEPollSettings `json:"poll-cpe,omitempty"`
	Radius          map[UUID]Radius          `json:"radius,omitempty"`
	LBSCPEInfo      map[UUID]LBSCPEInfo      `json:"lbs-cpe-info,omitempty"`
	LBSClientData   map[UUID]LBSClientData   `json:"lbs-client-data,omitempty"`
	LBSClientCoords map[UUID]LBSClientCoords `json:"lbs-client-coords,omitempty"`
}

func (self *DBDataObj) Reset() {
	*self = DBDataObj{}
}

type DBDataUUID struct {
	WLANs           []UUID `json:"wlan,omitempty"`
	CPEs            []UUID `json:"cpe,omitempty"`
	Stats           []UUID `json:"stat,omitempty"`
	ClientStats     []UUID `json:"client-stat,omitempty"`
	Events          []UUID `json:"event,omitempty"`
	StatEventRules  []UUID `json:"stat-event-rule,omitempty"`
	PollCPE         []UUID `json:"poll-cpe,omitempty"`
	Radius          []UUID `json:"radius,omitempty"`
	LBSCPEInfo      []UUID `json:"lbs-cpe-info,omitempty"`
	LBSClientData   []UUID `json:"lbs-client-data,omitempty"`
	LBSClientCoords []UUID `json:"lbs-client-coords,omitempty"`
}

func (self *DBDataUUID) Reset() {
	*self = DBDataUUID{}
}

type DBDataMasks struct {
	WLANs           *WLANMask            `json:"wlan,omitempty"`
	CPEs            *CPEMask             `json:"cpe,omitempty"`
	Stats           *StatsMask           `json:"stat,omitempty"`
	ClientStats     *SimpleMask          `json:"client-stat,omitempty"`
	Events          *EventMask           `json:"event,omitempty"`
	StatEventRules  *SimpleMask          `json:"stat-event-rule,omitempty"`
	PollCPE         *SimpleMask          `json:"poll-cpe,omitempty"`
	Radius          *SimpleMask          `json:"radius,omitempty"`
	LBSCPEInfo      *LBSCPEInfoMask      `json:"lbs-cpe-info,omitempty"`
	LBSClientData   *LBSClientDataMask   `json:"lbs-client-data,omitempty"`
	LBSClientCoords *LBSClientCoordsMask `json:"lbs-client-coords,omitempty"`
}

func (self *DBDataMasks) Reset() {
	*self = DBDataMasks{}
}

type DBRequestC DBDataObj
type DBRequestU DBDataObj
type DBRequestR DBDataMasks
type DBRequestD DBDataMasks

func (self *DBRequestC) UnmarshalJSON(b []byte) error {
	(*DBDataObj)(self).Reset()

	return json.Unmarshal(b, (*DBDataObj)(self))
}

func (self *DBRequestR) UnmarshalJSON(b []byte) error {
	(*DBDataMasks)(self).Reset()

	return json.Unmarshal(b, (*DBDataMasks)(self))
}

func (self *DBRequestU) UnmarshalJSON(b []byte) error {
	(*DBDataObj)(self).Reset()

	return json.Unmarshal(b, (*DBDataObj)(self))
}

func (self *DBRequestD) UnmarshalJSON(b []byte) error {
	(*DBDataMasks)(self).Reset()

	return json.Unmarshal(b, (*DBDataMasks)(self))
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

func (self *DBResponseC) UnmarshalJSON(b []byte) error {
	self.DBDataUUID.Reset()

	return json.Unmarshal(b, (*DBResponseUUID)(self))
}

func (self *DBResponseR) UnmarshalJSON(b []byte) error {
	self.DBDataObj.Reset()

	return json.Unmarshal(b, (*DBResponseObj)(self))
}

func (self *DBResponseU) UnmarshalJSON(b []byte) error {
	self.DBDataUUID.Reset()

	return json.Unmarshal(b, (*DBResponseUUID)(self))
}

func (self *DBResponseD) UnmarshalJSON(b []byte) error {
	self.DBDataUUID.Reset()

	return json.Unmarshal(b, (*DBResponseUUID)(self))
}

type ConnectorInfo struct {
	DbType    string   `json:"db_type"`
	DbServers []string `json:"db_servers"`
	Models    []string `json:"models"`
}

type ModuleStatus struct {
	Version     string      `json:"version"`
	IsConnected bool        `json:"connected"`
	Metadata    interface{} `json:"meta,omitempty"`
}

func (self ModuleStatus) Connected() ModuleStatus {
	var v = self
	v.IsConnected = true

	return v
}

func (self ModuleStatus) Disconnected() ModuleStatus {
	var v = self
	v.IsConnected = false

	return v
}

func (self ModuleStatus) String() string {
	var s, sErr = json.Marshal(self)
	if sErr != nil {
		panic(sErr)
	}

	return string(s)
}
