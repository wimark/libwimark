package libwimark

import (
	"fmt"
	"strings"
)

type Module int
type Direction int
type MessageType int
type Operation int

const (
	ModuleBackend = Module(0)
	ModuleConfig  = Module(1)
	ModuleDB      = Module(2)
	ModuleCPE     = Module(3)
)

func moduleP(v Module) *Module { return &v }

func parseModuleString(v_p *string) *Module {
	if v_p != nil {
		var v = *v_p
		switch v {
		case "BACKEND":
			return moduleP(ModuleBackend)
		case "CONFIG":
			return moduleP(ModuleConfig)
		case "DB":
			return moduleP(ModuleDB)
		case "CPE":
			return moduleP(ModuleCPE)
		}
	}
	return nil
}

const (
	OperationCreate = Operation(0)
	OperationRead   = Operation(1)
	OperationUpdate = Operation(2)
	OperationDelete = Operation(3)
)

func operationP(v Operation) *Operation { return &v }

func parseOpString(v_p *string) *Operation {
	if v_p != nil {
		var v = *v_p
		switch v {
		case "C":
			return operationP(OperationCreate)
		case "R":
			return operationP(OperationRead)
		case "U":
			return operationP(OperationUpdate)
		case "D":
			return operationP(OperationDelete)
		}
	}
	return nil
}

const (
	DirectionBroadcast = Direction(0)
	DirectionUnicast   = Direction(1)
)

const (
	MessageRequest  = MessageType(0)
	MessageResponse = MessageType(1)
)

type Topic struct {
	Dir            Direction
	SenderModule   Module
	SenderID       string
	ReceiverModule *Module `json:",omitempty"`
	ReceiverID     *string `json:",omitempty"`
	Type           *MessageType `json:",omitempty"`
	RequestID      *string `json:",omitempty"`
	Operation      *Operation `json:",omitempty"`
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
