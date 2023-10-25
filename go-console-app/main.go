package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappuchino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("-----")
	for i := 1; i <= len(coffees); i++ {
		fmt.Println(i, " - ", coffees[i])
	}
	fmt.Println("Q - Quit the pogram")

	char := ' '

	for char != 'q' && char != 'Q' {
		char, _, err = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))

		el, ok := coffees[i]
		if ok {
			fmt.Printf("You chose %s\n", el)
		}
	}

	fmt.Println("Program exiting")

}
