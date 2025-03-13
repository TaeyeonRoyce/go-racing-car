package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCar(t *testing.T) {
	testCases := []struct {
		name             string
		expectedPosition int
		expectedError    error
	}{
		{"a", 0, nil},
		{"abcde", 0, nil},
		{"abcdef", 0, _invalidCarNameLengthError},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			car, err := NewCar(test.name)
			assert.ErrorIs(t, err, test.expectedError)

			if err == nil {
				assert.Equal(t, test.expectedPosition, car.Position)
			}
		})
	}
}

func TestMoveForwardByNumber(t *testing.T) {
	testCases := []struct {
		car              Car
		expectedPosition int
		randomNumber     int
	}{
		{Car{Name: "a", Position: 0}, 1, 4},
		{Car{Name: "a", Position: 0}, 0, 3},
	}

	for _, test := range testCases {
		t.Run("", func(t *testing.T) {
			car := test.car
			car.MoveForwardByNumber(test.randomNumber)
			assert.Equal(t, test.expectedPosition, car.Position)
		})
	}
}
