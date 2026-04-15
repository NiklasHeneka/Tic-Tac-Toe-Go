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
	currentPlayer := X
	for {
		printBoard(board)
		if currentPlayer == X && userStarts || currentPlayer == O && !userStarts {
			userMove(&board, currentPlayer, name)
		} else {
			bot.MakeMove(&board, currentPlayer)
			fmt.Println("The Bot made its move.")
		}

		currentPlayer = changePlayer(currentPlayer)
		gameOver, winner := isGameOver(board)
		if gameOver {
			printBoard(board)
			printWinner(winner, userStarts, name)
			if playAgain() {
				board = Board{}
				currentPlayer = X
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
	fields     [3][3]Mark
	funnyField Position
	funnyMark  Mark
}

type Bot interface {
	MakeMove(board *Board, mark Mark)
}

type RandomBot struct{}

func (b RandomBot) MakeMove(board *Board, mark Mark) {
	pos := randomMove(board)
	board.fields[pos.x][pos.y] = mark
}

func randomMove(board *Board) Position {
	possiblePosition := getUnoccupiedFields(*board)
	random := rand.IntN(len(possiblePosition))
	return Position{possiblePosition[random].x, possiblePosition[random].y}
}

type SmartBot struct{}

func (b SmartBot) MakeMove(board *Board, mark Mark) {
	pos := smartMove(board, mark)
	board.fields[pos.x][pos.y] = mark
}

func smartMove(board *Board, mark Mark) Position {
	win, winPos := canWin(*board, mark)
	preWin, preWinPos := canPreventLoss(*board, mark)

	if win {
		return winPos
	} else if preWin {
		return preWinPos
	} else if board.fields[1][1] == None {
		return Position{1, 1}
	} else if board.fields[0][0] == None {
		return Position{0, 0}
	} else if board.fields[0][2] == None {
		return Position{0, 2}
	} else if board.fields[2][0] == None {
		return Position{2, 0}
	} else if board.fields[2][2] == None {
		return Position{2, 2}
	}
	return randomMove(board)
}

type AIBot struct{}

func (b AIBot) MakeMove(board *Board, mark Mark) {
	win, winPos := canWin(*board, mark)
	fakeWin, fakeWinPos := canFakeWin(*board, mark)
	if win {
		board.fields[winPos.x][winPos.y] = mark
	} else if fakeWin {
		board.funnyField = fakeWinPos
		board.funnyMark = mark
	} else {
		pos := smartMove(board, mark)
		board.fields[pos.x][pos.y] = mark
	}
}

func canWin(board Board, mark Mark) (bool, Position) {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board.fields[r][c] == None {
				newBoard := board
				newBoard.fields[r][c] = mark
				winner := checkWinner(newBoard)
				if winner == mark {
					return true, Position{r, c}
				}
			}
		}
	}
	return false, Position{}
}

func canFakeWin(board Board, mark Mark) (bool, Position) {
	fields := board.fields
	for i := 0; i < 3; i++ {
		if fields[i][0] == mark && fields[i][0] == fields[i][1] {
			return true, Position{i, -1}
		}
		if fields[i][2] == mark && fields[i][2] == fields[i][1] {
			return true, Position{i, 3}
		}
		if fields[0][i] == mark && fields[0][i] == fields[1][i] {
			return true, Position{-1, i}
		}
		if fields[2][i] == mark && fields[2][i] == fields[1][i] {
			return true, Position{3, i}
		}
	}
	if fields[0][0] == mark && fields[0][0] == fields[1][1] {
		return true, Position{-1, -1}
	} else if fields[2][2] == mark && fields[2][2] == fields[1][1] {
		return true, Position{3, 3}
	} else if fields[0][2] == mark && fields[0][2] == fields[1][1] {
		return true, Position{-1, 3}
	} else if fields[2][0] == mark && fields[2][0] == fields[1][1] {
		return true, Position{3, -1}
	}
	return false, Position{}
}

func canPreventLoss(board Board, mark Mark) (bool, Position) {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board.fields[r][c] == None {
				newBoard := board
				otherMark := changePlayer(mark)
				newBoard.fields[r][c] = otherMark
				winner := checkWinner(newBoard)
				if winner == otherMark {
					return true, Position{r, c}
				}
			}
		}
	}
	return false, Position{}
}

func changePlayer(mark Mark) Mark {
	if mark == X {
		return O
	}
	return X
}

func userMove(board *Board, mark Mark, name string) {
	for {
		var input string
		fmt.Printf("%v it is your turn to place an %v.\nYou can do this by entering the position e.g. 13:\n", name, mark.String())
		fmt.Scanln(&input)

		if len(input) != 2 {
			fmt.Println("Invalid format. Please enter exactly two digits (e.g. 13).")
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
			fmt.Println("Invalid input. Digits must be between 1 and 3.")
			continue
		}

		if board.fields[r][c] != None {
			fmt.Printf("Wrong field: %d%d is already taken by %s.\n", rowDigit, colDigit, board.fields[r][c])
			continue
		} else {
			board.fields[r][c] = mark
			return
		}
	}
}

func isGameOver(board Board) (bool, Mark) {
	empty := Position{0, 0}
	if board.funnyField != empty {
		return true, board.funnyMark
	}
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
	} else if fields[0][2] != None && fields[0][2] == fields[1][1] && fields[1][1] == fields[2][0] {
		return fields[0][2]
	}

	return None
}

func printWinner(winner Mark, userStarts bool, name string) {
	if winner == None {
		fmt.Println("This game ended in a draw. How boring...")
	} else if winner == X && userStarts || winner == O && !userStarts {
		fmt.Printf("Congratulations %v, you won!\n", name)
	} else {
		fmt.Println("You lost! Git Gud")
	}
}

func playAgain() bool {
	var playAgain string
	for {
		fmt.Print("Do you want to play again? (y/n): ")
		fmt.Scanln(&playAgain)

		if playAgain == "y" {
			return true
		} else if playAgain == "n" {
			return false
		}
	}
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
		fmt.Println("1) Easy Bot")
		fmt.Println("2) Smart Bot")
		fmt.Println("3) \"AI\" Bot thinking outside the box")
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
	f := board.fields
	empty := Position{0, 0}
	if board.funnyField == empty {
		fmt.Println()
		fmt.Printf("    -------------\n"+
			"    | %v | %v | %v |\n"+
			"    -------------\n"+
			"    | %v | %v | %v |\n"+
			"    -------------\n"+
			"    | %v | %v | %v |\n"+
			"    -------------\n", f[0][0], f[0][1], f[0][2], f[1][0], f[1][1], f[1][2], f[2][0], f[2][1], f[2][2])
		fmt.Println()
	} else {
		b := [5][5]Mark{}
		b[board.funnyField.x+1][board.funnyField.y+1] = board.funnyMark
		fmt.Printf("  %v   %v   %v   %v   %v\n"+
			"    -------------\n"+
			"  %v | %v | %v | %v | %v\n"+
			"    -------------\n"+
			"  %v | %v | %v | %v | %v\n"+
			"    -------------\n"+
			"  %v | %v | %v | %v | %v\n"+
			"    -------------\n"+
			"  %v   %v   %v   %v   %v\n",
			b[0][0], b[0][1], b[0][2], b[0][3], b[0][4],
			b[1][0], f[0][0], f[0][1], f[0][2], b[1][4],
			b[2][0], f[1][0], f[1][1], f[1][2], b[2][4],
			b[3][1], f[2][0], f[2][1], f[2][2], b[3][4],
			b[4][0], b[4][1], b[4][2], b[4][3], b[4][4])
	}
}
