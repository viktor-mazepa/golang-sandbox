package main

import "fmt"

type Animal struct {
	Name  string
	Sound string
	Legs  int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs\n", a.Name, a.Legs)
}

func main() {
	myTotal := sumMany(5, 6, 8, 0, 1, 6, 7, 55, -33)
	fmt.Println(myTotal)

	var dog Animal
	dog.Name = "Dog"
	dog.Sound = "woof"
	dog.Legs = 4
	dog.Says()
	dog.HowManyLegs()

	cat := Animal{
		Name:  "cat",
		Sound: "meow",
		Legs:  4,
	}
	cat.Says()
	cat.HowManyLegs()
}

// Naked return
func addTwoNumbers(x, y int) (sum int) {
	sum = x + y
	return
}

func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}
	return total
}
