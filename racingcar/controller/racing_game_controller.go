package controller

import (
	carModels "github.com/TaeyeonRoyce/go-racing-car/model"
	"github.com/TaeyeonRoyce/go-racing-car/ui"
)

func Run() {
	cars := readCars()
	trialCount := ui.ReadTrialCount()
}

func readCars() carModels.Cars {
	carNamesInput := ui.ReadCars()
	cars, err := carModels.NewCarsFromNames(carNamesInput)
	if err != nil {
		panic(err)
	}
	return cars
}
