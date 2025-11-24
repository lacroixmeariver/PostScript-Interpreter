package main

import (
	"testing"
)

func TestOpStrGet(t *testing.T) {
	i := CreateInterpreter()
	testIndex := 0
	testString := "hello"
	i.opStack.Push(testString)
	i.opStack.Push(testIndex)
	opGet(i)
	expected := int(testString[testIndex])
	checkStackTop(t, i, expected)
}

func TestOpStrGetInterval(t *testing.T) {
	i := CreateInterpreter()
	testIndex := 1
	testString := "hello"
	testCount := 3
	i.opStack.Push(testString)
	i.opStack.Push(testIndex)
	i.opStack.Push(testCount)
	opGetInterval(i)
	expected := string("ell")
	checkStackTop(t, i, expected)
}

func TestOpPutInterval(t *testing.T) {
	i := CreateInterpreter()
	testStr1 := "hello"
	testStr2 := "MOO"
	testCount := 1
	i.opStack.Push(testStr1)
	i.opStack.Push(testCount)
	i.opStack.Push(testStr2)
	opPutInterval(i)
	expected := string("hMOOo")
	checkStackTop(t, i, expected)
}