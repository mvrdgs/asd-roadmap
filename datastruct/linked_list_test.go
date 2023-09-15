package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Add(t *testing.T) {
	linkedList := NewLinkedList()

	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(5)

	assert.Equal(t, 5, linkedList.Length())
}

func TestLinkedList_Get(t *testing.T) {
	linkedList := NewLinkedList()

	v := linkedList.Get(0)
	assert.Nil(t, v)

	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)

	assert.Equal(t, 2, linkedList.Get(1))
}

func TestLinkedList_AddFirst(t *testing.T) {
	linkedList := NewLinkedList()

	linkedList.AddFirst(1)
	assert.Equal(t, 1, linkedList.Get(0))

	linkedList.Add(2)
	linkedList.AddFirst(3)
	assert.Equal(t, 3, linkedList.Get(0))
}

func TestLinkedList_Values(t *testing.T) {
	linkedList := NewLinkedList()

	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(5)

	assert.Equal(t, []any{1, 2, 3, 4, 5}, linkedList.Values())
}

func TestLinkedList_Remove(t *testing.T) {
	linkedList := NewLinkedList()

	linkedList.Remove(1)
	assert.Equal(t, []any{}, linkedList.Values())

	linkedList.Add(1)
	assert.Equal(t, []any{1}, linkedList.Values())
	linkedList.Remove(1)
	assert.Equal(t, []any{}, linkedList.Values())

	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(5)

	linkedList.Remove(1)
	assert.Equal(t, []any{2, 3, 4, 5}, linkedList.Values())

	linkedList.Remove(5)
	assert.Equal(t, []any{2, 3, 4}, linkedList.Values())
}

func TestLinkedList_Clear(t *testing.T) {
	linkedList := NewLinkedList()

	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(5)

	linkedList.Clear()
	assert.Equal(t, []any{}, linkedList.Values())
}

func TestLinkedList_IsEmpty(t *testing.T) {
	linkedList := NewLinkedList()

	assert.True(t, linkedList.IsEmpty())

	linkedList.Add(1)
	assert.False(t, linkedList.IsEmpty())
}

func TestLinkedList_Contains(t *testing.T) {
	linkedList := NewLinkedList()

	linkedList.Add(1)
	assert.False(t, linkedList.Contains(2))

	linkedList.Add(2)
	assert.True(t, linkedList.Contains(2))
}
