package simple_linked_list

import (
	"errors"

	"github.com/nxvl/mastering-algorithms-with-go/simple_list_elem"
)

type LinkedListError error

var (
	errEmptyList         = LinkedListError(errors.New("Invalid operation: list is empty"))
	errNoElementToRemove = LinkedListError(errors.New("Invalid operation: no element to remove"))
)

type SimpleLinkedList struct {
	size int
	head *simple_list_elem.SimpleListElem
	tail *simple_list_elem.SimpleListElem
}

func (list *SimpleLinkedList) InsNext(element *simple_list_elem.SimpleListElem, data interface{}) error {
	new_element := simple_list_elem.SimpleListElem{
		Data: data,
	}
	if element == nil {
		if list.Size() == 0 {
			list.tail = &new_element
		}
		new_element.Next = list.head
		list.head = &new_element
	} else {
		if element.Next == nil {
			list.tail = &new_element
		}
		new_element.Next = element.Next
		element.Next = &new_element
	}
	list.size++
	return nil
}

func (list *SimpleLinkedList) RemNext(element *simple_list_elem.SimpleListElem) (interface{}, error) {
	var data interface{}
	if list.Size() == 0 {
		return nil, errEmptyList
	}
	if element == nil {
		data = list.head.Data
		list.head = list.head.Next
		if list.Size() == 1 {
			list.tail = nil
		}
	} else {
		if element.Next == nil {
			return nil, errNoElementToRemove
		}
		data = element.Next.Data
		element.Next = element.Next.Next
		if element.Next == nil {
			list.tail = element
		}
	}
	list.size--
	return data, nil
}

func (list SimpleLinkedList) Size() int {
	return list.size
}

func (list SimpleLinkedList) Head() *simple_list_elem.SimpleListElem {
	return list.head
}

func (list SimpleLinkedList) Tail() *simple_list_elem.SimpleListElem {
	return list.tail
}

func (list SimpleLinkedList) IsHead(element *simple_list_elem.SimpleListElem) bool {
	return element == list.head
}

func (list SimpleLinkedList) IsTail(element *simple_list_elem.SimpleListElem) bool {
	return element == list.tail
}

func (list SimpleLinkedList) Data(element *simple_list_elem.SimpleListElem) interface{} {
	return element.Data
}

func (list SimpleLinkedList) Next(element *simple_list_elem.SimpleListElem) *simple_list_elem.SimpleListElem {
	return element.Next
}
