package main

import ("fmt")

func main(){
	name:="anu"
	favChain:="ethereum"
	fmt.Printf("My name is %s and my favourite chain is %s \n", name, favChain)

	multi:= `abcdefg jlj \n
			\n hiuyhiouou`
			
	multi2:= "abcdefgjkhujh \n jhjkh"

	fmt.Println(multi)
	fmt.Println(multi2)

	fmt.Println(name[0])
}