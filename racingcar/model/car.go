package model

type Car struct {
	Name     CarName
	Position int
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
