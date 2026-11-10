package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main(){
	fmt.Println("welcome to our pizza app")
	fmt.Println(("PLease rate our pizza between 1 and 5"))

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')  //this will take the input from the user and store it in the input variable

	fmt.Println("Thanks for the rating, ", input)

	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64) 
	if err != nil {
		fmt.Println("Error occurred while converting string to float")
	} else {
		fmt.Println("Your rating plus one is: ", numRating+1)
	}
}
