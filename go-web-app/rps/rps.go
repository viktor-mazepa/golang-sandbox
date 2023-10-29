package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
	Message        string `json:"message"`
}

var messageMap = map[int][]string{
	1: {"You are the best", "Win, win, win", "Wow! You are so lucky"},
	2: {"Bite my shiny metal a##!", "Hahahahaha", "The game’s over. I have all the money. Compare your lives to mine and ..."},
	3: {"Let's try againe... now!", "Next tme someone can win", "Ooooops"},
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	messageRand := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer wins!"
		winner = COMPUTERWINS
	}

	var result Round
	result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	result.Message = messageMap[winner][messageRand]
	return result
}
