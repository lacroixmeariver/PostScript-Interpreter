package main

import (
	"testing"
)

/*
	Disclosure: The tests in this file were reformatted using Generative AI to improve the structure
	for clarity and readability. Content of tests were written by me. 
*/

// ========================================
// STACK MANIPULATION OPERATORS
// ========================================

func TestOpDup(t *testing.T) {
	interp := CreateInterpreter()
	interp.opStack.Push(10)
	
	err := opDup(interp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if interp.opStack.StackCount() != 2 {
		t.Fatalf("expected 2 items on stack, got %d", interp.opStack.StackCount())
	}
	
	top, _ := interp.opStack.Pop()
	second, _ := interp.opStack.Peek()
	
	if top != second {
		t.Errorf("dup failed: top=%v, second=%v (should be equal)", top, second)
	}
	
	if top != 10 {
		t.Errorf("expected 10, got %v", top)
	}
}

func TestOpDupWithDifferentTypes(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"integer", 42},
		{"float", 3.14},
		{"string", "hello"},
		{"boolean", true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			interp.opStack.Push(tt.value)
			
			opDup(interp)
			
			if interp.opStack.StackCount() != 2 {
				t.Fatalf("expected 2 items, got %d", interp.opStack.StackCount())
			}
			
			top, _ := interp.opStack.Pop()
			second, _ := interp.opStack.Pop()
			
			if top != second || top != tt.value {
				t.Errorf("dup failed for %v", tt.value)
			}
		})
	}
}

func TestOpPop(t *testing.T) {
	interp := CreateInterpreter()
	interp.opStack.Push(5)
	
	initialCount := interp.opStack.StackCount()
	
	err := opPop(interp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	finalCount := interp.opStack.StackCount()
	
	if finalCount != initialCount-1 {
		t.Errorf("expected count to decrease by 1: initial=%d, final=%d", initialCount, finalCount)
	}
}

func TestOpPopMultipleItems(t *testing.T) {
	interp := CreateInterpreter()
	interp.opStack.Push(1)
	interp.opStack.Push(2)
	interp.opStack.Push(3)
	
	opPop(interp)
	opPop(interp)
	
	if interp.opStack.StackCount() != 1 {
		t.Errorf("expected 1 item remaining, got %d", interp.opStack.StackCount())
	}
	
	remaining, _ := interp.opStack.Pop()
	if remaining != 1 {
		t.Errorf("expected bottom item to be 1, got %v", remaining)
	}
}

func TestOpExch(t *testing.T) {
	interp := CreateInterpreter()
	bottom := 10
	top := 5
	
	interp.opStack.Push(bottom)
	interp.opStack.Push(top)
	
	err := opExch(interp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	newTop, _ := interp.opStack.Peek()
	
	if newTop != bottom {
		t.Errorf("expected top to be %v after exchange, got %v", bottom, newTop)
	}
	
	interp.opStack.Pop()
	newSecond, _ := interp.opStack.Pop()
	
	if newSecond != top {
		t.Errorf("expected second to be %v after exchange, got %v", top, newSecond)
	}
}

func TestOpExchWithDifferentTypes(t *testing.T) {
	interp := CreateInterpreter()
	interp.opStack.Push("hello")
	interp.opStack.Push(42)
	
	opExch(interp)
	
	top, _ := interp.opStack.Pop()
	second, _ := interp.opStack.Pop()
	
	if top != "hello" || second != 42 {
		t.Errorf("exch failed: got top=%v, second=%v", top, second)
	}
}

func TestOpClear(t *testing.T) {
	interp := CreateInterpreter()
	interp.opStack.Push(1)
	interp.opStack.Push(2)
	interp.opStack.Push(3)
	
	err := opClear(interp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if interp.opStack.StackCount() != 0 {
		t.Errorf("expected empty stack, got %d items", interp.opStack.StackCount())
	}
}

func TestOpClearEmptyStack(t *testing.T) {
	interp := CreateInterpreter()
	
	err := opClear(interp)
	if err != nil {
		t.Fatalf("unexpected error on empty stack: %v", err)
	}
	
	if interp.opStack.StackCount() != 0 {
		t.Errorf("expected stack to remain empty")
	}
}

func TestOpCount(t *testing.T) {
	tests := []struct {
		name      string
		itemCount int
	}{
		{"empty stack", 0},
		{"one item", 1},
		{"three items", 3},
		{"five items", 5},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interp := CreateInterpreter()
			
			for i := 0; i < tt.itemCount; i++ {
				interp.opStack.Push(i)
			}
			
			err := opCount(interp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			count, _ := interp.opStack.Pop()
			
			if count != tt.itemCount {
				t.Errorf("expected count %d, got %v", tt.itemCount, count)
			}
			
			// Verify original items still on stack
			if interp.opStack.StackCount() != tt.itemCount {
				t.Errorf("count modified stack: expected %d items, got %d", tt.itemCount, interp.opStack.StackCount())
			}
		})
	}
}

// ========================================
// INTEGRATION TESTS
// ========================================

func TestStackOpsChaining(t *testing.T) {
	// Test: 1 2 3 dup exch pop
	// Expected: [1, 2, 3] → [1, 2, 3, 3] → [1, 2, 3, 3] → [1, 2, 3]
	interp := CreateInterpreter()
	
	interp.opStack.Push(1)
	interp.opStack.Push(2)
	interp.opStack.Push(3)
	
	opDup(interp)   // [1, 2, 3, 3]
	opExch(interp)  // [1, 2, 3, 3] (no change, exchanges top two which are same)
	opPop(interp)   // [1, 2, 3]
	
	if interp.opStack.StackCount() != 3 {
		t.Errorf("expected 3 items, got %d", interp.opStack.StackCount())
	}
	
	top, _ := interp.opStack.Pop()
	if top != 3 {
		t.Errorf("expected top to be 3, got %v", top)
	}
}