package main

import "fmt"

func main() {
	fmt.Println("Lets learn Bytes and Runes in golang")

	// byte	One byte (8 bits)
	// rune	One Unicode code point

	//rune is always written in single quotes 
	//Rune is alias for int32. The value of a rune literal is the integer value of the Unicode code point for that character.

	var example rune = 'a'
	fmt.Println("Value of example is: ", example)


}