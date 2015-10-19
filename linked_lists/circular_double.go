package linked_list

type CircularDoublyLinkedList struct {
	size int
	head *DoubleListElem
}

func (list *CircularDoublyLinkedList) InsNext(element *DoubleListElem, data interface{}) error {
	new_element := &DoubleListElem{
		data: data,
	}
	if list.Size() == 0 {
		new_element.next = new_element
		new_element.prev = new_element
		list.head = new_element
	} else {
		if element == nil {
			return ErrInvalidInsert
		}
		new_element.next = element.next
		new_element.prev = element
		element.next.prev = new_element
		element.next = new_element
	}

	list.size++
	return nil
}

func (list *CircularDoublyLinkedList) InsPrev(element *DoubleListElem, data interface{}) error {
	new_element := &DoubleListElem{
		data: data,
	}
	if list.Size() == 0 {
		new_element.next = new_element
		new_element.prev = new_element
		list.head = new_element
	} else {
		if element == nil {
			return ErrInvalidInsert
		}
		new_element.prev = element.prev
		new_element.next = element
		element.prev.next = new_element
		element.prev = new_element
	}

	list.size++
	return nil
}

func (list *CircularDoublyLinkedList) Rem(element *DoubleListElem) (interface{}, error) {
	switch {
	case element == nil:
		return nil, ErrNoElementToRemove
	case list.Size() == 0:
		return nil, ErrEmptyList
	}
	data := list.Data(element)

	if list.Next(element) == element {
		list.head = nil
	} else {
		list.Next(element).prev = list.Prev(element)
		list.Prev(element).next = list.Next(element)
		if element == list.Head() {
			list.head = list.Next(element)
		}
	}

	list.size--
	return data, nil
}

func (list CircularDoublyLinkedList) Size() int {
	return list.size
}

func (list CircularDoublyLinkedList) Head() *DoubleListElem {
	return list.head
}

func (list CircularDoublyLinkedList) Data(element *DoubleListElem) interface{} {
	return element.data
}

func (list CircularDoublyLinkedList) Next(element *DoubleListElem) *DoubleListElem {
	return element.next
}

func (list CircularDoublyLinkedList) Prev(element *DoubleListElem) *DoubleListElem {
	return element.prev
}
