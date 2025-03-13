package model

import (
	"github.com/TaeyeonRoyce/go-racing-car/util"
)

type RacingGame struct {
	Cars       *Cars
	TrialCount int
}

func NewRacingGame(cars *Cars, trialCount int) *RacingGame {
	return &RacingGame{Cars: cars, TrialCount: trialCount}
}

func (r *RacingGame) PlaySingleRound() {
	for _, car := range *r.Cars {
		go car.MoveForwardByNumber(util.GenerateRandomNumber())
	}
}

func (r *RacingGame) FindWinners() *Cars {
	winnerPosition := r.findWinnerPosition()
	winners := make(Cars, 0)
	for _, car := range *r.Cars {
		if car.Position == winnerPosition {
			winners = append(winners, car)
		}
	}
	return &winners
}

func (r *RacingGame) findWinnerPosition() int {
	maxPosition := 0
	for _, car := range *r.Cars {
		if car.Position > maxPosition {
			maxPosition = car.Position
		}
	}
	return maxPosition
}
