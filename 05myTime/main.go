package main

import (
	"fmt"
	"time"
)


func main(){
	fmt.Println("Welcopme to time studay of golang")

	presentTime := time.Now()

	fmt.Println("Present time is: ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //mentioning the format in which we want to print the time

	createdDate := time.Date(2026, time.August, 14 ,8,0,0,0, time.UTC) //this is how we can create a date in golang
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday")) //mentioning the format in which we want to print the time

}