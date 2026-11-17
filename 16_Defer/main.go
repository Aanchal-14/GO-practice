package main 
import "fmt"
func main(){

	fmt.Println("Lets learn Defer in golang")

	// Defer is used to delay the execution of a function until the surrounding function returns.

	// Defer is often used to perform cleanup actions, such as closing a file or releasing a lock.


	//so bacially teh defer ststamnet wil run when teh code will be about to terminate
	defer fmt.Println("This will be printed at the end of the function")

	fmt.Println("This will be printed first")


	//if there are multiple defer statements, they will be executed in LIFO (Last In First Out) order.
	defer fmt.Println("This will be printed second")
	defer fmt.Println("This will be printed third")
	
	
	//even during crash of the program/ error, defer statements will be executed before the program terminates.

//NOTE : will learn defer wg.Done() and defer mu.Unlock() later when we learn go routines and mutexes

}