# Tic-Tac-Toe in Go 🎮

A terminal-based Tic-Tac-Toe game written in Go featuring multiple difficulty levels and an entertaining "AI" bot.

## Features

✨ **Multiple Bot Difficulties**
- **Easy Bot**: Makes random moves
- **Smart Bot**: Uses strategic decision-making (prioritizes winning, blocking opponent's winning moves, and occupying center/corners)
- **"AI" Bot**: An "AI" thinking outside the box

## Requirements

- Go 1.26 or later

## Installation

Clone the repository:
```bash
git clone https://github.com/NiklasHeneka/Tic-Tac-Toe-Go.git
cd Tic-Tac-Toe-Go
```

## Usage

Run the game:
```bash
go run main.go
```


### Move Input Format

The board is indexed from 1 to 3:
```
  1 2 3
1 • • •
2 • • •
3 • • •
```

Enter your move as two consecutive digits without spaces (e.g., `11`, `23`, `32`).

## Game Rules

Standard Tic-Tac-Toe rules apply:
- Players alternate turns placing X's and O's on a 3x3 grid
- First to get three in a row (horizontally, vertically, or diagonally) wins
- If the board fills without a winner, it's a draw
- Play as many rounds as you want!

## License

This project is open source and available under the MIT License.

## Author

[Niklas Heneka](https://github.com/NiklasHeneka)
