package main

import (
	"testing"
)

func TestIsWhitespace(t *testing.T) {
	result := IsWhitespace(' ')
	if result != true {
		t.Fatal("whitespace not detected")
	}
}

func TestIsComment(t *testing.T) {
	result := IsComment('%')
	if result != true {
		t.Fatal("comment not detected")
	}
}

func TestIsLetter(t *testing.T) {
	result := IsLetter('a')
	if result != true {
		t.Fatal("letter not detected")
	}
}

func TestTokenizeNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"42", 42},
		{"3.14", 3.14},
		{"-5", -5},
		{"-2.5", -2.5},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			tokenizer := CreateTokenizer(tt.input)
			tokens, err := tokenizer.Tokenize()

			if err != nil {
				t.Fatalf("tokenizer error")
			}
			if len(tokens) != 1 {
				t.Fatalf("expected 1 token")
			}
			if tokens[0].Value != tt.expected {
				t.Errorf("unexpected value")
			}

		})
	}

}

func TestTokenizeStrings(t *testing.T) {
	input := "(hello world)"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("tokenize error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Type != TOKEN_STRING {
		t.Errorf("expected STRING token")
	}
	if tokens[0].Value != "hello world" {
		t.Errorf("expected 'hello world', got '%v'", tokens[0].Value)
	}
}

func TestTokenizeOperators(t *testing.T) {
	input := "add sub mul"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("tokenize error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}

	expected := []string{"add", "sub", "mul"}
	for i, exp := range expected {
		if tokens[i].Value != exp {
			t.Errorf("token %d: expected %s, got %v", i, exp, tokens[i].Value)
		}
	}
}

func TestTokenizeSimpleExpression(t *testing.T) {
	input := "3 4 add"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("tokenize error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}

	// Check values
	if tokens[0].Value != 3 {
		t.Errorf("expected 3, got %v", tokens[0].Value)
	}
	if tokens[1].Value != 4 {
		t.Errorf("expected 4, got %v", tokens[1].Value)
	}
	if tokens[2].Value != "add" {
		t.Errorf("expected add, got %v", tokens[2].Value)
	}
}

func TestTokenizeProcedure(t *testing.T) {
	input := "{ dup mul }"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("tokenize error: %v", err)
	}

	if len(tokens) != 4 {
		t.Fatalf("expected 4 tokens, got %d", len(tokens))
	}

	if tokens[0].Type != TOKEN_BLOCK_START {
		t.Errorf("expected PROC_START")
	}
	if tokens[3].Type != TOKEN_BLOCK_END {
		t.Errorf("expected PROC_END")
	}
}

func TestTokenizeComment(t *testing.T) {
	input := "3 % this is a comment\n4 add"
	tokenizer := CreateTokenizer(input)
	tokens, err := tokenizer.Tokenize()

	if err != nil {
		t.Fatalf("tokenize error: %v", err)
	}

	// Should ignore comment, get: 3, 4, add
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens (comment ignored), got %d", len(tokens))
	}
}
