package linked_list

import "errors"

type LinkedListError error

var (
	ErrEmptyList         = LinkedListError(errors.New("Invalid operation: list is empty"))
	ErrInvalidInsert     = LinkedListError(errors.New("Invalid insert parameters"))
	ErrNoElementToRemove = LinkedListError(errors.New("Invalid operation: no element to remove"))
)
