package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press Enter when ready."

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	play(firstNumber, secondNumber, subtraction, answer)

}

func play(firstNumber, secondNumber, subtraction, answer int) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiple the results by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')
	fmt.Println("The answer is", answer)
}
