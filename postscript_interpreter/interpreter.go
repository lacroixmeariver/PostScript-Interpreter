package main

import (
	"fmt"
)

type Interpreter struct {
	opStack     *Stack                              // operand stack
	dictStack   []*PSDict                           // stack of dictionaries
	lexicalMode bool                                // for dynamic/lexical scoping
	operators   map[string]func(*Interpreter) error // map of operators and values
	quit        bool
}

// function acting like a constructor
func CreateInterpreter() *Interpreter {
	// initializing global dictionary
	globalDict := &PSDict{
		items:    make(map[string]PSConstant),
		capacity: 100,
	}

	// initializing interpreter
	interpreter := &Interpreter{
		opStack:     CreateStack(),
		dictStack:   []*PSDict{globalDict},
		lexicalMode: false,
		operators:   make(map[string]func(*Interpreter) error),
	}

	// populating operator dictionary with all the available operators
	interpreter.registerOperators()
	return interpreter
}

// creating dictionary of register operators and their associated functions
func (i *Interpreter) registerOperators() {

	// arithmetic
	i.operators["add"] = opAdd
	i.operators["sub"] = opSub
	i.operators["mul"] = opMul
	i.operators["div"] = opDiv
	i.operators["idiv"] = opIntdiv
	i.operators["mod"] = opMod
	i.operators["abs"] = opAbs
	i.operators["neg"] = opNeg
	i.operators["sqrt"] = opSqrt
	i.operators["ceiling"] = opCeil
	i.operators["floor"] = opFloor
	i.operators["round"] = opRound

	// stack manipulation
	i.operators["dup"] = opDup
	i.operators["pop"] = opPop
	i.operators["exch"] = opExch
	i.operators["clear"] = opClear
	i.operators["count"] = opCount

	// comparison
	i.operators["eq"] = opEq
	i.operators["ne"] = opNe
	i.operators["gt"] = opGt
	i.operators["ge"] = opGe
	i.operators["lt"] = opLt
	i.operators["le"] = opLe

	// boolean
	i.operators["and"] = opAnd
	i.operators["or"] = opOr
	i.operators["not"] = opNot
	i.operators["true"] = opTrue
	i.operators["false"] = opFalse

	// dictionary
	i.operators["dict"] = dOpDict
	i.operators["begin"] = dOpBegin
	i.operators["end"] = dOpEnd
	i.operators["def"] = dOpDef
	i.operators["length"] = dOpLength
	i.operators["maxlength"] = dOpMaxLength

	// flow control
	i.operators["if"] = opIf
	i.operators["ifelse"] = opIfElse
	i.operators["for"] = opFor
	i.operators["repeat"] = opRepeat
	i.operators["quit"] = opQuit
	i.operators["exec"] = opExec

	// input/output
	i.operators["print"] = opPrint
	i.operators["="] = opEquals
	i.operators["=="] = opEqualsEquals

	// string operations
	i.operators["get"] = opGet
	i.operators["getinterval"] = opGetInterval
	i.operators["putinterval"] = opPutInterval
}

// helper function to be able to search for a value through the dict stack
func (i *Interpreter) dictLookup(name string) (PSConstant, error) {
	index := len(i.dictStack) - 1
	for index >= 0 {
		if val, ok := i.dictStack[index].items[name]; ok {
			return val, nil
		}
		index--
	}
	return nil, fmt.Errorf("name undefined in dictionary stack")
}

// executes operation based on token type from list of tokens given as argument
func (i *Interpreter) Execute(tokens []Token) error {
	pos := 0

	for pos < len(tokens) {
		token := tokens[pos]

		if i.quit {
			break
		}
		switch token.Type {
		// if it's a value type, push it onto the stack
		case TOKEN_BOOL:
			i.opStack.Push(token.Value)

		case TOKEN_INT:
			i.opStack.Push(token.Value)

		case TOKEN_FLOAT:
			i.opStack.Push(token.Value)

		case TOKEN_STRING:
			i.opStack.Push(token.Value)

		case TOKEN_NAME:
			name := token.Value.(PSName)
			i.opStack.Push(name)

		// if it's an operator type, search for it in the dictionary
		case TOKEN_OPERATOR:
			val := token.Value.(string)
			opFunc, ok := i.operators[val]
			if ok {
				err := opFunc(i)
				if err != nil {
					return err
				}
			} else {
				value, err := i.dictLookup(val)
				if err != nil {
					return err
				}
				i.opStack.Push(value)
			}

		// if it's the start of a code block
		case TOKEN_BLOCK_START:
			procedure, newPos, err := i.buildProcedure(tokens, pos)
			if err != nil {
				return err
			}

			i.opStack.Push(procedure)
			pos = newPos

			continue

		// this shouldn't happen but if it hits the end of a code block an error is returned
		case TOKEN_BLOCK_END:
			return fmt.Errorf("how'd you get here?")
		}

		pos++
	}

	return nil
}

// the building of a code block/procedure 
func (i *Interpreter) buildProcedure(tokens []Token, startPos int) (PSBlock, int, error) {
	blockTokens := []Token{}

	depthCounter := 1 // for nested procedures
	pos := startPos + 1

	for depthCounter > 0 && pos < len(tokens) {
		currentToken := tokens[pos]

		if currentToken.Type == TOKEN_BLOCK_START {
			depthCounter++
		}

		if currentToken.Type == TOKEN_BLOCK_END {
			depthCounter--
			// depth counter being 0 means the procedure block is done
			if depthCounter == 0 {
				pos++
				break
			}
		}

		if depthCounter > 0 {
			blockTokens = append(blockTokens, currentToken)
		}
		pos++
	}

	if depthCounter > 0 {
		return PSBlock{}, pos, fmt.Errorf("unclosed procedure")
	}

	procedure := PSBlock{
		Body: blockTokens,
	}

	// adding snapshot to associated captured dictionary for lexical mode
	if i.lexicalMode {
		procedure.CapturedDict = i.dictStack[len(i.dictStack)-1]
	}

	return procedure, pos, nil
}
