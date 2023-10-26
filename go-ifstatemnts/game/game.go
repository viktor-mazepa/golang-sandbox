package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

var reader = bufio.NewReader(os.Stdin)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

func (g *Game) Rounds() {

	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
		}
	}
}

// clearScreen clears the screen
func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) PlayRound() bool {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	playerValue := -1
	g.DisplayChan <- ""
	g.DisplayChan <- fmt.Sprintf("Round %d", g.Round.RoundNumber)
	g.DisplayChan <- "-----------------"
	fmt.Print("Please enter rock, paper, or scissors ->")
	playerChoice, _ := reader.ReadString('\n')

	playerChoice = strings.Replace(playerChoice, "\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\r", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	g.DisplayChan <- ""
	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))

	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose: ROCK"
		break
	case PAPER:
		g.DisplayChan <- "Computer chose: PAPER"
		break
	case SCISSORS:
		g.DisplayChan <- "Computer chose: SCISSORS"
		break
	default:
	}

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw"
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.ComputerWins()
			} else {
				g.PlayersWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.ComputerWins()
			} else {
				g.PlayersWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.ComputerWins()
			} else {
				g.PlayersWins()
			}
			break
		default:
			g.DisplayChan <- "Invalid chose"
			return false
		}
	}

	return true

}

func (g *Game) ComputerWins() {
	g.DisplayChan <- "Computer wins"
	g.Round.ComputerScore++
}

func (g *Game) PlayersWins() {
	g.DisplayChan <- "Player wins"
	g.Round.PlayerScore++
}

func (g *Game) PrintSummary() {
	g.DisplayChan <- "Final Score"
	g.DisplayChan <- "-----------"
	g.DisplayChan <- fmt.Sprintf("Player: %d, Computer: %d", g.Round.PlayerScore, g.Round.ComputerScore)

	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChan <- "Player wins game!"
	} else {
		g.DisplayChan <- "Computer wins game!"
	}
}
