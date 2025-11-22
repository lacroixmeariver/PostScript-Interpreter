package main

import (
	"fmt"
)

// ================================ Boolean operators

// opAnd performs logical AND
func opAnd(i *Interpreter) error {
	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	boolX, okX := x.(bool)
	boolY, okY := y.(bool)

	if !okX || !okY {
		return fmt.Errorf("type mismatch, [and] requires two boolean values")
	}

	result := boolX && boolY

	i.opStack.Push(result)

	return nil
}

// opOr performs logical OR
func opOr(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow")
	}
	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	boolX, okX := x.(bool)
	boolY, okY := y.(bool)

	if !okX || !okY {
		return fmt.Errorf("type mismatch, [or] requires two boolean values")
	}

	result := boolX || boolY

	i.opStack.Push(result)

	return nil
}

// opNot performs logical NOT
func opNot(i *Interpreter) error {
	if i.opStack.StackCount() < 1 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}
	val, _ := i.opStack.Pop()

	boolVal, okVal := val.(bool)

	if !okVal {
		return fmt.Errorf("type mismatch, [not] requires a boolean value")
	}
	i.opStack.Push(!boolVal)

	return nil
}

// opTrue pushes true onto stack
func opTrue(i *Interpreter) error {
	i.opStack.Push(true)
	return nil
}

// opFalse pushes false onto stack
func opFalse(i *Interpreter) error {
	i.opStack.Push(false)
	return nil
}