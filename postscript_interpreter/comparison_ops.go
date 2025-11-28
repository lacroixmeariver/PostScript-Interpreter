package main

import (
	"fmt"
)

// ================================ comparison operations

// opEq pushes true if two items are equal
func opEq(i *Interpreter) error {
	// error handling
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in the stack")
	}

	x, _ := i.opStack.Pop()
	y, _ := i.opStack.Pop()

	numX, errX := ToNumber(x)
	numY, errY := ToNumber(y)

	if errX == nil && errY == nil {
		result := numX == numY
		i.opStack.Push(result)
		return nil
	}

	// trying as strings
	strA, errA := x.(string)
	strB, errB := y.(string)

	if errA && errB {
		result := strA == strB
		i.opStack.Push(result)
		return nil
	}

	// accounting for type mismatch
	return fmt.Errorf("type mismatch, please ensure both elements are the same type")
}

// opNe pushes true if two items are not equal
func opNe(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in the stack")
	}

	x, _ := i.opStack.Pop()
	y, _ := i.opStack.Pop()

	numX, errX := ToNumber(x)
	numY, errY := ToNumber(y)

	if errX == nil && errY == nil {
		result := numX != numY
		i.opStack.Push(result)
		return nil
	}

	// trying as strings
	strA, errA := x.(string)
	strB, errB := y.(string)

	if errA && errB {
		result := strA != strB
		i.opStack.Push(result)
		return nil
	}

	// accounting for type mismatch
	return fmt.Errorf("type mismatch, please ensure both elements are the same type")


}

// opGe pushes true if one item is greater than or equal to the other
func opGe(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in the stack")
	}

	// trying as numbers first
	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	numX, errX := ToNumber(x)
	numY, errY := ToNumber(y)

	if errX == nil && errY == nil {
		result := numX >= numY
		i.opStack.Push(result)
		return nil
	}

	// trying as strings
	strA, errA := x.(string)
	strB, errB := y.(string)

	if errA && errB {
		result := strA >= strB
		i.opStack.Push(result)
		return nil
	}

	// accounting for type mismatch
	return fmt.Errorf("type mismatch, please ensure both elements are the same type")
}

// opGt pushes true if one item is greater than the other
func opGt(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in the stack")
	}

	// trying as numbers first
	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	numX, errX := ToNumber(x)
	numY, errY := ToNumber(y)

	if errX == nil && errY == nil {
		result := numX > numY
		i.opStack.Push(result)
		return nil
	}

	// trying as strings
	strA, okA := x.(string)
	strB, okB := y.(string)

	if okA && okB {
		result := strA > strB
		i.opStack.Push(result)
		return nil
	}

	// accounting for type mismatch
	return fmt.Errorf("type mismatch, please ensure both elements are the same type")
}

// opLe pushes true if one item is less than or equal to the other
func opLe(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in the stack")
	}

	// trying as numbers first
	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	numX, errX := ToNumber(x)
	numY, errY := ToNumber(y)

	if errX == nil && errY == nil {
		result := numX <= numY
		i.opStack.Push(result)
		return nil
	}

	// trying as strings
	strA, okA := x.(string)
	strB, okB := y.(string)

	if okA && okB {
		result := strA <= strB
		i.opStack.Push(result)
		return nil
	}

	// accounting for type mismatch
	return fmt.Errorf("type mismatch, please ensure both elements are the same type")
}

// opLt pushes true if one item is less than the other
func opLt(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in the stack")
	}

	// trying as numbers first
	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	numX, errX := ToNumber(x)
	numY, errY := ToNumber(y)

	if errX == nil && errY == nil {
		result := numX < numY
		i.opStack.Push(result)
		return nil
	}

	// trying as strings
	strA, okA := x.(string)
	strB, okB := y.(string)

	if okA && okB {
		result := strA < strB
		i.opStack.Push(result)
		return nil
	}

	// accounting for type mismatch
	return fmt.Errorf("type mismatch, please ensure both elements are the same type")
}
