package main

import (
	"fmt"
	"strconv"
)

type TokenType int

// defining the types of token types for the interpreter
// essentially defining what each thing is and how to handle it
const (
	TOKEN_INT TokenType = iota // iota for numbering
	TOKEN_FLOAT
	TOKEN_STRING
	TOKEN_NAME
	TOKEN_OPERATOR
	TOKEN_BLOCK_START
	TOKEN_BLOCK_END
	TOKEN_BOOL
)

type Token struct {
	Type  TokenType
	Value interface{}
}

type Tokenizer struct {
	input string // input string
	pos   int    // individual character or position
}

func CreateTokenizer(input string) *Tokenizer {
	return &Tokenizer{input: input, pos: 0}
}

func (t *Tokenizer) Tokenize() ([]Token, error) {
	tokens := []Token{}

	for t.pos < len(t.input) { // position < end of the input
		t.skipWhitespace()
		if t.pos >= len(t.input) {
			break
		}

		ch := t.input[t.pos] // current character

		// depending on token type
		switch {
		case ch == '%': // ignore comment
			t.skipComment()
		case ch == '(': // string
			token, err := t.readString()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
		case ch == '{':
			// code block
			t.pos++
			tokens = append(tokens, Token{Type: TOKEN_BLOCK_START})
		case ch == '}':
			t.pos++
			tokens = append(tokens, Token{Type: TOKEN_BLOCK_END})
		case ch == '/':
			// variable logic
			token := t.readName()
			tokens = append(tokens, token)
		case IsDigit(ch) || (ch == '-' && t.pos+1 < len(t.input) && IsDigit(t.input[t.pos+1])):
			token := t.readNumber()
			tokens = append(tokens, token)

		case IsLetter(ch):
			token := t.readOperator()
			tokens = append(tokens, token)
		default:
			t.pos++
		}
	}

	return tokens, nil
}

// helper funcs

func (t *Tokenizer) skipWhitespace() {
	for t.pos < len(t.input) && IsWhitespace(t.input[t.pos]) {
		t.pos++
	}
}

func (t *Tokenizer) skipComment() {
	for t.pos < len(t.input) && t.input[t.pos] != '\n' {
		t.pos++
	}
}

func (t *Tokenizer) readString() (Token, error) {
	t.pos++        // for skipping the initial '('
	start := t.pos // starting place for the string
	for t.pos < len(t.input) && t.input[t.pos] != ')' {
		t.pos++ // moving forward until the closing parenthesis
	}

	if t.pos >= len(t.input) {
		return Token{}, fmt.Errorf("string unterminated")
	}

	// slicing the input
	value := t.input[start:t.pos]
	t.pos++

	return Token{Type: TOKEN_STRING, Value: value}, nil
}

func (t *Tokenizer) readName() Token {
	t.pos++        // skip initial '/'
	start := t.pos // start of character

	for t.pos < len(t.input) && IsLetter(t.input[t.pos]) || IsDigit(t.input[t.pos]) {
		t.pos++
	}
	name := t.input[start:t.pos]
	return Token{Type: TOKEN_NAME, Value: PSName(name)}
}


func (t *Tokenizer) readNumber() Token {
	start := t.pos
	hasDecimal := false

	if t.input[t.pos] == '-' {
		t.pos++
	}

	for t.pos < len(t.input) && (IsDigit(t.input[t.pos]) || t.input[t.pos] == '.') {
		if t.input[t.pos] == '.' {
			hasDecimal = true
		}
		t.pos++
	}
	numStr := t.input[start:t.pos]
	if hasDecimal {
		val, _ := strconv.ParseFloat(numStr, 64)
		return Token{Type: TOKEN_FLOAT, Value: val}
	} else {
		val, _ := strconv.Atoi(numStr)
		return Token{Type: TOKEN_INT, Value: val}
	}

}

// parses through operator and assigns value from the name
func (t *Tokenizer) readOperator() Token {
	start := t.pos

	for t.pos < len(t.input) && IsLetter(t.input[t.pos]) {
		t.pos++
	}

	op := t.input[start:t.pos]
	if op == "true" {
		return Token{Type: TOKEN_BOOL, Value: true}
	}
	if op == "false" {
		return Token{Type: TOKEN_BOOL, Value: false}
	}

	return Token{Type: TOKEN_OPERATOR, Value: op}
}

// Some helper functions to determine if a symbol is a comment, number, alpha value, or whitespace

func IsComment(ch byte) bool {
	return ch == '%'
}

func IsWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r'
}

func IsLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
