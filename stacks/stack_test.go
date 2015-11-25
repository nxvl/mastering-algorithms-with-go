package stacks

import "testing"

func newStack(t *testing.T) Stack {
	var stack = Stack{}
	switch {
	case stack.Size() != 0:
		t.Fatalf("Initial stack size is: %d", stack.Size())
	case stack.Head() != nil:
		t.Fatal("Top of the stack is not empty")
	}

	return stack
}

func newFilledStack(t *testing.T) Stack {
	var stack = newStack(t)
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	switch {
	case stack.Size() != 5:
		t.Fatalf("Initial filled stack size is: %d, expected 5", stack.Size())
	case stack.Data(stack.Head()) != 4:
		t.Fatalf("Wrong top of the stack value: %d", stack.Data(stack.Head()))
	case stack.Data(stack.Next(stack.Head())) != 3:
		t.Fatalf("Wrong 2nd value: %d", stack.Data(stack.Next(stack.Head())))
	case stack.Data(stack.Next(stack.Next(stack.Head()))) != 2:
		t.Fatalf("Wrong 3rd value: %d", stack.Data(stack.Next(stack.Next(stack.Head()))))
	case stack.Data(stack.Next(stack.Next(stack.Next(stack.Head())))) != 1:
		t.Fatalf("Wrong 4th value: %d", stack.Data(stack.Next(stack.Next(stack.Next(stack.Head())))))
	case stack.Data(stack.Next(stack.Next(stack.Next(stack.Next(stack.Head()))))) != 0:
		t.Fatalf("Wrong 5th value: %d", stack.Data(stack.Next(stack.Next(stack.Next(stack.Next(stack.Head()))))))
	}

	return stack
}

func TestStackPush(t *testing.T) {
	var stack = newStack(t)

	stack.Push(15)
	switch {
	case stack.Size() != 1:
		t.Errorf("Stack size is %d, expected 1", stack.Size())
	case stack.Data(stack.Head()) != 15:
		t.Errorf("Wrong top of the stack value: %d", stack.Data(stack.Head()))
	}

	stack.Push(10)
	switch {
	case stack.Size() != 2:
		t.Errorf("Stack size is %d, expected 1", stack.Size())
	case stack.Data(stack.Head()) != 10:
		t.Errorf("Wrong top of the stack value: %d", stack.Data(stack.Head()))
	case stack.Data(stack.Next(stack.Head())) != 15:
		t.Errorf("Wrong next of the top value: %d", stack.Data(stack.Head()))
	}
}

func TestFilledStackPop(t *testing.T) {
	var stack = newFilledStack(t)

	data, err := stack.Pop()
	switch {
	case err != nil:
		t.Error("Error has been returned")
	case stack.Size() != 4:
		t.Errorf("Stack size is %d, expected 4", stack.Size())
	case data != 4:
		t.Errorf("Wrong value returned: %d, expected 4", data)
	case stack.Data(stack.Head()) != 3:
		t.Errorf("Wrong new top value: %d, expected 3", stack.Data(stack.Head()))
	}
}

func TestEmptyStackPop(t *testing.T) {
	var stack = newStack(t)

	data, err := stack.Pop()
	switch {
	case err == nil:
		t.Error("No error returned from poping from empty list")
	case err != ErrEmptyStack:
		t.Error("Wrong error returned from poping from empty list")
	case data != nil:
		t.Error("Data is not nil")
	}
}

func TestFilledStackPeek(t *testing.T) {
	var stack = newFilledStack(t)

	data, err := stack.Peek()
	switch {
	case err != nil:
		t.Error("Error has been returned")
	case stack.Size() != 5:
		t.Errorf("Stack size is %d, expected 5", stack.Size())
	case data != 4:
		t.Errorf("Wrong value returned: %d, expected 4", data)
	case stack.Data(stack.Head()) != 4:
		t.Errorf("Wrong new top value: %d, expected 4", stack.Data(stack.Head()))
	}

}

func TestEmptyStackPeek(t *testing.T) {
	var stack = newStack(t)

	data, err := stack.Peek()
	switch {
	case err != ErrEmptyStack:
		t.Error("No error returned from poping from empty list")
	case data != nil:
		t.Error("Data is not nil")
	}
}
