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
