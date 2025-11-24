package main

import (
	"testing"
)

func TestOpIf(t *testing.T) {
	// testing if block
	// Expected block to execute is {1 2 add}
	i := CreateInterpreter()

	tokens := []Token{
		{Type: TOKEN_BOOL, Value: true},
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "add"},
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_OPERATOR, Value: "if"},
	}
	i.Execute(tokens)
	checkStackTop(t, i, 3.0)
}

func TestOpIfElse(t *testing.T) {
	// testing ifelse block
	// Expected block to execute is {1 2 sub}
	i := CreateInterpreter()

	tokens := []Token{
		// bool value is false
		{Type: TOKEN_BOOL, Value: false},
		// {1 2 add} as the if block/procedure
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "add"},
		// {1 2 sub} as else block procedure
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "sub"},
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_OPERATOR, Value: "ifelse"},
	}
	i.Execute(tokens)
	checkStackTop(t, i, -1.0)
}

func TestOpFor(t *testing.T) {
	// testing for loop
	// Expected result should be 3 at the top of the stack
	i := CreateInterpreter()

	tokens := []Token{
		// {0 1 3}
		{Type: TOKEN_INT, Value: 0}, // starting value
		{Type: TOKEN_INT, Value: 1}, // step value
		{Type: TOKEN_INT, Value: 3}, // ending value
		// { } - empty block should just push counter onto the stack
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_OPERATOR, Value: "for"},
	}
	i.Execute(tokens)
	checkStackTop(t, i, 3)
}

func TestOpRepeat(t *testing.T) {
	// testing repeat function
	// Expected result should be 2 at the top of the stack
	i := CreateInterpreter()

	tokens := []Token{
		// {1 2 3 4 5}
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_INT, Value: 4},
		{Type: TOKEN_INT, Value: 5},
		// 3 {=} repeat - should execute print procedure 3 times
		{Type: TOKEN_INT, Value: 3}, // n
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_OPERATOR, Value: "="}, // proc
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_OPERATOR, Value: "repeat"},
	}
	i.Execute(tokens)
	checkStackTop(t, i, 2)
}

func TestOpQuit(t *testing.T) {
	// testing the quit function
	// Expected result is that interpreter should not make it executing the third line
	i := CreateInterpreter()
	tokens := []Token{
		{Type: TOKEN_INT, Value: 1}, // pushes to stack
		{Type: TOKEN_OPERATOR, Value: "quit"},
		{Type: TOKEN_INT, Value: 2}, // should not execute

	}
	i.Execute(tokens)
	// if !i.quit {
	// 	t.Fatal("quit flag expected to be true")
	// }
	checkStackCount(t, i, 1)
	checkStackTop(t, i, 1)

}
