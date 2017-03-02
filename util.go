package libwimark

import (
	"fmt"
	"strings"
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
	DirectionBroadcast = direction(1)
	DirectionUnicast   = direction(2)
)

func (self direction) toString() string {
	var v, ok = map[direction]string{
		DirectionBroadcast: "B",
		DirectionUnicast:   "U",
	}[self]

	if !ok {
		panic(self)
	}

	return v
}

const (
	MessageRequest  = messageType(1)
	MessageResponse = messageType(2)
)

const (
	TOPIC_B_FORMAT     = `B/%s/%s`
	TOPIC_U_REQ_FORMAT = `U/%s/%s/%s/%s/REQ/%s/%s`
	TOPIC_U_RSP_FORMAT = `U/%s/%s/%s/%s/RSP/%s`
)

type Topic struct {
	Dir            direction
	SenderModule   module
	SenderID       string
	ReceiverModule *module      `json:",omitempty"`
	ReceiverID     *string      `json:",omitempty"`
	Type           *messageType `json:",omitempty"`
	RequestID      *string      `json:",omitempty"`
	Operation      *operation   `json:",omitempty"`
}

func (self *Topic) ToString() *string {
	var sender_module = self.SenderModule.toString()
	var sender_id = self.SenderID

	switch self.Dir {
	case DirectionBroadcast:
		var v = fmt.Sprintf(TOPIC_B_FORMAT, sender_module, sender_id)
		return &v
	case DirectionUnicast:
		if self.Type != nil && self.ReceiverModule != nil && self.ReceiverID != nil && self.RequestID != nil {
			var t = *self.Type
			var rmodule = self.ReceiverModule.toString()
			var rid = *self.ReceiverID
			var reqid = *self.RequestID
			switch t {
			case MessageRequest:
				if self.Operation != nil {
					var v = fmt.Sprintf(TOPIC_U_REQ_FORMAT, sender_module, sender_id, rmodule, rid, reqid, self.Operation.toString())
					return &v
				}
			case MessageResponse:
				var v = fmt.Sprintf(TOPIC_U_RSP_FORMAT, sender_module, sender_id, rmodule, rid, reqid)
				return &v
			}
		}
	}

	return nil
}

func MakeFullTopic(d direction, sm module, sid string, rm module, rid string, t messageType, reqid string, op operation) Topic {
	return Topic{
		Dir:            d,
		SenderModule:   sm,
		SenderID:       sid,
		ReceiverModule: &rm,
		ReceiverID:     &rid,
		Type:           &t,
		RequestID:      &reqid,
		Operation:      &op,
	}
}

// Generates topic for outgoing messages
func MakeRspTopic(reqTopic Topic) Topic {
	var t_dir = DirectionUnicast
	var t_smodule = *reqTopic.ReceiverModule
	var t_sid string
	if reqTopic.ReceiverID != nil {
		t_sid = *reqTopic.ReceiverID
	}
	var t_rmodule = reqTopic.SenderModule
	var t_rid = reqTopic.SenderID
	var t_type = MessageResponse
	var t_reqid = reqTopic.RequestID

	var rspTopic = Topic{Dir: t_dir, SenderModule: t_smodule, SenderID: t_sid, ReceiverModule: &t_rmodule, ReceiverID: &t_rid, Type: &t_type, RequestID: t_reqid}

	return rspTopic
}

const STRING_PLACEHOLDER = "STRING_PLACEHOLDER"

func ParseTopic(topic_string string) *Topic {
	if !strings.Contains(topic_string, " ") {
		var dir_s string
		var smodule_s string
		var sid string
		var rmodule_s string
		var rid string
		var msgtype_s string
		var reqid string
		var op_s string

		topic_string = strings.Replace(topic_string, "//", fmt.Sprintf("/%s/", STRING_PLACEHOLDER), -1)

		fmt.Sscanf(strings.Replace(topic_string, "/", " ", -1), "%s %s %s %s %s %s %s %s", &dir_s, &smodule_s, &sid, &rmodule_s, &rid, &msgtype_s, &reqid, &op_s)

		if sid == STRING_PLACEHOLDER {
			sid = ""
		}

		if rid == STRING_PLACEHOLDER {
			rid = ""
		}

		if dir_s == STRING_PLACEHOLDER || smodule_s == STRING_PLACEHOLDER {
			return nil
		}

		var smodule_opt = parseModuleString(smodule_s)
		if smodule_opt != nil {
			var smodule = *smodule_opt
			if false {
			} else if dir_s == "B" {
				return &Topic{Dir: DirectionBroadcast, SenderModule: smodule, SenderID: sid}
			} else if dir_s == "U" {
				if (msgtype_s == "REQ" || msgtype_s == "RSP") && reqid != STRING_PLACEHOLDER {
					var rmodule_opt = parseModuleString(rmodule_s)
					if rmodule_opt != nil {
						switch msgtype_s {
						case "REQ":
							var op_opt = parseOpString(op_s)
							if op_opt != nil {
								var t = MessageRequest
								return &Topic{Dir: DirectionUnicast, SenderModule: smodule, SenderID: sid, ReceiverModule: rmodule_opt, ReceiverID: &rid, Type: &t, RequestID: &reqid, Operation: op_opt}
							}
						case "RSP":
							var t = MessageResponse
							return &Topic{Dir: DirectionUnicast, SenderModule: smodule, SenderID: sid, ReceiverModule: rmodule_opt, ReceiverID: &rid, Type: &t, RequestID: &reqid}
						}
					}
				}
			}
		}
	}

	return nil
}
