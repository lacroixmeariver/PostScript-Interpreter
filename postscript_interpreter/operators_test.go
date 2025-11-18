package main

import (
	"testing"
)

func TestOpDupe(t *testing.T){

	i := CreateInterpreter()
	i.opStack.Push(10)
	opDup(i)
	if i.opStack.itemCount != 2 {
		t.Fatal("stack top not pushed to top", "count: ", i.opStack.itemCount)
	}

	result, _ := i.opStack.Pop()
	topStack, _ := i.opStack.Peek()
	if result != topStack {
		t.Fatal("stack top does not match original value", topStack)
	}

}

func TestOpPop (t *testing.T) {
	i := CreateInterpreter()
	i.opStack.Push(5)
	expected := i.opStack.itemCount - 1
	opPop(i)

	if i.opStack.itemCount != expected {
		t.Fatal("stack pop not executed")
	}

}

func TestOpExch (t *testing.T) {
	i := CreateInterpreter()
	a := 10
	b := 5
	i.opStack.Push(a)
	i.opStack.Push(b)
	
	opExch(i)

	topStack, _ := i.opStack.Peek()
	if topStack != a {
		t.Fatal("exchange not executed correctly", topStack)
	}
}


// ========== Helper Functions ==========

// toFloat converts interface{} to float64 for easy comparison
func toFloat(val interface{}) float64 {
    switch v := val.(type) {
    case int:
        return float64(v)
    case float64:
        return v
    default:
        return 0
    }
}

// ========== Addition Tests ==========

func TestOpAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     interface{}
        expected float64
    }{
        {"two positive ints", 3, 4, 7.0},
        {"two negative ints", -5, -3, -8.0},
        {"positive and negative", 10, -3, 7.0},
        {"with zero", 0, 5, 5.0},
        {"both zero", 0, 0, 0.0},
        {"two floats", 3.5, 4.2, 7.7},
        {"int and float", 3, 4.5, 7.5},
        {"float and int", 2.5, 5, 7.5},
        {"negative floats", -2.5, -1.5, -4.0},
        {"large numbers", 1000000, 2000000, 3000000.0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opAdd(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            if interp.opStack.StackCount() != 1 {
                t.Fatalf("expected 1 item on stack, got %d", interp.opStack.StackCount())
            }
            
            result, _ := interp.opStack.Pop()
            resultNum := toFloat(result)
            
            // Use small epsilon for float comparison
            if diff := resultNum - tt.expected; diff < -0.0001 || diff > 0.0001 {
                t.Errorf("expected %v, got %v", tt.expected, resultNum)
            }
        })
    }
}

func TestOpAddStackUnderflow(t *testing.T) {
    tests := []struct {
        name       string
        stackItems []interface{}
    }{
        {"empty stack", []interface{}{}},
        {"one item", []interface{}{5}},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            for _, item := range tt.stackItems {
                interp.opStack.Push(item)
            }
            
            err := opAdd(interp)
            if err == nil {
                t.Error("expected error for stack underflow, got nil")
            }
        })
    }
}

func TestOpAddTypeError(t *testing.T) {
    tests := []struct {
        name string
        a, b interface{}
    }{
        {"string and number", "hello", 5},
        {"number and string", 5, "world"},
        {"two strings", "hello", "world"},
        {"bool and number", true, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opAdd(interp)
            if err == nil {
                t.Error("expected type error, got nil")
            }
        })
    }
}

// ========== Subtraction Tests ==========

func TestOpSub(t *testing.T) {
    tests := []struct {
        name     string
        a, b     interface{}
        expected float64
    }{
        {"positive result", 10, 3, 7.0},
        {"negative result", 3, 10, -7.0},
        {"subtract zero", 5, 0, 5.0},
        {"from zero", 0, 5, -5.0},
        {"same numbers", 5, 5, 0.0},
        {"negative numbers", -5, -3, -2.0},
        {"with floats", 10.5, 3.2, 7.3},
        {"negative minus positive", -5, 3, -8.0},
        {"positive minus negative", 5, -3, 8.0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opSub(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            result, _ := interp.opStack.Pop()
            resultNum := toFloat(result)
            
            if diff := resultNum - tt.expected; diff < -0.0001 || diff > 0.0001 {
                t.Errorf("expected %v, got %v", tt.expected, resultNum)
            }
        })
    }
}

func TestOpSubStackUnderflow(t *testing.T) {
    interp := CreateInterpreter()
    interp.opStack.Push(5)
    
    err := opSub(interp)
    if err == nil {
        t.Error("expected error for stack underflow, got nil")
    }
}

// ========== Multiplication Tests ==========

func TestOpMul(t *testing.T) {
    tests := []struct {
        name     string
        a, b     interface{}
        expected float64
    }{
        {"positive ints", 3, 4, 12.0},
        {"with zero", 5, 0, 0.0},
        {"with one", 5, 1, 5.0},
        {"two negatives", -3, -4, 12.0},
        {"positive and negative", 3, -4, -12.0},
        {"floats", 2.5, 4.0, 10.0},
        {"decimal result", 3, 0.5, 1.5},
        {"large numbers", 1000, 1000, 1000000.0},
        {"small decimals", 0.1, 0.2, 0.02},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opMul(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            result, _ := interp.opStack.Pop()
            resultNum := toFloat(result)
            
            if diff := resultNum - tt.expected; diff < -0.0001 || diff > 0.0001 {
                t.Errorf("expected %v, got %v", tt.expected, resultNum)
            }
        })
    }
}

func TestOpMulStackUnderflow(t *testing.T) {
    interp := CreateInterpreter()
    
    err := opMul(interp)
    if err == nil {
        t.Error("expected error for empty stack, got nil")
    }
}

// ========== Division Tests ==========

func TestOpDiv(t *testing.T) {
    tests := []struct {
        name     string
        a, b     interface{}
        expected float64
    }{
        {"even division", 10, 2, 5.0},
        {"decimal result", 10, 3, 3.333333},
        {"divide by one", 5, 1, 5.0},
        {"divide one", 1, 2, 0.5},
        {"negative dividend", -10, 2, -5.0},
        {"negative divisor", 10, -2, -5.0},
        {"both negative", -10, -2, 5.0},
        {"floats", 7.5, 2.5, 3.0},
        {"fraction", 0.5, 0.25, 2.0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opDiv(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            result, _ := interp.opStack.Pop()
            resultNum := toFloat(result)
            
            // Use larger epsilon for division due to floating point precision
            if diff := resultNum - tt.expected; diff < -0.001 || diff > 0.001 {
                t.Errorf("expected %v, got %v", tt.expected, resultNum)
            }
        })
    }
}

func TestOpDivByZero(t *testing.T) {
    tests := []struct {
        name string
        a, b interface{}
    }{
        {"int zero", 10, 0},
        {"float zero", 5.0, 0.0},
        {"zero divided by zero", 0, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opDiv(interp)
            if err == nil {
                t.Error("expected error for division by zero, got nil")
            }
        })
    }
}

func TestOpDivStackUnderflow(t *testing.T) {
    interp := CreateInterpreter()
    interp.opStack.Push(5)
    
    err := opDiv(interp)
    if err == nil {
        t.Error("expected error for stack underflow, got nil")
    }
}

// ========== Integration Tests ==========

func TestArithmeticChaining(t *testing.T) {
    // Test: 3 4 add 2 mul 10 sub  →  (3+4)*2-10 = 4
    interp := CreateInterpreter()
    
    interp.opStack.Push(3)
    interp.opStack.Push(4)
    opAdd(interp)  // Stack: [7]
    
    interp.opStack.Push(2)
    opMul(interp)  // Stack: [14]
    
    interp.opStack.Push(10)
    opSub(interp)  // Stack: [4]
    
    result, _ := interp.opStack.Pop()
    if toFloat(result) != 4.0 {
        t.Errorf("expected 4, got %v", result)
    }
}

func TestComplexArithmetic(t *testing.T) {
    // Test: 100 25 sub 5 div 3 mul  →  (100-25)/5*3 = 45
    interp := CreateInterpreter()
    
    interp.opStack.Push(100)
    interp.opStack.Push(25)
    opSub(interp)  // Stack: [75]
    
    interp.opStack.Push(5)
    opDiv(interp)  // Stack: [15]
    
    interp.opStack.Push(3)
    opMul(interp)  // Stack: [45]
    
    result, _ := interp.opStack.Pop()
    if toFloat(result) != 45.0 {
        t.Errorf("expected 45, got %v", result)
    }
}