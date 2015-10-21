package queues

import "errors"

type QueueError error

var (
	ErrEmptyQueue = QueueError(errors.New("Invalid operation: Queue is empty"))
)
