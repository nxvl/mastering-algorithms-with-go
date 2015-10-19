package linked_list

import (
	"testing"
)

func TestCircularLinkedList(t *testing.T) {
	var ll = CircularLinkedList{}
	if ll.size != 0 {
		t.Error("Size doesn't start at 0")
	}
	if ll.head != nil {
		t.Error("Head not nil")
	}
}

func TestCircularLinkedListInsNext(t *testing.T) {
	var list = new(CircularLinkedList)
	// Insert first element
	list.InsNext(nil, 5)
	switch {
	case list.Size() != 1:
		t.Error("Size is not 1")
	case list.Head() != list.Next(list.Head()):
		t.Error("Head doesn't point to self")
	}

	// Insert second element
	list.InsNext(list.Head(), 15)
	value_head, found_head := list.Data(list.Head()).(int)
	value_new, found_new := list.Data(list.Next(list.Head())).(int)
	switch {
	case list.Size() != 2:
		t.Error("Size was not increased")
	case list.Next(list.Next(list.Head())) != list.Head():
		t.Error("Not circular")
	case !found_head || !found_new:
		t.Error("Couldn't convert to int")
	case value_head != 5:
		t.Error("Wrong head value")
	case value_new != 15:
		t.Error("Wrong new value")
	}

	// Insert element after head
	list.InsNext(list.Head(), 10)
	values := make([]int, 3)
	last_elem := list.Next(list.Next(list.Head()))
	values[0], _ = list.Data(list.Head()).(int)
	values[1], _ = list.Data(list.Next(list.Head())).(int)
	values[2], _ = list.Data(last_elem).(int)
	switch {
	case list.Size() != 3:
		t.Error("Size was not increased")
	case values[0] != 5 && values[1] != 10 && values[2] != 15:
		t.Error("Incorrect values")
	case list.Next(last_elem) != list.Head():
		t.Error("Not circular")
	}

	// Insert after last
	list.InsNext(last_elem, 20)
	values = make([]int, 4)
	last_elem = list.Next(list.Next(list.Next(list.Head())))
	values[0], _ = list.Data(list.Head()).(int)
	values[1], _ = list.Data(list.Next(list.Head())).(int)
	values[2], _ = list.Data(list.Next(list.Next(list.Head()))).(int)
	values[3], _ = list.Data(last_elem).(int)
	switch {
	case list.Size() != 4:
		t.Error("Size was not increased")
	case values[0] != 5 && values[1] != 10 && values[2] != 15 && values[3] != 20:
		t.Error("Incorrect values")
	case list.Next(last_elem) != list.Head():
		t.Error("Not circular")
	}
}

func TestCircularLinkedListRemNext(t *testing.T) {
	list := new(CircularLinkedList)

	// Remove from empty list
	data, err := list.RemNext(nil)
	switch {
	case err == nil:
		t.Error("Doesn't return error")
	case data != nil:
		t.Error("Data returned")
	case list.Size() != 0:
		t.Error("Wrong list size")
	}

	for _, i := range []int{5, 20, 15, 10} {
		list.InsNext(list.Head(), i)
	}
	switch {
	case list.Size() != 4:
		t.Error("Wrong Size")
	case list.Data(list.Head()) != 5:
		t.Error("Wrong head")
	}

	// Remove after last elem
	last_elem := list.Next(list.Next(list.Next(list.Head())))
	data, err = list.RemNext(last_elem)
	switch {
	case list.Size() != 3:
		t.Error("List size didn't decreased")
	case err != nil:
		t.Error("Error returned")
	case data != 5:
		t.Error("Wrong value returned")
	case list.Data(list.Head()) != 10:
		t.Error("Wrong new head")
	case list.Next(last_elem) != list.Head():
		t.Error("Not circular")
	}

	// Remove element after head
	last_elem = list.Next(list.Next(list.Head()))
	data, err = list.RemNext(list.Head())
	switch {
	case list.Size() != 2:
		t.Error("List size didn't decreased")
	case err != nil:
		t.Error("Error returned")
	case data != 15:
		t.Error("Wrong value returned")
	case list.Next(list.Head()) != last_elem:
		t.Error("Head next wasn't updated")
	}

	// Leave list with 1 elem
	data, err = list.RemNext(list.Head())
	switch {
	case list.Size() != 1:
		t.Error("List size didn't decreased")
	case err != nil:
		t.Error("Error returned")
	case data != 20:
		t.Error("Wrong value returned")
	case list.Head() != list.Next(list.Head()):
		t.Error("Not circular")
	}

	// Remove last value
	data, err = list.RemNext(list.Head())
	switch {
	case list.Size() != 0:
		t.Error("List size didn't decreased")
	case err != nil:
		t.Error("Error returned")
	case data != 10:
		t.Error("Wrong value returned")
	case list.Head() != nil:
		t.Error("Head not set to nil")
	}

	// Remove from empty list
	data, err = list.RemNext(nil)
	switch {
	case err == nil:
		t.Error("Doesn't return error")
	case data != nil:
		t.Error("Data returned")
	case list.Size() != 0:
		t.Error("Wrong list size")
	}
}

func TestCircularLinkedListSize(t *testing.T) {
	var list = new(CircularLinkedList)
	for i := 0; i < 5; i++ {
		list.InsNext(list.Head(), i)
	}
	switch {
	case list.Size() != 5:
		t.Error("Wrong Size")
	case list.size != list.Size():
		t.Error("Size() and size differ")
	}
}

func TestCircularLinkedListHead(t *testing.T) {
	var list = new(CircularLinkedList)
	list.InsNext(nil, 10)
	switch {
	case list.Data(list.Head()) != 10:
		t.Error("Wrong value")
	case list.head != list.Head():
		t.Error("Head() and head don't match")
	}
}

func TestCircularLinkedListData(t *testing.T) {
	var list = new(CircularLinkedList)
	list.InsNext(nil, 10)
	if list.Data(list.Head()) != 10 {
		t.Error("Not the correct data")
	}
}

func TestCircularLinkedListNext(t *testing.T) {
	var list = new(CircularLinkedList)
	list.InsNext(nil, 10)
	list.InsNext(list.Head(), 20)
	switch {
	case list.Size() != 2:
		t.Error("Wrong size")
	case list.Next(list.Head()).data != 20:
		t.Error("20 is not next to head in size 2 list")
	}
}
