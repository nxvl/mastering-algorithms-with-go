package double_list

import (
	. "github.com/nxvl/mastering-algorithms-with-go/linked_list_errors"
)

type DoubleLinkedList struct {
	size int
	head *DoubleListElem
	tail *DoubleListElem
}

func (list *DoubleLinkedList) InsNext(element *DoubleListElem, data interface{}) error {
	if element == nil && list.Size() != 0 {
		return ErrInvalidInsert
	}
	new_element := &DoubleListElem{
		data: data,
	}

	if list.Size() == 0 {
		list.head = new_element
		list.tail = new_element
	} else {
		new_element.next = element.next
		new_element.prev = element
		if element.next == nil {
			list.tail = new_element
		} else {
			element.next.prev = new_element
		}
		element.next = new_element
	}

	list.size++
	return nil
}

func (list *DoubleLinkedList) InsPrev(element *DoubleListElem, data interface{}) error {
	if element == nil && list.Size() != 0 {
		return ErrInvalidInsert
	}
	new_element := &DoubleListElem{
		data: data,
	}

	if list.Size() == 0 {
		list.head = new_element
		list.tail = new_element
	} else {
		new_element.prev = element.prev
		new_element.next = element
		if element.prev == nil {
			list.head = new_element
		} else {
			element.prev.next = new_element
		}
		element.prev = new_element
	}

	list.size++
	return nil
}

func (list *DoubleLinkedList) Rem(element *DoubleListElem) (interface{}, error) {
	switch {
	case element == nil:
		return nil, ErrNoElementToRemove
	case list.Size() == 0:
		return nil, ErrEmptyList
	}
	data := list.Data(element)

	if list.IsHead(element) {
		list.head = list.Next(element)
		if list.Head() == nil {
			list.tail = nil
		} else {
			element.next.prev = nil
		}
	} else {
		element.prev.next = element.next
		if element.next == nil {
			list.tail = element.prev
		} else {
			element.next.prev = element.prev
		}
	}

	list.size--
	return data, nil
}

func (list DoubleLinkedList) Size() int {
	return list.size
}

func (list DoubleLinkedList) Head() *DoubleListElem {
	return list.head
}

func (list DoubleLinkedList) Tail() *DoubleListElem {
	return list.tail
}

func (list DoubleLinkedList) IsHead(element *DoubleListElem) bool {
	return element == list.head
}

func (list DoubleLinkedList) IsTail(element *DoubleListElem) bool {
	return element == list.tail
}

func (list DoubleLinkedList) Data(element *DoubleListElem) interface{} {
	return element.data
}

func (list DoubleLinkedList) Next(element *DoubleListElem) *DoubleListElem {
	return element.next
}

func (list DoubleLinkedList) Prev(element *DoubleListElem) *DoubleListElem {
	return element.prev
}
