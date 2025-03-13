package model

const (
	_moveThreshold = 4
)

type Car struct {
	Name     CarName
	Position int
}

func (c *Car) MoveForwardByNumber(number int) {
	if number >= _moveThreshold {
		c.Position++
	}
}

func NewCar(name string) (*Car, error) {
	carName, err := NewCarName(name)
	if err != nil {
		return nil, err
	}
	return &Car{
		Name:     carName,
		Position: 0,
	}, nil
}
