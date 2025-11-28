package main

import "fmt"

// defining the structure of a stack
type Stack struct {
	items     []PSConstant // items in stack
	itemCount int // number of items in stack
}

// constructor, creates instance of stack 
// *Stack: pointer to a stack 
func CreateStack() *Stack {
	return &Stack{
		items: make([]PSConstant, 0),
	}
}

// adds a PSConstant item to the top of the stack and increases the item count
func (s *Stack) Push(item PSConstant) {
	s.items = append(s.items, item)
	s.itemCount++
}

// removes and returns PSConstant item from the top of the stack and decreases the item count
func (s *Stack) Pop() (PSConstant, error) {
	if len(s.items) <= 0 {
		return nil, fmt.Errorf("no items in stack")
	}

	// hanging on to the item removed to be returned 
	item := s.items[len(s.items)-1]

	// slicing at item index to shave off removed item
	s.items = s.items[:len(s.items)-1]
	s.itemCount--

	return item, nil
}

// returns PSConstant at the top of the stack without removing it
func (s *Stack) Peek() (PSConstant, error) {
	if len(s.items) <= 0 {
		return nil, fmt.Errorf("stack underflow, no items in stack")
	}
	return s.items[len(s.items)-1], nil
}

// returns boolean value indicating whether stack is empty
// note: returns true if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) <= 0
}

// helper function to return stack count without giving direct access to stack top
func (s *Stack) StackCount() int {
	return s.itemCount
}
