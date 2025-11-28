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

func TestOpDict(t *testing.T) {
	// testing the creation of a dictionary with cap of 10
	// expected value: dictionary with capacity of 10
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "dict"},
	}

	testInterpreter := executeTest(t, tokens)

	// checking to see if stack has 1 item
	if testInterpreter.opStack.StackCount() != 1 {
		t.Errorf("Expected 1 item on stack, got %d", testInterpreter.opStack.StackCount())
	}

	// checking to see that item is a dictionary
	top, _ := testInterpreter.opStack.Peek()
	dict, ok := top.(*PSDict)
	if !ok {
		t.Fatalf("Expected *PSDict on top of stack, got %T", top)
	}

	// it is a dictionary - what's it's capacity?
	if dict.capacity != 10 {
		t.Errorf("Expected capacity 10, got %d", dict.capacity)
	}

	// checking that the map got initialized
	if dict.items == nil {
		t.Error("Expected items map to be initialized, got nil")
	}

	// making sure the dictionary at this stage is empty
	if len(dict.items) != 0 {
		t.Errorf("Expected empty dictionary, got %d items", len(dict.items))
	}
}

func TestOpBegin(t *testing.T) {
	// testing the addition of new dict to dict stack
	// expected value: dictionary with cap 10 on the dict stack
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "dict"},
		{Type: TOKEN_OPERATOR, Value: "begin"},
	}

	testInterpreter := CreateInterpreter()

	// initial dict stack just has one dictionary (the global one)
	initialDictStackSize := len(testInterpreter.dictStack)

	err := testInterpreter.Execute(tokens)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// executing tokens should add one to stack
	if len(testInterpreter.dictStack) != initialDictStackSize+1 {
		t.Errorf("Expected dict stack size %d, got %d", initialDictStackSize+1, len(testInterpreter.dictStack))
	}

	// dictionary added should have been taken from the opstack
	if testInterpreter.opStack.StackCount() != 0 {
		t.Errorf("Expected empty operand stack, got %d items", testInterpreter.opStack.StackCount())
	}
}

func TestOpEnd(t *testing.T) {
	// testing the definition of the end of a dictionary on dict stack
	// dict stack should be +1 then back to initial size
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "dict"},
		{Type: TOKEN_OPERATOR, Value: "begin"},
		{Type: TOKEN_OPERATOR, Value: "end"},
	}

	testInterpreter := CreateInterpreter()
	initialDictStackSize := len(testInterpreter.dictStack) // stack: [global]

	err := testInterpreter.Execute(tokens) // stack: [newDict, global] >> [global]
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// dict stack should be back to original size
	if len(testInterpreter.dictStack) != initialDictStackSize {
		t.Errorf("Expected dict stack size %d, got %d", initialDictStackSize, len(testInterpreter.dictStack))
	}
}

func TestOpDef(t *testing.T) {
	// testing the definition of a variable or name binding
	// expected value: x should be defined in the dictionary as 5
	// note: converting to PSName in test cases bypassing tokenizer
	tokens := []Token{
		{Type: TOKEN_NAME, Value: PSName("x")}, // pushing /x as a name
		{Type: TOKEN_INT, Value: 5},            // pushing 5
		{Type: TOKEN_OPERATOR, Value: "def"},   // def operator setting x = 5
	}

	testInterpreter := executeTest(t, tokens)

	// opstack should be empty at this point (after execution)
	if testInterpreter.opStack.StackCount() != 0 {
		t.Errorf("Expected empty operand stack after def, got %d items", testInterpreter.opStack.StackCount())
	}

	// top of dict stack (the global in this case)
	currentDict := testInterpreter.dictStack[len(testInterpreter.dictStack)-1]

	// checking to see if x exists in the current dictionary
	val, exists := currentDict.items["x"]
	if !exists {
		t.Fatal("Expected 'x' to be defined in dictionary, but it wasn't found")
	}

	// verifying the value of x is 5
	if val != 5 {
		t.Errorf("Expected x = 5, got x = %v", val)
	}
}

func TestOpLength(t *testing.T) {
	// test: 10 dict /x 5 def /y 10 def length
	// Dictionary with 2 items should return length 2
	tokens := []Token{
		// newDict(10) entry: [x = 5], [y = 10]
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "dict"},
		{Type: TOKEN_OPERATOR, Value: "dup"},
		{Type: TOKEN_OPERATOR, Value: "begin"},
		{Type: TOKEN_NAME, Value: PSName("x")},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "def"},
		{Type: TOKEN_NAME, Value: PSName("y")},
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "def"},
		{Type: TOKEN_OPERATOR, Value: "end"},
		{Type: TOKEN_OPERATOR, Value: "length"},
	}

	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 2)
}
