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

// basic arithmetic tests ============================================================
// add, sub, mul, div

func TestOpAdd(t *testing.T) {

	// populating list of tests
	tests := []struct {
		name     string  // name of the test being run
		x, y     any     // operands
		expected float64 // expected value
	}{
		{"positive integer addition", 3, 4, 7},
		{"negative integer addition", -5, -3, -8},
		{"mixed sign integer addition", 10, -3, 7},
		{"zero addition", 5, 0, 5},
		{"floating point addition", 3.2, 2.5, 5.7},
		{"floating point and integer addition", 5, 2.5, 7.5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.x)
			testInterpreter.opStack.Push(test.y)

			// error present
			err := opAdd(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected add error: %v", err)
			}

			// obtaining result from top of stack
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("add(%v, %v): expected %v, got %v", test.x, test.y, test.expected, result)
			}
		})
	}
}

func TestOpSub(t *testing.T) {
	tests := []struct {
		name     string
		x, y     any
		expected float64
	}{
		{"positive integer subtraction", 10, 3, 7},
		{"mixed sign integer subtraction", 3, 10, -7},
		{"negative integer subtraction", -5, -3, -2},
		{"zero subtraction", 5, 0, 5.0},
		{"floating point subtraction", 5.7, 2.5, 3.2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.x)
			testInterpreter.opStack.Push(test.y)

			// error occurred
			err := opSub(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected sub error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("sub(%v, %v): expected %v, got %v", test.x, test.y, test.expected, result)
			}
		})
	}
}

func TestOpMul(t *testing.T) {
	tests := []struct {
		name     string
		x, y     any
		expected float64
	}{
		{"integer multiplication", 5, 5, 25},
		{"mixed sign multiplication", -2, 5, -10},
		{"negative integer multiplication", -1, -5, 5},
		{"floating point multiplication", 2.5, 2, 5.0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.x)
			testInterpreter.opStack.Push(test.y)
			err := opMul(testInterpreter)

			//error
			if err != nil {
				t.Fatalf("unexpected mul error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("mul(%v, %v): expected: %v, got: %v", test.x, test.y, test.expected, result)
			}
		})
	}
}

func TestOpDiv(t *testing.T) {
	tests := []struct {
		name     string
		x, y     any
		expected float64
	}{
		{"integer division", 10, 2, 5},
		{"integer division with remainder", 10, 3, 3.3333333333333335},
		{"negative dividend", -10, 2, -5},
		{"negative divisor", 10, -2, -5},
		{"both negative", -10, -2, 5},
		{"floating point division", 7.5, 2.5, 3.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.x)
			testInterpreter.opStack.Push(test.y)

			// error
			err := opDiv(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected div error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("div(%v, %v): expected %v, got %v", test.x, test.y, test.expected, result)
			}
		})
	}
}

func TestOpDivByZero(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push(10)
	testInterpreter.opStack.Push(0)

	err := opDiv(testInterpreter)
	if err == nil {
		t.Error("expected division by zero error")
	}
}

// integer division tests =====================================================

func TestOpIntdiv(t *testing.T) {
	tests := []struct {
		name     string
		x, y     any
		expected int
	}{
		{"simple division", 10, 3, 3},
		{"exact division", 20, 5, 4},
		{"division with remainder", 7, 2, 3},
		{"negative dividend", -10, 3, -3},
		{"negative divisor", 10, -3, -3},
		{"both negative", -10, -3, 3},
		{"float inputs", 10.8, 3.2, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.x)
			testInterpreter.opStack.Push(test.y)

			// error
			err := opIntdiv(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected idiv error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("idiv(%v, %v): expected %v, got %v", test.x, test.y, test.expected, result)
			}
		})
	}
}

func TestOpIntdivByZero(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push(10)
	testInterpreter.opStack.Push(0)

	// error
	err := opIntdiv(testInterpreter)
	if err == nil {
		t.Error("expected division by zero error")
	}
}

func TestOpMod(t *testing.T) {
	tests := []struct {
		name     string
		x, y     any
		expected int
	}{
		{"simple modulo", 10, 3, 1},
		{"exact division", 20, 5, 0},
		{"small remainder", 7, 2, 1},
		{"larger remainder", 8, 3, 2},
		{"negative dividend", -10, 3, -1},
		{"large numbers", 100, 7, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.x)
			testInterpreter.opStack.Push(test.y)

			// error
			err := opMod(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected mod error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("mod(%v, %v): expected %v, got %v", test.x, test.y, test.expected, result)
			}
		})
	}
}

func TestOpModByZero(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push(10)
	testInterpreter.opStack.Push(0)

	// error
	err := opMod(testInterpreter)
	if err == nil {
		t.Error("expected division by zero error")
	}
}

// unary operations =============================================

func TestOpAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected float64
	}{
		{"positive integer", 5, 5},
		{"negative integer", -5, 5},
		{"zero", 0, 0},
		{"positive float", 3.14, 3.14},
		{"negative float", -3.14, 3.14},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.input)

			// error 
			err := opAbs(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected abs error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("abs(%v): expected %v, got %v", test.input, test.expected, result)
			}
		})
	}
}

func TestOpNeg(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected float64
	}{
		{"positive integer", 5, -5},
		{"negative integer", -5, 5},
		{"zero", 0, 0.0},
		{"positive float", 3.14, -3.14},
		{"negative float", -3.14, 3.14},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.input)

			// error
			err := opNeg(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected neg error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("neg(%v): expected %v, got %v", test.input, test.expected, result)
			}
		})
	}
}

func TestOpSqrt(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected float64
	}{
		{"perfect square 9", 9, 3},
		{"perfect square 16", 16, 4},
		{"perfect square 25", 25, 5},
		{"non-perfect square", 2, 1.4142135623730951},
		{"large number", 100, 10},
		{"zero", 0, 0},
		{"decimal", 0.25, 0.5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.input)

			// error
			err := opSqrt(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected sqrt error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			resultFloat, _ := result.(float64)
			if resultFloat != test.expected {
				t.Errorf("sqrt(%v): expected %v, got %v", test.input, test.expected, result)
			}
		})
	}
}

// rounding operations ===============================================

func TestOpCeiling(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected float64
	}{
		{"round up low decimal", 3.2, 4},
		{"round up high decimal", 3.8, 4},
		{"already integer", 3.0, 3.0},
		{"negative round up", -4.8, -4},
		{"negative low decimal", -4.2, -4},
		{"small positive", 0.1, 1},
		{"small negative", -0.1, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.input)

			// error 
			err := opCeil(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected ceiling error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("ceiling(%v): expected %v, got %v", test.input, test.expected, result)
			}
		})
	}
}

func TestOpFloor(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected float64
	}{
		{"round down low decimal", 3.2, 3},
		{"round down high decimal", 3.8, 3},
		{"already integer", 3.0, 3},
		{"negative round down", -4.8, -5},
		{"negative high decimal", -4.2, -5},
		{"small positive", 0.9, 0},
		{"small negative", -0.1, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.input)

			// error
			err := opFloor(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected floor error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("floor(%v): expected %v, got %v", test.input, test.expected, result)
			}
		})
	}
}

func TestOpRound(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected float64
	}{
		{"round down", 3.2, 3},
		{"round up at half", 3.5, 4},
		{"round up", 3.8, 4},
		{"already integer", 3.0, 3},
		{"negative round down", -4.8, -5},
		{"negative round up", -4.2, -4},
		{"negative at half", -4.5, -5},
		{"positive at half", 0.5, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.input)

			// error
			err := opRound(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected round error: %v", err)
			}

			// result checking
			result, _ := testInterpreter.opStack.Pop()
			if result != test.expected {
				t.Errorf("round(%v): expected %v, got %v", test.input, test.expected, result)
			}
		})
	}
}

// error handling for stack underflow errors

func TestArithmeticStackUnderflow(t *testing.T) {
	tests := []struct {
		name string
		op   func(*Interpreter) error
	}{
		{"add underflow", opAdd},
		{"sub underflow", opSub},
		{"mul underflow", opMul},
		{"div underflow", opDiv},
		{"idiv underflow", opIntdiv},
		{"mod underflow", opMod},
		{"abs underflow", opAbs},
		{"neg underflow", opNeg},
		{"sqrt underflow", opSqrt},
		{"ceiling underflow", opCeil},
		{"floor underflow", opFloor},
		{"round underflow", opRound},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()

			err := test.op(testInterpreter)
			if err == nil {
				t.Error("expected stack underflow error")
			}
		})
	}
}
