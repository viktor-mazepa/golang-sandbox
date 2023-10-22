package main

import "log"

var myInt int

// positive value or zero
var myUint uint

var myFloat float32
var myFloat64 float64

func main() {

	myInt = 10
	myUint = 20

	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Viktor"

	log.Println(myString)

	myString = "John"

	var myBool = true

	log.Println(myBool)

}
