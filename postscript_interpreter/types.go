package main

type PSConstant interface{} // represents any PS value

// for literal names like \x
type PSName string

// for code blocks
type PSBlock struct {
	Body      []Token
	DictStack []*PSDict
}

// defining the dictionary
type PSDict struct {
	items    map[string]PSConstant
	capacity int
}

type PSOperator func(*Interpreter) error
