package datastruct

type Array struct {
	elements []any
}

func NewArray() *Array {
	return &Array{make([]any, 0)}
}

func (a *Array) Add(e any) {
	a.elements = append(a.elements, e)
}

func (a *Array) Get(index int) any {
	return a.elements[index]
}

func (a *Array) Set(index int, e any) {
	a.elements[index] = e
}

func (a *Array) Remove(index int) any {
	e := a.elements[index]
	a.elements = append(a.elements[:index], a.elements[index+1:]...)
	return e
}

func (a *Array) IndexOf(e any) int {
	for i, v := range a.elements {
		if v == e {
			return i
		}
	}
	return -1
}

func (a *Array) Contains(e any) bool {
	return a.IndexOf(e) != -1
}

func (a *Array) IsEmpty() bool {
	return len(a.elements) == 0
}

func (a *Array) Size() int {
	return len(a.elements)
}

func (a *Array) Clear() {
	a.elements = make([]any, 0)
}

func (a *Array) Values() []any {
	return a.elements
}
