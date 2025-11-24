package main

import "fmt"

// ======================================== Flow control operators

func opIf(i *Interpreter) error {

	p, _ := i.opStack.Pop()
	b, _ := i.opStack.Pop()
	proc := p.(PSBlock)
	conditionBool := b.(bool)
	if conditionBool {
		err := i.Execute(proc.Body)
		if err != nil {
			return fmt.Errorf("opif failed: %v", err)
		}
	}
	return nil
}

func opIfElse(i *Interpreter) error {
	p2, _ := i.opStack.Pop()
	p1, _ := i.opStack.Pop()
	b, _ := i.opStack.Pop()
	proc1 := p1.(PSBlock)
	proc2 := p2.(PSBlock)
	conditionBool := b.(bool)
	if conditionBool {
		err := i.Execute(proc1.Body)
		if err != nil {
			return fmt.Errorf("opifelse failed: %v", err)
		}
	} else {
		err := i.Execute(proc2.Body)
		if err != nil {
			return fmt.Errorf("opifelse failed: %v", err)
		}

	}
	return nil
}

func opFor(i *Interpreter) error {

	_proc, _ := i.opStack.Pop()
	_end, _ := i.opStack.Pop()
	_step, _ := i.opStack.Pop()
	_start, _ := i.opStack.Pop()

	proc := _proc.(PSBlock)
	step := _step.(int)
	start := _start.(int)
	end := _end.(int)
	counter := start

	if step > 0 {
		for counter <= end {
			i.opStack.Push(counter)
			err := i.Execute(proc.Body)
			if err != nil {
				return fmt.Errorf("opfor failed: %v", err)
			}
			counter = counter + step
		}
	} else {
		for counter >= end {
			i.opStack.Push(counter)
			err := i.Execute(proc.Body)
			if err != nil {
				return fmt.Errorf("opifelse failed: %v", err)
			}
			counter = counter + step
		}
	}

	return nil
}

func opRepeat(i *Interpreter) error {

	p, _ := i.opStack.Pop()
	n, _ := i.opStack.Pop()

	counter := 1
	proc := p.(PSBlock)

	for counter <= n.(int) {
		err := i.Execute(proc.Body)
		if err != nil {
			return fmt.Errorf("opifelse failed: %v", err)
		}
		counter++
	}

	return nil
}

func opQuit(i *Interpreter) error {

	i.quit = true
	return nil
}
