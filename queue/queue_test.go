package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	list := Queue[int]{}

	list.enqueue(5)
	list.enqueue(7)
	list.enqueue(9)

	if *list.deque() != 5 { t.Fail() }
	if list.length != 2 { t.Fail() }

	list.enqueue(11)
	if *list.deque() != 7 { t.Fail() }
	if *list.deque() != 9 { t.Fail() }
	if *list.peek() != 11 { t.Fail() }
	if *list.deque() != 11 { t.Fail() }
	if list.deque() != nil { t.Fail() }
	if list.length != 0 { t.Fail() }

	list.enqueue(69)
	if *list.peek() != 69 { t.Fail() }
	if list.length != 1 { t.Fail() }
}
