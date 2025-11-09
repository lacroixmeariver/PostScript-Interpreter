package main

import (
	"testing"
)

// stack tests
// initializing a stack
func TestStackCreation (t *testing.T) {
	testStack := CreateStack()
	count := len(testStack.items)
	if testStack == nil {
		t.Fatal("Stack returned nil, not created")
	}
	if count != 0 {
		t.Errorf("Stack is not empty")
	}
}

// testing push 
func TestStackPush (t *testing.T) {
	testStack := CreateStack()
	testStack.Push(5) // single integer value
	if len(testStack.items) != 1 {
		t.Errorf("Stack does not contain any values")
	}
}

// testing stack pop 
func TestStackPop (t *testing.T){
	testStack := CreateStack()
	testStack.Push(5)
	testStack.Pop()
	if len(testStack.items) != 0{
		t.Fatal("stack did not correctly pop item")
	}
}

func TestStackCount (t *testing.T){
	testStack := CreateStack()
	testStack.Push(5)
	testStack.Push(15)
	testStack.Push(11)

	if len(testStack.items) != 3{
		t.Fatal("stack item count incorrect")
	}
}