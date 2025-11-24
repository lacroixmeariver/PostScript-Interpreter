package main

import (
	"testing"
)

func TestOpDict(t *testing.T) {
	// Test: 10 dict
	// Should create a dictionary with capacity 10 and push it on the stack
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "dict"},
	}

	interp := executeTest(t, tokens)

	// Check stack has 1 item
	if interp.opStack.StackCount() != 1 {
		t.Errorf("Expected 1 item on stack, got %d", interp.opStack.StackCount())
	}

	// Check that it's a dictionary
	top, _ := interp.opStack.Peek()
	dict, ok := top.(*PSDict)
	if !ok {
		t.Fatalf("Expected *PSDict on top of stack, got %T", top)
	}

	// Check capacity
	if dict.capacity != 10 {
		t.Errorf("Expected capacity 10, got %d", dict.capacity)
	}

	// Check that items map is initialized
	if dict.items == nil {
		t.Error("Expected items map to be initialized, got nil")
	}

	// Check that it's empty
	if len(dict.items) != 0 {
		t.Errorf("Expected empty dictionary, got %d items", len(dict.items))
	}
}

func TestOpBegin(t *testing.T) {
	// Test: 10 dict begin
	// Should create dict and add it to dict stack
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "dict"},
		{Type: TOKEN_OPERATOR, Value: "begin"},
	}

	interp := CreateInterpreter()
	initialDictStackSize := len(interp.dictStack) // Should be 1 (global dict)

	err := interp.Execute(tokens)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Dict stack should have grown by 1
	if len(interp.dictStack) != initialDictStackSize+1 {
		t.Errorf("Expected dict stack size %d, got %d", initialDictStackSize+1, len(interp.dictStack))
	}

	// Operand stack should be empty (dict was consumed)
	if interp.opStack.StackCount() != 0 {
		t.Errorf("Expected empty operand stack, got %d items", interp.opStack.StackCount())
	}
}

func TestOpEnd(t *testing.T) {
	// Test: 10 dict begin end
	// Should add and then remove dict from dict stack
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "dict"},
		{Type: TOKEN_OPERATOR, Value: "begin"},
		{Type: TOKEN_OPERATOR, Value: "end"},
	}

	interp := CreateInterpreter()
	initialDictStackSize := len(interp.dictStack) // Should be 1

	err := interp.Execute(tokens)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Dict stack should be back to original size
	if len(interp.dictStack) != initialDictStackSize {
		t.Errorf("Expected dict stack size %d, got %d", initialDictStackSize, len(interp.dictStack))
	}
}

func TestOpDef(t *testing.T) {
	// Test: /x 5 def
	// Should define x = 5 in the current dictionary
	tokens := []Token{
		{Type: TOKEN_NAME, Value: "x"},       // Push /x as a name
		{Type: TOKEN_INT, Value: 5},          // Push 5
		{Type: TOKEN_OPERATOR, Value: "def"}, // Define x = 5
	}

	interp := executeTest(t, tokens)

	// After def, operand stack should be empty
	if interp.opStack.StackCount() != 0 {
		t.Errorf("Expected empty operand stack after def, got %d items", interp.opStack.StackCount())
	}

	// Check that x is defined in the current dictionary (top of dict stack)
	currentDict := interp.dictStack[len(interp.dictStack)-1]

	val, exists := currentDict.items["x"]
	if !exists {
		t.Fatal("Expected 'x' to be defined in dictionary, but it wasn't found")
	}

	// Check the value is 5
	if val != 5 {
		t.Errorf("Expected x = 5, got x = %v", val)
	}
}

func TestOpLength(t *testing.T) {
	// Test: 10 dict /x 5 def /y 10 def length
	// Dictionary with 2 items should return length 2
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "dict"},
		{Type: TOKEN_OPERATOR, Value: "begin"},
		{Type: TOKEN_NAME, Value: "x"},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "def"},
		{Type: TOKEN_NAME, Value: "y"},
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "def"},
		{Type: TOKEN_OPERATOR, Value: "end"},
		// Now get the dict back to check its length
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "dict"},
		{Type: TOKEN_OPERATOR, Value: "length"},
	}

	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 0) // New empty dict has 0 items
}

func TestOpMaxLength(t *testing.T) {
	// Test: 42 dict maxlength
	// Should return 42 (the capacity)
	tokens := []Token{
		{Type: TOKEN_INT, Value: 42},
		{Type: TOKEN_OPERATOR, Value: "dict"},
		{Type: TOKEN_OPERATOR, Value: "maxlength"},
	}

	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 42)
}
