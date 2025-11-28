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

// defining the structure of a token
type Token struct {
	Type  TokenType
	Value any
}

// defining structure of actual tokenizer
type Tokenizer struct {
	input string // input string
	pos   int    // individual position within token
}

// constructor
func CreateTokenizer(input string) *Tokenizer {
	return &Tokenizer{input: input, pos: 0}
}

// tokenize function for breaking up input into 
// tokens interpreter will recognize
func (t *Tokenizer) Tokenize() ([]Token, error) {
	tokens := []Token{}

	for t.pos < len(t.input) { // position < end of the input
		t.skipWhitespace()
		if t.pos >= len(t.input) {
			break
		}

		currentChar := t.input[t.pos]

		// depending on token type
		switch {

		case currentChar == '%': // ignore comment
			t.skipComment() // helper function to skip it

		case currentChar == '(': // string
			token, err := t.readString() // helper function to read strings

			if err != nil {
				return nil, err
			}

			tokens = append(tokens, token)

		case currentChar == '{': // start of code block
			t.pos++
			// recognized as start token and appended
			tokens = append(tokens, Token{Type: TOKEN_BLOCK_START}) 

		case currentChar == '}': // end of code block
			t.pos++
			// recognized and appended
			tokens = append(tokens, Token{Type: TOKEN_BLOCK_END})

		case currentChar == '/':
			// variable logic
			token := t.readName() // helper function to read name without '\'
			tokens = append(tokens, token)

		case currentChar == '=': // recognizing '=' and '==' as operators 
			t.pos++
			if t.pos < len(t.input) && t.input[t.pos] == '=' {
				t.pos++
				tokens = append(tokens, Token{Type: TOKEN_OPERATOR, Value: "=="})
			} else {
				tokens = append(tokens, Token{Type: TOKEN_OPERATOR, Value: "="})
			}

		case IsDigit(currentChar) || (currentChar == '-' && t.pos+1 < len(t.input) && IsDigit(t.input[t.pos+1])):
			token := t.readNumber()
			tokens = append(tokens, token)

		case IsLetter(currentChar):
			token := t.readWord()
			tokens = append(tokens, token)

		default:
			t.pos++
		}
	}

	return tokens, nil
}

// tokenizer helper functions =================================================

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

	// returning it as a string type token
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

// parses through word and assigns value from the name
func (t *Tokenizer) readWord() Token {
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

// some additional helper functions to determine if a symbol is a comment, number, alpha value, or whitespace

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
