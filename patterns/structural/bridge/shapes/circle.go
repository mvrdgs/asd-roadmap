package shapes

import "github.com/mvrdgs/asd-roadmap/patterns/structural/bridge/colors"

type circle struct {
	color colors.Color
}

func Circle(color colors.Color) Shape {
	return &circle{
		color: color,
	}
}

func (c *circle) Draw() string {
	return "Drawing a circle with color " + c.color.Fill()
}
