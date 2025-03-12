package ui

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadCars() []string {
	fmt.Println("자동차 이름을 입력하세요 (이름은 쉼표(,)로 구분)")
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(errors.New("입력을 읽을 수 없습니다"))
	}
	return splitCarsByDelimiter(line, ",")
}

func splitCarsByDelimiter(line, delimiter string) []string {
	cars := make([]string, 0)
	for _, carName := range strings.Split(line, delimiter) {
		cars = append(cars, strings.TrimSpace(carName))
	}
	return cars
}
