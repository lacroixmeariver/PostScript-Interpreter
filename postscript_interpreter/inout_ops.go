package main

import "fmt"

// ======================================== input/output operators

// writes characters of string to stdout
func opPrint(i *Interpreter) error {
	v, _ := i.opStack.Pop()
	str := v.(string)
	fmt.Print(str)

	return nil
}

// writes text representation of any to stdout
func opEquals(i *Interpreter) error {
	v, _ := i.opStack.Pop()
	fmt.Println(v)
	
	return nil
}

// destructive display of top of stack
func opEqualsEquals(i *Interpreter) error {
	v, _ := i.opStack.Pop()
	if str, ok := v.(string); ok {
		fmt.Printf("(%s)\n", str)
	} else {
		fmt.Println(v)
	}
	
	return nil
}