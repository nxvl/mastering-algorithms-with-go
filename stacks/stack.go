package stacks

// Stack item struct, contains an interface to the data and a pointer
// to the next element in stack
type StackItem struct {
	data interface{}
	next *StackItem
}

// Stack struct, contains a pointer to the stack item at the top
// of the stack
type Stack struct {
	top  *StackItem
	size int
}

// Pushes the data passed into the stack, optionally returns an error
func (stack *Stack) Push(data interface{}) error {
	new_elem := &StackItem{
		data: data,
		next: stack.top,
	}
	stack.top = new_elem
	stack.size++
	return nil
}

// Pops the element at the top of the list, returns the data at the top
// of the stack and an optional error. Removes the element from the top.
func (stack *Stack) Pop() (interface{}, error) {
	if stack.Size() == 0 {
		return nil, ErrEmptyStack
	}
	data := stack.top.data
	stack.top = stack.top.next
	stack.size--
	return data, nil
}

// Peeks the element at the top of the list, returns the data at the top
// of the stack and an optional error. Does not remove the element at
// the top
func (stack Stack) Peek() (interface{}, error) {
	if stack.top == nil {
		return nil, ErrEmptyStack
	} else {
		return stack.top.data, nil
	}
}

// Returns the number of elements in the stack
func (stack Stack) Size() int {
	return stack.size
}
