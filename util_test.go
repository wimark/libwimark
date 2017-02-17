package libwimark

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	REQ_JSON_FORMAT = `{"Dir":%d,"SenderModule":%d,"SenderID":"%s","ReceiverModule":%d,"ReceiverID":"%s","Type":%d,"RequestID":"%s","Operation":%d}`
	RSP_JSON_FORMAT = `{"Dir":%d,"SenderModule":%d,"SenderID":"%s","ReceiverModule":%d,"ReceiverID":"%s","Type":%d,"RequestID":"%s"}`
	B_JSON_FORMAT   = `{"Dir":%d,"SenderModule":%d,"SenderID":"%s"}`
)

func TestParseRequest(t *testing.T) {
	var fixture = "U/CONFIG/A001/DB/B002/REQ/G45134513/R"
	var expectation = fmt.Sprintf(REQ_JSON_FORMAT, DirectionUnicast, ModuleConfig, "A001", ModuleDB, "B002", MessageRequest, "G45134513", OperationRead)
	var j, _ = json.Marshal(ParseTopic(fixture))
	var result = string(j)

	assert.Equal(t, expectation, result)
}

func TestParseResponse(t *testing.T) {
	var fixture = "U/CONFIG/A001/DB/B002/RSP/G45134513"
	var expectation = fmt.Sprintf(RSP_JSON_FORMAT, DirectionUnicast, ModuleConfig, "A001", ModuleDB, "B002", MessageResponse, "G45134513")
	var j, _ = json.Marshal(ParseTopic(fixture))
	var result = string(j)

	assert.Equal(t, expectation, result)
}

func TestParseBroadcast(t *testing.T) {
	var fixture = "B/CPE/AP001"
	var expectation = fmt.Sprintf(B_JSON_FORMAT, DirectionBroadcast, ModuleCPE, "AP001")
	var j, _ = json.Marshal(ParseTopic(fixture))
	var result = string(j)

	assert.Equal(t, expectation, result)
}

func TestSerializeRequest(t *testing.T) {
	var dir = DirectionUnicast
	var smodule = ModuleConfig
	var sid = "A001"
	var rmodule = ModuleDB
	var rid = "B002"
	var msgtype = MessageRequest
	var reqid = "G45134513"
	var op = OperationRead
	var fixture = Topic{Dir: dir, SenderModule: smodule, SenderID: sid, ReceiverModule: &rmodule, ReceiverID: &rid, Type: &msgtype, RequestID: &reqid, Operation: &op}
	var expectation = "U/CONFIG/A001/DB/B002/REQ/G45134513/R"
	var result = *fixture.ToString()

	assert.Equal(t, expectation, result)
}

func TestSerializeResponse(t *testing.T) {
	var dir = DirectionUnicast
	var smodule = ModuleConfig
	var sid = "A001"
	var rmodule = ModuleDB
	var rid = "B002"
	var msgtype = MessageResponse
	var reqid = "G45134513"
	var fixture = Topic{Dir: dir, SenderModule: smodule, SenderID: sid, ReceiverModule: &rmodule, ReceiverID: &rid, Type: &msgtype, RequestID: &reqid}
	var expectation = "U/CONFIG/A001/DB/B002/RSP/G45134513"
	var result = *fixture.ToString()

	assert.Equal(t, expectation, result)
}

func TestSerializeBroadcast(t *testing.T) {
	var dir = DirectionBroadcast
	var smodule = ModuleCPE
	var sid = "A001"
	var fixture = Topic{Dir: dir, SenderModule: smodule, SenderID: sid}
	var expectation = "B/CPE/A001"
	var result = *fixture.ToString()

	assert.Equal(t, expectation, result)
}
