package main 

import "fmt"

//This function intSeq returns another function, which we define anonymously in the body of intSeq. 
// The returned function closes over the variable i to form a closure.

func intSeq() func() int {

	i:=0;

	return func() int{
		i++;
		fmt.Println(i)
		return i
	}
}



func main(){

	//Go supports anonymous functions, which can form closures. 
	// Anonymous functions are useful when you want to define a function inline without having to name it.

	nextInt:= intSeq()

	fmt.Println("Lets learn closures in golang")

	//Each call to nextInt generates a new unique function with its own i value. 
	// We can call nextInt multiple times and each time it will return a different value of i.

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//To demonstrate that the state is unique to that particular function, we can create a new one and see that it does not share state with the previous one.
	
	newInts:= intSeq()
	fmt.Println("-=--=-=-",newInts())


}