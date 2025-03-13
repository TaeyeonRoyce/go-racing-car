package ui

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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

func ReadTrialCount() int {
	fmt.Println("시도할 회수는 몇 회인가요?")
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(errors.New("입력을 읽을 수 없습니다"))
	}
	return parseTrialCount(strings.TrimSpace(line))
}

func parseTrialCount(line string) int {
	trialCount, err := strconv.Atoi(line)
	if err != nil {
		panic(errors.New("회수는 숫자로 입력해주세요"))
	}
	if trialCount <= 0 {
		panic(errors.New("회수는 0보다 커야합니다"))
	}
	return trialCount
}
