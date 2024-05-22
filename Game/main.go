package main

import (
	"fmt"
)

const (
	Empty = iota
	PlayerX
	PlayerO
)

var board [3][3]int

func printBoard() {
	for _, row := range board {
		for _, cell := range row {
			switch cell {
			case PlayerX:
				fmt.Print("X ")
			case PlayerO:
				fmt.Print("O ")
			default:
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func checkWin() int {
	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] != Empty && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if board[0][i] != Empty && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}
	// Check diagonals
	if board[0][0] != Empty && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != Empty && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}

	return Empty
}

func main() {
	var currentPlayer = PlayerX
	var row, col int
	var moves = 0

	for {
		printBoard()
		fmt.Printf("Player %s, enter row and column (0, 1, 2): ", func() string {
			if currentPlayer == PlayerX {
				return "X"
			}
			return "O"
		}())

		fmt.Scanf("%d %d", &row, &col)

		if row < 0 || row > 2 || col < 0 || col > 2 || board[row][col] != Empty {
			fmt.Println("Invalid move! Try again.")
			continue
		}

		board[row][col] = currentPlayer
		moves++

		if winner := checkWin(); winner != Empty {
			printBoard()
			fmt.Printf("Player %s wins!\n", func() string {
				if winner == PlayerX {
					return "X"
				}
				return "O"
			}())
			break
		}

		if moves == 9 {
			printBoard()
			fmt.Println("It's a draw!")
			break
		}

		if currentPlayer == PlayerX {
			currentPlayer = PlayerO
		} else {
			currentPlayer = PlayerX
		}
	}
}
