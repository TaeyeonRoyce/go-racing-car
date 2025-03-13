package model

import "errors"

const (
	_minNameLength = 1
	_maxNameLength = 5
)

var _invalidCarNameLengthError = errors.New("자동차 이름은 1자 이상, 5자 이하만 가능합니다")

type CarName string

func NewCarName(name string) (CarName, error) {
	nameLength := len(name)
	if nameLength < _minNameLength || nameLength > _maxNameLength {
		return "", _invalidCarNameLengthError
	}
	return CarName(name), nil
}

func (c CarName) Value() string {
	return string(c)
}
