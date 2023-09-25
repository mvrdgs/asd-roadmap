package factory

import (
	"errors"

	"github.com/mvrdgs/asd-roadmap/patterns/creational/factory/circle"
	"github.com/mvrdgs/asd-roadmap/patterns/creational/factory/rectangle"
)

var ErrInvalidShapeType = errors.New("invalid shape type")

type Shape interface {
	Area() float64
}

type ShapeFactory struct{}

func New() *ShapeFactory {
	return &ShapeFactory{}
}

func (ShapeFactory) CreateShape(shapeType string, values ...float64) (Shape, error) {
	switch shapeType {
	case "circle":
		return circle.New(values)
	case "rectangle":
		return rectangle.New(values)
	default:
		return nil, ErrInvalidShapeType
	}
}
