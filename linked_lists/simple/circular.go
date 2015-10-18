package simple_list

import (
	. "github.com/nxvl/mastering-algorithms-with-go/linked_lists/ll_errors"
)

type CircularLinkedList struct {
	size int
	head *SimpleListElem
}

func (list *CircularLinkedList) InsNext(element *SimpleListElem, data interface{}) error {
	new_element := SimpleListElem{
		data: data,
	}
	if list.Size() == 0 {
		new_element.next = &new_element
		list.head = &new_element
	} else {
		new_element.next = element.next
		element.next = &new_element
	}

	list.size++
	return nil
}

func (list *CircularLinkedList) RemNext(element *SimpleListElem) (interface{}, error) {
	if list.Size() == 0 {
		return nil, ErrEmptyList
	}
	data := list.Data(list.Next(element))
	if list.Next(element) == element {
		list.head = nil
	} else {
		elem_to_remove := list.Next(element)
		element.next = list.Next(elem_to_remove)
		if elem_to_remove == list.Head() {
			list.head = list.Next(elem_to_remove)
		}
	}

	list.size--
	return data, nil
}

func (list CircularLinkedList) Size() int {
	return list.size
}

func (list CircularLinkedList) Head() *SimpleListElem {
	return list.head
}

func (list CircularLinkedList) Data(element *SimpleListElem) interface{} {
	return element.data
}

func (list CircularLinkedList) Next(element *SimpleListElem) *SimpleListElem {
	return element.next
}
