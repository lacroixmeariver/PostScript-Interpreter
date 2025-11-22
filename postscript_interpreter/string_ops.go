package main

// ======================================== String operators

// opLength returns the length of a string (or dictionary)
// Note: This will need to handle both strings and dictionaries
func opLength(i *Interpreter) error {
	val, _ := i.opStack.Pop()
	str := val.(string)
	result := len(str)
	i.opStack.Push(result)
	return nil
}

// TODO: Add these string operators:
// - opGet: get character at index
// - opGetInterval: get substring
// - opPutInterval: replace substring