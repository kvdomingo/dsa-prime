package stack

type Node[T any] struct {
	value T
	prev *Node[T]
}

type Stack[T any] struct {
	length uint
	head *Node[T]
}

func (s *Stack[T]) push(item T) {
	node := Node[T]{value: item}
	s.length++

	if s.head == nil {
		s.head = &node
		return
	}

	node.prev = s.head
	s.head = &node
}

func (s *Stack[T]) pop() *T {
	if s.head == nil { return nil }
	s.length--

	if s.length == 0 {
		head := *s.head
		s.head = nil
		return &head.value
	}

	head := *s.head
	s.head = &*head.prev
	head.prev = nil
	return &head.value
}

func (s *Stack[T]) peek() *T {
	if s.head == nil { return nil }
	return &s.head.value
}
