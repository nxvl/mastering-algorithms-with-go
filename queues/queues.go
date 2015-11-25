package queues

import "github.com/nxvl/mastering-algorithms-with-go/linked_lists"

// Queue struct, contains pointers to the elements at the front and back
// of the queue
type Queue struct {
	linked_list.SimpleLinkedList
}

// Enqueues the data passed into the queue, optionally returns an error
func (q *Queue) Enqueue(data interface{}) error {
	return q.InsNext(q.Tail(), data)
}

// Dequeues and item from the front of the queue, returns the data at the
// front and an optional error. Removes the element from the front of the
// queue
func (q *Queue) Dequeue() (interface{}, error) {
	data, err := q.RemNext(nil)
	if err == linked_list.ErrEmptyList {
		return data, ErrEmptyQueue
	}
	return data, err
}

// Peeks the element in the front of the queue. Returns the data at the front
// and an optional error. Does not alter the queue in any way.
func (q Queue) Peek() (interface{}, error) {
	if q.Size() == 0 {
		return nil, ErrEmptyQueue
	} else {
		return q.Data(q.Head()), nil
	}
}
