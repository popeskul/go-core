package point

import "errors"

type Point struct {
	X, Y float64
}

func New(x, y float64) (*Point, error) {
	if x < 0 || y < 0 {
		return nil, errors.New("coordinates cannot be less than zero")
	}

	return &Point{x, y}, nil
}
