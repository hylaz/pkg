package stringex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	r1 := Rand(1, 100)
	assert.GreaterOrEqual(t, r1, int64(1))
	assert.LessOrEqual(t, r1, int64(100))
}
