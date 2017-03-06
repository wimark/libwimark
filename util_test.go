package libwimark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequest(t *testing.T) {
	var fixture = `U/CONFIG/A001/DB/B002/REQ/G45134513/R`
	var expectation = fixture
	var result = ParseTopicPath(fixture).TopicPath()

	assert.Equal(t, expectation, result)
}

func TestResponse(t *testing.T) {
	var fixture = `U/CONFIG/A001/DB/B002/RSP/G45134513`
	var expectation = fixture
	var result = ParseTopicPath(fixture).TopicPath()

	assert.Equal(t, expectation, result)
}

func TestBroadcast(t *testing.T) {
	var fixture = `B/CPE/AP001`
	var expectation = fixture
	var result = ParseTopicPath(fixture).TopicPath()

	assert.Equal(t, expectation, result)
}
