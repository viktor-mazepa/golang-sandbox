package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	simpleOperators()
	precedence()
	modulusOperator()
	condOperator()
	shortCircuitEval()
	assignmentOperator()
}

func assignmentOperator() {
	x := 12
	x++
	fmt.Println(x)

	y := 10
	y *= 2
	fmt.Println(y)

	y /= 4
	fmt.Println(y)

}

func shortCircuitEval() {
	a := 12
	b := 0
	c, err := divideTwoNumber(a, b)

	if err != nil {
		fmt.Println(err)
	} else {
		if c == 2 {
			fmt.Println("We found two")
		}
	}

}

func divideTwoNumber(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

func condOperator() {
	second := 31
	minute := 1

	if minute < 59 && second+1 > 59 {
		minute++
	}
	fmt.Println(minute)
}

func simpleOperators() {
	answer := 7 + 3*4

	fmt.Println("Answer:", answer)

	answer = (7 + 3) * 4
	fmt.Println("Answer:", answer)

	var radius = 12.0

	area := math.Pi * math.Pow(radius, 2)
	fmt.Println("Area:", area)

	half := 1 / 2
	fmt.Println("Half with int division:", half)

	halfFloat := 1.0 / 2.0
	fmt.Println("Half with float division:", halfFloat)

	reminder := 50 % 3
	fmt.Println("Reminder:", reminder)

	x := 3
	x++
	fmt.Println("x is now:", x)
}

func precedence() {
	a := 12.0 * 3.0 / 4.0

	b := (12.0 * 3.0) / 4.0

	c := 12.0 * (3.0 / 4.0)

	fmt.Println("a:", a, "b:", b, "c:", c)

	unclear := 12 * (3 / 4)
	fmt.Println("Unclear:", unclear)

	f := 12.0 / 3.0 / 4.0
	fmt.Println("f:", f)

	f = 12.0 / (3.0 / 4.0)
	fmt.Println("f:", f)

	fmt.Println()

	x := 12 + 3 - 4
	y := (12 + 3) - 4
	z := 12 + (3 - 4)
	fmt.Println("x:", x, "y:", y, "z:", z)

	x = 12 + 3*4
	y = (12 + 3) * 4
	z = 12 + (3 * 4)
	fmt.Println("x:", x, "y:", y, "z:", z)
}

func modulusOperator() {
	x := 12
	y := 5
	if x%y == 0 {
		fmt.Println(y, "devides exactly into", x)
	} else {
		fmt.Println(y, "does not divide exactly into", x)
	}

	for m := 1; m <= 12; m++ {
		fmt.Println("The month after", m, "is", m%12+1)
	}

}
