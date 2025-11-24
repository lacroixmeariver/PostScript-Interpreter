package main

import (
	"testing"
)

/*
	Disclosure: The tests in this file were written using Generative AI.
*/

// Helper function to create interpreter and execute tokens
func executeTest(t *testing.T, tokens []Token) *Interpreter {
	interp := CreateInterpreter()
	err := interp.Execute(tokens)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	return interp
}

// Helper to check top of stack
func checkStackTop(t *testing.T, interp *Interpreter, expected interface{}) {
	if interp.opStack.StackCount() == 0 {
		t.Fatalf("Stack is empty, expected %v", expected)
	}
	top, _ := interp.opStack.Peek()
	if top != expected {
		t.Errorf("Expected %v on top of stack, got %v", expected, top)
	}
}

// Helper to check stack count
func checkStackCount(t *testing.T, interp *Interpreter, expected int) {
	count := interp.opStack.StackCount()
	if count != expected {
		t.Errorf("Expected stack count %d, got %d", expected, count)
	}
}

// ============================================ Arithmetic Tests

func TestAdd(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "add"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 8.0)
}

func TestSub(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "sub"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 7.0)
}

func TestMul(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 4},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "mul"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 20.0)
}

func TestDiv(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 20},
		{Type: TOKEN_INT, Value: 4},
		{Type: TOKEN_OPERATOR, Value: "div"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 5.0)
}

func TestIdiv(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 7},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "idiv"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 3)
}

func TestMod(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "mod"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 1)
}

func TestAbs(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: -5},
		{Type: TOKEN_OPERATOR, Value: "abs"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 5.0)
}

func TestNeg(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "neg"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, -5.0)
}

func TestSqrt(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 16},
		{Type: TOKEN_OPERATOR, Value: "sqrt"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 4.0)
}

func TestCeiling(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_FLOAT, Value: 3.2},
		{Type: TOKEN_OPERATOR, Value: "ceiling"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 4.0)
}

func TestFloor(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_FLOAT, Value: 3.8},
		{Type: TOKEN_OPERATOR, Value: "floor"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 3.0)
}

func TestRound(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_FLOAT, Value: 3.5},
		{Type: TOKEN_OPERATOR, Value: "round"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 4.0)
}

// ============================================ Stack Tests

func TestDup(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 42},
		{Type: TOKEN_OPERATOR, Value: "dup"},
	}
	interp := executeTest(t, tokens)
	checkStackCount(t, interp, 2)
	checkStackTop(t, interp, 42)
}

func TestPop(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "pop"},
	}
	interp := executeTest(t, tokens)
	checkStackCount(t, interp, 2)
	checkStackTop(t, interp, 2)
}

func TestExch(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "exch"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 1)
	interp.opStack.Pop()
	checkStackTop(t, interp, 2)
}

func TestClear(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "clear"},
	}
	interp := executeTest(t, tokens)
	checkStackCount(t, interp, 0)
}

func TestCount(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "count"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 3)
	checkStackCount(t, interp, 4) // 3 original items + count result
}

// ============================================ Comparison Tests

func TestEq(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "eq"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, true)
}

func TestNe(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "ne"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, true)
}

func TestGt(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "gt"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, true)
}

func TestGe(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "ge"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, true)
}

func TestLt(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "lt"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, true)
}

func TestLe(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "le"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, true)
}

// ============================================ Boolean Tests

func TestAnd(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "true"},
		{Type: TOKEN_OPERATOR, Value: "false"},
		{Type: TOKEN_OPERATOR, Value: "and"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, false)
}

func TestOr(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "true"},
		{Type: TOKEN_OPERATOR, Value: "false"},
		{Type: TOKEN_OPERATOR, Value: "or"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, true)
}

func TestNot(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "true"},
		{Type: TOKEN_OPERATOR, Value: "not"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, false)
}

func TestTrue(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "true"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, true)
}

func TestFalse(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "false"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, false)
}

// ============================================ Complex Tests

func TestComplexArithmetic(t *testing.T) {
	// (3 + 5) * 2 = 16
	tokens := []Token{
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "add"},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "mul"},
	}
	interp := executeTest(t, tokens)
	checkStackTop(t, interp, 16.0)
}

func TestStackManipulation(t *testing.T) {
	// 1 2 3 exch pop -> leaves 1 2
	tokens := []Token{
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_OPERATOR, Value: "exch"},
		{Type: TOKEN_OPERATOR, Value: "pop"},
	}
	interp := executeTest(t, tokens)
	checkStackCount(t, interp, 2)
	checkStackTop(t, interp, 3)
}

func TestDivideByZeroError(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 0},
		{Type: TOKEN_OPERATOR, Value: "div"},
	}
	interp := CreateInterpreter()
	err := interp.Execute(tokens)
	if err == nil {
		t.Error("Expected divide by zero error")
	}
}

func TestStackUnderflowError(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_OPERATOR, Value: "add"}, // Need 2 operands, only have 1
	}
	interp := CreateInterpreter()
	err := interp.Execute(tokens)
	if err == nil {
		t.Error("Expected stack underflow error")
	}
}

func TestUndefinedOperatorError(t *testing.T) {
	tokens := []Token{
		{Type: TOKEN_OPERATOR, Value: "nonexistent"},
	}
	interp := CreateInterpreter()
	err := interp.Execute(tokens)
	if err == nil {
		t.Error("Expected undefined operator error")
	}
}

// ============================================ buildProcedure Tests

func TestBuildProcedureSimple(t *testing.T) {
	// Test: { 1 2 add }
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "add"},
		{Type: TOKEN_BLOCK_END},
	}

	interp := CreateInterpreter()
	proc, endPos, err := interp.buildProcedure(tokens, 0)

	if err != nil {
		t.Fatalf("buildProcedure failed: %v", err)
	}

	if endPos != 5 {
		t.Errorf("Expected endPos 5, got %d", endPos)
	}

	if len(proc.Body) != 3 {
		t.Errorf("Expected 3 tokens in procedure body, got %d", len(proc.Body))
	}

	// Check the tokens are correct
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
	// Test: { { 1 2 } 3 }
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_BLOCK_END},
	}

	interp := CreateInterpreter()
	proc, endPos, err := interp.buildProcedure(tokens, 0)

	if err != nil {
		t.Fatalf("buildProcedure failed: %v", err)
	}

	if endPos != 7 {
		t.Errorf("Expected endPos 7, got %d", endPos)
	}

	if len(proc.Body) != 5 {
		t.Errorf("Expected 4 tokens in procedure body, got %d", len(proc.Body))
	}

	// Check that nested block markers are preserved
	if proc.Body[0].Type != TOKEN_BLOCK_START {
		t.Errorf("Expected first token to be BLOCK_START")
	}
	if proc.Body[4].Type != TOKEN_INT || proc.Body[4].Value != 3 {
		t.Errorf("Expected last token to be INT 3")
	}
}

func TestBuildProcedureDeepNesting(t *testing.T) {
	// Test: { { { 1 } } }
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_BLOCK_END},
	}

	interp := CreateInterpreter()
	proc, endPos, err := interp.buildProcedure(tokens, 0)

	if err != nil {
		t.Fatalf("buildProcedure failed: %v", err)
	}

	if endPos != 7 {
		t.Errorf("Expected endPos 7, got %d", endPos)
	}

	if len(proc.Body) != 5 {
		t.Errorf("Expected 5 tokens in procedure body, got %d", len(proc.Body))
	}
}

func TestBuildProcedureEmpty(t *testing.T) {
	// Test: { }
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_BLOCK_END},
	}

	interp := CreateInterpreter()
	proc, endPos, err := interp.buildProcedure(tokens, 0)

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
	// Test: { 1 2 (missing closing brace)
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
	}

	interp := CreateInterpreter()
	_, _, err := interp.buildProcedure(tokens, 0)

	if err == nil {
		t.Error("Expected error for unclosed procedure, got nil")
	}

	if err.Error() != "unclosed procedure" {
		t.Errorf("Expected 'unclosed procedure' error, got: %v", err)
	}
}

func TestBuildProcedureWithMultipleStatements(t *testing.T) {
	// Test: { 1 2 add 3 4 mul }
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_OPERATOR, Value: "add"},
		{Type: TOKEN_INT, Value: 3},
		{Type: TOKEN_INT, Value: 4},
		{Type: TOKEN_OPERATOR, Value: "mul"},
		{Type: TOKEN_BLOCK_END},
	}

	interp := CreateInterpreter()
	proc, endPos, err := interp.buildProcedure(tokens, 0)

	if err != nil {
		t.Fatalf("buildProcedure failed: %v", err)
	}

	if endPos != 8 {
		t.Errorf("Expected endPos 8, got %d", endPos)
	}

	if len(proc.Body) != 6 {
		t.Errorf("Expected 6 tokens in procedure body, got %d", len(proc.Body))
	}
}

func TestExecuteWithProcedure(t *testing.T) {
	// Test that procedures are correctly pushed onto the stack
	// { 5 10 add } should push the procedure onto the stack
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 5},
		{Type: TOKEN_INT, Value: 10},
		{Type: TOKEN_OPERATOR, Value: "add"},
		{Type: TOKEN_BLOCK_END},
	}

	interp := executeTest(t, tokens)

	if interp.opStack.StackCount() != 1 {
		t.Errorf("Expected 1 item on stack, got %d", interp.opStack.StackCount())
	}

	top, _ := interp.opStack.Peek()
	proc, ok := top.(PSBlock)
	if !ok {
		t.Fatalf("Expected PSBlock on top of stack, got %T", top)
	}

	if len(proc.Body) != 3 {
		t.Errorf("Expected procedure with 3 tokens, got %d", len(proc.Body))
	}
}

func TestExecuteWithNestedProcedures(t *testing.T) {
	// Test: { { 1 } { 2 } }
	tokens := []Token{
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 1},
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_BLOCK_START},
		{Type: TOKEN_INT, Value: 2},
		{Type: TOKEN_BLOCK_END},
		{Type: TOKEN_BLOCK_END},
	}

	interp := executeTest(t, tokens)

	if interp.opStack.StackCount() != 1 {
		t.Errorf("Expected 1 item on stack, got %d", interp.opStack.StackCount())
	}

	top, _ := interp.opStack.Peek()
	proc, ok := top.(PSBlock)
	if !ok {
		t.Fatalf("Expected PSBlock on top of stack, got %T", top)
	}

	// Should have 6 tokens: { 1 } { 2 }
	if len(proc.Body) != 6 {
		t.Errorf("Expected procedure with 6 tokens, got %d", len(proc.Body))
	}
}
