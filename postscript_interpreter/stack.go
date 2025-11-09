package main

import "fmt"

type Stack struct{
	items []PSConstant
	itemCount int
}

// Constructor 
// *Stack - return type, returns pointer to stack created
func CreateStack() *Stack {
	return &Stack{
		items: make([]PSConstant, 0),
	}
} 

func (s *Stack) Push (item PSConstant) {
	s.items = append(s.items, item)
	s.itemCount ++
}

func (s *Stack) Pop () (PSConstant, error){
	if len(s.items) <= 0 {
		return nil, fmt.Errorf("no items in stack")
	}
	item := s.items[len(s.items) - 1]
	// slicing at this index
	s.items = s.items[:len(s.items) - 1]
	s.itemCount --
	return item, nil
}

func (s *Stack) Peek() (item PSConstant){
	item = s.items[:len(s.items)]
	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) <= 0 
}

func (s *Stack) StackCount() int {
	return s.itemCount
}