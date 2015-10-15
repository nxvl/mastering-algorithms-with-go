package linked_list_errors

import "errors"

type LinkedListError error

var (
	ErrEmptyList         = LinkedListError(errors.New("Invalid operation: list is empty"))
	ErrInvalidInsert     = LinkedListError(errors.New("Invalid insert parameters"))
	ErrNoElementToRemove = LinkedListError(errors.New("Invalid operation: no element to remove"))
)
