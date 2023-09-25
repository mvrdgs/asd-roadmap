package circle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircle_New(t *testing.T) {
	tests := []struct {
		name      string
		values    []float64
		expect    *Circle
		expectErr error
	}{
		{
			name:      "no values",
			values:    []float64{},
			expect:    nil,
			expectErr: InputErr,
		},
		{
			name:      "more than 1 value",
			values:    []float64{1, 1},
			expect:    nil,
			expectErr: InputErr,
		},
		{
			name:      "invalid radius",
			values:    []float64{-1},
			expect:    nil,
			expectErr: InputErr,
		},
		{
			name:      "valid radius",
			values:    []float64{1},
			expect:    &Circle{radius: 1},
			expectErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := New(tt.values)
			assert.ErrorIs(t, err, tt.expectErr)
			assert.Equal(t, tt.expect, c)
		})
	}
}

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name   string
		values []float64
		expect float64
	}{
		{
			name:   "radius 1",
			values: []float64{1},
			expect: 3.14,
		},
		{
			name:   "radius 2",
			values: []float64{2},
			expect: 12.56,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := New(tt.values)
			assert.NoError(t, err)
			assert.Equal(t, tt.expect, c.Area())
		})
	}
}
