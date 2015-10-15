package simple_linked_list

import "testing"

func TestSimpleLinkedList(t *testing.T) {
	var ll = SimpleLinkedList{}
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

func TestSimpleLinkedListInsNext(t *testing.T) {
	var list = new(SimpleLinkedList)
	// Insert first element
	list.InsNext(nil, 10)
	if list.size != 1 {
		t.Error("Size not increased")
	}
	// Head and tail must be the same
	if list.Head() != list.tail {
		t.Error("Head and tail not equal for size 1 list")
	}

	// Head data should the one given
	value, found := list.Data(list.Head()).(int)
	if !found {
		t.Error("data could not be converted back to int")
	}
	if value != 10 {
		t.Error("data value incorrect")
	}

	// Insert second element at head
	list.InsNext(nil, 20)
	if list.Size() != 2 {
		t.Error("Size not increased")
	}
	// Head and tail must be different
	if list.Head() == list.Tail() {
		t.Error("Head and tail equal for size 2 list")
	}

	// Head data should be the new value given
	value, found = list.Data(list.Head()).(int)
	if !found {
		t.Error("data could not be converted back to int")
	}
	if value != 20 {
		t.Error("head data value incorrect")
	}
	// Tail data should be the old value given
	value, found = list.Data(list.Tail()).(int)
	if !found {
		t.Error("data could not be converted back to int")
	}
	if value != 10 {
		t.Error("data value incorrect")
	}

	// Insert element after head
	list.InsNext(list.Head(), 15)
	if list.size != 3 {
		t.Error("Size not increased")
	}

	// Our data should be inserted after head
	value, found = list.Data(list.Next(list.Head())).(int)
	if !found {
		t.Error("data could not be converted back to int")
	}
	if value != 15 {
		t.Error("head data value incorrect")
	}

	// Head data should still be the same
	value, found = list.Data(list.Head()).(int)
	if !found {
		t.Error("data could not be converted back to int")
	}
	if value != 20 {
		t.Error("head data value incorrect")
	}
	// Tail data should still be the same
	value, found = list.Data(list.Tail()).(int)
	if !found {
		t.Error("data could not be converted back to int")
	}
	if value != 10 {
		t.Error("data value incorrect")
	}

	// Insert element after tail
	list.InsNext(list.Tail(), 5)
	if list.Size() != 4 {
		t.Error("Size not increased")
	}
	value, found = list.Data(list.Tail()).(int)
	if !found {
		t.Error("data could not be converted back to int")
	}
	if value != 5 {
		t.Error("data value incorrect")
	}
}

func TestSimpleLinkedListRemNext(t *testing.T) {
	var list = new(SimpleLinkedList)
	// Insert 5 elements
	for i := 5; i < 25; i += 5 {
		list.InsNext(nil, i)
	}
	if list.Size() != 4 {
		t.Error("Wring size")
	}

	// Try to remove after tail
	value, err := list.RemNext(list.Tail())
	if err == nil {
		t.Error("No error from removing after tail.")
	}
	_, ok := err.(LinkedListError)
	if !ok {
		t.Error("Wrong error type")
	}
	if value != nil {
		t.Error("Returned a value")
	}

	// Remove head
	value, err = list.RemNext(nil)
	if err != nil {
		t.Error("Couldn't remove head")
	}
	if value != 20 {
		t.Error("Wrong value")
	}
	if list.Data(list.Head()) != 15 {
		t.Error("Wrong head")
	}
	if list.Data(list.Tail()) != 5 {
		t.Error("Wrong tail")
	}
	if list.Size() != 3 {
		t.Error("Size not decreased")
	}

	// Remove an element
	value, err = list.RemNext(list.head)
	if err != nil {
		t.Error("Couldn't remove element")
	}
	if value != 10 {
		t.Error("Wrong value")
	}
	if list.Data(list.Head()) != 15 {
		t.Error("Wrong head")
	}
	if list.Data(list.Tail()) != 5 {
		t.Error("Wrong tail")
	}
	if list.Size() != 2 {
		t.Error("Size not decreased")
	}

	// Remove tail
	value, err = list.RemNext(list.head)
	if err != nil {
		t.Error("Couldn't remove tail")
	}
	if value != 5 {
		t.Error("Wrong value")
	}
	if list.Data(list.Head()) != 15 {
		t.Error("Wrong head")
	}
	if list.Tail() != list.Head() {
		t.Error("tail and head are not equal")
	}
	if list.Size() != 1 {
		t.Error("Size not decreased")
	}

	// Remove last element
	value, err = list.RemNext(nil)
	if err != nil {
		t.Error("Couldn't remove last element")
	}
	if value != 15 {
		t.Error("Wrong value")
	}
	if list.Head() != nil {
		t.Error("Head not set to nil")
	}
	if list.Tail() != list.Head() {
		t.Error("tail and head are not equal")
	}
	if list.Size() != 0 {
		t.Error("Size not decreased")
	}

	// Remove from empty list
	value, err = list.RemNext(nil)
	if err == nil {
		t.Error("Didn't return error removing from empty list")
	}
	_, ok = err.(LinkedListError)
	if !ok {
		t.Error("Wrong error type")
	}
	if value != nil {
		t.Error("Returned a value")
	}
}

func TestSimpleLinkedListSize(t *testing.T) {
	var list = new(SimpleLinkedList)
	// Insert 5 elements
	for i := 0; i < 5; i++ {
		list.InsNext(nil, 10)
	}
	if list.size != list.Size() {
		t.Error("Size() and size differ")
	}
	if list.Size() != 5 {
		t.Error("Wrong size")
	}
}

func TestSimpleLinkedListHead(t *testing.T) {
	var list = new(SimpleLinkedList)
	list.InsNext(nil, 10)
	if list.Data(list.Head()) != 10 {
		t.Error("Wrong value")
	}
	if list.head != list.Head() {
		t.Error("Head() and head don't match")
	}
}

func TestSimpleLinkedListTail(t *testing.T) {
	var list = new(SimpleLinkedList)
	list.InsNext(nil, 10)
	if list.Data(list.Tail()) != 10 {
		t.Error("Wrong value")
	}
	if list.tail != list.Tail() {
		t.Error("Tail() and tail don't match")
	}
}

func TestSimpleLinkedListIsHead(t *testing.T) {
	var list = new(SimpleLinkedList)
	list.InsNext(nil, 10)
	if !list.IsHead(list.Head()) {
		t.Error("Head() is not head")
	}
}

func TestSimpleLinkedListIsTail(t *testing.T) {
	var list = new(SimpleLinkedList)
	list.InsNext(nil, 10)
	if !list.IsTail(list.Tail()) {
		t.Error("Tail() is not tail")
	}
}

func TestSimpleLinkedListData(t *testing.T) {
	var list = new(SimpleLinkedList)
	list.InsNext(nil, 10)
	if list.Data(list.Head()) != 10 {
		t.Error("Not the correct data")
	}
}

func TestSimpleLinkedListNext(t *testing.T) {
	var list = new(SimpleLinkedList)
	list.InsNext(nil, 10)
	list.InsNext(nil, 20)
	if list.Size() != 2 {
		t.Error("Wrong size")
	}
	if list.Next(list.Head()) != list.Tail() {
		t.Error("Tail is not next to head in size 2 list")
	}
}
