package sets

import "errors"

type SetError error

var (
	ErrNothingToRemove = SetError(errors.New("Invalid operation: data not found in set"))
)
