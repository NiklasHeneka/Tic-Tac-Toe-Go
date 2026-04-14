package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello and welcome to a game of Tic-Tac-Toe!")
	name := getPlayerName()
	userStarts := doesUserStart()
	bot := selectBot()
	fmt.Println("Let the game begin! *Imagine some epic music playing*")

	board := Board{}
	round := 1
	printBoard(board)

	for {
		if round%2 == 1 && userStarts {
			userMove(&board, X)
		} else if round%2 == 2 && !userStarts {
			userMove(&board, O)
		} else if round%2 == 1 && !userStarts {
			bot.MakeMove(&board, X)
		} else {
			bot.MakeMove(&board, O)
		}

		printBoard(board)
		round++

		gameOver, winner := isGameOver(board)
		if gameOver {
			if winner == None {
				fmt.Println("This game ended in a draw. How boring...")
			} else if winner == X && userStarts || winner == O && !userStarts {
				fmt.Printf("Congratulations %v, you won!\n", name)
			} else {
				fmt.Println("You lost!")
			}

			var playAgain string
			for {
				fmt.Print("Do you want to play again? (y/n): ")
				fmt.Scanln(&playAgain)

				if playAgain == "y" {
					board = Board{}
					round = 1
					break
				} else if playAgain == "n" {
					return
				}
			}
		}
	}

}

type Mark int

const (
	None Mark = iota
	X
	O
)

func (p Mark) String() string {
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
	fields [3][3]Mark
}

type Bot interface {
	MakeMove(board *Board, mark Mark)
}

type RandomBot struct{}

func (b RandomBot) MakeMove(board *Board, mark Mark) {
	possiblePosition := getUnoccupiedFields(*board)
	random := rand.IntN(len(possiblePosition))
	board.fields[possiblePosition[random].x][possiblePosition[random].y] = mark
}

type SmartBot struct{}

func (b SmartBot) MakeMove(board *Board, mark Mark) {}

type AIBot struct{}

func (b AIBot) MakeMove(board *Board, mark Mark) {}

func userMove(board *Board, mark Mark) {
	for {
		var input string
		fmt.Print("Enter your field (e.g., 11, 23, 31): ")
		fmt.Scanln(&input)

		if len(input) != 2 {
			fmt.Println("Invalid format. Please enter exactly two digits (e.g., 13).")
			continue
		}

		rowDigit, errR := strconv.Atoi(string(input[0]))
		colDigit, errC := strconv.Atoi(string(input[1]))

		if errR != nil || errC != nil {
			fmt.Println("Invalid input. Please enter numbers only.")
			continue
		}

		r, c := rowDigit-1, colDigit-1

		if r < 0 || r > 2 || c < 0 || c > 2 {
			fmt.Println("Error: Digits must be between 1 and 3.")
			continue
		}

		if board.fields[r][c] != None {
			fmt.Printf("Error: %d%d is already taken by %s.\n", rowDigit, colDigit, board.fields[r][c])
			continue
		} else {
			board.fields[r][c] = mark
			return
		}
	}
}

func isGameOver(board Board) (bool, Mark) {
	winner := checkWinner(board)
	if winner != None {
		return true, winner
	}

	if isBoardFull(board) {
		return true, None
	}
	return false, None
}

func checkWinner(board Board) Mark {
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
	return len(getUnoccupiedFields(board)) == 0
}

type Position struct {
	x, y int
}

func getUnoccupiedFields(board Board) []Position {
	var positions []Position
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board.fields[r][c] == None {
				positions = append(positions, Position{r, c})
			}
		}
	}
	return positions
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
		fmt.Print("Do you want to make the first move? (y/n): ")
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
	fmt.Println()
}
