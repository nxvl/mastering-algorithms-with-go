package linked_list

type SimpleLinkedList struct {
	size int
	head *SimpleListElem
	tail *SimpleListElem
}

func (list *SimpleLinkedList) InsNext(element *SimpleListElem, data interface{}) error {
	new_element := SimpleListElem{
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

func (list *SimpleLinkedList) RemNext(element *SimpleListElem) (interface{}, error) {
	var data interface{}
	if list.Size() == 0 {
		return nil, ErrEmptyList
	}
	if element == nil {
		data = list.head.data
		list.head = list.head.next
		if list.Size() == 1 {
			list.tail = nil
		}
	} else {
		if element.next == nil {
			return nil, ErrNoElementToRemove
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
	return element.data
}

func (list SimpleLinkedList) Next(element *SimpleListElem) *SimpleListElem {
	return element.next
}
