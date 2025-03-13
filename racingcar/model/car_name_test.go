package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCarName(t *testing.T) {
	testCases := []struct {
		name     string
		expected error
	}{
		{"a", nil},
		{"abcde", nil},
		{"abcdef", _invalidCarNameLengthError},
		{"", _invalidCarNameLengthError},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewCarName(test.name)
			assert.Equal(t, test.expected, err)
		})
	}
}
