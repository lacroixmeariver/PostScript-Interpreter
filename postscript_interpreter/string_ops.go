package main

import "fmt"

// ======================================== string operations

// opLength returns the length of a string/dictionary
func opLength(i *Interpreter) error {
	val, _ := i.opStack.Pop()

	// try as string
	if str, ok := val.(string); ok {
		result := len(str)
		i.opStack.Push(result)
		return nil
	}

	// try as dict
	if dict, ok := val.(*PSDict); ok {
		result := len(dict.items)
		i.opStack.Push(result)
		return nil
	}

	return fmt.Errorf("length requires string or dictionary, got: %T", val)
}

// opGet gets returns the ASCII value of the character at an index
func opGet(i *Interpreter) error {

	indexVal, _ := i.opStack.Pop() // desired index
	strVal, _ := i.opStack.Pop() // string to be indexed

	// converting to usable types
	index := indexVal.(int)
	
	str := strVal.(string)

	if index >= len(str) || index < 0 {
		return fmt.Errorf("out of bounds index")
	}
	result := str[index]
	i.opStack.Push(int(result))

	return nil 
}

// opGetInterval returns substring of given string from index to index + count
func opGetInterval(i *Interpreter) error {

	countVal, _ := i.opStack.Pop() // count
	indexVal, _ := i.opStack.Pop() // starting index
	strVal, _ := i.opStack.Pop() // string 

	// conversions
	count := countVal.(int)
	index := indexVal.(int)
	str := strVal.(string)


	if index >= len(str) || index < 0 {
		return fmt.Errorf("out of bounds index")
	}
	if count < 0 {
		return fmt.Errorf("count cannot be negative number")
	}
	if index + count > len(str) {
		return fmt.Errorf("substring goes beyond original string length")
	}
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
	// addressing the issue of str1 < str2
	if len(str2) > len(str1) {
		remainder := (index + len(str2)) - len(runes)
		j := len(str2) - remainder
		for j < len(str2){
			runes = append(runes, rune(str2[j]))
			j++
		}
	}
	result := string(runes)

	i.opStack.Push(string(result))
	return nil 
}
