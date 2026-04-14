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

	board := Board{}
	printBoard(board)

}

type Player int

const (
	None Player = iota
	X
	O
)

func (p Player) String() string {
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
	fields [3][3]Player
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

func isGameOver(board Board) (bool, Player) {
	winner := checkWinner(board)
	if winner != None {
		return true, winner
	}

	if isBoardFull(board) {
		return true, None
	}
	return false, None
}

func checkWinner(board Board) Player {
	fields := board.fields
	for i := 0; i < 3; i++ {
		if fields[i][0] != None && fields[i][0] == fields[i][1] && fields[i][1] == fields[i][2] {
			return fields[i][0]
		}
		if fields[0][i] != None && fields[0][i] == fields[1][i] && fields[1][i] == fields[2][i] {
			return fields[0][i]
		}
	}

	if fields[0][0] != None && fields[0][0] == fields[1][1] && fields[1][1] == fields[2][2] {
		return fields[0][0]
	}

	if fields[0][2] != None && fields[0][2] == fields[1][1] && fields[1][1] == fields[2][0] {
		return fields[0][2]
	}

	return None
}

func isBoardFull(board Board) bool {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board.fields[r][c] == None {
				return false
			}
		}
	}
	return true
}

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

func printBoard(board Board) {
	fmt.Println()
	for ind, row := range board.fields {
		for _, cell := range row {
			fmt.Printf("|%s", cell)
			fmt.Print("|")
		}
		fmt.Println()
		if ind != len(board.fields)-1 {
			fmt.Println("---------")
		}
	}
}
