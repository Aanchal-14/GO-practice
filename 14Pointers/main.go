package main

import "fmt"

func zeroVal(num int) {
	num = 0 
}

func zeroPtr(num *int){
	*num = 0
}


func main(){
	x:=5

	zeroVal(x)
	fmt.Println("Value of x after zeroVal function: ", x)

	zeroPtr(&x)
	fmt.Println("Value of x after zeroPtr function: ", x)

	//creating pointer using new function. new function returns a pointer to a newly allocated zero value of the specified type.
	p:= new(int)
	fmt.Println("Value of p is: ", p)
	fmt.Println("Value of address stored in p is: ", *p)
	fmt.Println("Address of p is: ", &p )



}