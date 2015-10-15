package double_linked_list

import (
	. "github.com/nxvl/mastering-algorithms-with-go/double_list_elem"
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
		Data: data,
	}

	if list.Size() == 0 {
		list.head = new_element
		list.tail = new_element
	} else {
		new_element.Next = element.Next
		new_element.Prev = element
		if element.Next == nil {
			list.tail = new_element
		} else {
			element.Next.Prev = new_element
		}
		element.Next = new_element
	}

	list.size++
	return nil
}

func (list *DoubleLinkedList) InsPrev(element *DoubleListElem, data interface{}) error {
	if element == nil && list.Size() != 0 {
		return ErrInvalidInsert
	}
	new_element := &DoubleListElem{
		Data: data,
	}

	if list.Size() == 0 {
		list.head = new_element
		list.tail = new_element
	} else {
		new_element.Prev = element.Prev
		new_element.Next = element
		if element.Prev == nil {
			list.head = new_element
		} else {
			element.Prev.Next = new_element
		}
		element.Prev = new_element
	}

	list.size++
	return nil
}

func (list *DoubleLinkedList) RemNext(element *DoubleListElem) (interface{}, error) {
	return nil, nil
}

func (list *DoubleLinkedList) RemPrev(element *DoubleListElem) (interface{}, error) {
	return nil, nil
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
	return element.Data
}

func (list DoubleLinkedList) Next(element *DoubleListElem) *DoubleListElem {
	return element.Next
}

func (list DoubleLinkedList) Prev(element *DoubleListElem) *DoubleListElem {
	return element.Prev
}
