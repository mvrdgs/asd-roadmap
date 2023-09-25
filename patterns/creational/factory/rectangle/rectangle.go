package rectangle

import "errors"

var InputErr = errors.New("invalid input")

type Rectangle struct {
	height float64
	width  float64
}

func New(values []float64) (*Rectangle, error) {
	if len(values) != 2 {
		return nil, InputErr
	}
	height, width := values[0], values[1]
	if height <= 0 || width <= 0 {
		return nil, InputErr
	}
	return &Rectangle{height: height, width: width}, nil
}

func (r *Rectangle) Area() float64 {
	return r.height * r.width
}
