package controller

import (
	carModels "github.com/TaeyeonRoyce/go-racing-car/model"
	"github.com/TaeyeonRoyce/go-racing-car/ui"
)

func Run() {
	cars := readCars()
	trialCount := ui.ReadTrialCount()
	game := carModels.NewRacingGame(cars, trialCount)
	playGame(game)
}

func readCars() *carModels.Cars {
	carNamesInput := ui.ReadCars()
	cars, err := carModels.NewCarsFromNames(carNamesInput)
	if err != nil {
		panic(err)
	}
	return cars
}

func playGame(game carModels.RacingGame) {
	for i := 0; i < game.TrialCount; i++ {
		game.PlaySingleRound()
	}
}
