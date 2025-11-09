package main

import "fmt"

type Stack struct{
	items []any
	itemCount int
}

// Constructor 
// *Stack - return type, returns pointer to stack created
func CreateStack() *Stack {
	return &Stack{
		items: make([]any, 0),
	}
} 

func (s *Stack) Push (item any){
	s.items = append(s.items, item)
	s.itemCount ++
	fmt.Println("Added to stack!")
	fmt.Println("Stack has: ", s.itemCount, " items.")
}

func (s *Stack) Pop () (any){
	item := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	s.itemCount --
	fmt.Println("Popped off stack!")
	fmt.Println("Stack has: ", s.itemCount, " items.")
	return item
}