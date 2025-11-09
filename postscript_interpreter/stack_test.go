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
	tests := []struct {
		name string
		values []PSConstant
		expected int
	}{
		{"empty", []PSConstant{}, 0},
		{"single item", []PSConstant{5}, 1},
		{"three items", []PSConstant{"hi", "hello", "how you doin'"}, 3},
		{"different types", []PSConstant{99, "red balloons", '!'}, 3},
		{"nil", []PSConstant{nil}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			testStack := CreateStack()

			for _, val := range tt.values {
				testStack.Push(val)
			}

			result := len(testStack.items)
			if result != tt.expected {
				t.Fatal("pushing results in unexpected stack count", 
				tt.expected, result)
			}
		})
	}
}

// testing pop
func TestStackPop (t *testing.T) {
	tests := []struct {
		name string
		values []PSConstant
		remove int
		expected int
	}{
		{"empty", []PSConstant{}, 1, 0},
		{"single item", []PSConstant{5, 10, 15}, 1, 2},
		{"three items", []PSConstant{"hi", "hello", "how you doin'"}, 3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			testStack := CreateStack()

			for _, val := range tt.values {
				testStack.Push(val)
			}
			
			i := 0
			for i < tt.remove{
				testStack.Pop()
				i++
			}
			
			result := len(testStack.items)
			if result != tt.expected {
				t.Fatal("pushing results in unexpected stack count", 
				tt.expected, result)
			}
		})
	}
}

// testing count 
func TestStackCount (t *testing.T) {
	tests := []struct {
		name string
		values []PSConstant
		expected int
	}{
		{"empty", []PSConstant{}, 0},
		{"single item", []PSConstant{5}, 1},
		{"three items", []PSConstant{"hi", "hello", "how you doin'"}, 3},
		{"four items of different types", []PSConstant{99, "red balloons", '!', 1001}, 4},
		{"nil", []PSConstant{nil}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			testStack := CreateStack()

			for _, val := range tt.values {
				testStack.Push(val)
			}

			result := testStack.StackCount()
			if result != tt.expected {
				t.Fatal("pushing results in unexpected stack count", 
				tt.expected, result)
			}
		})
	}
}

func TestIsEmpty (t *testing.T){
		tests := []struct {
		name string
		values []PSConstant
		expected bool
	}{
		{"empty", []PSConstant{}, true},
		{"not empty", []PSConstant{5}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			testStack := CreateStack()

			for _, val := range tt.values {
				testStack.Push(val)
			}

			result := testStack.IsEmpty()
			if result != tt.expected {
				t.Fatal("pushing results in unexpected stack count", 
				tt.expected, result)
			}
		})
	}
}