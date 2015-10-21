package queues

// Queue item struct, contains an interface to the data and a pointer
// to the next element in queue
type QueueItem struct {
	data interface{}
	next *QueueItem
}

// Queue struct, contains pointers to the elements at the front and back
// of the queue
type Queue struct {
	size        int
	front, back *QueueItem
}

// Enqueues the data passed into the queue, optionally returns an error
func (q *Queue) Enqueue(data interface{}) error {
	new_item := &QueueItem{
		data: data,
	}
	if q.size == 0 {
		q.front = new_item
	} else {
		q.back.next = new_item
	}
	q.back = new_item
	q.size++
	return nil
}

// Dequeues and item from the front of the queue, returns the data at the
// front and an optional error. Removes the element from the front of the
// queue
func (q *Queue) Dequeue() (interface{}, error) {
	switch q.size {
	case 0:
		return nil, ErrEmptyQueue
	case 1:
		q.back = nil
	}
	data := q.front.data
	q.front = q.front.next

	q.size--
	return data, nil
}

// Peeks the element in the front of the queue. Returns the data at the front
// and an optional error. Does not alter the queue in any way.
func (q Queue) Peek() (interface{}, error) {
	if q.size == 0 {
		return nil, ErrEmptyQueue
	} else {
		return q.front.data, nil
	}
}

func (q Queue) Size() int {
	return q.size
}
