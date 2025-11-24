package main

import "fmt"

// ================================== Dictionary operations

// dOpDict creates a PSDict  with given capacity and pushes it onto the opStack
func dOpDict(i *Interpreter) error {

	capacity, _ := i.opStack.Pop()
	cap := capacity.(int)

	dictionary := &PSDict{
		items:    make(map[string]PSConstant),
		capacity: cap, // useful for max length
	}

	i.opStack.Push(dictionary)

	return nil
}

// dOpBegin defines the start point of a new dictionary on the DictStack
func dOpBegin(i *Interpreter) error {

	val, err := i.opStack.Pop()
	if err != nil {
		return fmt.Errorf("stack underflow, not enough elements in the stack")
	}

	dict, ok := val.(*PSDict)
	if !ok {
		return fmt.Errorf("dictionary conversion failed")
	}

	dictCap := i.dictStack[len(i.dictStack)-1].capacity
	currentLength := len(i.dictStack[len(i.dictStack)-1].items)
	if currentLength >= dictCap {
		return fmt.Errorf("dictionary capacity exceeded")
	}
	i.dictStack = append(i.dictStack, dict)
	return nil
}

// dOpEnd defines the end point of a new dictionary on the DictStack
func dOpEnd(i *Interpreter) error {

	if len(i.dictStack) <= 1 {
		return fmt.Errorf("dict stack underflow")
	}

	i.dictStack = i.dictStack[:len(i.dictStack)-1]
	return nil
}

// dOpDef associates a key value pair for the current dictionary
func dOpDef(i *Interpreter) error {

	if i.opStack.StackCount() < 2 {
		return fmt.Errorf("stack underflow, not enough elements in the stack")
	}
	value, _ := i.opStack.Pop()
	k, _ := i.opStack.Pop()

	var key string
	switch val := k.(type) {
	case PSName:
		key = string(val)
	case string:
		key = val
	default:
		return fmt.Errorf("string or constant expected")
	}

	currentDict := i.dictStack[len(i.dictStack)-1]
	currentDict.items[key] = value

	return nil
}

// dOpLength pushes the length of the current dictionary onto the opStack
func dOpLength(i *Interpreter) error {

	val, _ := i.opStack.Pop()
	dict := val.(*PSDict)
	result := len(dict.items)
	i.opStack.Push(result)
	return nil
}

// dOpMaxLength pushes the capacity of the current dictionary onto the stack
func dOpMaxLength(i *Interpreter) error {

	if len(i.dictStack) < 1 {
		return fmt.Errorf("dict stack underflow, not enough elements in the stack")
	}

	val, _ := i.opStack.Pop()
	currentDict := val.(*PSDict)
	i.opStack.Push(currentDict.capacity)
	return nil
}
