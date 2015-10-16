package double_linked_list

import "testing"

func TestDoubleLinkedList(t *testing.T) {
	var ll = DoubleLinkedList{}
	if ll.size != 0 {
		t.Error("Size doesn't start at 0")
	}
	if ll.Head() != nil {
		t.Error("Head not nil")
	}
	if ll.Tail() != nil {
		t.Error("Tail not nil")
	}
}

func TestDoubleLinkedListInsNext(t *testing.T) {
	var list = new(DoubleLinkedList)
	/*
		Insert first element
	*/
	list.InsNext(nil, 10)
	switch {
	case list.size != 1:
		t.Error("Size not increased")
	case list.Head() != list.Tail():
		// Head and tail must be the same
		t.Error("Head and tail not equal for size 1 list")
	case list.Next(list.Head()) != nil || list.Prev(list.Head()) != nil:
		t.Error("Head should not have a Next or Prev value")
	}

	// Head data should the one given
	value, found := list.Data(list.Head()).(int)
	switch {
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
	switch {
	case list.size != 2:
		t.Error("Size not increased")
	case list.Next(list.Head()) != list.Tail():
		t.Error("Head next value is not set to tail")
	case list.Prev(list.Head()) != nil:
		t.Error("Head prev value is not nil")
	case list.Prev(list.Tail()) != list.Head():
		t.Error("Tail prev value is not set to Head")
	case list.Next(list.Tail()) != nil:
		t.Error("Tail next is not set to nil")
	}

	// Head data should still be the same
	value, found = list.Data(list.Head()).(int)
	switch {
	case !found:
		t.Error("data could not be converted back to int")
	case value != 10:
		t.Error("head data value incorrect")
	}
	// Tail data should be the new element
	value, found = list.Data(list.Tail()).(int)
	switch {
	case !found:
		t.Error("data could not be converted back to int")
	case value != 20:
		t.Error("head data value incorrect")
	}

	/*
	   Insert another element after head
	*/
	list.InsNext(list.Head(), 15)
	new_elem := list.Next(list.Head())
	switch {
	case list.size != 3:
		t.Error("Size not increased")
	case list.Next(list.Head()) != list.Prev(list.Tail()):
		t.Error("Head next value is not set to tail Prev")
	case list.Prev(list.Head()) != nil:
		t.Error("Head prev value is not nil")
	case list.Next(list.Head()) != new_elem:
		t.Error("Head next value is not new elem")
	case list.Prev(new_elem) != list.Head():
		t.Error("New elem Prev is not head")
	case list.Next(new_elem) != list.Tail():
		t.Error("New elem Next is not tail")
	case list.Prev(list.Tail()) != new_elem:
		t.Error("Tail prev value is not set to new value")
	case list.Next(list.Tail()) != nil:
		t.Error("Tail next is not set to nil")
	}

	/*
		Insert element after tail
	*/
	list.InsNext(list.Tail(), 5)
	if list.Size() != 4 {
		t.Error("Size not increased")
	}
	value, found = list.Data(list.Tail()).(int)
	switch {
	case !found:
		t.Error("data could not be converted back to int")
	case value != 5:
		t.Error("data value incorrect")
	}
}

func TestDoubleLinkedListInsPrev(t *testing.T) {
	var list = new(DoubleLinkedList)
	/*
		Insert first element
	*/
	list.InsPrev(nil, 10)
	switch {
	case list.Size() != 1:
		t.Error("Size not increased")
	case list.Head() != list.Tail():
		// Head and tail must be the same
		t.Error("Head and tail not equal for size 1 list")
	case list.Next(list.Head()) != nil || list.Prev(list.Head()) != nil:
		t.Error("Head should not have a Next or Prev value")
	}

	// Head data should the one given
	value, found := list.Data(list.Head()).(int)
	switch {
	case !found:
		t.Error("data could not be converted back to int")
	case value != 10:
		t.Error("data value incorrect")
	}

	/*
		Does not allow to insert another element with nil param
	*/
	err := list.InsPrev(nil, 20)
	if err == nil {
		t.Error("Doesn't return an error")
	}

	/*
		Insert element before Head
	*/
	list.InsPrev(list.Head(), 20)
	switch {
	case list.size != 2:
		t.Error("Size not increased")
	case list.Next(list.Head()) != list.Tail():
		t.Error("Head next value is not set to tail")
	case list.Prev(list.Head()) != nil:
		t.Error("Head prev value is not nil")
	case list.Prev(list.Tail()) != list.Head():
		t.Error("Tail prev value is not set to Head")
	case list.Next(list.Tail()) != nil:
		t.Error("Tail next is not set to nil")
	}

	// Tail data should still be the same
	value, found = list.Data(list.Tail()).(int)
	switch {
	case !found:
		t.Error("data could not be converted back to int")
	case value != 10:
		t.Error("tail data value incorrect")
	}
	// Head data should be the new element
	value, found = list.Data(list.Head()).(int)
	switch {
	case !found:
		t.Error("data could not be converted back to int")
	case value != 20:
		t.Error("head data value incorrect")
	}

	/*
	   Insert element before Tail
	*/
	list.InsPrev(list.Tail(), 15)
	new_elem := list.Prev(list.Tail())
	switch {
	case list.size != 3:
		t.Error("Size not increased")
	case list.Next(list.Head()) != list.Prev(list.Tail()):
		t.Error("Head next value is not set to tail Prev")
	case list.Prev(list.Head()) != nil:
		t.Error("Head prev value is not nil")
	case list.Next(list.Head()) != new_elem:
		t.Error("Head next value is not new elem")
	case list.Prev(new_elem) != list.Head():
		t.Error("New elem Prev is not head")
	case list.Next(new_elem) != list.Tail():
		t.Error("New elem Next is not tail")
	case list.Prev(list.Tail()) != new_elem:
		t.Error("Tail prev value is not set to new value")
	case list.Next(list.Tail()) != nil:
		t.Error("Tail next is not set to nil")
	}

	/*
		Insert element before head
	*/
	list.InsPrev(list.Head(), 5)
	if list.Size() != 4 {
		t.Error("Size not increased")
	}
	value, found = list.Data(list.Head()).(int)
	switch {
	case !found:
		t.Error("data could not be converted back to int")
	case value != 5:
		t.Error("data value incorrect")
	}
}

func TestDoubleLinkedListRem(t *testing.T) {
	var list = new(DoubleLinkedList)
	// Insert 5 elements
	for i := 1; i < 5; i++ {
		list.InsNext(list.Head(), 5*i)
	}

	if list.Size() != 4 {
		t.Error("Wrong size")
	}
	data, err := list.Rem(list.Next(list.Head()))
	switch {
	case err != nil:
		t.Error("Returned error")
	case data != 20:
		t.Error("Wrong value removed")
	case list.Size() != 3:
		t.Error("Wrong size")
	case list.Data(list.Head()) != 5 && list.Data(list.Next(list.Head())) != 15:
		t.Error("Wrong values")
	}

	data, err = list.Rem(list.Head())
	switch {
	case err != nil:
		t.Error("Returned error")
	case data != 5:
		t.Error("Wrong value removed")
	case list.Size() != 2:
		t.Error("Wrong size")
	case list.Data(list.Head()) != 15 && list.Data(list.Next(list.Head())) != 10:
		t.Error("Wrong values")
	}

	data, err = list.Rem(list.Tail())
	switch {
	case err != nil:
		t.Error("Returned error")
	case data != 10:
		t.Error("Wrong value removed")
	case list.Size() != 1:
		t.Error("Wrong size")
	case list.Data(list.Head()) != 15 && list.Data(list.Tail()) != 15:
		t.Error("Wrong values")
	case list.Head() != list.Tail():
		t.Error("Tail and Head are not the same on size 1 list")
	}
}

func TestDoubleLinkedListSize(t *testing.T) {
	var list = new(DoubleLinkedList)
	// Insert 5 elements
	for i := 0; i < 5; i++ {
		list.InsNext(list.Head(), 10)
	}
	if list.size != list.Size() {
		t.Error("Size() and size differ")
	}
	if list.Size() != 5 {
		t.Error("Wrong size")
	}
}

func TestDoubleLinkedListHead(t *testing.T) {
	var list = new(DoubleLinkedList)
	list.InsNext(nil, 10)
	if list.Data(list.Head()) != 10 {
		t.Error("Wrong value")
	}
	if list.head != list.Head() {
		t.Error("Head() and head don't match")
	}
}

func TestDoubleLinkedListTail(t *testing.T) {
	var list = new(DoubleLinkedList)
	list.InsNext(nil, 10)
	if list.Data(list.Tail()) != 10 {
		t.Error("Wrong value")
	}
	if list.tail != list.Tail() {
		t.Error("Tail() and tail don't match")
	}
}

func TestDoubleLinkedListIsHead(t *testing.T) {
	var list = new(DoubleLinkedList)
	list.InsNext(nil, 10)
	if !list.IsHead(list.Head()) {
		t.Error("Head() is not head")
	}
}

func TestDoubleLinkedListIsTail(t *testing.T) {
	var list = new(DoubleLinkedList)
	list.InsNext(nil, 10)
	if !list.IsTail(list.Tail()) {
		t.Error("Tail() is not tail")
	}
}

func TestDoubleLinkedListData(t *testing.T) {
	var list = new(DoubleLinkedList)
	list.InsNext(nil, 10)
	if list.Data(list.Head()) != 10 {
		t.Error("Not the correct data")
	}
}

func TestDoubleLinkedListNext(t *testing.T) {
	var list = new(DoubleLinkedList)
	list.InsNext(nil, 10)
	list.InsNext(list.Head(), 20)
	if list.Size() != 2 {
		t.Error("Wrong size")
	}
	if list.Next(list.Head()) != list.Tail() {
		t.Error("Tail is not next to head in size 2 list")
	}
}

func TestDoubleLinkedListPrev(t *testing.T) {
	var list = new(DoubleLinkedList)
	list.InsNext(nil, 10)
	list.InsNext(list.Head(), 20)
	if list.Size() != 2 {
		t.Error("Wrong size")
	}
	if list.Prev(list.Tail()) != list.Head() {
		t.Error("head is not prev to head in size 2 list")
	}
}
