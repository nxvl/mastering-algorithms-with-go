package linked_list

import "testing"

func TestCircularDoublyLinkedList(t *testing.T) {
	var ll = CircularDoublyLinkedList{}
	switch {
	case ll.size != 0:
		t.Error("Size doesn't start at 0")
	case ll.Head() != nil:
		t.Error("Head not nil")
	}
}

func TestCircularDoublyLinkedListInsNext(t *testing.T) {
	var list = new(CircularDoublyLinkedList)
	/*
		Insert first element
	*/
	list.InsNext(nil, 10)
	value, found := list.Data(list.Head()).(int)
	switch {
	case list.size != 1:
		t.Error("Size not increased")
	case list.Head() != list.Next(list.Head()) && list.Head() != list.Prev(list.Head()):
		t.Error("Not circular")
	case !found:
		t.Error("data could not be converted back to int")
	case value != 10:
		t.Error("data value incorrect")
	}

	/*
		Does not allow to insert another element with nil param
	*/
	err := list.InsNext(nil, 20)
	if err == nil {
		t.Error("Doesn't return an error")
	}

	/*
		Insert element after head
	*/
	list.InsNext(list.Head(), 20)
	new_elem := list.Next(list.Head())
	switch {
	case list.size != 2:
		t.Error("Size not increased")
	case list.Data(new_elem) != 20:
		t.Error("New element value is wrong")
	case list.Data(list.Head()) != 10:
		t.Error("Bad head data")
	case list.Prev(list.Head()) != nil:
		t.Error("Head prev value is not nil")
	case list.Prev(list.Head()) != new_elem && list.Next(new_elem) != list.Head():
		t.Error("Not circular")
	}

	/*
	   Insert another element after head
	*/
	list.InsNext(list.Head(), 15)
	new_elem = list.Next(list.Head())
	switch {
	case list.size != 3:
		t.Error("Size not increased")
	case list.Prev(list.Head()) != nil:
		t.Error("Head prev value is not nil")
	case list.Prev(list.Head()) != list.Next(new_elem):
		t.Error("Not circular")
	case list.Prev(new_elem) != list.Head():
		t.Error("New elem Prev is not head")
	}
}
