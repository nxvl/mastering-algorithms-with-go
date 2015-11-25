package sets

import "errors"

type StackError error

var (
	ErrNothingToRemove = StackError(errors.New("Invalid operation: data not found in set"))
)
