// Entry point 
package main
import "fmt"

func main(){
	
	testStack := CreateStack()

	var num int 

	fmt.Println("Enter a number")
	fmt.Scan(&num)
	testStack.Push(num)
	testStack.Pop()
}

