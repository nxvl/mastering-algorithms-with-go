package simple_linked_list

import (
	. "github.com/nxvl/mastering-algorithms-with-go/linked_list_errors"
	. "github.com/nxvl/mastering-algorithms-with-go/simple_list_elem"
)

type SimpleLinkedList struct {
	size int
	head *SimpleListElem
	tail *SimpleListElem
}

func (list *SimpleLinkedList) InsNext(element *SimpleListElem, data interface{}) error {
	new_element := SimpleListElem{
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

func (list *SimpleLinkedList) RemNext(element *SimpleListElem) (interface{}, error) {
	var data interface{}
	if list.Size() == 0 {
		return nil, ErrEmptyList
	}
	if element == nil {
		data = list.head.Data
		list.head = list.head.Next
		if list.Size() == 1 {
			list.tail = nil
		}
	} else {
		if element.Next == nil {
			return nil, ErrNoElementToRemove
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

func (list SimpleLinkedList) Head() *SimpleListElem {
	return list.head
}

func (list SimpleLinkedList) Tail() *SimpleListElem {
	return list.tail
}

func (list SimpleLinkedList) IsHead(element *SimpleListElem) bool {
	return element == list.head
}

func (list SimpleLinkedList) IsTail(element *SimpleListElem) bool {
	return element == list.tail
}

func (list SimpleLinkedList) Data(element *SimpleListElem) interface{} {
	return element.Data
}

func (list SimpleLinkedList) Next(element *SimpleListElem) *SimpleListElem {
	return element.Next
}
