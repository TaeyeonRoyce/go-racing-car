package model

import "errors"

var _duplicateCarNameError = errors.New("중복된 자동차 이름이 존재합니다")

type Cars []*Car

func NewCarsFromNames(names []string) (*Cars, error) {
	cars := make(Cars, len(names))
	for i, name := range names {
		car, err := NewCar(name)
		if err != nil {
			return nil, err
		}
		cars[i] = car
	}
	if err := validateDuplicateCarNames(cars); err != nil {
		return nil, err
	}
	return &cars, nil
}

func validateDuplicateCarNames(cars Cars) error {
	uniqueCarNames := make(map[CarName]bool)
	for _, car := range cars {
		if _, ok := uniqueCarNames[car.Name]; ok {
			return _duplicateCarNameError
		}
		uniqueCarNames[car.Name] = true
	}
	return nil
}
