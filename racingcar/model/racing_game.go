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
