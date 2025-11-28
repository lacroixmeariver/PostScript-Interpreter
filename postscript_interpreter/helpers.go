package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"testing"
)

// ====================================== helper functions to assist in conversions/test executions

// allows for interface types to be converted into numbers to perform operations on
func convertToNumber(num PSConstant) (float64, error) {
	switch val := num.(type) {
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	default:
		return math.NaN(), fmt.Errorf("incorrect input")
	}
}

// helper to capture stdout for input/output operations tests
func captureOutput(f func()) string {
	ogStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	// redirecting stdout to writer
	os.Stdout = writer

	// executing function that will write to stdout
	f()

	// closing to indicate information no longer incoming
	writer.Close()
	// redirecting stdout to original stdout
	os.Stdout = ogStdout

	// reading what was captured
	var buf bytes.Buffer
	buf.ReadFrom(reader)
	return buf.String()
}

// helper function to create interpreter and execute tokens
func executeTest(t *testing.T, tokens []Token) *Interpreter {
	testInterpreter := CreateInterpreter()
	err := testInterpreter.Execute(tokens)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	return testInterpreter
}

// helper to cross check top of stack with expected value
func compareStackTop(t *testing.T, testInterpreter *Interpreter, expected any) {
	if testInterpreter.opStack.StackCount() == 0 {
		t.Fatalf("Stack is empty, expected %v", expected)
	}
	top, _ := testInterpreter.opStack.Peek()
	if top != expected {
		t.Errorf("Expected %v on top of stack, got %v", expected, top)
	}
}

// helper to cross check stack count with expected value
func compareStackCount(t *testing.T, interp *Interpreter, expected int) {
	count := interp.opStack.StackCount()
	if count != expected {
		t.Errorf("Expected stack count %d, got %d", expected, count)
	}
}
