package main

import "fmt"

func main() {
	fmt.Println("Lets learn if else in golang")



	//if else is used to execute a block of code based on a condition. 
	// if the condition is true, the block of code inside the if statement will be executed. 
	// if the condition is false, the block of code inside the else statement will be executed. 
	// we can also use else if to check multiple conditions. we can also use switch case to check multiple conditions.

	if 7%2 == 0 {
		fmt.Println("7 is even number")
	} else {
		fmt.Println("7 is odd number")
	}

	//IMPT -- make sure else is in the same line as the closing bracket of if statement. otherwise it will give error.
	
	//Nested if else is used to check multiple conditions. we can use nested if else to check multiple conditions. we can also use switch case to check multiple conditions.
	var num = 10

	if num%2 == 0 {
		fmt.Println("10 is even number")
	} else if num%3 == 0 {
		fmt.Println("10 is divisible by 3")
	} else {
		fmt.Println("10 is odd number and not divisible by 3")
	}

	// we can also declare a variable inside the if statement. the scope of the variable will be limited to the if statement. 
	if num := 15; num%5 == 0 {
		fmt.Println("15 is divisible by 5")
	} else {
		fmt.Println("15 is not divisible by 5")
	}

}