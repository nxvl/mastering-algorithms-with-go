package linked_list

import "errors"

type LinkedListError error

var (
	errEmptyList         = LinkedListError(errors.New("Invalid operation: list is empty"))
	errNoElementToRemove = LinkedListError(errors.New("Invalid operation: no element to remove"))
)

type ListElem struct {
	data interface{}
	next *ListElem
}

type LinkedList struct {
	size int
	head *ListElem
	tail *ListElem
}

func (list *LinkedList) InsNext(element *ListElem, data interface{}) error {
	new_element := ListElem{
		data: data,
	}
	if element == nil {
		if list.Size() == 0 {
			list.tail = &new_element
		}
		new_element.next = list.head
		list.head = &new_element
	} else {
		if element.next == nil {
			list.tail = &new_element
		}
		new_element.next = element.next
		element.next = &new_element
	}
	list.size++
	return nil
}

func (list *LinkedList) RemNext(element *ListElem) (interface{}, error) {
	var data interface{}
	if list.Size() == 0 {
		return nil, errEmptyList
	}
	if element == nil {
		data = list.head.data
		list.head = list.head.next
		if list.Size() == 1 {
			list.tail = nil
		}
	} else {
		if element.next == nil {
			return nil, errNoElementToRemove
		}
		data = element.next.data
		element.next = element.next.next
		if element.next == nil {
			list.tail = element
		}
	}
	list.size--
	return data, nil
}

func (list LinkedList) Size() int {
	return list.size
}

func (list LinkedList) Head() *ListElem {
	return list.head
}

func (list LinkedList) Tail() *ListElem {
	return list.tail
}

func (list LinkedList) IsHead(element *ListElem) bool {
	return element == list.head
}

func (list LinkedList) IsTail(element *ListElem) bool {
	return element == list.tail
}

func (list LinkedList) Data(element *ListElem) interface{} {
	return element.data
}

func (list LinkedList) Next(element *ListElem) *ListElem {
	return element.next
}
