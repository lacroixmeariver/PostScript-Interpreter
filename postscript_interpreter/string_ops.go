package main

// ======================================== String operators

// opLength returns the length of a string (or dictionary)
// TODO: Modify to handle both strings and dictionaries
func opLength(i *Interpreter) error {
	val, _ := i.opStack.Pop()
	str := val.(string)
	result := len(str)
	i.opStack.Push(result)
	return nil
}

// opGet gets returns the ASCII value of the character at an index
// TODO: Error handling for out of index bounds
func opGet(i *Interpreter) error {

	indexVal, _ := i.opStack.Pop() // desired index
	strVal, _ := i.opStack.Pop() // string to be indexed

	// converting to usable types
	index := indexVal.(int)
	str := strVal.(string)

	result := str[index]
	i.opStack.Push(int(result))

	return nil 
}

// opGetInterval returns substring of given string from index to index + count
// TODO: Error handling for out of index bounds
func opGetInterval(i *Interpreter) error {

	countVal, _ := i.opStack.Pop() // count
	indexVal, _ := i.opStack.Pop() // starting index
	strVal, _ := i.opStack.Pop() // string 

	// conversions
	count := countVal.(int)
	index := indexVal.(int)
	str := strVal.(string)

	// slice for result
	result := str[index:(count + index)]

	i.opStack.Push(string(result))
	return nil 
}

// opPutInterval replaces string with substring up to given index of a given string
func opPutInterval(i *Interpreter) error {

	s2, _ := i.opStack.Pop()
	ind, _ := i.opStack.Pop()
	s1, _ := i.opStack.Pop()

	index := ind.(int)
	str1 := s1.(string)
	str2 := s2.(string)

	runes := []rune(str1)

	pos := 0
	
	for pos < len(str2) && index + pos < len(runes) {
		runes[index + pos] = rune(str2[pos])
		pos ++
	}

	result := string(runes)

	i.opStack.Push(string(result))
	return nil 
}
