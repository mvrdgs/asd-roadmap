package datastruct

type Node struct {
	value any
	next  *Node
}

type LinkedList struct {
	head *Node
	last *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		last: nil,
		size: 0,
	}
}

func (l *LinkedList) Add(value any) {
	node := &Node{value, nil}
	if l.last != nil {
		l.last.next = node
		l.last = node
	} else {
		l.head = node
		l.last = node
	}
	l.size++
}

func (l *LinkedList) Get(index int) any {
	if index < 0 || index >= l.size {
		return nil
	}
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value
}

func (l *LinkedList) AddFirst(value any) {
	node := &Node{value, nil}
	if l.head != nil {
		node.next = l.head
		l.head = node
	} else {
		l.head = node
		l.last = node
	}
	l.size++
}

func (l *LinkedList) Length() int {
	return l.size
}

func (l *LinkedList) Values() []any {
	values := make([]any, 0)
	node := l.head
	for node != nil {
		values = append(values, node.value)
		node = node.next
	}
	return values
}

func (l *LinkedList) Remove(value any) {
	if l.head == nil {
		return
	}
	if l.head.value == value {
		if l.head == l.last {
			l.head = nil
			l.last = nil
			l.size--
			return
		}
		l.head = l.head.next
		l.size--
		return
	}
	prev := l.head
	for prev.next != nil {
		if prev.next.value == value {
			prev.next = prev.next.next
			l.size--
			return
		}
		prev = prev.next
	}
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.last = nil
	l.size = 0
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Contains(value any) bool {
	node := l.head
	for node != nil {
		if node.value == value {
			return true
		}
		node = node.next
	}
	return false
}
