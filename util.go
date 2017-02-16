package libwimark

import (
	"fmt"
	"strings"
)

type module int
type direction int
type messageType int
type operation int

const (
	ModuleBackend = module(0)
	ModuleConfig  = module(1)
	ModuleDB      = module(2)
	ModuleCPE     = module(3)
)

func (self module) toString() string {
	var v, ok = map[module]string{
		ModuleBackend: "BACKEND",
		ModuleConfig:  "CONFIG",
		ModuleDB:      "DB",
		ModuleCPE:     "CPE",
	}[self]

	if !ok {
		panic(self)
	}

	return v
}

func parseModuleString(v_p *string) *module {
	if v_p != nil {
		var v = *v_p
		var ret, ok = map[string]module{
			"BACKEND": ModuleBackend,
			"CONFIG":  ModuleConfig,
			"DB":      ModuleDB,
			"CPE":     ModuleCPE,
		}[v]
		if ok {
			return &ret
		}
	}
	return nil
}

const (
	OperationCreate = operation(0)
	OperationRead   = operation(1)
	OperationUpdate = operation(2)
	OperationDelete = operation(3)
)

func parseOpString(v_p *string) *operation {
	if v_p != nil {
		var v = *v_p
		var ret, ok = map[string]operation{
			"C": OperationCreate,
			"R": OperationRead,
			"U": OperationUpdate,
			"D": OperationDelete,
		}[v]
		if ok {
			return &ret
		}
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
	DirectionBroadcast = direction(0)
	DirectionUnicast   = direction(1)
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
	MessageRequest  = messageType(0)
	MessageResponse = messageType(1)
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
		var v = fmt.Sprintf("B/%s/%s", sender_module, sender_id)
		return &v
	case DirectionUnicast:
		if self.Type != nil && self.ReceiverModule != nil && self.ReceiverID != nil && self.RequestID != nil {
			var t = *self.Type
			var rmodule = *self.ReceiverModule
			var rid = *self.ReceiverID
			var reqid = *self.RequestID
			switch t {
			case MessageRequest:
				if self.Operation != nil {
					var v = fmt.Sprintf("U/%s/%s/%s/%s/REQ/%s/%s", sender_module, sender_id, rmodule, rid, reqid, self.Operation.toString())
					return &v
				}
			case MessageResponse:
				var v = fmt.Sprintf("U/%s/%s/%s/%s/RSP/%s", sender_module, sender_id, rmodule, rid, reqid)
				return &v
			}
		}
	}

	return nil
}

func ParseTopic(topic_string string) *Topic {
	if !strings.Contains(topic_string, " ") {
		var dir_opt string
		var smodule_s_opt string
		var sid_opt string
		var rmodule_opt string
		var rid_opt string
		var t_opt string
		var reqid_opt string
		var op_opt string

		fmt.Sscanf(strings.Replace(topic_string, "/", " ", -1), "%s %s %s %s %s %s %s %s", &dir_opt, &smodule_s_opt, &sid_opt, &rmodule_opt, &rid_opt, &t_opt, &reqid_opt, &op_opt)

		var smodule_opt = parseModuleString(&smodule_s_opt)
		{
			var dir_opt = map[bool]*string{true: &dir_opt, false: nil}[len(dir_opt) > 0]
			var sid_opt = map[bool]*string{true: &sid_opt, false: nil}[len(sid_opt) > 0]
			var rmodule_opt = map[bool]*string{true: &rmodule_opt, false: nil}[len(rmodule_opt) > 0]
			var rid_opt = map[bool]*string{true: &rid_opt, false: nil}[len(rid_opt) > 0]
			var t_opt = map[bool]*string{true: &t_opt, false: nil}[len(t_opt) > 0]
			var reqid_opt = map[bool]*string{true: &reqid_opt, false: nil}[len(reqid_opt) > 0]
			var op_opt = map[bool]*string{true: &op_opt, false: nil}[len(op_opt) > 0]
			if dir_opt != nil && smodule_opt != nil && sid_opt != nil {
				var dir = *dir_opt
				var smodule = *smodule_opt
				var sid = *sid_opt
				if false {
				} else if dir == "B" {
					return &Topic{Dir: DirectionBroadcast, SenderModule: smodule, SenderID: sid}
				} else if dir == "U" {
					if t_opt != nil {
						var uniType = *t_opt
						if uniType == "REQ" || uniType == "RSP" {
							var rmodule_opt = parseModuleString(rmodule_opt)
							var op_opt = parseOpString(op_opt)
							if false {
							} else if uniType == "REQ" {
								if rmodule_opt != nil && rid_opt != nil && reqid_opt != nil && op_opt != nil {
									var t = MessageRequest
									return &Topic{Dir: DirectionUnicast, SenderModule: smodule, SenderID: sid, ReceiverModule: rmodule_opt, ReceiverID: rid_opt, Type: &t, RequestID: reqid_opt, Operation: op_opt}
								}
							} else if uniType == "RSP" {
								if rmodule_opt != nil && rid_opt != nil && reqid_opt != nil {
									var t = MessageResponse
									return &Topic{Dir: DirectionUnicast, SenderModule: smodule, SenderID: sid, ReceiverModule: rmodule_opt, ReceiverID: rid_opt, Type: &t, RequestID: reqid_opt}
								}
							}
						}
					}
				}
			}
		}
	}

	return nil
}
