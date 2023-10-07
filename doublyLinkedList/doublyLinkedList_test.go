package doublyLinkedList

import "testing"

func TestDoublyLinkedList(t *testing.T) {
	list := DoublyLinkedList[int]{}

	list.append(5)
	list.append(7)
	list.append(9)

	if *list.get(2) != 9 {
		t.Fail()
	}
	if *list.removeAt(1) != 7 {
		t.Fail()
	}
	if list.length != 2 {
		t.Fail()
	}

	list.append(11)
	if *list.removeAt(1) != 9 {
		t.Fail()
	}
	if list.remove(9) != nil {
		t.Fail()
	}
	if *list.removeAt(0) != 5 {
		t.Fail()
	}
	if *list.removeAt(0) != 11 {
		t.Fail()
	}
	if list.length != 0 {
		t.Fail()
	}

	list.prepend(5)
	list.prepend(7)
	list.prepend(9)

	if *list.get(2) != 5 {
		t.Fail()
	}
	if *list.get(0) != 9 {
		t.Fail()
	}
	if *list.remove(9) != 9 {
		t.Fail()
	}
	if list.length != 2 {
		t.Fail()
	}
	if *list.get(0) != 7 {
		t.Fail()
	}
}
