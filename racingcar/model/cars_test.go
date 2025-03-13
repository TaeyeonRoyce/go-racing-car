package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCarsFromNames(t *testing.T) {
	testCases := []struct {
		names            []string
		expectedCarCount int
		expectedError    error
	}{
		{[]string{"a", "b"}, 2, nil},
		{[]string{"a", "b", "c", "d", "e"}, 5, nil},
		{[]string{"abcedfg", "f"}, 0, _invalidCarNameLengthError},
	}

	for _, test := range testCases {
		t.Run("", func(t *testing.T) {
			cars, err := NewCarsFromNames(test.names)
			assert.ErrorIs(t, err, test.expectedError)

			if err == nil {
				assert.Equal(t, test.expectedCarCount, len(*cars))
			}
		})
	}
}

func TestValidateDuplicationCarNames(t *testing.T) {
	testCases := []struct {
		names         []string
		expectedError error
	}{
		{[]string{"a", "a"}, _duplicateCarNameError},
		{[]string{"a", "b"}, nil},
	}

	for _, test := range testCases {
		t.Run("", func(t *testing.T) {
			cars, err := NewCarsFromNames(test.names)
			assert.ErrorIs(t, err, test.expectedError)

			if err == nil {
				assert.Equal(t, len(*cars), len(*cars))
			}
		})
	}
}
