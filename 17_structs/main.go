package main

import "fmt"

//ways to define struct

//1. Using type keyword
	type Person struct {
		name string
		age int
	}

//2. Using struct keyword  -- that is anonympus struct without type keyword, we can use it to create a struct without a name and assign it to a variable
	var Employee struct {
		name string
		age int
	}


//returning through value, means it will return a copy of the struct, not the original struct
func newPerson(name string, age int) Person {
	p:= Person{
		name: name,
		age: age,
	}
	return p
}

//returnign through pointer, means it will return the address of the struct, so any changes made to the struct through the pointer will affect the original struct
func newPersonPointer(name string, age int) *Person {
	p:= &Person{
		name: name,
		age: age,
	}
	fmt.Println("Address of p is: ", p)  // this will print like &{Jane 25} as not hexadecimal address 
	//itss Go's formatting representation of a pointer to a struct. It is not the actual hexadecimal memory address.

    //If you specifically want to print the actual address, use: %p
	fmt.Printf("Actual address of p is: %p\n", p)  // this will print the actual address in hexadecimal format
	return p
}



func main() {
	fmt.Println("Lets learn Structs in golang")

	//Structs are used to create complex data types that group together zero or more fields with varying data types.

	person1:= newPerson("John", 30)
	fmt.Println("Person1: ", person1)

	person2:= newPersonPointer("Jane", 25)
	fmt.Println("Person2: ", person2)

	//accessing struct fields
	fmt.Println("Person1 Name: ", person1.name)
	fmt.Println("Person1 Age: ", person1.age)
	
	//accessing struct fields through pointer
	fmt.Println("Person2 Name: ", person2.name)
	fmt.Println("Person2 Age: ", person2.age)

	//Omitted fields will be zero-valued.
	person3:= Person{
		name: "Alice",
	}
	fmt.Println("Person3: ", person3)


	//accessing and modifying struct fields
	Employee.name = "Alice"
	Employee.age = 25


	//Anonymous structs
	//here  what we are doing is we are creating a struct without a name and assigning it to a variable person4. 
	// This is useful when you need a one-off struct type that you don't want to define globally.
	//so note that we are not using the type keyword here, we are directly using the struct keyword to define the struct and then assigning it to the variable person4.
	person4:= struct {
		name string
		age int
	}{
		name: "Bob",
		age: 40,
	}
	fmt.Println("Person4: ", person4)

	//practicing 
	person5 := struct{name string
	city string}{
		name: "Charlie",
		city: "New York",
	}
	fmt.Println("Person5: ", person5)
	
	

}