package main
import "fmt"


func main(){

	//Maps are used to store key-value pairs, where the keys are unique and are used to access the corresponding values. 

	//we access the values using the keys.

	//When we add a key-value pair with key that already exists in the map, the new value we set will replace the old value of the corresponding key.

	//. While arrays have a fixed size, in maps, we can add as many as key-value pairs as we want.


	//DECLARATION OF MAPS 

	//1.USING MAP LITERAL SYNTAX
	var person = map[string]string{"name": "John", "age": "30", "city": "New York"}
	fmt.Println("Person map is: ", person)
	fmt.Println("Length of person map is: ", len(person))

	//2.USING MAKE FUNCTION
	var person2 = make(map[string]string)
	person2["name"] = "Jane"
	person2["age"] = "25"
	person2["city"] = "Los Angeles"
	
	fmt.Println("Person2 map is: ", person2)
	fmt.Println("Length of person2 map is: ", len(person2))

	//3.USING VAR KEYWORD
	var person3 map[string]string
	person3 = make(map[string]string)
	person3["name"] = "Mike"
	person3["age"] = "35"
	person3["city"] = "Chicago"
	
	fmt.Println("Person3 map is: ", person3)
	fmt.Println("Length of person3 map is: ", len(person3))

	//ACCESSING VALUES FROM MAPS
	fmt.Println("Name of person is: ", person["name"])
	fmt.Println("Age of person is: ", person["age"])
	fmt.Println("City of person is: ", person["city"])

	//UPDATING VALUES IN MAPS
	person["age"] = "31"
	fmt.Println("Updated age of person is: ", person["age"])

	//DELETING VALUES FROM MAPS
	delete(person, "city")
	fmt.Println("Person map after deleting city is: ", person)
	fmt.Println("Length of person map after deleting city is: ", len(person))

	//CHECKING IF A KEY EXISTS IN MAPS
	
	//ok is a boolean value that indicates whether the key exists in the map or not. if the key exists, ok will be true and value will be the corresponding value of the key. 
	// if the key does not exist, ok will be false and value will be the zero value of the type of the value.
	value, ok := person["city"]
	if ok {
		fmt.Println("City of person is: ", value)
	} else {
		fmt.Println("City not found in person map") //value will be empty string because the key does not exist in the map
	}

	//ITERATING OVER MAPS
	for key, value := range person {
		fmt.Println("key -", key, "| value - ", value)
	}

	//NESTED MAPS
	var person4 = map[string]map[string]string{
		"person1": {"name": "John", "age": "30", "city": "New York"},
		"person2": {"name": "Jane", "age": "25", "city": "Los Angeles"},
		"person3": {"name": "Mike", "age": "35", "city": "Chicago"},
	}

	fmt.Println("Person4 map is: ", person4)
	fmt.Println("Length of person4 map is: ", len(person4))

	for key, value := range person4 {
		fmt.Println("key -", key, "| value - ", value)
	}

	//NOTE: MAPS ARE NOT SAFE FOR CONCURRENT USE. IF YOU WANT TO USE MAPS IN CONCURRENT ENVIRONMENT, YOU NEED TO USE SYNC.MAP OR USE MUTEX LOCKS.
	//NOTE: MAPS ARE NOT ORDERED. THE ORDER OF THE ELEMENTS IN THE MAP IS NOT GUARANTEED. IF YOU WANT TO MAINTAIN THE ORDER OF THE ELEMENTS, YOU NEED TO USE SLICE OR ARRAY.

}