package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Hello and welcome to a game of Tic-Tac-Toe!")
	getPlayerName()
	doesUserStart()
	selectBot()
	fmt.Println("Let the game begin! *Imagine some epic music playing*")

}

type Marker int

const (
	None Marker = iota
	X
	O
)

func (p Marker) String() string {
	switch p {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return " "
	}
}

type Board struct {
	fields [3][3]Marker
}

type Bot interface {
	MakeMove(board Board) Move
}

type Move struct {
	Row, Col int
}
type RandomBot struct{}

func (b RandomBot) MakeMove(board Board) Move { return Move{} }

type SmartBot struct{}

func (b SmartBot) MakeMove(board Board) Move { return Move{} }

type AIBot struct{}

func (b AIBot) MakeMove(board Board) Move { return Move{} }

func getPlayerName() string {
	scanner := bufio.NewScanner(os.Stdin)
	var playerName string

	for {
		fmt.Print("Enter your name:")

		if scanner.Scan() {
			input := scanner.Text()
			playerName = strings.TrimSpace(input)

			if len(playerName) > 0 {
				fmt.Printf("Welcome %s\n", playerName)
				break
			}
		}

		fmt.Println("Invalid input. Please try again.")
	}
	return playerName
}

func selectBot() Bot {
	for {
		fmt.Println("\nChoose the bot you want to play again:")
		fmt.Println("1) Random positioning")
		fmt.Println("2) Smart positioning")
		fmt.Println("3) \"AI\" thinking outside the box")
		fmt.Print("Selection (1-3): ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			return RandomBot{}
		case "2":
			return SmartBot{}
		case "3":
			return AIBot{}
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

func doesUserStart() bool {
	var userStarts string

	for {
		fmt.Print("Do you want to start the game? (y/n): ")
		fmt.Scanln(&userStarts)

		if userStarts == "y" {
			return true
		} else if userStarts == "n" {
			return false
		}
	}
}
