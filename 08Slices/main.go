package main

import "fmt"
import "sort"

func main()  {
	fmt.Println("Hello to class of Slices in golang");
	
	//slices are used to store multiple values of same type in a single variable. slices are dynamic in nature. we can add or remove elements from a slice. 
	// slices are built on top of arrays. slices are more flexible than arrays. slices are reference types. when we pass a slice to a function, we are passing the reference of the slice and not the copy of the slice.
	//different from arrays, slices do not have a fixed size. we can create a slice using the make function or by slicing an array. we can also create a slice using the []type{} syntax.
	
	var fruitSlice = []string{} //no length but do initialise
	 
	fmt.Printf("Type of fruitSlice is : %T \n", fruitSlice)

	var fruitSlice2 = []string{"Apple", "Mango", "Grapes"} //no length but do initialise
	 
	fmt.Printf("Type of fruitSlice is : %T \n", fruitSlice2)


	//no initialisuing like an array , but we can add elements to the slice using the append function. append function takes a slice and one or more values and returns a new slice with the values added to the end of the original slice.
	fruitSlice = append(fruitSlice, "Apple")
	fruitSlice = append(fruitSlice, "Mango")
	fruitSlice = append(fruitSlice, "Grapes")

	fmt.Println("Fruit slice is: ", fruitSlice)
	fmt.Println("Length of fruit slice is: ", len(fruitSlice))

	//another way to create a slice is by slicing an array. we can create a slice by specifying the start and end index of the array. the start index is inclusive and the end index is exclusive. the length of the slice is end index - start index.

	var fruitArray = [5]string{"Apple", "Mango", "Grapes", "Banana", "Orange"}

	var fruitSlice3 = fruitArray[1:4] //this will create a slice from index 1 to 3 of the array

	fmt.Println("Fruit slice 3 is: ", fruitSlice3)
	fmt.Println("Length of fruit slice 3 is: ", len(fruitSlice3))

	//we can also create a slice by specifying only the start index. the end index will be the length of the array. the length of the slice will be length of the array - start index.

	var fruitSlice4 = fruitArray[2:] //this will create a slice from index 2 to end of the array

	fmt.Println("Fruit slice 4 is: ", fruitSlice4)
	fmt.Println("Length of fruit slice 4 is: ", len(fruitSlice4))

	//we can also create a slice by specifying only the end index. the start index will be 0. the length of the slice will be :end index - 0.

	var fruitSlice5 = fruitArray[:3] //this will create a slice from index 0 to 2 of the array

	fmt.Println("Fruit slice 5 is: ", fruitSlice5)
	fmt.Println("Length of fruit slice 5 is: ", len(fruitSlice5))

	//note: last index is exclusive, so it will not include the element at index 3 but start index is inclusive, so it will include the element at index 2. so the length of the slice will be 3-0 = 3.



	//Using make() 
	var fruitSlice6 = make([]string, 4) //this will create a slice of length 4 and capacity 4. the length of the slice is the number of elements in the slice and the capacity of the slice is the number of elements that can be added to the slice before it needs to be resized.
	//so slice will get allocation for 4 elements 
	fruitSlice6[0] = "Apple"
	fruitSlice6[1] = "Mango"
	fruitSlice6[2] = "Grapes"
	fruitSlice6[3] = "Banana"

	//fruitSlice6[4] = "Orange" //this will give an error as the length of the slice is 4 and we are trying to add an element at index 4 which is out of bounds. so we need to use append function to add elements to the slice.

	//but using append we can add elements because it automatcially reallcoate the memory again

	fruitSlice6 = append(fruitSlice6, "Orange") //this will add the element to the slice and increase the length of the slice to 5 and capacity to 8. so it will reallocate the memory for 8 elements.

	fmt.Println("Fruit slice 6 is: ", fruitSlice6)
	fmt.Println("Length of fruit slice 6 is: ", len(fruitSlice6))
	fmt.Println("Capacity of fruit slice 6 is: ", cap(fruitSlice6)) //capacity is the number of elements that can be added to the slice before it needs to be resized.



	//using sort 
	sort.Strings(fruitSlice6) //this will sort the slice in ascending order. we need to import "sort" package to use this function.
	fmt.Println(fruitSlice6) //this will print the sorted slice

	//REMOVE an element from slice: THIS IS DONE USING APPEND 

	var subjects = []string{"English", "Hindi", "Maths", "Science"}
	var index int = 2
	subjects = append(subjects[:index], subjects[index+1:]...)
	fmt.Println(subjects)
}
