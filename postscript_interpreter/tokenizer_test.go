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

// testing some of the helper functions ==============

func TestIsWhitespace(t *testing.T) {
	whitespaceChars := []byte{' ', '\n', '\t', '\r'}
	for _, ch := range whitespaceChars {
		if !IsWhitespace(ch) {
			t.Errorf("expected '%c' to be whitespace character", ch)
		}
	}

	// additionally testing non-whitespace characters
	nonWhitespaceChars := []byte{'a', '5', '%', '{'}
	for _, ch := range nonWhitespaceChars {
		if IsWhitespace(ch) {
			t.Errorf("'%c' is not whitespace character", ch)
		}
	}
}

func TestIsComment(t *testing.T) {
	// making sure % registers as a comment
	if !IsComment('%') {
		t.Error("expected '%' to be a comment character")
	}
	
	// additionally testing non-comment characters
	nonCommentChars := []byte{'a', ' ', '5', '\n'}
	for _, ch := range nonCommentChars {
		if IsComment(ch) {
			t.Errorf("'%c' is not comment character", ch)
		}
	}
}

func TestIsLetter(t *testing.T) {
	// testing expected letter characters
	letters := []byte{'a', 'z', 'A', 'Z', 'm'}
	for _, ch := range letters {
		if !IsLetter(ch) {
			t.Errorf("Expected '%c' to be an alphabetic character", ch)
		}
	}
	
	// testing non-letter characters
	nonLetters := []byte{'5', ' ', '%', '{', '0', '9'}
	for _, ch := range nonLetters {
		if IsLetter(ch) {
			t.Errorf("'%c' is not alphabetic character", ch)
		}
	}
}

func TestIsDigit(t *testing.T) {
	digits := []byte{'0', '5', '9'}
	for _, ch := range digits {
		if !IsDigit(ch) {
			t.Errorf("Expected '%c' to be a digit", ch)
		}
	}
	
	// testing non-digits
	nonDigits := []byte{'a', ' ', '%', '{'}
	for _, ch := range nonDigits {
		if IsDigit(ch) {
			t.Errorf("'%c' is not a digit", ch)
		}
	}
}

// testing tokenizing ============================

// parsing numbers
func TestTokenizeNumber(t *testing.T) {
	tests := []struct {
		input        string
		expectedType TokenType
		expectedVal  any
	}{
		{"42", TOKEN_INT, 42},
		{"3.14", TOKEN_FLOAT, 3.14},
		{"-5", TOKEN_INT, -5},
		{"-2.5", TOKEN_FLOAT, -2.5},
		{"0", TOKEN_INT, 0},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			tokenizer := CreateTokenizer(test.input)
			tokens, err := tokenizer.Tokenize()

			if err != nil {
				t.Fatalf("Tokenize error: %v", err)
			}
			if len(tokens) != 1 {
				t.Fatalf("Expected 1 token, got %d", len(tokens))
			}
			if tokens[0].Type != test.expectedType {
				t.Errorf("Expected type %v, got %v", test.expectedType, tokens[0].Type)
			}
			if tokens[0].Value != test.expectedVal {
				t.Errorf("Expected value %v, got %v", test.expectedVal, tokens[0].Value)
			}
		})
	}
}

// parsing strings
func TestTokenizeStrings(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"(hello world)", "hello world"},
		{"(test)", "test"},
		{"()", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			tokenizer := CreateTokenizer(test.input)
			tokens, err := tokenizer.Tokenize()

			if err != nil {
				t.Fatalf("Tokenize error: %v", err)
			}
			if len(tokens) != 1 {
				t.Fatalf("Expected 1 token, got %d", len(tokens))
			}
			if tokens[0].Type != TOKEN_STRING {
				t.Errorf("Expected TOKEN_STRING, got %v", tokens[0].Type)
			}
			if tokens[0].Value != test.expected {
				t.Errorf("Expected '%s', got '%v'", test.expected, tokens[0].Value)
			}
		})
	}
}

// parsing operators
func TestTokenizeOperators(t *testing.T) {
	input := "add sub mul"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("Tokenize error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("Expected 3 tokens, got %d", len(tokens))
	}

	expected := []string{"add", "sub", "mul"}
	for i, exp := range expected {
		if tokens[i].Type != TOKEN_OPERATOR {
			t.Errorf("Token %d: expected TOKEN_OPERATOR, got %v", i, tokens[i].Type)
		}
		if tokens[i].Value != exp {
			t.Errorf("Token %d: expected %s, got %v", i, exp, tokens[i].Value)
		}
	}
}

// parsing expressions
func TestTokenizeSimpleExpression(t *testing.T) {
	input := "3 4 add"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("Tokenize error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("Expected 3 tokens, got %d", len(tokens))
	}

	// checking first number
	if tokens[0].Type != TOKEN_INT {
		t.Errorf("Expected TOKEN_INT for first token, got %v", tokens[0].Type)
	}
	if tokens[0].Value != 3 {
		t.Errorf("Expected 3, got %v", tokens[0].Value)
	}

	// checking second number
	if tokens[1].Type != TOKEN_INT {
		t.Errorf("Expected TOKEN_INT for second token, got %v", tokens[1].Type)
	}
	if tokens[1].Value != 4 {
		t.Errorf("Expected 4, got %v", tokens[1].Value)
	}

	// checking third number
	if tokens[2].Type != TOKEN_OPERATOR {
		t.Errorf("Expected TOKEN_OPERATOR for third token, got %v", tokens[2].Type)
	}
	if tokens[2].Value != "add" {
		t.Errorf("Expected 'add', got %v", tokens[2].Value)
	}
}

// parsing procedures
func TestTokenizeProcedure(t *testing.T) {
	input := "{ dup mul }"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("Tokenize error: %v", err)
	}
	if len(tokens) != 4 { // braces count in tokens
		t.Fatalf("Expected 4 tokens, got %d", len(tokens))
	}

	if tokens[0].Type != TOKEN_BLOCK_START {
		t.Errorf("Expected TOKEN_BLOCK_START, got %v", tokens[0].Type)
	}
	if tokens[1].Type != TOKEN_OPERATOR || tokens[1].Value != "dup" {
		t.Errorf("Expected 'dup' operator")
	}
	if tokens[2].Type != TOKEN_OPERATOR || tokens[2].Value != "mul" {
		t.Errorf("Expected 'mul' operator")
	}
	if tokens[3].Type != TOKEN_BLOCK_END {
		t.Errorf("Expected TOKEN_BLOCK_END, got %v", tokens[3].Type)
	}
}

// parsing comments
func TestTokenizeComment(t *testing.T) {
	input := "3 % this is a comment\n4 add"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("Tokenize error: %v", err)
	}

	// comment should be ignored, resulting in: 3, 4, add
	if len(tokens) != 3 {
		t.Fatalf("Expected 3 tokens (comment ignored), got %d", len(tokens))
	}

	// cross checking expected tokens
	if tokens[0].Type != TOKEN_INT || tokens[0].Value != 3 {
		t.Errorf("Expected INT 3, got %v %v", tokens[0].Type, tokens[0].Value)
	}
	if tokens[1].Type != TOKEN_INT || tokens[1].Value != 4 {
		t.Errorf("Expected INT 4, got %v %v", tokens[1].Type, tokens[1].Value)
	}
	if tokens[2].Type != TOKEN_OPERATOR || tokens[2].Value != "add" {
		t.Errorf("Expected OPERATOR 'add', got %v %v", tokens[2].Type, tokens[2].Value)
	}
}

// parsing names
func TestTokenizeNames(t *testing.T) {
	input := "/x /myvar /test123"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("Tokenize error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("Expected 3 tokens, got %d", len(tokens))
	}

	expected := []PSName{"x", "myvar", "test123"}
	for i, exp := range expected {
		if tokens[i].Type != TOKEN_NAME {
			t.Errorf("Token %d: expected TOKEN_NAME, got %v", i, tokens[i].Type)
		}
		if tokens[i].Value != exp {
			t.Errorf("Token %d: expected %v, got %v", i, exp, tokens[i].Value)
		}
	}
}

// parsing booleans
func TestTokenizeBooleans(t *testing.T) {
	input := "true false"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("Tokenize error: %v", err)
	}
	if len(tokens) != 2 {
		t.Fatalf("Expected 2 tokens, got %d", len(tokens))
	}

	// true
	if tokens[0].Type != TOKEN_BOOL {
		t.Errorf("Expected TOKEN_BOOL for 'true', got %v", tokens[0].Type)
	}
	if tokens[0].Value != true {
		t.Errorf("Expected true, got %v", tokens[0].Value)
	}

	// false
	if tokens[1].Type != TOKEN_BOOL {
		t.Errorf("Expected TOKEN_BOOL for 'false', got %v", tokens[1].Type)
	}
	if tokens[1].Value != false {
		t.Errorf("Expected false, got %v", tokens[1].Value)
	}
}

func TestTokenizeEquals(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"=", "="},
		{"==", "=="},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			tokenizer := CreateTokenizer(test.input)
			tokens, err := tokenizer.Tokenize()

			if err != nil {
				t.Fatalf("Tokenize error: %v", err)
			}
			if len(tokens) != 1 {
				t.Fatalf("Expected 1 token, got %d", len(tokens))
			}
			if tokens[0].Type != TOKEN_OPERATOR {
				t.Errorf("Expected TOKEN_OPERATOR, got %v", tokens[0].Type)
			}
			if tokens[0].Value != test.expected {
				t.Errorf("Expected '%s', got '%v'", test.expected, tokens[0].Value)
			}
		})
	}
}
