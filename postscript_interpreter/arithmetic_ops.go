package main

import (
	"fmt"
	"math"
)

// ========================================= Arithmetic operators

// opAdd adds 2 operands
func opAdd(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}

	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	xToNum, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	yToNum, err := ToNumber(y)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	result := xToNum + yToNum
	i.opStack.Push(result)

	return nil
}

// opSub subtracts one operand from the other
func opSub(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}

	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	xToNum, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	yToNum, err := ToNumber(y)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	result := xToNum - yToNum
	i.opStack.Push(result)

	return nil
}

// opMul multiplies one operand by the other
func opMul(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}

	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	xToNum, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	yToNum, err := ToNumber(y)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	result := xToNum * yToNum
	i.opStack.Push(result)

	return nil
}

// opDiv divides one operand by the other
func opDiv(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}

	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	xToNum, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	yToNum, err := ToNumber(y)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	if yToNum == 0 {
		return fmt.Errorf("divide by zero error")
	}

	result := xToNum / yToNum
	i.opStack.Push(result)

	return nil
}

// integer division operators 

// opIntdiv performs integer division
func opIntdiv(i *Interpreter) error {

	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}

	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	xToNum, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	yToNum, err := ToNumber(y)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	if yToNum == 0 {
		return fmt.Errorf("divide by zero error")
	}

	result := xToNum / yToNum
	i.opStack.Push(int(result))

	return nil
}

// opMod performs modulo operation
func opMod(i *Interpreter) error {
	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}

	y, _ := i.opStack.Pop()
	x, _ := i.opStack.Pop()

	xToNum, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	yToNum, err := ToNumber(y)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	if yToNum == 0 {
		return fmt.Errorf("divide by zero error")
	}

	result := int(xToNum) % int(yToNum)
	i.opStack.Push(result)

	return nil
}

// opSqrt calculates square root
func opSqrt(i *Interpreter) error {
	if i.opStack.StackCount() < 1 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}
	x, _ := i.opStack.Pop()

	xToNum, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	result := math.Sqrt(xToNum)

	i.opStack.Push(result)
	return nil
}

// opCeil returns ceiling of number
func opCeil(i *Interpreter) error {
	if i.opStack.StackCount() < 1 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}

	x, _ := i.opStack.Pop()

	xToNum, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	result := math.Ceil(xToNum)

	i.opStack.Push(result)
	return nil
}

// opFloor returns floor of number
func opFloor(i *Interpreter) error {
	if i.opStack.StackCount() < 1 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}

	x, _ := i.opStack.Pop()

	xToNum, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	result := math.Floor(xToNum)

	i.opStack.Push(result)
	return nil
}

// opRound rounds to nearest integer
func opRound(i *Interpreter) error {
	if i.opStack.StackCount() < 1 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}

	x, _ := i.opStack.Pop()

	xToNum, err := ToNumber(x)
	if err != nil {
		return fmt.Errorf("operand error")
	}

	result := math.Round(xToNum)

	i.opStack.Push(result)
	return nil
}

// opAbs takes absolute value of given number
func opAbs(i *Interpreter) error {
	if i.opStack.StackCount() < 1 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
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

// opNeg negates a number
func opNeg(i *Interpreter) error {
	if i.opStack.StackCount() < 1 {
		return fmt.Errorf("stack underflow, not enough elements in stack")
	}

	val, _ := i.opStack.Pop()
	num, err := ToNumber(val)
	if err != nil {
		return err
	}

	i.opStack.Push(-num)
	return nil
}
