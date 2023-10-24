package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

func (v *Vehicle) showDetails() {
	fmt.Printf("Number of passengers: %d, Number of wheels: %d", v.NumberOfPassengers, v.NumberOfWheels)
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (c *Car) show() {
	fmt.Printf("Make: %s, Model: %s, Year: %d, Is electirc: %t, Is Hybrid: %t, ", c.Make, c.Model, c.Year, c.IsElectric, c.IsHybrid)
	c.Vehicle.showDetails()
}
func main() {

	car := Car{
		Make:       "Toyota",
		Model:      "RAV-4",
		Year:       2019,
		IsElectric: false,
		IsHybrid:   true,
		Vehicle: Vehicle{
			NumberOfWheels:     4,
			NumberOfPassengers: 5,
		},
	}
	car.show()
}
