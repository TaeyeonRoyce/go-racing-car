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
