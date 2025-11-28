package main

// =================================== stack operations

// opDup duplicates top of stack and pushes it to top of stack
func opDup(i *Interpreter) error {
	topStack, _ := i.opStack.Peek()
	if topStack != nil {
		dup := topStack
		i.opStack.Push(dup)
	}

	return nil
}

// opPop pops the top of the stack
func opPop(i *Interpreter) error {
	if i.opStack.StackCount() > 0 {
		i.opStack.Pop()
	}

	return nil
}

// opExch swaps the 2 top-most elements of the stack
func opExch(i *Interpreter) error {
	if i.opStack.StackCount() >= 2 {
		b, _ := i.opStack.Pop()
		a, _ := i.opStack.Pop()

		i.opStack.Push(b)
		i.opStack.Push(a)
	}

	return nil
}

// opClear clears the stack
func opClear(i *Interpreter) error {
	for i.opStack.StackCount() > 0 {
		i.opStack.Pop()
	}
	
	return nil
}

// opCount pushes the number of elements in stack onto the stack
func opCount(i *Interpreter) error {
	count := i.opStack.StackCount()
	i.opStack.Push(count)

	return nil
}
