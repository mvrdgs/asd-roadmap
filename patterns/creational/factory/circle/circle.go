package circle

import "errors"

var InputErr = errors.New("invalid input")

type Circle struct {
	radius float64
}

func New(values []float64) (*Circle, error) {
	if len(values) != 1 {
		return nil, InputErr
	}
	radius := values[0]
	if radius <= 0 {
		return nil, InputErr
	}
	return &Circle{radius: radius}, nil
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
