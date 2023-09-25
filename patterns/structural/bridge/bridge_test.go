package bridge

import (
	"testing"

	"github.com/mvrdgs/asd-roadmap/patterns/structural/bridge/colors"
	"github.com/mvrdgs/asd-roadmap/patterns/structural/bridge/shapes"
	"github.com/stretchr/testify/assert"
)

func TestBridge(t *testing.T) {
	tests := []struct {
		name  string
		color colors.Color
		shape func(colors.Color) shapes.Shape
		want  string
	}{
		{
			name:  "Draw a red square",
			color: colors.Red(),
			shape: shapes.Square,
			want:  "Drawing a square with color red",
		},
		{
			name:  "Draw a blue circle",
			color: colors.Blue(),
			shape: shapes.Circle,
			want:  "Drawing a circle with color blue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shape := tt.shape(tt.color)
			assert.Equal(t, tt.want, shape.Draw())
		})
	}
}
