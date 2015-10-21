package queues

import "testing"

func newQueue(t *testing.T) Queue {
	var q = Queue{}
	switch {
	case q.size != 0:
		t.Fatalf("Bad empty Queue size: %d", q.size)
	case q.front != nil || q.back != nil:
		t.Fatal("Head or Tail contain values in an empty Queue")
	}

	return q
}

func newFilledQueue(t *testing.T) Queue {
	var q = Queue{}
	for i := 0; i < 3; i++ {
		q.Enqueue(i)
	}
	switch {
	case q.size != 3:
		t.Fatalf("Bad empty Queue size: %d", q.size)
	case q.front.data != 0:
		t.Fatal("Head contains wrong value")
	case q.back.data != 2:
		t.Fatal("Tail contains wrong value")
	}

	return q
}

func TestQueueEnqueue(t *testing.T) {
	var q = newQueue(t)

	q.Enqueue(5)
	switch {
	case q.Size() != 1:
		t.Errorf("Wrong queue size: %d, expecting 1", q.Size())
	case q.front.data != 5:
		t.Errorf("Wrong front data value: %d, expecting 5", q.front.data)
	case q.back.data != 5:
		t.Errorf("Wrong back data value: %d, expecting 5", q.back.data)
	}

	q.Enqueue(10)
	switch {
	case q.Size() != 2:
		t.Errorf("Wrong queue size: %d, expecting 2", q.Size())
	case q.front.data != 5:
		t.Errorf("Wrong front data value: %d, expecting 5", q.front.data)
	case q.back.data != 10:
		t.Errorf("Wrong back data value: %d, expecting 10", q.back.data)
	}

	q.Enqueue(15)
	switch {
	case q.Size() != 3:
		t.Errorf("Wrong queue size: %d, expecting 3", q.Size())
	case q.front.data != 5:
		t.Errorf("Wrong front data value: %d, expecting 5", q.front.data)
	case q.front.next.data != 10:
		t.Errorf("Wrong front next data value: %d, expecting 10", q.front.data)
	case q.back.data != 15:
		t.Errorf("Wrong back data value: %d, expecting 15", q.back.data)
	}
}

func TestFilledQueueDequeue(t *testing.T) {
	var q = newFilledQueue(t)

	data, err := q.Dequeue()
	switch {
	case err != nil:
		t.Error("Error was returned")
	case q.Size() != 2:
		t.Errorf("Wrong queue size: %d, expecting 2", q.Size())
	case data != 0:
		t.Errorf("Wrong value returned: %d, expecting 0", data)
	case q.front.data != 1:
		t.Errorf("Wrong queue front: %d, expected 1", q.front.data)
	case q.back.data != 2:
		t.Errorf("Wrong queue back: %d, expected 2", q.front.data)
	}

	data, err = q.Dequeue()
	switch {
	case err != nil:
		t.Error("Error was returned")
	case q.Size() != 1:
		t.Errorf("Wrong queue size: %d, expecting 1", q.Size())
	case data != 1:
		t.Errorf("Wrong value returned: %d, expecting 1", data)
	case q.front.data != 2:
		t.Errorf("Wrong queue front: %d, expected 2", q.front.data)
	case q.back.data != 2:
		t.Errorf("Wrong queue back: %d, expected 2", q.front.data)
	case q.back != q.front:
		t.Error("Tail and front are not equal in a size 1 queue")
	}

	data, err = q.Dequeue()
	switch {
	case err != nil:
		t.Error("Error was returned")
	case q.Size() != 0:
		t.Errorf("Wrong queue size: %d, expecting 0", q.Size())
	case data != 2:
		t.Errorf("Wrong value returned: %d, expecting 2", data)
	case q.front != nil:
		t.Error("Queue front is not nil")
	case q.back != nil:
		t.Error("Queue back is not nil")
	}
}

func TestEmptyQueueDequeue(t *testing.T) {
	var q = newQueue(t)

	data, err := q.Dequeue()
	switch {
	case data != nil:
		t.Error("Data is not empty")
	case err != ErrEmptyQueue:
		t.Error("Wrong error")
	}
}

func TestFilledQueuePeek(t *testing.T) {
	var q = newFilledQueue(t)

	data, err := q.Peek()
	switch {
	case q.Size() != 3:
		t.Errorf("Wrong queue size: %d, expecting 3", q.Size())
	case data != 0:
		t.Errorf("Wrong value returned: %d, expecting 0", data)
	case err != nil:
		t.Error("Error returned")
	}
}

func TestEmptyQueuePeek(t *testing.T) {
	var q = newQueue(t)

	data, err := q.Peek()
	switch {
	case data != nil:
		t.Error("Data returned, expected nil")
	case err != ErrEmptyQueue:
		t.Error("Wrong error returned")
	}
}
