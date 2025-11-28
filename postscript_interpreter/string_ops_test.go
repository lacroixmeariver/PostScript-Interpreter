package main

import (
	"testing"
)

/*
 -----------------------------------------------------------------------------
	Note: Parts of these tests were drafted with the use of Generative AI.
	All test content and logic has been reviewed and verified manually.
 -----------------------------------------------------------------------------
*/

func TestOpStrGet(t *testing.T) {
	i := CreateInterpreter()
	testIndex := 0
	testString := "hello"
	i.opStack.Push(testString)
	i.opStack.Push(testIndex)
	opGet(i)
	expected := int(testString[testIndex])
	compareStackTop(t, i, expected)
}

func TestOpStrGetLastChar(t *testing.T) {
	// Test getting last character
	i := CreateInterpreter()
	testString := "hello"
	i.opStack.Push(testString)
	i.opStack.Push(4) // last index
	opGet(i)
	expected := int('o') // 111
	compareStackTop(t, i, expected)
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
	compareStackTop(t, i, expected)
}

func TestOpStrGetIntervalFullString(t *testing.T) {
	// Test getting entire string
	i := CreateInterpreter()
	testString := "world"
	i.opStack.Push(testString)
	i.opStack.Push(0) // start at beginning
	i.opStack.Push(5) // get all 5 chars
	opGetInterval(i)
	compareStackTop(t, i, "world")
}

func TestOpStrGetIntervalSingleChar(t *testing.T) {
	// Test getting single character substring
	i := CreateInterpreter()
	i.opStack.Push("hello")
	i.opStack.Push(2) // index 2
	i.opStack.Push(1) // count 1
	opGetInterval(i)
	compareStackTop(t, i, "l")
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
	compareStackTop(t, i, expected)
}

func TestOpPutIntervalAtStart(t *testing.T) {
	// Test replacing at start of string
	i := CreateInterpreter()
	i.opStack.Push("hello")
	i.opStack.Push(0) // start at beginning
	i.opStack.Push("XY")
	opPutInterval(i)
	compareStackTop(t, i, "XYllo")
}

func TestOpPutIntervalAtEnd(t *testing.T) {
	// Test replacing at end of string
	i := CreateInterpreter()
	i.opStack.Push("hello")
	i.opStack.Push(3) // index 3
	i.opStack.Push("AB")
	opPutInterval(i)
	compareStackTop(t, i, "helAB")
}

func TestOpPutIntervalSingleChar(t *testing.T) {
	// Test replacing single character
	i := CreateInterpreter()
	i.opStack.Push("hello")
	i.opStack.Push(2) // index 2
	i.opStack.Push("X")
	opPutInterval(i)
	compareStackTop(t, i, "heXlo")
}

func TestOpPutIntervalLongerReplacement(t *testing.T) {
	// Test when replacement is longer than remaining string
	i := CreateInterpreter()
	i.opStack.Push("hi")
	i.opStack.Push(1) // index 1
	i.opStack.Push("WORLD")
	opPutInterval(i)
	compareStackTop(t, i, "hWORLD")
}
