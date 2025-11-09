package main

type PSConstant interface{} // represents any PS value

type PSName string

type PSCodeblock struct {
	Body []string
	DictStack []PSDict
}

// defining the dictionary 
type PSDict map[string]PSConstant