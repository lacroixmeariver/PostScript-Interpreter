package main

import "fmt"

// ======================================== I/O operators

// TODO: Implement these I/O operators:
// - opEquals: write text representation to stdout (=)
// - opEqualsEquals: write PostScript representation to stdout (==)

func opPrint(i *Interpreter) error {
	v, _ := i.opStack.Pop()

	fmt.Print(v)
	return nil
}
