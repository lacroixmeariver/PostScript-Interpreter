package main

import (
	"fmt"
	"math"
)

// Helper functions

// Allows for interface types to be converted into numbers to perform operations on
func ToNumber(num PSConstant) (float64, error) {

	switch val := num.(type) {
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	default:
		return math.NaN(), fmt.Errorf("bad input")
	}
}

// ========================================= Arithmetic operators

// adds 2 operands
func opAdd(i *Interpreter) error {

	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("not enough items")
	}

	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	numX, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	numY, err := ToNumber(y)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	result := numX + numY
	i.opStack.Push(result)

	return nil
}

// subtracts one operand from the other
func opSub(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("not enough items")
	}

	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	numX, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	numY, err := ToNumber(y)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	result := numX - numY
	i.opStack.Push(result)

	return nil
}

// multiplies one operand from the other
func opMul(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("not enough items")
	}

	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	numX, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	numY, err := ToNumber(y)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	result := numX * numY
	i.opStack.Push(result)

	return nil
}

// divides one operand from the other
func opDiv(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("not enough items")
	}

	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	numX, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	numY, err := ToNumber(y)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	if numY == 0 {
		return fmt.Errorf("divide by zero error")
	}

	result := numX / numY
	i.opStack.Push(result)

	return nil
}

// =================================== Stack operations

// duplicates top of stack and pushes it to top of stack
func opDup(i *Interpreter) error {

	topStack, _ := i.opStack.Peek()
	if topStack != nil {
		dup := topStack
		i.opStack.Push(dup)
	}

	return nil
}

// pops the top of the stack
func opPop(i *Interpreter) error {
	if i.opStack.StackCount() > 0 {
		i.opStack.Pop()
	}

	return nil
}

// swaps the 2 top-most elements of the stack
func opExch(i *Interpreter) error {
	if i.opStack.StackCount() >= 2 {
		b, _ := i.opStack.Pop()
		a, _ := i.opStack.Pop()

		i.opStack.Push(b)
		i.opStack.Push(a)
	}

	return nil
}

// clears the stack
func opClear(i *Interpreter) error {
	for i.opStack.StackCount() > 0 {
		i.opStack.Pop()
	}
	return nil
}

// pushes the number of elements in stack onto the stack
func opCount(i *Interpreter) error {
	count := i.opStack.StackCount()
	i.opStack.Push(count)
	return nil
}

// takes top element on stack and pushes absolute value in its place
func opAbs(i *Interpreter) error {
	if i.opStack.StackCount() < 1 {
		return fmt.Errorf("not enough elements in stack")
	}

	val, _ := i.opStack.Pop()
	num, err := ToNumber(val)
	if err != nil {
		return err
	}

	if num < 0 {
		num = -num
	}

	i.opStack.Push(num)

	return nil
}

// takes element at top of stack and pushes its negative value onto the stack
func opNeg(i *Interpreter) error {
	if i.opStack.StackCount() < 1 {
		return fmt.Errorf("not enough elements in the stack")
	}

	val, _ := i.opStack.Pop()
	num, err := ToNumber(val)
	if err != nil {
		return err
	}

	i.opStack.Push(-num)
	return nil
}

// ================================ Boolean operators

// pushes true if two items are equal
func opEq(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow")
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

	result := x == y
	i.opStack.Push(result)

	return nil
}

// pushes true if two items are not equal
func opNe(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow")
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

	result := x != y
	i.opStack.Push(result)

	return nil
}

// pushes true if one item is greater than or equal to the other
func opGe(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow")
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
	return fmt.Errorf("type mismatch")
}

// pushes true if one item is greater than the other
func opGt(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow")
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
	return fmt.Errorf("type mismatch")
}

// pushes true if one item is less than or equal to the other
func opLe(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow")
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
	return fmt.Errorf("type mismatch")
}

// pushes true if one item is less than the other
func opLt(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow")
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
	return fmt.Errorf("type mismatch")
}

func opAnd(i *Interpreter) error {
	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	boolX, okX := x.(bool)
	boolY, okY := y.(bool)

	if !okX || !okY {
		return fmt.Errorf("requires two boolean values")
	}

	result := boolX && boolY

	i.opStack.Push(result)

	return nil
}

func opOr(i *Interpreter) error {
	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	boolX, _ := x.(bool)
	boolY, _ := y.(bool)

	result := boolX || boolY

	i.opStack.Push(result)

	return nil
}

func opNot(i *Interpreter) error {
	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	boolX, okX := x.(bool)
	boolY, okY := y.(bool)

	if okX || okY {
		return fmt.Errorf("requires two boolean values")
	}

	result := !boolX && !boolY

	i.opStack.Push(result)

	return nil
}

func opTrue(i *Interpreter) error {
	x, _ := i.opStack.Pop()

	boolX := x.(bool)
	if boolX {
		i.opStack.Push(PSConstant(true))
	}
	return nil
}

func opFalse(i *Interpreter) error {
	x, _ := i.opStack.Pop()

	boolX := x.(bool)
	if !boolX {
		i.opStack.Push(PSConstant(false))
	}
	return nil
}

// ======================================== String manipulation

func opLength(i *Interpreter) error {
	val, _ := i.opStack.Pop()
	str := val.(string)
	result := len(str)
	i.opStack.Push(result)
	return nil
}
