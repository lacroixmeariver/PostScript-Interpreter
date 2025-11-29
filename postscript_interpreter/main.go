// Entry point
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	lexicalFlag := flag.Bool("lex", false, "Use lexical scoping") // for switching to lexical mode 
	flag.Parse()

	mainInterpreter := CreateInterpreter()
	mainInterpreter.lexicalMode = *lexicalFlag
	// scoping mode for displaying on startup
	scopingMode := "Dynamic scoping mode"
	if *lexicalFlag {
		scopingMode = "Lexical scoping mode"
	}

	// ascii art to make it look fancy
	printWelcomeScreen()
	printScopingMode(scopingMode)

	// scan std input
	scanner := bufio.NewScanner(os.Stdin)

	// the actual REPL loop
	for {
		fmt.Printf("\nPS (%d)> ", mainInterpreter.opStack.StackCount()) // for displaying stack count
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		
		// stylized help guide for available commands
		if input == "commands" {
			printREPLCommands()
			continue
		}

		// input string as argument
		tokenizer := CreateTokenizer(input)
		tokens, err := tokenizer.Tokenize()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		err = mainInterpreter.Execute(tokens)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		// catching the quit flag 
		if mainInterpreter.quit {
			fmt.Println("\nExiting...")
			break
		}
	}
}

func printScopingMode(mode string) {

	// print current scoping mode to terminal
	fmt.Println("        ╭─────────────────────────────────────────────────────────────╮")
    fmt.Printf("        │           Currently in: %-36s│\n", mode)
    fmt.Println("        ╰─────────────────────────────────────────────────────────────╯")
    fmt.Println()

}

func printWelcomeScreen() {
	fmt.Println(`
	╭─────────────────────────────────────────────────────────────╮
	│                                                             │
	│   ██████╗ ███████╗                                          │
	│   ██╔══██╗██╔════╝          PostScript Interpreter          │
	│   ██████╔╝███████╗          ----------------------          │
	│   ██╔═══╝ ╚════██║                                          │
	│   ██║     ███████║                                          │
	│   ╚═╝     ╚══════╝                                          │
	│                                                             │
	│   Ingrid Llorente                                           │
	│   Washington State University - CptS 355 - Fall 2025        │
	│   - Type 'quit' to exit                                     │
	│   - To enable lexical scoping mode, run: 'go run . -lex'    │
	│   - For command reference type 'commands'                   │
	│                                                             │
	╰─────────────────────────────────────────────────────────────╯
	`)
}

// stylized help page for quick user reference 
func printREPLCommands() {
    fmt.Println(`
	╭─────────────────────────────────────────────────────────────╮
	│                   AVAILABLE COMMANDS                        │
	╰─────────────────────────────────────────────────────────────╯

	ARITHMETIC OPERATORS (12):
	add          num1 num2 → sum           5 3 add = → 8
	sub          num1 num2 → difference    10 3 sub = → 7
	mul          num1 num2 → product       4 5 mul = → 20
	div          num1 num2 → quotient      20 4 div = → 5.0
	idiv         int1 int2 → quotient      7 2 idiv = → 3
	mod          int1 int2 → remainder     10 3 mod = → 1
	abs          num → |num|               -5 abs = → 5
	neg          num → -num                5 neg = → -5
	sqrt         num → √num                16 sqrt = → 4
	ceiling      num → ⌈num⌉               3.2 ceiling = → 4.0
	floor        num → ⌊num⌋               3.8 floor = → 3.0
	round        num → rounded             3.5 round = → 4.0

	STACK MANIPULATION (5):
	dup          any → any any            5 dup → [5, 5]
	pop          any → -                  5 pop → []
	exch         a b → b a                1 2 exch → [2, 1]
	clear        any... → -               Clear entire stack
	count        any... → any... n        Push stack size

	COMPARISON OPERATORS (6):
	eq           a b → bool               5 5 eq = → true
	ne           a b → bool               5 3 ne = → true
	gt           a b → bool               5 3 gt = → true
	ge           a b → bool               5 5 ge = → true
	lt           a b → bool               3 5 lt = → true
	le           a b → bool               3 5 le = → true

	BOOLEAN OPERATORS (5):
	and          bool1 bool2 → bool       true false and = → false
	or           bool1 bool2 → bool       true false or = → true
	not          bool → bool              true not = → false
	true         - → true                 Push true
	false        - → false                Push false

	DICTIONARY OPERATIONS (6):
	dict         int → dict               10 dict (create dict)
	begin        dict → -                 Start using dictionary
	end          - → -                    Stop using dictionary
	def          key val → -              /x 5 def (define x=5)
	length       dict → int               dict length = (entry count)
	maxlength    dict → int               dict maxlength = (capacity)

	STRING OPERATIONS (3):
	get          str idx → int            (hello) 0 get = → 104
	getinterval  str idx cnt → substr     (hello) 1 3 getinterval =
	putinterval  str1 idx str2 → str      (hello) 1 (XY) putinterval =

	FLOW CONTROL (6):
	if           bool proc → -            5 3 gt {(yes) print} if
	ifelse       bool p1 p2 → -           true {1} {2} ifelse exec =
	for          j k l proc → -           0 1 5 {} for (0 to 5)
	repeat       n proc → -               3 {(hi) print} repeat
	exec         proc → -                 {1 2 add} exec = → 3
	quit         - → -                    Exit interpreter

	I/O OPERATIONS (3):
	print        str → -                  (hello) print
	=            any → -                  42 = (print with newline)
	==           any → -                  (test) == (show as (test))

	SPECIAL COMMANDS:
	commands     Show this command list
	quit         Exit the interpreter

	EXAMPLES:

	Variables:
		/x 5 def              Define x = 5
		/y 10 def             Define y = 10
		x y add =             Prints 15

	Procedures:
		/square {dup mul} def
		5 square exec =       Prints 25

	Conditionals:
		5 3 gt {(bigger)} {(smaller)} ifelse print

	Loops:
		1 1 10 {dup mul =} for    Squares from 1 to 10

	Scoping:
		/x 1 def
		/show {x =} def
		10 dict begin /x 2 def show exec end
		(Dynamic: 2, Lexical: 1)

	╰─────────────────────────────────────────────────────────────╯
	`)
}