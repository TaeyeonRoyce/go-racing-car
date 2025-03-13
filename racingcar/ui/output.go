package ui

import (
	"fmt"
	"github.com/TaeyeonRoyce/go-racing-car/model"
	"strings"
)

func PrintCurrentPositions(cars *model.Cars) {
	for _, car := range *cars {
		fmt.Println(car.Name, ":", strings.Repeat("-", car.Position))
	}
	fmt.Println()
}

func PrintWinners(winners *model.Cars) {
	var winnerNames []string
	for _, winner := range *winners {
		winnerNames = append(winnerNames, winner.Name.Value())
	}
	fmt.Println("최종 우승자:", strings.Join(winnerNames, ", "))
}
