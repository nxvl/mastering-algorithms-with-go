package stacks

import "errors"

type StackError error

var (
	ErrEmptyStack = StackError(errors.New("Invalid operation: stack is empty"))
)
