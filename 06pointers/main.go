package main

import "fmt"

func main(){

	fmt.Println("Welcome to class on pointers")

	//pointers are used to store the address of a variable. its a direct reference to the memory location of a variable. we can use pointers to change the value of a variable in a function. we can also use pointers to pass large structs or arrays to functions without copying them.

	// var ptr *int 
	// fmt.Println("Value of pointer is: ", ptr) //this will print the value of the pointer which is nil as it is not pointing to any variable

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of pointer is: ", ptr) //this will print the address of the variable
	fmt.Println("Value of the variable is: ", *ptr) //this will print the value of the variable

	*ptr = *ptr+2
	fmt.Println("Value of the variable after increment is: ", *ptr)

}
