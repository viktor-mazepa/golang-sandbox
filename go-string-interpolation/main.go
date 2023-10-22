package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is our favorite number?")
	user.OwnsADog = readBool("Do you have a dog (y/n)?")
	//fmt.Println("Your name is", userName+". You are", userAge, "years old")

	//fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old", userName, userAge))

	fmt.Printf("Your name is %s. You are %d years old. Your favorite number is %.4f. Owns a dog:%t\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
		user.OwnsADog)

}

func prompt() {
	fmt.Print("->")
}

func readString(question string) string {
	for {
		fmt.Println(question)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a valid name")
		} else {
			return userInput
		}
	}
}

func readInt(question string) int {
	for {
		fmt.Println(question)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a valid number")
		} else {
			return num
		}
	}
}

func readFloat(question string) float64 {
	for {
		fmt.Println(question)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a valid number")
		} else {
			return num
		}
	}
}

func readBool(question string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(question)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'y' {
			return true
		}
	}
}
