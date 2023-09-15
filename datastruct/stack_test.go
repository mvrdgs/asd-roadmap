package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	assert.True(t, s.IsEmpty())

	s.Push(1)
	assert.Equal(t, 1, s.Size())
	assert.Equal(t, 1, s.Peek())
	assert.Equal(t, 1, s.Pop())
	assert.True(t, s.IsEmpty())

	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Size())
	assert.Equal(t, 3, s.Peek())
	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Size())
	assert.Equal(t, 2, s.Peek())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Size())
	assert.Equal(t, 1, s.Peek())
	assert.Equal(t, 1, s.Pop())
	assert.True(t, s.IsEmpty())

	s.Push(3)
	s.Push(4)
	s.Push(5)
	assert.Equal(t, 3, s.Size())

	s.Clear()
	assert.True(t, s.IsEmpty())
}
