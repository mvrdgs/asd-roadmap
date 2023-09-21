package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArray_Add(t *testing.T) {
	arr := NewArray()
	require.Len(t, arr.elements, 0)

	arr.Add("element 1")
	require.Len(t, arr.elements, 1)
	assert.Equal(t, "element 1", arr.elements[0])
}

func TestArray_Get(t *testing.T) {
	arr := NewArray()
	arr.Add("element 1")
	arr.Add("element 2")

	assert.Equal(t, "element 1", arr.Get(0))
	assert.Equal(t, "element 2", arr.Get(1))
}

func TestArray_Set(t *testing.T) {
	arr := NewArray()
	arr.Add("element 1")
	arr.Add("element 2")

	arr.Set(0, "element 3")
	assert.Equal(t, "element 3", arr.Get(0))
}

func TestArray_Remove(t *testing.T) {
	arr := NewArray()
	arr.Add("element 1")
	arr.Add("element 2")
	arr.Add("element 3")

	assert.Equal(t, "element 2", arr.Remove(1))
	assert.Equal(t, "element 3", arr.Get(1))
}

func TestArray_IndexOf(t *testing.T) {
	arr := NewArray()
	arr.Add("element 1")
	arr.Add("element 2")
	arr.Add("element 3")

	assert.Equal(t, 1, arr.IndexOf("element 2"))
	assert.Equal(t, -1, arr.IndexOf("element 4"))
}

func TestArray_Contains(t *testing.T) {
	arr := NewArray()
	arr.Add("element 1")
	arr.Add("element 2")
	arr.Add("element 3")

	assert.True(t, arr.Contains("element 2"))
	assert.False(t, arr.Contains("element 4"))
}

func TestArray_IsEmpty(t *testing.T) {
	arr := NewArray()
	assert.True(t, arr.IsEmpty())

	arr.Add("element 1")
	assert.False(t, arr.IsEmpty())
}

func TestArray_Size(t *testing.T) {
	arr := NewArray()
	assert.Equal(t, 0, arr.Size())

	arr.Add("element 1")
	assert.Equal(t, 1, arr.Size())
}

func TestArray_Clear(t *testing.T) {
	arr := NewArray()
	arr.Add("element 1")
	arr.Add("element 2")
	arr.Add("element 3")

	arr.Clear()
	assert.Equal(t, 0, arr.Size())
}

func TestArray_Values(t *testing.T) {
	arr := NewArray()
	arr.Add("element 1")
	arr.Add("element 2")
	arr.Add("element 3")

	assert.Equal(t, []any{"element 1", "element 2", "element 3"}, arr.Values())
}
