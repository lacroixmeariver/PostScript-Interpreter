# PostScript-Interpreter

## Project description
A PostScript interpreter implementation written in Go. Supports a subset of PostScript commands with both dynamic and lexical scoping modes. Intended to run as a REPL for a command line application. 

## Some features
- **Commands include:**
- Stack manipulation
- Arithmetic operations
- Dictionary operations
- String operations
- Boolean operations
- Flow control
- Input/output

- **Dual scoping**
- Dynamic by default
- Lexical by use of flag

## Requirements
- Go 1.23 or higher

## Build and Run 
Ensure you are in the correct `postscript_interpreter` directory by entering: `cd postscript_interpreter` \
then to build: `go build` \
and run using: `go run .` - for dynamic scoping (default setting) \
`go run . -lex` for lexical scoping

## General REPL info
The number displayed in REPL parenthesis: `PS (#)>` represents number of items in operand stack. \
To **exit** the REPL, type `quit`

## Run Tests
Ensure you are in the correct `postscript_interpreter` directory by entering: `cd postscript_interpreter`\
then to run all tests in verbose mode: `go test -v` \
to check test coverage: `go test -cover`

## Supported commands
Coming soon.. 

