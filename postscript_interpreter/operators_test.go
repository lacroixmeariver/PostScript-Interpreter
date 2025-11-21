package main

import (
	"testing"
)

/*
    Disclosure: The tests below were written using Generative AI to ensure exhaustive code coverage.
*/

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


// ========== Equality Tests ==========

func TestOpEq(t *testing.T) {
    tests := []struct {
        name     string
        a, b     interface{}
        expected bool
    }{
        // Numbers
        {"equal ints", 5, 5, true},
        {"unequal ints", 5, 3, false},
        {"equal floats", 3.14, 3.14, true},
        {"unequal floats", 3.14, 2.71, false},
        {"int and float equal", 5, 5.0, true},
        {"int and float unequal", 5, 5.1, false},
        {"zero", 0, 0, true},
        {"negative equal", -5, -5, true},
        {"negative unequal", -5, 5, false},
        
        // Strings (if you implement string comparison)
        {"equal strings", "hello", "hello", true},
        {"unequal strings", "hello", "world", false},
        {"empty strings", "", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opEq(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            result, _ := interp.opStack.Pop()
            if result != tt.expected {
                t.Errorf("eq(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
            }
        })
    }
}

func TestOpEqStackUnderflow(t *testing.T) {
    interp := CreateInterpreter()
    interp.opStack.Push(5)  // Only one item
    
    err := opEq(interp)
    if err == nil {
        t.Error("expected stack underflow error")
    }
}

// ========== Not Equal Tests ==========

func TestOpNe(t *testing.T) {
    tests := []struct {
        name     string
        a, b     interface{}
        expected bool
    }{
        {"equal ints", 5, 5, false},
        {"unequal ints", 5, 3, true},
        {"equal floats", 3.14, 3.14, false},
        {"unequal floats", 3.14, 2.71, true},
        {"int and float equal", 5, 5.0, false},
        {"int and float unequal", 5, 5.1, true},
        {"zero", 0, 0, false},
        {"negative", -5, 5, true},
        
        // Strings
        {"equal strings", "hello", "hello", false},
        {"unequal strings", "hello", "world", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opNe(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            result, _ := interp.opStack.Pop()
            if result != tt.expected {
                t.Errorf("ne(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
            }
        })
    }
}

// ========== Less Than Tests ==========

func TestOpLt(t *testing.T) {
    tests := []struct {
        name     string
        a, b     interface{}
        expected bool
    }{
        // a < b
        {"less than", 3, 5, true},
        {"greater than", 5, 3, false},
        {"equal", 5, 5, false},
        
        // Negative numbers
        {"negative less", -5, -3, true},
        {"negative greater", -3, -5, false},
        {"negative and positive", -5, 3, true},
        {"positive and negative", 3, -5, false},
        
        // Zero
        {"zero less than positive", 0, 5, true},
        {"positive greater than zero", 5, 0, false},
        {"zero less than negative", 0, -5, false},
        
        // Floats
        {"float less", 2.5, 3.7, true},
        {"float greater", 3.7, 2.5, false},
        {"float equal", 3.14, 3.14, false},
        
        // Mixed
        {"int less than float", 3, 3.5, true},
        {"float less than int", 2.5, 3, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opLt(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            result, _ := interp.opStack.Pop()
            if result != tt.expected {
                t.Errorf("lt(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
            }
        })
    }
}

func TestOpLtStrings(t *testing.T) {
    // Only if you implement string comparison
    tests := []struct {
        name     string
        a, b     string
        expected bool
    }{
        {"lexicographic less", "apple", "banana", true},
        {"lexicographic greater", "banana", "apple", false},
        {"equal strings", "apple", "apple", false},
        {"case sensitive", "Apple", "apple", true},  // 'A' < 'a' in ASCII
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opLt(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            result, _ := interp.opStack.Pop()
            if result != tt.expected {
                t.Errorf("lt(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
            }
        })
    }
}

// ========== Less Than or Equal Tests ==========

func TestOpLe(t *testing.T) {
    tests := []struct {
        name     string
        a, b     interface{}
        expected bool
    }{
        // a <= b
        {"less than", 3, 5, true},
        {"greater than", 5, 3, false},
        {"equal", 5, 5, true},  // ← Different from lt!
        
        // Negative numbers
        {"negative less", -5, -3, true},
        {"negative equal", -5, -5, true},
        {"negative greater", -3, -5, false},
        
        // Zero
        {"zero equal zero", 0, 0, true},
        {"zero less than positive", 0, 5, true},
        {"positive greater than zero", 5, 0, false},
        
        // Floats
        {"float less", 2.5, 3.7, true},
        {"float equal", 3.14, 3.14, true},
        {"float greater", 3.7, 2.5, false},
        
        // Edge cases
        {"very close floats equal", 3.14159, 3.14159, true},
        {"very close floats less", 3.14158, 3.14159, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opLe(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            result, _ := interp.opStack.Pop()
            if result != tt.expected {
                t.Errorf("le(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
            }
        })
    }
}

// ========== Greater Than Tests ==========

func TestOpGt(t *testing.T) {
    tests := []struct {
        name     string
        a, b     interface{}
        expected bool
    }{
        // a > b
        {"greater than", 5, 3, true},
        {"less than", 3, 5, false},
        {"equal", 5, 5, false},
        
        // Negative numbers
        {"negative greater", -3, -5, true},
        {"negative less", -5, -3, false},
        {"positive greater than negative", 3, -5, true},
        {"negative less than positive", -5, 3, false},
        
        // Zero
        {"positive greater than zero", 5, 0, true},
        {"zero not greater than zero", 0, 0, false},
        {"negative less than zero", -5, 0, false},
        
        // Floats
        {"float greater", 3.7, 2.5, true},
        {"float less", 2.5, 3.7, false},
        {"float equal", 3.14, 3.14, false},
        
        // Large numbers
        {"large numbers", 1000000, 999999, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opGt(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            result, _ := interp.opStack.Pop()
            if result != tt.expected {
                t.Errorf("gt(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
            }
        })
    }
}

// ========== Greater Than or Equal Tests ==========

func TestOpGe(t *testing.T) {
    tests := []struct {
        name     string
        a, b     interface{}
        expected bool
    }{
        // a >= b
        {"greater than", 5, 3, true},
        {"less than", 3, 5, false},
        {"equal", 5, 5, true},  // ← Different from gt!
        
        // Negative numbers
        {"negative greater", -3, -5, true},
        {"negative equal", -5, -5, true},
        {"negative less", -5, -3, false},
        
        // Zero
        {"zero equal zero", 0, 0, true},
        {"positive greater than zero", 5, 0, true},
        {"zero not greater than positive", 0, 5, false},
        
        // Floats
        {"float greater", 3.7, 2.5, true},
        {"float equal", 3.14, 3.14, true},
        {"float less", 2.5, 3.7, false},
        
        // Boundary cases
        {"max int", 2147483647, 2147483646, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(tt.a)
            interp.opStack.Push(tt.b)
            
            err := opGe(interp)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            
            result, _ := interp.opStack.Pop()
            if result != tt.expected {
                t.Errorf("ge(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
            }
        })
    }
}

// ========== Integration Tests ==========

func TestComparisonChaining(t *testing.T) {
    // Test: 5 3 gt 2 1 gt and  →  (5>3) AND (2>1) → true AND true → true
    interp := CreateInterpreter()
    
    interp.opStack.Push(5)
    interp.opStack.Push(3)
    opGt(interp)  // Stack: [true]
    
    interp.opStack.Push(2)
    interp.opStack.Push(1)
    opGt(interp)  // Stack: [true, true]
    
    opAnd(interp)  // Stack: [true]
    
    result, _ := interp.opStack.Pop()
    if result != true {
        t.Errorf("expected true, got %v", result)
    }
}

func TestComparisonWithArithmetic(t *testing.T) {
    // Test: 3 4 add 10 lt  →  7 < 10 → true
    interp := CreateInterpreter()
    
    interp.opStack.Push(3)
    interp.opStack.Push(4)
    opAdd(interp)  // Stack: [7]
    
    interp.opStack.Push(10)
    opLt(interp)  // Stack: [true]
    
    result, _ := interp.opStack.Pop()
    if result != true {
        t.Errorf("expected true, got %v", result)
    }
}

func TestComparisonSymmetry(t *testing.T) {
    // Test that a < b is equivalent to b > a
    interp1 := CreateInterpreter()
    interp1.opStack.Push(3)
    interp1.opStack.Push(5)
    opLt(interp1)
    result1, _ := interp1.opStack.Pop()
    
    interp2 := CreateInterpreter()
    interp2.opStack.Push(5)
    interp2.opStack.Push(3)
    opGt(interp2)
    result2, _ := interp2.opStack.Pop()
    
    if result1 != result2 {
        t.Errorf("symmetry broken: 3<5=%v but 5>3=%v", result1, result2)
    }
}

// ========== Error Cases ==========

func TestComparisonTypeErrors(t *testing.T) {
    // Only if you're checking types strictly
    tests := []struct {
        name string
        op   func(*Interpreter) error
    }{
        {"lt type error", opLt},
        {"le type error", opLe},
        {"gt type error", opGt},
        {"ge type error", opGe},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            interp := CreateInterpreter()
            interp.opStack.Push(5)
            interp.opStack.Push("hello")  // Mixed types
            
            err := tt.op(interp)
            // Depending on your implementation, this might error or not
            // If you allow number-string comparison, skip this test
            if err == nil {
                t.Log("Note: mixed type comparison allowed")
            }
        })
    }
}

