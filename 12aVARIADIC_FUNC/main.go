package main

import "fmt"


//Variadic functions can be called with any number of trailing arguments. 
// For example, fmt.Println is a common variadic function.

func multiply(nums ...int) int {
	fmt.Println("Numbers are: ", nums)

	result:= 1

	for _,num :=range nums{
		result*=num
	}

	fmt.Println("Result is: ", result)
	return result
}

func main(){
	//1. Variadic functions can be called in the usual way with individual arguments.

	multiply(1,2,3,4,5)

	multiply(1,2,3)

	//2. Variadic functions can also be called with a slice of arguments using the ... operator.

	nums := []int{1,2,3,4,5}

	multiply(nums...)	
}