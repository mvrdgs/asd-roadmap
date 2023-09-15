package datastruct

type Stack struct {
	elements []any
}

func NewStack() *Stack {
	return &Stack{make([]any, 0)}
}

func (s *Stack) Push(e any) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() any {
	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return e
}

func (s *Stack) Peek() any {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Clear() {
	s.elements = make([]any, 0)
}
