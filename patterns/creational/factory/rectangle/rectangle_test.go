package rectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangle_New(t *testing.T) {
	tests := []struct {
		name      string
		values    []float64
		expect    *Rectangle
		expectErr error
	}{
		{
			name:      "less than 2 values",
			values:    []float64{1},
			expect:    nil,
			expectErr: InputErr,
		},
		{
			name:      "more than 2 values",
			values:    []float64{1, 1, 1},
			expect:    nil,
			expectErr: InputErr,
		},
		{
			name:      "invalid height",
			values:    []float64{-1, 1},
			expect:    nil,
			expectErr: InputErr,
		},
		{
			name:      "invalid width",
			values:    []float64{1, -1},
			expect:    nil,
			expectErr: InputErr,
		},
		{
			name:      "valid height and width",
			values:    []float64{1, 1},
			expect:    &Rectangle{height: 1, width: 1},
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

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name   string
		values []float64
		expect float64
	}{
		{
			name:   "height 1 and width 1",
			values: []float64{1, 1},
			expect: 1,
		},
		{
			name:   "height 2 and width 2",
			values: []float64{2, 2},
			expect: 4,
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
