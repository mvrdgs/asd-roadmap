package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()

	assert.Equal(t, 0, q.Size())

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 3, q.Size())
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 2, q.Peek())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 3, q.Peek())
	assert.Equal(t, 3, q.Dequeue())
	assert.Equal(t, true, q.IsEmpty())

	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	assert.Equal(t, []any{4, 5, 6}, q.Values())

	q.Clear()
	assert.Equal(t, true, q.IsEmpty())
}
