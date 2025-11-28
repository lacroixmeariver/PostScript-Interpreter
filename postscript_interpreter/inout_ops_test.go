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


func TestOpPrint(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push("hello")

	output := captureOutput(func() {
		opPrint(testInterpreter)
	})

	if output != "hello" {
		t.Errorf("Expected 'hello', got '%s'", output)
	}
}

func TestOpEquals(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push(42)

	output := captureOutput(func() {
		opEquals(testInterpreter)
	})

	if output != "42\n" {
		t.Errorf("Expected '42\\n', got '%s'", output)
	}
}

func TestOpEqualsEqualsString(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push("hello")

	output := captureOutput(func() {
		opEqualsEquals(testInterpreter)
	})

	if output != "(hello)\n" {
		t.Errorf("Expected '(hello)\\n', got '%s'", output)
	}
}

func TestOpEqualsEqualsNumber(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push(42)

	output := captureOutput(func() {
		opEqualsEquals(testInterpreter)
	})

	if output != "42\n" {
		t.Errorf("Expected '42\\n', got '%s'", output)
	}
}
