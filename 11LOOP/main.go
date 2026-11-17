package main 

import "fmt"

func main(){

	fmt.Println("Lets learn loops in golang")

	//loops are used to execute a block of code multiple times. we can use loops to iterate over a collection of data. 
	// we can also use loops to execute a block of code until a condition is met. 
	// we can also use loops to execute a block of code for a specific number of times.

	//declatrion of loop in golang is different from other programming languages. 

	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is: ", i)
	}

	//using range to iterate over arrays, slices, maps, strings, etc. 
	// range returns two values, the index and the value of the element at that index. we can use the index to access the value of the element at that index. 
	// we can also use the value to access the value of the element at that index.

	//using range to iterate over arrays
	var fruitsArray = [4]string{"Apple", "Mango", "Grapes", "Banana"}

	for index, value := range fruitsArray {
		fmt.Println("Index is: ", index, "Value is: ", value)
	}

	//using range to iterate over slices
	var fruitsSlice = []string{"Apple", "Mango", "Grapes", "Banana"}

	for index, value := range fruitsSlice {
		fmt.Println("Index is: ", index, "Value is: ", value)
	}
	
	//using range to iterate over maps
	var fruitsMap = map[string]string{"name": "Apple", "color": "Red", "taste": "Sweet"}	

	for key, value := range fruitsMap {
		fmt.Println("Key is: ", key, "Value is: ", value)
	}		

	//using range to iterate over strings
	var fruitsString = "Apple"

	for index, value := range fruitsString {
		fmt.Println("Index is: ", index, "Value is: ", string(value))
	}

	//We can ignore one or both the fields in a for-range loop by giving an _ instead of giving a variable name
	for index, _ := range fruitsArray {
		fmt.Println("Index is: ", index)
	}

	for _, value := range fruitsArray {
		fmt.Println("Value is: ", value)
	}
	

	//using break and continue in loops
	for i := 0; i < 5; i++ {
		if i == 3 {
			break //this will break the loop when i is 3
		}
		fmt.Println("Value of i is: ", i)
	}


	for i := 0; i < 5; i++ {
		if i == 3 {
			continue //this will skip the iteration when i is 3
		}
		fmt.Println("Value of i is: ", i)
	}

	//using nested loops
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("Value of i is: ", i, "Value of j is: ", j)
		}
	}

	//using goto in loops
	for i := 0; i < 5; i++ {
		if i == 3 {
			goto end //this will jump to the label end when i is 3
		}
		fmt.Println("Value of i is: ", i)
	}
end:
	fmt.Println("End of loop")

	



}