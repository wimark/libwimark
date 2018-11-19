package libtc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	assert.NoError(t, ExecTest())
}
