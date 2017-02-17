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
