package datastruct

type Queue struct {
	elements []any
}

func NewQueue() *Queue {
	return &Queue{make([]any, 0)}
}

func (q *Queue) Enqueue(e any) {
	q.elements = append(q.elements, e)
}

func (q *Queue) Dequeue() any {
	e := q.elements[0]
	q.elements = q.elements[1:]
	return e
}

func (q *Queue) Peek() any {
	return q.elements[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) Clear() {
	q.elements = make([]any, 0)
}

func (q *Queue) Values() []any {
	return q.elements
}
