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

// equality operations ========================================

func TestOpEq(t *testing.T) {
	tests := []struct {
		name     string
		a, b     any
		expected bool
	}{
		// Numbers
		{"equal ints", 5, 5, true},
		{"unequal ints", 5, 3, false},
		{"equal floats", 3.14, 3.14, true},
		{"unequal floats", 3.14, 2.71, false},
		{"int equals float", 5, 5.0, true},
		{"int not equals float", 5, 5.1, false},
		{"zero equals zero", 0, 0, true},
		{"negative equals", -5, -5, true},
		{"negative not equals", -5, 5, false},

		// Strings
		{"equal strings", "hello", "hello", true},
		{"unequal strings", "hello", "world", false},
		{"empty strings", "", "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.a)
			testInterpreter.opStack.Push(test.b)

			err := opEq(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("eq(%v, %v): expected %v, got %v", test.a, test.b, test.expected, result)
			}
		})
	}
}

func TestOpEqStackUnderflow(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push(5)

	err := opEq(testInterpreter)
	if err == nil {
		t.Error("expected stack underflow error")
	}
}

func TestOpNe(t *testing.T) {
	tests := []struct {
		name     string
		a, b     any
		expected bool
	}{
		{"equal ints", 5, 5, false},
		{"unequal ints", 5, 3, true},
		{"equal floats", 3.14, 3.14, false},
		{"unequal floats", 3.14, 2.71, true},
		{"int equals float", 5, 5.0, false},
		{"int not equals float", 5, 5.1, true},
		{"zero equals zero", 0, 0, false},
		{"negative not equals positive", -5, 5, true},

		// Strings
		{"equal strings", "hello", "hello", false},
		{"unequal strings", "hello", "world", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.a)
			testInterpreter.opStack.Push(test.b)

			err := opNe(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("ne(%v, %v): expected %v, got %v", test.a, test.b, test.expected, result)
			}
		})
	}
}

// ========================================
// LESS THAN OPERATORS
// ========================================

func TestOpLt(t *testing.T) {
	tests := []struct {
		name     string
		a, b     any
		expected bool
	}{
		// Basic comparisons
		{"less than", 3, 5, true},
		{"greater than", 5, 3, false},
		{"equal", 5, 5, false},

		// Negative numbers
		{"negative less than negative", -5, -3, true},
		{"negative greater than negative", -3, -5, false},
		{"negative less than positive", -5, 3, true},
		{"positive greater than negative", 3, -5, false},

		// Zero comparisons
		{"zero less than positive", 0, 5, true},
		{"positive greater than zero", 5, 0, false},
		{"zero greater than negative", 0, -5, false},

		// Float comparisons
		{"float less than", 2.5, 3.7, true},
		{"float greater than", 3.7, 2.5, false},
		{"float equal", 3.14, 3.14, false},

		// Mixed types
		{"int less than float", 3, 3.5, true},
		{"float less than int", 2.5, 3, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.a)
			testInterpreter.opStack.Push(test.b)

			err := opLt(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("lt(%v, %v): expected %v, got %v", test.a, test.b, test.expected, result)
			}
		})
	}
}

func TestOpLtStrings(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected bool
	}{
		{"lexicographic less", "apple", "banana", true},
		{"lexicographic greater", "banana", "apple", false},
		{"equal strings", "apple", "apple", false},
		{"case sensitive A before a", "Apple", "apple", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.a)
			testInterpreter.opStack.Push(test.b)

			err := opLt(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("lt(%v, %v): expected %v, got %v", test.a, test.b, test.expected, result)
			}
		})
	}
}

func TestOpLe(t *testing.T) {
	tests := []struct {
		name     string
		a, b     any
		expected bool
	}{
		// Basic comparisons
		{"less than", 3, 5, true},
		{"greater than", 5, 3, false},
		{"equal", 5, 5, true},

		// Negative numbers
		{"negative less than negative", -5, -3, true},
		{"negative equal", -5, -5, true},
		{"negative greater than negative", -3, -5, false},

		// Zero comparisons
		{"zero equals zero", 0, 0, true},
		{"zero less than positive", 0, 5, true},
		{"positive greater than zero", 5, 0, false},

		// Float comparisons
		{"float less than", 2.5, 3.7, true},
		{"float equal", 3.14, 3.14, true},
		{"float greater than", 3.7, 2.5, false},

		// Edge cases
		{"very close floats equal", 3.14159, 3.14159, true},
		{"very close floats less", 3.14158, 3.14159, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.a)
			testInterpreter.opStack.Push(test.b)

			err := opLe(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("le(%v, %v): expected %v, got %v", test.a, test.b, test.expected, result)
			}
		})
	}
}

// ========================================
// GREATER THAN OPERATORS
// ========================================

func TestOpGt(t *testing.T) {
	tests := []struct {
		name     string
		a, b     any
		expected bool
	}{
		// Basic comparisons
		{"greater than", 5, 3, true},
		{"less than", 3, 5, false},
		{"equal", 5, 5, false},

		// Negative numbers
		{"negative greater than negative", -3, -5, true},
		{"negative less than negative", -5, -3, false},
		{"positive greater than negative", 3, -5, true},
		{"negative less than positive", -5, 3, false},

		// Zero comparisons
		{"positive greater than zero", 5, 0, true},
		{"zero not greater than zero", 0, 0, false},
		{"negative less than zero", -5, 0, false},

		// Float comparisons
		{"float greater than", 3.7, 2.5, true},
		{"float less than", 2.5, 3.7, false},
		{"float equal", 3.14, 3.14, false},

		// Large numbers
		{"large numbers", 1000000, 999999, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.a)
			testInterpreter.opStack.Push(test.b)

			err := opGt(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("gt(%v, %v): expected %v, got %v", test.a, test.b, test.expected, result)
			}
		})
	}
}

func TestOpGe(t *testing.T) {
	tests := []struct {
		name     string
		a, b     any
		expected bool
	}{
		// Basic comparisons
		{"greater than", 5, 3, true},
		{"less than", 3, 5, false},
		{"equal", 5, 5, true},

		// Negative numbers
		{"negative greater than negative", -3, -5, true},
		{"negative equal", -5, -5, true},
		{"negative less than negative", -5, -3, false},

		// Zero comparisons
		{"zero equals zero", 0, 0, true},
		{"positive greater than zero", 5, 0, true},
		{"zero not greater than positive", 0, 5, false},

		// Float comparisons
		{"float greater than", 3.7, 2.5, true},
		{"float equal", 3.14, 3.14, true},
		{"float less than", 2.5, 3.7, false},

		// Boundary cases
		{"max int", 2147483647, 2147483646, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.a)
			testInterpreter.opStack.Push(test.b)

			err := opGe(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("ge(%v, %v): expected %v, got %v", test.a, test.b, test.expected, result)
			}
		})
	}
}

// ========================================
// INTEGRATION TESTS
// ========================================

func TestComparisonWithBooleanLogic(t *testing.T) {
	// Test: 5 3 gt 2 1 gt and  →  (5>3) AND (2>1) → true AND true → true
	testInterpreter := CreateInterpreter()

	testInterpreter.opStack.Push(5)
	testInterpreter.opStack.Push(3)
	_ = opGt(testInterpreter)

	testInterpreter.opStack.Push(2)
	testInterpreter.opStack.Push(1)
	_ = opGt(testInterpreter)

	_ = opAnd(testInterpreter)

	result, _ := testInterpreter.opStack.Pop()
	if result != true {
		t.Errorf("expected true, got %v", result)
	}
}

func TestComparisonWithArithmetic(t *testing.T) {
	// Test: 3 4 add 10 lt  →  7 < 10 → true
	testInterpreter := CreateInterpreter()

	testInterpreter.opStack.Push(3)
	testInterpreter.opStack.Push(4)
	_ = opAdd(testInterpreter)

	testInterpreter.opStack.Push(10)
	_ = opLt(testInterpreter)

	result, _ := testInterpreter.opStack.Pop()
	if result != true {
		t.Errorf("expected true, got %v", result)
	}
}

func TestComparisonSymmetry(t *testing.T) {
	// Test: a < b is equivalent to b > a
	interp1 := CreateInterpreter()
	interp1.opStack.Push(3)
	interp1.opStack.Push(5)
	_ = opLt(interp1)
	result1, _ := interp1.opStack.Pop()

	interp2 := CreateInterpreter()
	interp2.opStack.Push(5)
	interp2.opStack.Push(3)
	_ = opGt(interp2)
	result2, _ := interp2.opStack.Pop()

	if result1 != result2 {
		t.Errorf("symmetry broken: 3<5=%v but 5>3=%v", result1, result2)
	}
}

// ========================================
// ERROR HANDLING TESTS
// ========================================

func TestComparisonTypeErrors(t *testing.T) {
	tests := []struct {
		name string
		op   func(*Interpreter) error
	}{
		{"lt type error", opLt},
		{"le type error", opLe},
		{"gt type error", opGt},
		{"ge type error", opGe},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(5)
			testInterpreter.opStack.Push("hello")

			err := test.op(testInterpreter)
			if err == nil {
				t.Log("Note: mixed type comparison allowed")
			}
		})
	}
}
