package main

import (
	"fmt"
)

func main() {

	x := 10

	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 156

	fmt.Println("x is now", x)

	changeValueOfPointer(myFirstPointer)
	fmt.Println("x is now", x)

}

func changeValueOfPointer(num *int) {
	*num = 25
}
