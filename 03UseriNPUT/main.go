package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Hello to User Input"
	fmt.Println(welcome)


	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")
	input , _ := reader.ReadString('\n')  //this will take the input from the user and store it in the input variable
	//the _ is used to ignore the error value returned by the ReadString function
	//if we had not used the _ then we would have to handle the error value returned by the ReadString function
	
	fmt.Println("Thanks for the rating, ", input)

	fmt.Printf("Type of input is %T \n",input)

}
