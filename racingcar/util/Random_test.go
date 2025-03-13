package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		got := GenerateRandomNumber()
		assert.GreaterOrEqual(t, got, 0)
		assert.LessOrEqual(t, got, 9)
	}
}
