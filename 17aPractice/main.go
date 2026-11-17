package main

import "fmt"

// Make a struct Block with Number uint64, Hash string and TxCount int. Create one and print it with fmt.Printf("%+v\n", b).
// Write var empty Block and print it. What is in each compartment?
// Write b2 := b, change b2.Number, and print both. Did b change?


type Block struct {
	Number uint64
	Hash uint64
	TxCount int


}


func createBlock(Number uint64, Hash uint64, TxCount int) *Block {
	T:=&Block{
		Number: Number,
		Hash: Hash,
		TxCount: TxCount,
	}
	return T
}
	

func main(){
	b := createBlock(1, 0x1234567890abcdef, 10)
	fmt.Printf("%+v\n", *b) 
	fmt.Println("Address of b is: ", b) // Prints the struct pointer using Go's default formatting (e.g. &{1 ... 10})
	fmt.Println("Address ", &b) // Address where the pointer variable b itself is stored
	fmt.Printf("Actual address of b is: %p\n", b) // Prints the hexadecimal address of the struct that b points to

	c:= Block{}
	fmt.Printf("%+v\n", c)


	b2:=b
	b2.Number=222 //changed b as well
	fmt.Printf("b2: %+v\n", *b2)
}












