package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateShape(t *testing.T) {
	f := New()

	tests := []struct {
		name      string
		shape     string
		input     []float64
		expect    float64
		expectErr error
	}{
		{
			name:   "rectangle",
			shape:  "rectangle",
			input:  []float64{2, 3},
			expect: 6,
		},
		{
			name:   "circle",
			shape:  "circle",
			input:  []float64{2},
			expect: 12.56,
		},
		{
			name:      "invalid shape",
			shape:     "invalid",
			expectErr: ErrInvalidShapeType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shape, err := f.CreateShape(tt.shape, tt.input...)
			assert.ErrorIs(t, err, tt.expectErr)
			if shape != nil {
				assert.Equal(t, tt.expect, shape.Area())
			}
		})
	}
}
