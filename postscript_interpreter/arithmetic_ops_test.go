package main

import (
	"testing"
)

/*
	Disclosure: The tests in this file were written using Generative AI.
*/

// ========================================
// BASIC ARITHMETIC OPERATORS
// ========================================

func TestOpAdd(t *testing.T) {
	tests := []struct {
		name     string
		x, y     interface{}
		expected float64
	}{
		{"positive integers", 3, 4, 7.0},
		{"negative integers", -5, -3, -8.0},
		{"mixed signs", 10, -3, 7.0},
		{"with zero", 5, 0, 5.0},
		{"floats", 3.5, 2.5, 6.0},
		{"int and float", 5, 2.5, 7.5},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.x)
			interp.opStack.Push(tt.y)
			
			err := opAdd(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("add(%v, %v): expected %v, got %v", tt.x, tt.y, tt.expected, result)
			}
		})
	}
}

func TestOpSub(t *testing.T) {
	tests := []struct {
		name     string
		x, y     interface{}
		expected float64
	}{
		{"positive integers", 10, 3, 7.0},
		{"negative result", 3, 10, -7.0},
		{"negative integers", -5, -3, -2.0},
		{"with zero", 5, 0, 5.0},
		{"floats", 5.5, 2.5, 3.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.x)
			interp.opStack.Push(tt.y)
			
			err := opSub(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("sub(%v, %v): expected %v, got %v", tt.x, tt.y, tt.expected, result)
			}
		})
	}
}

func TestOpMul(t *testing.T) {
	tests := []struct {
		name     string
		x, y     interface{}
		expected float64
	}{
		{"positive integers", 5, 3, 15.0},
		{"with zero", 5, 0, 0.0},
		{"negative integers", -5, 3, -15.0},
		{"both negative", -5, -3, 15.0},
		{"floats", 2.5, 4.0, 10.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.x)
			interp.opStack.Push(tt.y)
			
			err := opMul(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("mul(%v, %v): expected %v, got %v", tt.x, tt.y, tt.expected, result)
			}
		})
	}
}

func TestOpDiv(t *testing.T) {
	tests := []struct {
		name     string
		x, y     interface{}
		expected float64
	}{
		{"simple division", 10, 2, 5.0},
		{"division with remainder", 10, 3, 3.3333333333333335},
		{"negative dividend", -10, 2, -5.0},
		{"negative divisor", 10, -2, -5.0},
		{"both negative", -10, -2, 5.0},
		{"floats", 7.5, 2.5, 3.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.x)
			interp.opStack.Push(tt.y)
			
			err := opDiv(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("div(%v, %v): expected %v, got %v", tt.x, tt.y, tt.expected, result)
			}
		})
	}
}

func TestOpDivByZero(t *testing.T) {
	interp := CreateInterpreter()
	interp.opStack.Push(10)
	interp.opStack.Push(0)
	
	err := opDiv(interp)
	if err == nil {
		t.Error("expected division by zero error")
	}
}

// ========================================
// INTEGER ARITHMETIC OPERATORS
// ========================================

func TestOpIdiv(t *testing.T) {
	tests := []struct {
		name     string
		x, y     interface{}
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
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.x)
			interp.opStack.Push(tt.y)
			
			err := opIdiv(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("idiv(%v, %v): expected %v, got %v", tt.x, tt.y, tt.expected, result)
			}
		})
	}
}

func TestOpIdivByZero(t *testing.T) {
	interp := CreateInterpreter()
	interp.opStack.Push(10)
	interp.opStack.Push(0)
	
	err := opIdiv(interp)
	if err == nil {
		t.Error("expected division by zero error")
	}
}

func TestOpMod(t *testing.T) {
	tests := []struct {
		name     string
		x, y     interface{}
		expected int
	}{
		{"simple modulo", 10, 3, 1},
		{"exact division", 20, 5, 0},
		{"small remainder", 7, 2, 1},
		{"larger remainder", 8, 3, 2},
		{"negative dividend", -10, 3, -1},
		{"large numbers", 100, 7, 2},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.x)
			interp.opStack.Push(tt.y)
			
			err := opMod(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("mod(%v, %v): expected %v, got %v", tt.x, tt.y, tt.expected, result)
			}
		})
	}
}

func TestOpModByZero(t *testing.T) {
	interp := CreateInterpreter()
	interp.opStack.Push(10)
	interp.opStack.Push(0)
	
	err := opMod(interp)
	if err == nil {
		t.Error("expected division by zero error")
	}
}

// ========================================
// UNARY ARITHMETIC OPERATORS
// ========================================

func TestOpAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected float64
	}{
		{"positive integer", 5, 5.0},
		{"negative integer", -5, 5.0},
		{"zero", 0, 0.0},
		{"positive float", 3.14, 3.14},
		{"negative float", -3.14, 3.14},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.input)
			
			err := opAbs(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("abs(%v): expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}
}

func TestOpNeg(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected float64
	}{
		{"positive integer", 5, -5.0},
		{"negative integer", -5, 5.0},
		{"zero", 0, 0.0},
		{"positive float", 3.14, -3.14},
		{"negative float", -3.14, 3.14},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.input)
			
			err := opNeg(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("neg(%v): expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}
}

func TestOpSqrt(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected float64
	}{
		{"perfect square 9", 9, 3.0},
		{"perfect square 16", 16, 4.0},
		{"perfect square 25", 25, 5.0},
		{"non-perfect square", 2, 1.4142135623730951},
		{"large number", 100, 10.0},
		{"zero", 0, 0.0},
		{"decimal", 0.25, 0.5},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.input)
			
			err := opSqrt(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			resultFloat, _ := result.(float64)
			if resultFloat != tt.expected {
				t.Errorf("sqrt(%v): expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}
}

// ========================================
// ROUNDING OPERATORS
// ========================================

func TestOpCeiling(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected float64
	}{
		{"round up low decimal", 3.2, 4.0},
		{"round up high decimal", 3.8, 4.0},
		{"already integer", 3.0, 3.0},
		{"negative round up", -4.8, -4.0},
		{"negative low decimal", -4.2, -4.0},
		{"small positive", 0.1, 1.0},
		{"small negative", -0.1, 0.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.input)
			
			err := opCeil(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("ceiling(%v): expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}
}

func TestOpFloor(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected float64
	}{
		{"round down low decimal", 3.2, 3.0},
		{"round down high decimal", 3.8, 3.0},
		{"already integer", 3.0, 3.0},
		{"negative round down", -4.8, -5.0},
		{"negative high decimal", -4.2, -5.0},
		{"small positive", 0.9, 0.0},
		{"small negative", -0.1, -1.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.input)
			
			err := opFloor(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("floor(%v): expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}
}

func TestOpRound(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected float64
	}{
		{"round down", 3.2, 3.0},
		{"round up at half", 3.5, 4.0},
		{"round up", 3.8, 4.0},
		{"already integer", 3.0, 3.0},
		{"negative round down", -4.8, -5.0},
		{"negative round up", -4.2, -4.0},
		{"negative at half", -4.5, -5.0},
		{"positive at half", 0.5, 1.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.input)
			
			err := opRound(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			result, _ := interp.opStack.Pop()
			if result != tt.expected {
				t.Errorf("round(%v): expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}
}

// ========================================
// ERROR HANDLING TESTS
// ========================================

func TestArithmeticStackUnderflow(t *testing.T) {
	tests := []struct {
		name string
		op   func(*Interpreter) error
	}{
		{"add underflow", opAdd},
		{"sub underflow", opSub},
		{"mul underflow", opMul},
		{"div underflow", opDiv},
		{"idiv underflow", opIdiv},
		{"mod underflow", opMod},
		{"abs underflow", opAbs},
		{"neg underflow", opNeg},
		{"sqrt underflow", opSqrt},
		{"ceiling underflow", opCeil},
		{"floor underflow", opFloor},
		{"round underflow", opRound},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			
			err := tt.op(interp)
			if err == nil {
				t.Error("expected stack underflow error")
			}
		})
	}
}