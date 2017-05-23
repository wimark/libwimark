package libwimark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequest(t *testing.T) {
	var fixture = `REQ/CONFIG/A001/DB/B002/G45134513/R`
	var expectation = fixture
	var result = ParseTopicPath(fixture).TopicPath()

	assert.Equal(t, expectation, result)
}

func TestResponse(t *testing.T) {
	var fixture = `RSP/CONFIG/A001/DB/B002/G45134513`
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

func TestEvent(t *testing.T) {
	var fixture = `EVENT/STAT/STAT001/STAT_RULE_VIOLATION`
	var expectation = fixture
	var result = ParseTopicPath(fixture).TopicPath()

	assert.Equal(t, expectation, result)
}
