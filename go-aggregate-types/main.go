package main

import "fmt"

type Car struct {
	NumberOfTires int
	Luxury        bool
	Make          string
	Model         string
	Year          int
}

func main() {
	// Array
	var myStrings [3]string

	myStrings[0] = "Cat"
	myStrings[1] = "Dog"
	myStrings[2] = "Fish"

	fmt.Println("First element in array myStrings is", myStrings[0])

	var myCar Car
	myCar.NumberOfTires = 4
	myCar.Luxury = false
	myCar.Make = "Volkswagen"

	someCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("Some car is a %d %s %s\n", someCar.Year, someCar.Make, someCar.Model)

}
