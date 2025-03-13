package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindWinners(t *testing.T) {
	testCases := []struct {
		name       string
		RacingGame *RacingGame
		winnerCars *Cars
	}{
		{
			"multiple winners",
			&RacingGame{
				Cars: &Cars{
					{Name: "a", Position: 1},
					{Name: "b", Position: 2},
					{Name: "c", Position: 2},
				},
				TrialCount: 1,
			},
			&Cars{
				{Name: "b", Position: 2},
				{Name: "c", Position: 2},
			},
		},
		{
			"single winners",
			&RacingGame{
				Cars: &Cars{
					{Name: "a", Position: 1},
					{Name: "b", Position: 2},
					{Name: "c", Position: 3},
				},
				TrialCount: 1,
			},
			&Cars{
				{Name: "c", Position: 3},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			winners := test.RacingGame.FindWinners()
			assert.EqualValues(t, *test.winnerCars, *winners)
		})
	}
}
