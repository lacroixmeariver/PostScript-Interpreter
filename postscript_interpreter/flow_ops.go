package main

import "fmt"

// ======================================== flow control operators

// executes procedure if leading bool val is true
func opIf(i *Interpreter) error {
	proc, _ := i.opStack.Pop()
	boolVar, _ := i.opStack.Pop()

	// converting stack object to procedure and bool var to boolean value
	procedure := proc.(PSBlock)
	conditionalBool := boolVar.(bool)

	if conditionalBool {

		// check for lexical scoping mode
		// if true saves + uses current dict and executes procedure with captured environment
		if i.lexicalMode && procedure.CapturedDict != nil{ // check to make sure env isn't empty
			savedStack := i.dictStack 
			i.dictStack = []*PSDict{procedure.CapturedDict} 
			err := i.Execute(procedure.Body)
			i.dictStack = savedStack 
			return err
		// if not in lexical mode just execute procedure looking through most recent dictionary
		} else {  
			return i.Execute(procedure.Body)
		}
	}

	return nil
}

// executes procedure #1 if leading bool val is true, executes procedure #2 otherwise
func opIfElse(i *Interpreter) error {
	proc2, _ := i.opStack.Pop()
	proc1, _ := i.opStack.Pop()
	boolVar, _ := i.opStack.Pop()

	// converting values
	procedure1 := proc1.(PSBlock)
	procedure2 := proc2.(PSBlock)
	conditionalBool := boolVar.(bool)

	// first conditional block
	if conditionalBool { 
		if i.lexicalMode && procedure1.CapturedDict != nil{ 
			// lexical mode
			savedStack := i.dictStack
			i.dictStack = []*PSDict{procedure1.CapturedDict}
			err := i.Execute(procedure1.Body)
			i.dictStack = savedStack
			return err
		// dynamic scoping
		} else { 
			return i.Execute(procedure1.Body) 
		}
	// second conditional block
	} else {
		if i.lexicalMode && procedure2.CapturedDict != nil{
			savedStack := i.dictStack
			i.dictStack = []*PSDict{procedure2.CapturedDict}
			err := i.Execute(procedure2.Body)
			i.dictStack = savedStack
			return err
		} else {
			return i.Execute(procedure2.Body)
		}
	}
}

// executes procedure in a loop according to a start/stop/step index
func opFor(i *Interpreter) error {
	proc, _ := i.opStack.Pop()
	endVar, _ := i.opStack.Pop()
	stepVar, _ := i.opStack.Pop()
	startVar, _ := i.opStack.Pop()

	// converting + initializing counter variable
	procedure := proc.(PSBlock) // func to be executed
	step := stepVar.(int) // the number by which count is incremented
	start := startVar.(int) // starting index
	end := endVar.(int) // ending index
	counter := start

	if step > 0 {
        for counter <= end { // for loops inclusive to end number in PS
            i.opStack.Push(counter)
            
			// lexical mode 
            if i.lexicalMode && procedure.CapturedDict != nil {
                savedStack := i.dictStack
                i.dictStack = []*PSDict{procedure.CapturedDict}
                err := i.Execute(procedure.Body)
                i.dictStack = savedStack
                if err != nil {
                    return fmt.Errorf("opfor failed: %v", err)
                }
			// dynamic mode
            } else {
                err := i.Execute(procedure.Body)
                if err != nil {
                    return fmt.Errorf("opfor failed: %v", err)
                }
            }
            
            counter = counter + step
        }
	// accounting for step values < 0
    } else {
        for counter >= end {
            i.opStack.Push(counter)
            
			// lexical mode
            if i.lexicalMode && procedure.CapturedDict != nil {
                savedStack := i.dictStack
                i.dictStack = []*PSDict{procedure.CapturedDict}
                err := i.Execute(procedure.Body)
                i.dictStack = savedStack
                if err != nil {
                    return fmt.Errorf("opfor failed: %v", err)
                }
			// dynamic mode
            } else {
                err := i.Execute(procedure.Body)
                if err != nil {
                    return fmt.Errorf("opfor failed: %v", err)
                }
            }
            
            counter = counter + step
        }
	}
		
	return nil
}

// repeats a procedure n times 
func opRepeat(i *Interpreter) error {

	proc, _ := i.opStack.Pop()
	n, _ := i.opStack.Pop()

	counter := 1 // counting loops
	procedure := proc.(PSBlock) // func to be executed
	num := n.(int) // stop index

	for counter <= num {

		// lexical mode
		if i.lexicalMode && procedure.CapturedDict != nil {
			savedStack := i.dictStack
			i.dictStack = []*PSDict{procedure.CapturedDict}
			err := i.Execute(procedure.Body)
			i.dictStack = savedStack
			if err != nil {
				return fmt.Errorf("opfor failed: %v", err)
			}
		// dynamic mode
		} else {
			err := i.Execute(procedure.Body)
			if err != nil {
				return fmt.Errorf("oprepeat failed: %v", err)
			}
		}
		counter ++
	}

	return nil
}

// quits the application
func opQuit(i *Interpreter) error {

	i.quit = true
	return nil
}

// executes some arbitrary object/procedure 
func opExec(i* Interpreter) error {
	proc, _ := i.opStack.Pop()

    procedure, ok := proc.(PSBlock)
    if !ok {
        return fmt.Errorf("exec requires procedure")
    }
    
	// lexical mode
    if i.lexicalMode && procedure.CapturedDict != nil {
        savedStack := i.dictStack
        i.dictStack = []*PSDict{procedure.CapturedDict}
        err := i.Execute(procedure.Body)
        i.dictStack = savedStack
        return err
	// dynamic mode
    } else {
        return i.Execute(procedure.Body)
    }
}