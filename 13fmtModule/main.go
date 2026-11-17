package main 

import "fmt"

func main()  {
	fmt.Println("Hello to class of fmt module in golang");	

	//fmt package is used to format the output. it is used to print the output to the console.
	//  it is also used to read the input from the console. it is also used to format the string. it is also used to format the number. 
		//  it is also used to format the channel value. it is also used to format the error value.

	fmt.Println("Hello to class of fmt module in golang");	//Println is used to print the output to the console. it adds a new line at the end of the output. it is used to print the output to the console. it adds a new line at the end of the output.
	fmt.Print("Hello to class of fmt module in golang");	//Print is used to print the output to the console. it does not add a new line at the end of the output. it is used to print the output to the console. it does not add a new line at the end of the output.

	//printf is formatting function. it is used to format the output. 
	name:="John"
	age:=30
	fmt.Printf("My name is %s and I am %d years old", name, age);

	//fmt.Sprintf() is used to format the string and return it. it is used to format the string without printing it.
	result := fmt.Sprintf("My name is %s and I am %d years old", name, age);
	fmt.Println(result);

	//input is read using fmt.Scanf()
	var inputName string
	var inputAge int
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &inputName)
	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &inputAge)
	fmt.Printf("My name is %s and I am %d years old", inputName, inputAge)

	//more ways ti take input from user using fmt.Scan() and fmt.Scanln()
	var inputName2 string
	var inputAge2 int
	fmt.Print("Enter your name: ")
	fmt.Scan(&inputName2)
	fmt.Print("Enter your age: ")
	fmt.Scan(&inputAge2)
	fmt.Printf("My name is %s and I am %d years old", inputName2, inputAge2)

	var inputName3 string
	var inputAge3 int
	fmt.Print("Enter your name: ")
	fmt.Scanln(&inputName3)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&inputAge3)
	fmt.Printf("My name is %s and I am %d years old", inputName3, inputAge3)
}