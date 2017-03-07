package libwimark

import (
	"fmt"
	"regexp"
)

const MQTT_ANY_WILDCARD = "+"

type module int
type direction int
type messageType int
type operation int

const (
	ModuleAny     = module(-1)
	ModuleBackend = module(1)
	ModuleConfig  = module(2)
	ModuleDB      = module(3)
	ModuleCPE     = module(4)
	ModuleStat    = module(5)
)

func (self module) toString() string {
	var v, ok = map[module]string{
		ModuleAny:     MQTT_ANY_WILDCARD,
		ModuleBackend: "BACKEND",
		ModuleConfig:  "CONFIG",
		ModuleDB:      "DB",
		ModuleCPE:     "CPE",
		ModuleStat:    "STAT",
	}[self]

	if !ok {
		panic(self)
	}

	return v
}

func parseModuleString(v string) *module {
	var ret, ok = map[string]module{
		"BACKEND": ModuleBackend,
		"CONFIG":  ModuleConfig,
		"DB":      ModuleDB,
		"CPE":     ModuleCPE,
		"STAT":    ModuleStat,
	}[v]
	if ok {
		return &ret
	}
	return nil
}

const (
	OperationCreate = operation(1)
	OperationRead   = operation(2)
	OperationUpdate = operation(3)
	OperationDelete = operation(4)
)

func parseOpString(v string) *operation {
	var ret, ok = map[string]operation{
		"C": OperationCreate,
		"R": OperationRead,
		"U": OperationUpdate,
		"D": OperationDelete,
	}[v]
	if ok {
		return &ret
	}

	return nil
}

func (self operation) toString() string {
	var v, ok = map[operation]string{
		OperationCreate: "C",
		OperationRead:   "R",
		OperationUpdate: "U",
		OperationDelete: "D",
	}[self]

	if !ok {
		panic(self)
	}

	return v
}

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
	SenderModule module
	SenderID     string
}

func ParseBroadcastTopic(s string) *BroadcastTopic {
	var r = regexp.MustCompile(TOPIC_B_REGEXP)
	var ds = r.FindAllStringSubmatch(s, -1)
	if ds != nil && len(ds) == 1 {
		var data = ds[0]
		if data != nil && len(data) == 2+1 {
			var smodule = parseModuleString(data[1])
			if smodule != nil {
				var v BroadcastTopic

				v.SenderModule = *smodule
				v.SenderID = data[2]

				return &v
			}
		}
	}

	return nil
}

func (self *BroadcastTopic) TopicPath() string {
	return fmt.Sprintf(TOPIC_B_FORMAT, self.SenderModule.toString(), self.SenderID)
}

type RequestTopic struct {
	SenderModule   module
	SenderID       string
	ReceiverModule module
	ReceiverID     string
	RequestID      string
	Operation      operation
}

func (self *RequestTopic) TopicPath() string {
	return fmt.Sprintf(TOPIC_REQ_FORMAT, self.SenderModule.toString(), self.SenderID, self.ReceiverModule.toString(), self.ReceiverID, self.RequestID, self.Operation.toString())
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
			var smodule = parseModuleString(data[1])
			var rmodule = parseModuleString(data[3])
			var op = parseOpString(data[6])

			if smodule != nil && rmodule != nil && op != nil {
				var v RequestTopic

				v.SenderModule = *smodule
				v.SenderID = data[2]
				v.ReceiverModule = *rmodule
				v.ReceiverID = data[4]
				v.RequestID = data[5]
				v.Operation = *op

				return &v
			}
		}
	}

	return nil
}

type ResponseTopic struct {
	SenderModule   module
	SenderID       string
	ReceiverModule module
	ReceiverID     string
	RequestID      string
}

func (self *ResponseTopic) TopicPath() string {
	return fmt.Sprintf(TOPIC_RSP_FORMAT, self.SenderModule.toString(), self.SenderID, self.ReceiverModule.toString(), self.ReceiverID, self.RequestID)
}

func ParseResponseTopic(s string) *ResponseTopic {
	var r = regexp.MustCompile(TOPIC_RSP_REGEXP)
	var ds = r.FindAllStringSubmatch(s, -1)
	if ds != nil && len(ds) == 1 {
		var data = ds[0]
		if data != nil && len(data) == 5+1 {
			var smodule = parseModuleString(data[1])
			var rmodule = parseModuleString(data[3])

			if smodule != nil && rmodule != nil {
				var v ResponseTopic

				v.SenderModule = *smodule
				v.SenderID = data[2]
				v.ReceiverModule = *rmodule
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
