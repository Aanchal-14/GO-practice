
package main

import "fmt"

//declaring a function in golang is different from other programming languages.
func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) (result int) {
	result = a * b
	return
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func power(base int, exponent int) (result int) {
	result = 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return
}

func sum(numbers ...int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

// we can also specify closures in golang. closures are functions that are defined inside another function. closures can access the variables of the outer function. closures can also be returned from a function.
func outerFunction() func(int) int {
	return func(x int) int {
		return x * x
	}
}

// we can also specify recursive functions in golang. recursive functions are functions that call themselves.
// recursive functions can be used to solve problems that can be broken down into smaller subproblems.
// recursive functions can also be used to solve problems that can be solved using divide and conquer approach.

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// we can also specify higher order functions in golang.
// higher order functions are functions that take other functions as parameters or return other functions as return values.

func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// we can also specify function types in golang. function types are used to define the type of a function. we can use function types to define the type of a function that takes other functions as parameters or returns other functions as return values.

type operation func(int, int) int

func applyOperation2(a int, b int, op operation) int {
	return op(a, b)
}

//VERY IMPT -- WE CAANOT DEFINE A FUNCTION INSIDE ANOTHER FUNCTION IN GOLANG. 
				// WE CAN ONLY DEFINE A FUNCTION INSIDE ANOTHER FUNCTION USING ANONYMOUS FUNCTIONS OR CLOSURES.

func main() {
	fmt.Println("Hello to class of functions in golang")

	//functions are used to execute a block of code multiple times. we can use functions to execute a block of code until a condition is met.

	result := add(5, 10)
	fmt.Println("Result of addition is: ", result)

	// we can declare a function using the func keyword followed by the name of the function and the parameters of the function.

	result2 := subtract(10, 5)
	fmt.Println("Result of subtraction is: ", result2)

	// we can also specify the return type of the function.

	result3 := multiply(5, 10)
	fmt.Println("Result of multiplication is: ", result3)

	// we can also specify multiple return types of the function.

	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient is: ", quotient, "Remainder is: ", remainder)

	// we can also specify named return types of the function.

	result4 := power(2, 3)
	fmt.Println("Result of power is: ", result4)

	// we can also specify variadic parameters of the function.

	result5 := sum(1, 2, 3, 4, 5)
	fmt.Println("Result of sum is: ", result5)

	// we can also specify anonymous functions in golang.
	func() {
		fmt.Println("Hello from anonymous function")
	}() 
	
	//this is called immediately invoked function expression (IIFE) in golang. we can also assign an anonymous function to a variable and call it later.

	square := outerFunction()
	result6 := square(5)
	fmt.Println("Result of square is: ", result6)

	result7 := factorial(5)
	fmt.Println("Result of factorial is: ", result7)

	result8 := fibonacci(10)
	fmt.Println("Result of fibonacci is: ", result8)

	addition := func(a int, b int) int {
		return a + b
	}

	result9 := applyOperation(5, 10, addition)
	fmt.Println("Result of addition using higher order function is: ", result9)

	subtraction := func(a int, b int) int {
		return a - b
	}

	result10 := applyOperation2(10, 5, subtraction)
	fmt.Println("Result of subtraction using higher order function is: ", result10)

	// we can also specify function literals in golang. function literals are functions that are defined without a name. we can use function literals to define functions that are used only once. we can also use function literals to define functions that are used as parameters or return values of other functions.

	result11 := applyOperation2(5, 10, func(a int, b int) int {
		return a * b
	})

	fmt.Println("Result of multiplication using function literal is: ", result11)

	// we can also specify function expressions in golang. function expressions are functions that are defined as expressions. we can use function expressions to define functions that are used only once. we can also use function expressions to define functions that are used as parameters or return values of other functions.

	result12 := applyOperation2(10, 5, func(a int, b int) int {
		return a / b
	})

	fmt.Println("Result of division using function expression is: ", result12)
}
