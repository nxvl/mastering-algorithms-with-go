package stacks

import "github.com/nxvl/mastering-algorithms-with-go/linked_lists"

// Stack item struct, contains an interface to the data and a pointer
// to the next element in stack
type StackItem struct {
	data interface{}
	next *StackItem
}

// Stack struct, contains a pointer to the stack item at the top
// of the stack
type Stack struct {
	linked_list.SimpleLinkedList
}

// Pushes the data passed into the stack, optionally returns an error
func (stack *Stack) Push(data interface{}) error {
	return stack.InsNext(nil, data)
}

// Pops the element at the top of the list, returns the data at the top
// of the stack and an optional error. Removes the element from the top.
func (stack *Stack) Pop() (interface{}, error) {
	data, err := stack.RemNext(nil)
	switch {
	case err == nil:
		return data, nil
	case err == linked_list.ErrEmptyList:
		return data, ErrEmptyStack
	default:
		return data, err
	}
}

// Peeks the element at the top of the list, returns the data at the top
// of the stack and an optional error. Does not remove the element at
// the top
func (stack Stack) Peek() (interface{}, error) {
	if stack.Head() == nil {
		return nil, ErrEmptyStack
	} else {
		return stack.Data(stack.Head()), nil
	}
}
