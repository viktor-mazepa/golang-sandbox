package main

import (
	"bufio"
	"fmt"
	"loopsApp/mylogger"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i <= 10; i++ {
		fmt.Println("i is", i)
	}

	cond := 1000

	for cond > 100 {
		cond = rand.Intn(1000) + 1
		fmt.Println("cond is", cond)
		if cond > 100 {
			fmt.Println("cond is", cond, "so loops keeps going")

		}
	}

	fmt.Println("Got", cond, "and broke out of loop")

	reader := bufio.NewReader(os.Stdin)

	ch := make(chan string)

	go mylogger.ListenForLog(ch)

	fmt.Println("Enter something")

	for i := 0; i < 5; i++ {
		fmt.Print("->")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
