package libtc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	assert.NoError(t, ApiTest())
}
