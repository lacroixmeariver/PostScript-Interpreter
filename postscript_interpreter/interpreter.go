package main

import (
	"fmt"
)

type Interpreter struct {
	opStack     *Stack                              // operand stack
	dictStack   []*PSDict                           // stack of dictionaries
	lexicalMode bool                                // for dynamic/lexical scoping
	operators   map[string]func(*Interpreter) error // map of operators and values
}

// function acting like a constructor
func CreateInterpreter() *Interpreter {
	globalDict := & PSDict{
		items: make(map[string]PSConstant),
		capacity: 100,
	}

	interpreter := &Interpreter{
		opStack:     CreateStack(),
		dictStack:   []*PSDict{globalDict},
		lexicalMode: false,
		operators:   make(map[string]func(*Interpreter) error),
	}

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
    i.operators["idiv"] = opIdiv
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
	
	// TODO: Add missing operators 
}

func (i *Interpreter) Execute(tokens []Token) error {
	pos := 0
	for pos < len(tokens) {
		token := tokens[pos]
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
			val := token.Value
			name := PSName(val.(string))
			i.opStack.Push(name)

		// if it's an operator type, search for it in the dictionary
		case TOKEN_OPERATOR:
			val := token.Value.(string)
			opFunc, ok := i.operators[val]
			if !ok {
				return fmt.Errorf("operator not defined")
			}
			err := opFunc(i)
			if err != nil {
				return err
			}

		// if it's the start of a code block
		case TOKEN_BLOCK_START:
			proc, newPos, err := i.buildProcedure(tokens, pos)
			if err != nil {
				return err
			}

			i.opStack.Push(proc)
			pos = newPos

		// this shouldn't happen but if it hits the end of a code block
		case TOKEN_BLOCK_END:
			return fmt.Errorf("how'd you get here")
		}

		pos++
	}

	return nil
}

func (i *Interpreter) buildProcedure(tokens []Token, startPos int) (PSBlock, int, error) {
	blockTokens := []Token{}
	depthCounter := 1
	pos := startPos + 1

	for depthCounter > 0 && pos < len(tokens) {
		currentToken := tokens[pos]

		if currentToken.Type == TOKEN_BLOCK_START {
			depthCounter++
		}

		if currentToken.Type == TOKEN_BLOCK_END {
			depthCounter--
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

	proc := PSBlock{
		Body: blockTokens,
	}

	return proc, pos, nil
}
