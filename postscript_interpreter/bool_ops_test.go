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

// boolean constant operators ========================================

func TestOpTrue(t *testing.T) {
	testInterpreter := CreateInterpreter()

	err := opTrue(testInterpreter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if testInterpreter.opStack.StackCount() != 1 {
		t.Errorf("expected 1 item on stack, got %d", testInterpreter.opStack.StackCount())
	}

	result, _ := testInterpreter.opStack.Pop()
	if result != true {
		t.Errorf("expected true, got %v", result)
	}
}

func TestOpTrueMultiplePushes(t *testing.T) {
	testInterpreter := CreateInterpreter()

	err1 := opTrue(testInterpreter)
	if err1 != nil {
		t.Fatalf("opTrue failed: %v", err1)
	}
	err2 := opTrue(testInterpreter)
	if err2 != nil {
		t.Fatalf("opTrue failed: %v", err2)
	}
	err3 := opTrue(testInterpreter)
	if err3 != nil {
		t.Fatalf("opTrue failed: %v", err3)
	}

	if testInterpreter.opStack.StackCount() != 3 {
		t.Errorf("expected 3 items on stack, got %d", testInterpreter.opStack.StackCount())
	}

	for i := 0; i < 3; i++ {
		result, _ := testInterpreter.opStack.Pop()
		if result != true {
			t.Errorf("expected true at position %d, got %v", i, result)
		}
	}
}

func TestOpFalse(t *testing.T) {
	testInterpreter := CreateInterpreter()

	err := opFalse(testInterpreter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if testInterpreter.opStack.StackCount() != 1 {
		t.Errorf("expected 1 item on stack, got %d", testInterpreter.opStack.StackCount())
	}

	result, _ := testInterpreter.opStack.Pop()
	if result != false {
		t.Errorf("expected false, got %v", result)
	}
}

func TestOpFalseMultiplePushes(t *testing.T) {
	testInterpreter := CreateInterpreter()

	err := opFalse(testInterpreter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	opFalse(testInterpreter)

	if testInterpreter.opStack.StackCount() != 2 {
		t.Errorf("expected 2 items on stack, got %d", testInterpreter.opStack.StackCount())
	}

	for i := 0; i < 2; i++ {
		result, _ := testInterpreter.opStack.Pop()
		if result != false {
			t.Errorf("expected false at position %d, got %v", i, result)
		}
	}
}

// logical not operations =======================================

func TestOpNot(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected bool
	}{
		{"not true", true, false},
		{"not false", false, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.input)

			err := opNot(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("not(%v): expected %v, got %v", test.input, test.expected, result)
			}
		})
	}
}

func TestOpNotStackUnderflow(t *testing.T) {
	testInterpreter := CreateInterpreter()

	err := opNot(testInterpreter)
	if err == nil {
		t.Error("expected stack underflow error")
	}
}

func TestOpNotTypeError(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push(5) // not a boolean val

	err := opNot(testInterpreter)
	if err == nil {
		t.Error("expected type error for non-boolean value")
	}
}

// logical or operations =======================================

func TestOpOr(t *testing.T) {
	tests := []struct {
		name     string
		a, b     bool
		expected bool
	}{
		{"true or true", true, true, true},
		{"true or false", true, false, true},
		{"false or true", false, true, true},
		{"false or false", false, false, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.a)
			testInterpreter.opStack.Push(test.b)

			err := opOr(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("or(%v, %v): expected %v, got %v", test.a, test.b, test.expected, result)
			}
		})
	}
}

func TestOpOrStackUnderflow(t *testing.T) {
	tests := []struct {
		name       string
		stackItems int
	}{
		{"empty stack", 0},
		{"one item only", 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()

			for i := 0; i < test.stackItems; i++ {
				testInterpreter.opStack.Push(true)
			}

			err := opOr(testInterpreter)
			if err == nil {
				t.Error("expected stack underflow error")
			}
		})
	}
}

func TestOpOrTypeError(t *testing.T) {
	tests := []struct {
		name string
		a, b interface{}
	}{
		{"int and bool", 5, true},
		{"bool and int", true, 5},
		{"two ints", 5, 10},
		{"string and bool", "hello", true},
		{"bool and string", false, "world"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.a)
			testInterpreter.opStack.Push(test.b)

			err := opOr(testInterpreter)
			if err == nil {
				t.Error("expected type error")
			}
		})
	}
}

// boolean integration tests ===============================

func TestBooleanSimpleExpression(t *testing.T) {
	// testing logical or/not against true/false values
	// expected values: true, false
	testInterpreter := CreateInterpreter()

	// stack visual: [empty]
	opTrue(testInterpreter) // stack visual: [true]
	opFalse(testInterpreter) // stack visual: [true, false]
	opOr(testInterpreter) // stack visual: [true] - results in true
	opNot(testInterpreter) // stack visual: [false] - results in false

	result, _ := testInterpreter.opStack.Pop()
	if result != false {
		t.Errorf("expected false, got %v", result)
	}
}

func TestBooleanChainedOr(t *testing.T) {
	// testing logical or against true/false values
	// expected values: true, true
	testInterpreter := CreateInterpreter()

	// stack visual: [empty]
	opTrue(testInterpreter) // stack visual: [true]
	opTrue(testInterpreter) // stack visual: [true, true]
	opOr(testInterpreter) // stack visual: [true] - results in true
	opFalse(testInterpreter) // stack visual: [true, false]
	opOr(testInterpreter) // stack visual: [true] - results in true

	result, _ := testInterpreter.opStack.Pop()
	if result != true {
		t.Errorf("expected true, got %v", result)
	}
}

func TestBooleanNegationWithOr(t *testing.T) {
	// testing logical or/not against true/false values
	// expected values: true, true
	testInterpreter := CreateInterpreter()

	// stack visual: [empty]
	opFalse(testInterpreter) // stack visual: [false]
	opNot(testInterpreter) // stack visual: [true] - results in true
	opTrue(testInterpreter) // stack visual: [true, true]
	opOr(testInterpreter) // stack visual: [true] - results in true

	result, _ := testInterpreter.opStack.Pop()
	if result != true {
		t.Errorf("expected true, got %v", result)
	}
}

func TestBooleanComplexExpression(t *testing.T) {
	// testing logical or/not against true/false values
	// Expected: true, true
	testInterpreter := CreateInterpreter()

	// stack visual: [empty]
	opTrue(testInterpreter) // stack visual: [true]
	opFalse(testInterpreter) // stack visual: [true, false]
	opNot(testInterpreter) // stack visual: [true, true] - results in true 
	opOr(testInterpreter) // stack visual: [true] - results in true

	result, _ := testInterpreter.opStack.Pop()
	if result != true {
		t.Errorf("expected true, got %v", result)
	}

	// verifying stack emptiness
	if testInterpreter.opStack.StackCount() != 0 {
		t.Errorf("expected empty stack, got %d items", testInterpreter.opStack.StackCount())
	}
}
