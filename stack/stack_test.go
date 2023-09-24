package stack

import "testing"

func TestStack(t *testing.T) {
	list := Stack[int]{}

	list.push(5)
	list.push(7)
	list.push(9)

	if *list.pop() != 9 { t.Fail() }
	if list.length != 2 { t.Fail() }

	list.push(11)
	if *list.pop() != 11 { t.Fail() }
	if *list.pop() != 7 { t.Fail() }
	if *list.peek() != 5 { t.Fail() }
	if *list.pop() != 5 { t.Fail() }
	if list.pop() != nil { t.Fail() }

	list.push(69)
	if *list.peek() != 69 { t.Fail() }
	if list.length != 1 { t.Fail() }
}
