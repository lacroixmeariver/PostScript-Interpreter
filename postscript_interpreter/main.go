// Entry point 
package main
import "fmt"

func main(){
	
	testStack := CreateStack()

	var num int 

	fmt.Println("Enter a number")
	fmt.Scan(&num)
	testStack.Push(num)
	fmt.Println("Stack empty?", testStack.IsEmpty())
	testStack.Peek()
	testStack.Pop()
	fmt.Println("Stack empty?", testStack.IsEmpty())
}

