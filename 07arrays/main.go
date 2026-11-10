package main
//ARRAYS ARE USED VERY LESS IN GOLANG 
import "fmt"
func main()  {
	fmt.Println("Welcome to class on arrays")

	//arrays are used to store multiple values of the same type in a single variable. we can use arrays to store a list of values. we can also use arrays to store a list of structs or other complex data types.

	// var fruitsArray [4]string 

	// fruitsArray[0] = "Apple"
	// fruitsArray[1] = "Mango"
	// fruitsArray[3] = "Grapes" //missed index 2 so it will be empty

	// fmt.Println("Fruits array is: ", fruitsArray)
	// fmt.Println("Length of fruits array is: ", len(fruitsArray))

	// //another way to declare and initialize an array
	// var vegArray = [3]string{"Potato", "Tomato", "Onion"}

	// fmt.Println("Veggies array is: ", vegArray)
















	//when declaring array you have to provide the size of array compulsory
	var fruitList [4]string 
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Peach"
	// fruitList[3] = "Grapes"
	fmt.Println("FRuitList is: ", fruitList)
	//you will notice in the output that between each elemenmt tehre is a space
	//and after last element there is a space and then the closing bracket. this is because the array is of size 4 
	// and we have only assigned 3 values to it. so the last element is empty and it will be printed as a space.



	fmt.Println("length of fruitlist is :", len(fruitList)) 
	// impt: this will give the length of the array which is 4 in this case. even though we have only assigned 3 values to it, the length of the array is still 4 because we have declared it as an array of size 4.


	//SHORTHAND way to declare and initialize an array
	//ANother way to declare and initialize an array is to use the shorthand syntax. we can use the shorthand syntax to declare and initialize an array in a single line. we can also use the shorthand syntax to declare and initialize an array without specifying the size of the array. in this case, the size of the array will be determined by the number of values assigned to it.
	
	var vegList = []string{"potato","tomato","onion",12}
	fmt.Println("VegList is: ", vegList)
	fmt.Println("Length", len(vegList)) //this will give the length of the array which is 3 in this case. because we have assigned 3 values to it and we have not specified the size of the array. so the size of the array is determined by the number of values assigned to it.
}