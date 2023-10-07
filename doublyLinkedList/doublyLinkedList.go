package doublyLinkedList

type Node[T comparable] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	length uint
	head   *Node[T]
	tail   *Node[T]
}

type IDoublyLinkedList[T comparable] interface {
	prepend(item T)
	insertAt(item T, idx uint)
	append(item T)
	remove(item T) *T
	get(idx uint) *T
	removeAt(idx uint) *T
}

func (d *DoublyLinkedList[T]) getAt(idx uint) *Node[T] {
	curr := d.head
	for i := uint(0); curr != nil && i < idx; i++ {
		curr = curr.next
	}
	return curr
}

func (d *DoublyLinkedList[T]) removeNode(node *Node[T]) *T {
	d.length--
	if d.length == 0 {
		out := d.head.value
		d.head = nil
		d.tail = nil
		return &out
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == d.head {
		d.head = node.next
	}

	if node == d.tail {
		d.tail = node.prev
	}

	node.prev = nil
	node.next = nil

	return &node.value
}

func (d *DoublyLinkedList[T]) prepend(item T) {
	node := &Node[T]{value: item}
	d.length++
	if d.head == nil {
		d.head = node
		d.tail = node
		return
	}

	node.next = d.head
	d.head.prev = node
	d.head = node
}

func (d *DoublyLinkedList[T]) insertAt(item T, idx uint) {
	if idx > d.length {
		panic("oh no")
	} else if idx == d.length {
		d.append(item)
		return
	} else if idx == 0 {
		d.prepend(item)
		return
	}

	d.length++
	curr := d.head
	for i := uint(0); curr != nil && i < idx; i++ {
		curr = curr.next
	}

	node := &Node[T]{value: item}
	node.next = curr
	node.prev = curr.prev
	curr.prev = node
	if node.prev != nil {
		node.prev.next = curr
	}
}

func (d *DoublyLinkedList[T]) append(item T) {
	d.length++
	node := &Node[T]{value: item}
	if d.tail == nil {
		d.head = node
		d.tail = node
		return
	}

	node.prev = d.tail
	d.tail.next = node
	d.tail = node
}

func (d *DoublyLinkedList[T]) remove(item T) *T {
	curr := d.head
	for i := uint(0); curr != nil && i < d.length; i++ {
		if curr.value == item {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return nil
	}

	return &curr.value
}

func (d *DoublyLinkedList[T]) get(idx uint) *T {
	node := d.getAt(idx)
	return &node.value
}

func (d *DoublyLinkedList[T]) removeAt(idx uint) *T {
	node := d.getAt(idx)
	if node == nil {
		return nil
	}

	return d.removeNode(node)
}
