package main

import "fmt"

//jwtToken :=30000  //but here we are not allowed to use this as it is not inside the function
var jwtToken int = 30000 //this is allowed as it is outside the function


const LoginToken string = "qwertyuiop" //this is a constant and it is not allowed to change the value of this variable
									   //first letter is capital so it is exported and can be used outside the package

LoginToken = "qwertyuiopp" //this is a constant


func main(){

	var username string = "Aanchal"

	fmt.Println(username)

	fmt.Printf("Variable is of type : %T \n", username)


	var isLoggedIn bool = true;
	println(isLoggedIn)

	fmt.Printf("Variable is of type : %T \n", isLoggedIn)


	var smallVar uint8 = 255
	fmt.Println(smallVar)
	fmt.Printf("Variable is of type : %T \n", smallVar)

	smallVar = smallVar + 1 //this will NOT give an error but it will give the output as 0 because the maximum value of uint8 is 255 and when we add 1 to it, it will overflow and give the output as 0
	fmt.Printf("-=-=-==-Variable is of type : %d \n", smallVar)

	var smallFLoat float64 = 25.4555555559000000
	fmt.Println(smallFLoat)
	fmt.Printf("Variable is of type : %T \n", smallFLoat)

	var smallFloat2 float64 = 25.4555555559000000
	fmt.Printf("Variable value with 2 precision %.2f \n", smallFloat2) //this will give the output with 2 decimal places


	//default value and some aliases

	var anotherVAr int   //if you keep it undefined it will always giev 0 
	fmt.Println(anotherVAr)
	fmt.Printf("Variable is of type : %T \n", anotherVAr)

	//implicit type of declaring the variable
	var website = "www.aanchal.com"   //lexer will automatically assign the type of variable based on the value assigned to it
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)

	

	// website=3 //not allowed now

	//no var style
	numberOfUsers := 3005300.0123 //no variable used no type used
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type : %T \n", numberOfUsers)

	LoginToken := "qwertyuiip" //this is a constant
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}