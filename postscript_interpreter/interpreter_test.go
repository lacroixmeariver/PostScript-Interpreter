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

// ============================================ arithmetic interpreter tests

func TestAdd(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "add"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 8.0)
}

func TestSub(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "sub"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 7.0)
}

func TestMul(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 4},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "mul"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 20.0)
}

func TestDiv(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 20},
		{Type: TOKEN_INT, Value: 4},
		{Type: TOKEN_OPERATOR, Value: "div"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 5.0)
}

func TestIdiv(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 7},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "idiv"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 3)
}

func TestDivideByZeroError(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 0},
		{Type: TOKEN_OPERATOR, Value: "div"},
	}
	testInterpreter := CreateInterpreter()
	err := testInterpreter.Execute(tokens)
	if err == nil {
		t.Error("Expected divide by zero error")
	}
}

func TestMod(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "mod"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 1)
}

func TestAbs(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: -5},
		{Type: TOKEN_OPERATOR, Value: "abs"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 5.0)
}

func TestNeg(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "neg"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, -5.0)
}

func TestSqrt(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 16},
		{Type: TOKEN_OPERATOR, Value: "sqrt"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 4.0)
}

func TestCeiling(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_FLOAT, Value: 3.2},
		{Type: TOKEN_OPERATOR, Value: "ceiling"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 4.0)
}

func TestFloor(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_FLOAT, Value: 3.8},
		{Type: TOKEN_OPERATOR, Value: "floor"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 3.0)
}

func TestRound(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_FLOAT, Value: 3.5},
		{Type: TOKEN_OPERATOR, Value: "round"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 4.0)
}

func TestComplexArithmetic(t *testing.T) {
	tokens := []Token{
		// stack: [3, 5, add, 2, mul]
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "add"},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "mul"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 16.0)
}

// ============================================ stack manipulation tests

func TestDup(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 42},
		{Type: TOKEN_OPERATOR, Value: "dup"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackCount(t, testInterpreter, 2)
	compareStackTop(t, testInterpreter, 42)
}

func TestPop(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "pop"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackCount(t, testInterpreter, 2)
	compareStackTop(t, testInterpreter, 2)
}

func TestExch(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "exch"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 1)
	testInterpreter.opStack.Pop()
	compareStackTop(t, testInterpreter, 2)
}

func TestClear(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "clear"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackCount(t, testInterpreter, 0)
}

func TestCount(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "count"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, 3)
	compareStackCount(t, testInterpreter, 4)
}


func TestStackManipulation(t *testing.T) {
	tokens := []Token{
		//stack: [1, 2, 3, exch, pop]
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "exch"},
		{Type: TOKEN_OPERATOR, Value: "pop"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackCount(t, testInterpreter, 2)
	compareStackTop(t, testInterpreter, 3)
}

// ============================================ comparison tests

func TestEq(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "eq"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, true)
}

func TestNe(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "ne"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, true)
}

func TestGt(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "gt"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, true)
}

func TestGe(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "ge"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, true)
}

func TestLt(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "lt"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, true)
}

func TestLe(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "le"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, true)
}

// ============================================ boolean/logical tests

func TestAnd(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "true"},
		{Type: TOKEN_OPERATOR, Value: "false"},
		{Type: TOKEN_OPERATOR, Value: "and"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, false)
}

func TestOr(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "true"},
		{Type: TOKEN_OPERATOR, Value: "false"},
		{Type: TOKEN_OPERATOR, Value: "or"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, true)
}

func TestNot(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "true"},
		{Type: TOKEN_OPERATOR, Value: "not"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, false)
}

func TestTrue(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "true"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, true)
}

func TestFalse(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "false"},
	}
	testInterpreter := executeTest(t, tokens)
	compareStackTop(t, testInterpreter, false)
}

// ============================================ error handling

func TestStackUnderflowError(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "add"}, // not enough operands to complete this op
	}
	testInterpreter := CreateInterpreter()
	err := testInterpreter.Execute(tokens)
	if err == nil {
		t.Error("Expected stack underflow error")
	}
}

func TestUndefinedOperatorError(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "nonexistent"},
	}
	testInterpreter := CreateInterpreter()
	err := testInterpreter.Execute(tokens)
	if err == nil {
		t.Error("Expected undefined operator error")
	}
}

// ============================================ building procedure tests

func TestBuildProcedureSimple(t *testing.T) {
	// building: {1 2 add}
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "add"},
		{Type: TOKEN_BLOCK_END},
	}

	testInterpreter := CreateInterpreter()
	proc, endPos, err := testInterpreter.buildProcedure(tokens, 0)

	if err != nil {
		t.Fatalf("buildProcedure failed: %v", err)
	}

	if endPos != 5 {
		t.Errorf("Expected endPos 5, got %d", endPos)
	}

	if len(proc.Body) != 3 {
		t.Errorf("Expected 3 tokens in procedure body, got %d", len(proc.Body))
	}

	// token check
	if proc.Body[0].Type != TOKEN_INT || proc.Body[0].Value != 1 {
		t.Errorf("Expected first token to be INT 1")
	}
	if proc.Body[1].Type != TOKEN_INT || proc.Body[1].Value != 2 {
		t.Errorf("Expected second token to be INT 2")
	}
	if proc.Body[2].Type != TOKEN_OPERATOR || proc.Body[2].Value != "add" {
		t.Errorf("Expected third token to be OPERATOR add")
	}
}

func TestBuildProcedureNested(t *testing.T) {
	// building: { { 1 2 } 3 }
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_BLOCK_END},
	}

	testInterpreter := CreateInterpreter()
	proc, endPos, err := testInterpreter.buildProcedure(tokens, 0)

	if err != nil {
		t.Fatalf("buildProcedure failed: %v", err)
	}

	if endPos != 7 {
		t.Errorf("Expected endPos 7, got %d", endPos)
	}

	if len(proc.Body) != 5 {
		t.Errorf("Expected 4 tokens in procedure body, got %d", len(proc.Body))
	}

	// checking nested block structure
	if proc.Body[0].Type != TOKEN_BLOCK_START {
		t.Errorf("Expected first token to be BLOCK_START")
	}
	if proc.Body[4].Type != TOKEN_INT || proc.Body[4].Value != 3 {
		t.Errorf("Expected last token to be INT 3")
	}
}

func TestBuildProcedureEmpty(t *testing.T) {
	// building: { }
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_BLOCK_END},
	}

	testInterpreter := CreateInterpreter()
	proc, endPos, err := testInterpreter.buildProcedure(tokens, 0)

	if err != nil {
		t.Fatalf("buildProcedure failed: %v", err)
	}

	if endPos != 2 {
		t.Errorf("Expected endPos 2, got %d", endPos)
	}

	if len(proc.Body) != 0 {
		t.Errorf("Expected empty procedure body, got %d tokens", len(proc.Body))
	}
}

func TestBuildProcedureUnclosed(t *testing.T) {
	// building: { 1 2 
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		// no closing brace
	}

	testInterpreter := CreateInterpreter()
	_, _, err := testInterpreter.buildProcedure(tokens, 0)

	if err == nil {
		t.Error("Expected error for unclosed procedure, got nil")
	}

	if err.Error() != "unclosed procedure" {
		t.Errorf("Expected 'unclosed procedure' error, got: %v", err)
	}
}
