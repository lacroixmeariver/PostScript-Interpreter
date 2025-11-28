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

func TestOpLengthString(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push("hello")
	
	err := opLength(testInterpreter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	compareStackTop(t, testInterpreter, 5)
}

func TestOpStrGet(t *testing.T) {
	i := CreateInterpreter()

	testIndex := 0
	testString := "hello"

	i.opStack.Push(testString)
	i.opStack.Push(testIndex)

	// getting char value at index
	err := opGet(i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	expected := int(testString[testIndex])
	compareStackTop(t, i, expected)
}

func TestOpStrGetLastChar(t *testing.T) {
	// test getting last character
	i := CreateInterpreter()

	testString := "hello"
	i.opStack.Push(testString)

	i.opStack.Push(4) // last index
	err := opGet(i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

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

	// creating substring
	err := opGetInterval(i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := string("ell")
	compareStackTop(t, i, expected)
}

func TestOpStrGetIntervalFullString(t *testing.T) {
	// test getting entire string
	i := CreateInterpreter()

	testString := "world"
	i.opStack.Push(testString)

	i.opStack.Push(0) // start at beginning
	i.opStack.Push(5) // get all 5 chars

	err := opGetInterval(i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	compareStackTop(t, i, "world")
}

func TestOpStrGetIntervalSingleChar(t *testing.T) {
	// test getting single character substring
	i := CreateInterpreter()

	i.opStack.Push("hello")
	i.opStack.Push(2) // index 2
	i.opStack.Push(1) // count 1

	err := opGetInterval(i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
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

	err := opPutInterval(i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	expected := string("hMOOo")
	compareStackTop(t, i, expected)
}

func TestOpPutIntervalAtStart(t *testing.T) {
	// test replacing at start of string
	i := CreateInterpreter()

	i.opStack.Push("hello")
	i.opStack.Push(0) // start at beginning
	i.opStack.Push("XY")

	err := opPutInterval(i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	compareStackTop(t, i, "XYllo")
}


func TestOpPutIntervalAtEnd(t *testing.T) {
	// test replacing at end of string
	i := CreateInterpreter()

	i.opStack.Push("hello")
	i.opStack.Push(3) // index 3
	i.opStack.Push("AB")

	err := opPutInterval(i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	compareStackTop(t, i, "helAB")
}

func TestOpPutIntervalSingleChar(t *testing.T) {
	// test replacing single character
	i := CreateInterpreter()

	i.opStack.Push("hello")
	i.opStack.Push(2) // index 2
	i.opStack.Push("X")

	err := opPutInterval(i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	compareStackTop(t, i, "heXlo")
}

func TestOpPutIntervalLongerReplacement(t *testing.T) {
	// test when replacement is longer than remaining string
	i := CreateInterpreter()
	
	i.opStack.Push("hi")
	i.opStack.Push(1) // index 1
	i.opStack.Push("WORLD")

	err := opPutInterval(i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	compareStackTop(t, i, "hWORLD")
}
