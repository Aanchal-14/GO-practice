package main

import "fmt"

//An interface is a list of methods. Any type that has all of them fits, automatically.

//declaration of interface type
type Speaker interface {
	Speak() string
}

func announce(s Speaker) {
    fmt.Println(s.Speak())
}

//the problem that interface fixes --
//FOr each struct /types you will need a sperate finction to print what each animal speak / that separet functions to call each speak fucntion

type Dog struct{}
func (d Dog) Speak() string { return "Woof" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow" }


//But with interface , we only need to define one function (here announce() ) that can call any animal's speak fucntion . 

//IT WILL RUN FOR ALL TYPS THAT HAS sPEAK AS ITS METHOD with same signature i,e same param,s and same retuns


func main(){
	fmt.Println("Lets learn Interfaces in golang")

	//USING INTERFACE FUNCTION

	//so with announce() you can now pass any type to announce and it will call the speaker' respective Speak fucntion

	announce(Dog{})
	announce(Cat{})

}