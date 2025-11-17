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
type PSDict map[string]PSConstant

//type PSOperator func (*Interpreter) error
