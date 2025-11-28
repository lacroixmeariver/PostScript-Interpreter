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

// stack manipulation operations =============================================

// testing to ensure dup works with different types of data
func TestOpDup(t *testing.T) {
	tests := []struct {
		name  string
		value any
	}{
		{"integer", 42},
		{"float", 3.14},
		{"string", "hello"},
		{"boolean", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()
			testInterpreter.opStack.Push(test.value)

			opDup(testInterpreter)

			if testInterpreter.opStack.StackCount() != 2 {
				t.Fatalf("expected 2 items, got %d", testInterpreter.opStack.StackCount())
			}

			first, _ := testInterpreter.opStack.Pop()
			second, _ := testInterpreter.opStack.Pop()

			if first != second || first != test.value {
				t.Errorf("dup failed for %v", test.value)
			}
		})
	}
}

// running multiple pop() operations
func TestOpPop(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push(1)
	testInterpreter.opStack.Push(2)
	testInterpreter.opStack.Push(3)

	opPop(testInterpreter)
	opPop(testInterpreter)

	if testInterpreter.opStack.StackCount() != 1 {
		t.Errorf("expected 1 item remaining, got %d", testInterpreter.opStack.StackCount())
	}

	remaining, _ := testInterpreter.opStack.Pop()
	if remaining != 1 {
		t.Errorf("expected bottom item to be 1, got %v", remaining)
	}
}

func TestOpExch(t *testing.T) {
	testInterpreter := CreateInterpreter()
	second := 10
	first := 5

	// stack: [first, second]
	testInterpreter.opStack.Push(second)
	testInterpreter.opStack.Push(first)

	// stack: [second, first]
	err := opExch(testInterpreter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// new first should be equal to second number
	newTop, _ := testInterpreter.opStack.Peek()
	if newTop != second {
		t.Errorf("expected first to be %v after exchange, got %v", second, newTop)
	}

	testInterpreter.opStack.Pop()
	newSecond, _ := testInterpreter.opStack.Pop()
	if newSecond != first {
		t.Errorf("expected second to be %v after exchange, got %v", first, newSecond)
	}
}

// ensuring exch works on different types of data
func TestOpExchWithDifferentTypes(t *testing.T) {
	testInterpreter := CreateInterpreter()
	testInterpreter.opStack.Push("hello")
	testInterpreter.opStack.Push(42)

	err := opExch(testInterpreter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	first, _ := testInterpreter.opStack.Pop()
	second, _ := testInterpreter.opStack.Pop()

	if first != "hello" || second != 42 {
		t.Errorf("exch failed: got first=%v, second=%v", first, second)
	}
}

func TestOpClear(t *testing.T) {
	testInterpreter := CreateInterpreter()

	// stack: [1, 2, 3]
	testInterpreter.opStack.Push(1)
	testInterpreter.opStack.Push(2)
	testInterpreter.opStack.Push(3)

	err := opClear(testInterpreter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if testInterpreter.opStack.StackCount() != 0 {
		t.Errorf("expected empty stack, got %d items", testInterpreter.opStack.StackCount())
	}
}

func TestOpClearEmptyStack(t *testing.T) {
	testInterpreter := CreateInterpreter()

	// stack: []
	err := opClear(testInterpreter)
	if err != nil {
		t.Fatalf("unexpected error on empty stack: %v", err)
	}

	if testInterpreter.opStack.StackCount() != 0 {
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInterpreter := CreateInterpreter()

			for i := 0; i < test.itemCount; i++ {
				testInterpreter.opStack.Push(i)
			}

			err := opCount(testInterpreter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			count, _ := testInterpreter.opStack.Pop()

			if count != test.itemCount {
				t.Errorf("expected count %d, got %v", test.itemCount, count)
			}

			// verifying original items still on stack
			if testInterpreter.opStack.StackCount() != test.itemCount {
				t.Errorf("count modified stack: expected %d items, got %d", test.itemCount, testInterpreter.opStack.StackCount())
			}
		})
	}
}

// stack operations integration testing ===================================================

func TestStackOpsChaining(t *testing.T) {
	// Test: 1 2 3 dup exch pop
	testInterpreter := CreateInterpreter()

	// stack: [1, 2, 3]
	testInterpreter.opStack.Push(1)
	testInterpreter.opStack.Push(2)
	testInterpreter.opStack.Push(3)

	opDup(testInterpreter)  // stack: [1, 2, 3, 3]
	opExch(testInterpreter) // stack: [1, 2, 3, 3] (no change, exchanges first two which are same)
	opPop(testInterpreter)  // stack: [1, 2, 3]

	if testInterpreter.opStack.StackCount() != 3 {
		t.Errorf("expected 3 items, got %d", testInterpreter.opStack.StackCount())
	}

	first, _ := testInterpreter.opStack.Pop()
	if first != 3 {
		t.Errorf("expected first to be 3, got %v", first)
	}
}
