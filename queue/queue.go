package queue

type Node[T any] struct {
	value T
	next *Node[T]
}

type Queue[T any] struct {
	length uint
	head *Node[T]
	tail *Node[T]
}

func (q *Queue[T]) enqueue(item T) {
	node := Node[T]{value: item}
	q.length++

	if q.tail == nil {
		q.head, q.tail = &node, &node
		return
	}

	q.tail.next = &node
	q.tail = &node
}

func (q *Queue[T]) deque() *T {
	if q.head == nil {
		return nil
	}

	q.length--
	head := q.head
	q.head = q.head.next
	head.next = nil

	if q.length == 0 {
		q.tail = nil
	}

	return &head.value
}

func (q *Queue[T]) peek() *T {
	if q.head == nil {
		return nil
	}

	return &q.head.value
}
