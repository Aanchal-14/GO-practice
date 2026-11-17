package main

import "fmt"


	// A method is a function with a special receiver argument.
	//A method is a function that belongs to a type. You call it with a dot

	// The receiver appears in its own argument list between the func keyword and the method name.

	//its used for structs and interfaces. It allows you to define behavior for your custom types like structs and interfaces

	type Wallet struct{
		Address string
		Balance float64

	}


	//defining method for struct type Wallet 

	func (w Wallet) GetBalance() float64 {
		return w.Balance  
	} 

	//(w Wallet)
	//The receiver. This is the only new part. It says "this function belongs to Wallet". 
	// Inside the method, w is the wallet you called it on. 
	// It works like "this", but you choose the name (usually the first letter of the type).

	//	WRITING METHODS WITH POINTER RECEIVERS
	//used pointer receiver to modify the original struct instance.
	func (w *Wallet) Deposit(amount float64) {
		w.Balance += amount
	}

func main(){
	
	p:= Wallet{
		Address: "0x1234567890abcdef",
		Balance: 100.0,
	}

	fmt.Println("Balance of the wallet is:", p.GetBalance())
	p.Deposit(50.0)
	fmt.Println("Balance of the wallet after deposit is:", p.GetBalance())
}