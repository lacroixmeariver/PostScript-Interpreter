package main

import (
	"testing"
)

/*
    Disclosure: The tests in this file were written using Generative AI.
*/

// ========================================
// BOOLEAN CONSTANT OPERATORS
// ========================================

func TestOpTrue(t *testing.T) {
	interp := CreateInterpreter()
	
	err := opTrue(interp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if interp.opStack.StackCount() != 1 {
		t.Errorf("expected 1 item on stack, got %d", interp.opStack.StackCount())
	}
	
	result, _ := interp.opStack.Pop()
	if result != true {
		t.Errorf("expected true, got %v", result)
	}
}

func TestOpTrueMultiplePushes(t *testing.T) {
	interp := CreateInterpreter()
	
	opTrue(interp)
	opTrue(interp)
	opTrue(interp)
	
	if interp.opStack.StackCount() != 3 {
		t.Errorf("expected 3 items on stack, got %d", interp.opStack.StackCount())
	}
	
	for i := 0; i < 3; i++ {
		result, _ := interp.opStack.Pop()
		if result != true {
			t.Errorf("expected true at position %d, got %v", i, result)
		}
	}
}

func TestOpFalse(t *testing.T) {
	interp := CreateInterpreter()
	
	err := opFalse(interp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if interp.opStack.StackCount() != 1 {
		t.Errorf("expected 1 item on stack, got %d", interp.opStack.StackCount())
	}
	
	result, _ := interp.opStack.Pop()
	if result != false {
		t.Errorf("expected false, got %v", result)
	}
}

func TestOpFalseMultiplePushes(t *testing.T) {
	interp := CreateInterpreter()
	
	opFalse(interp)
	opFalse(interp)
	
	if interp.opStack.StackCount() != 2 {
		t.Errorf("expected 2 items on stack, got %d", interp.opStack.StackCount())
	}
	
	for i := 0; i < 2; i++ {
		result, _ := interp.opStack.Pop()
		if result != false {
			t.Errorf("expected false at position %d, got %v", i, result)
		}
	}
}

// ========================================
// LOGICAL NOT OPERATOR
// ========================================

func TestOpNot(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected bool
	}{
		{"not true", true, false},
		{"not false", false, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.input)
			
			err := opNot(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("not(%v): expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}
}

func TestOpNotStackUnderflow(t *testing.T) {
	interp := CreateInterpreter()
	
	err := opNot(interp)
	if err == nil {
		t.Error("expected stack underflow error")
	}
}

func TestOpNotTypeError(t *testing.T) {
	interp := CreateInterpreter()
	interp.opStack.Push(5)  // Not a boolean
	
	err := opNot(interp)
	if err == nil {
		t.Error("expected type error for non-boolean")
	}
}

// ========================================
// LOGICAL OR OPERATOR
// ========================================

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
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.a)
			interp.opStack.Push(tt.b)
			
			err := opOr(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("or(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
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
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			
			for i := 0; i < tt.stackItems; i++ {
				interp.opStack.Push(true)
			}
			
			err := opOr(interp)
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
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.a)
			interp.opStack.Push(tt.b)
			
			err := opOr(interp)
			if err == nil {
				t.Error("expected type error")
			}
		})
	}
}

// ========================================
// BOOLEAN INTEGRATION TESTS
// ========================================

func TestBooleanSimpleExpression(t *testing.T) {
	// Test: true false or not
	// Expected: (true || false) = true, !true = false
	interp := CreateInterpreter()
	
	opTrue(interp)
	opFalse(interp)
	opOr(interp)
	opNot(interp)
	
	result, _ := interp.opStack.Pop()
	if result != false {
		t.Errorf("expected false, got %v", result)
	}
}

func TestBooleanChainedOr(t *testing.T) {
	// Test: true true or false or
	// Expected: (true || true) = true, (true || false) = true
	interp := CreateInterpreter()
	
	opTrue(interp)
	opTrue(interp)
	opOr(interp)
	opFalse(interp)
	opOr(interp)
	
	result, _ := interp.opStack.Pop()
	if result != true {
		t.Errorf("expected true, got %v", result)
	}
}

func TestBooleanNegationWithOr(t *testing.T) {
	// Test: false not true or
	// Expected: !false = true, (true || true) = true
	interp := CreateInterpreter()
	
	opFalse(interp)
	opNot(interp)
	opTrue(interp)
	opOr(interp)
	
	result, _ := interp.opStack.Pop()
	if result != true {
		t.Errorf("expected true, got %v", result)
	}
}

func TestBooleanComplexExpression(t *testing.T) {
	// Test: true false not or
	// Expected: !false = true, (true || true) = true
	interp := CreateInterpreter()
	
	opTrue(interp)
	opFalse(interp)
	opNot(interp)
	opOr(interp)
	
	result, _ := interp.opStack.Pop()
	if result != true {
		t.Errorf("expected true, got %v", result)
	}
	
	// Verify stack is empty after operation
	if interp.opStack.StackCount() != 0 {
		t.Errorf("expected empty stack, got %d items", interp.opStack.StackCount())
	}
}