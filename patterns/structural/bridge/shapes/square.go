package shapes

import (
	"fmt"

	"github.com/mvrdgs/asd-roadmap/patterns/structural/bridge/colors"
)

type square struct {
	color colors.Color
}

func Square(color colors.Color) Shape {
	return &square{
		color: color,
	}
}

func (s *square) Draw() string {
	return fmt.Sprintf("Drawing a square with color %s", s.color.Fill())
}
