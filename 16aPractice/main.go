package main 

import "fmt"

func main(){

	defer fmt.Println("Lets learn Defer in golang")
	defer fmt.Println("This will be printed second")
	defer fmt.Println("This will be printed first")

	a:=10

	if a>5{
		defer fmt.Println("This will be printed before all")
		return
	}

	defer fmt.Println("This will be printed fourth")

}