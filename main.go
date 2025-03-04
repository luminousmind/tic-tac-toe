package main

import (
	"fmt"
	"strconv"
	"time"
)

var winner string

var gameboard [3][3]rune = [3][3]rune{
	{' ', ' ', ' '},
	{' ', ' ', ' '},
	{' ', ' ', ' '},
}

var playerTurn rune = 'X'

func printBoard() {
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			fmt.Printf("%c ", gameboard[x][y])
		}

		fmt.Println()
	}
}

func gameStart() {
	winner = ""

	fmt.Println("Tic-tac-toe is a simple boardgame in which two players alternate marking the spaces in three-by-three grid with X and O.")
	fmt.Println("In a brief explanation of the rules of Tic-tac-toe:")
	fmt.Println()
	fmt.Println("The rules of Tic-tac-toe:")
	fmt.Println("1. The game is played on a three-by-three grid.")
	fmt.Println("2. Two players exchange turns marking a square on the grid with X or O.")
	fmt.Println("3. The first player to get three in a row horizontally, vertically, or diagonally wins.")
	fmt.Println("4. If all squares are filled without a winner, it is a draw.")
	fmt.Println()
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
}

func getCord(axis rune) int {
	fmt.Printf("Enter the %c coordinate to place an %c: ", axis, playerTurn)

	var input string

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return -1
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
		return -1
	}

	return num
}

func getMove() [2]int {
	if playerTurn == 'X' {
		fmt.Println("It is Player 1's turn.")
	} else {
		fmt.Println("It is Player 2's turn.")
	}

	var coords [2]int

	for {
		coords[0] = getCord('X')
		coords[1] = getCord('Y')

		if coords[0] == -1 || coords[1] == -1 {
			fmt.Println("\n\n🚫 Error reading input, please repeat typing the coordinates.")
		} else if gameboard[coords[0]][coords[1]] != ' ' {
			fmt.Println("\n\n🚫 That tile is already taken, please try again.")
		} else {
			return coords
		}
	}
}

func win(player rune) {
	if player == 'X' {
		winner = "Player 1"
	} else if player == 'O' {
		winner = "Player 2"
	}
}

func test(x1, y1, x2, y2, x3, y3 int) {
	if gameboard[x1][y1] == gameboard[x2][y2] && gameboard[x1][y1] == gameboard[x3][y3] {
		win(gameboard[x1][y1])
	}
}

func testWin() {
	for i := 0; i < 3; i++ {
		test(i, 0, i, 1, i, 2)
		test(0, i, 1, i, 2, i)
	}

	// diagonals
	test(0, 0, 1, 1, 2, 2)
	test(2, 0, 1, 1, 0, 2)
}

func main() {
	gameStart()

	for {
		move := getMove()

		gameboard[move[0]][move[1]] = playerTurn

		if playerTurn == 'X' {
			playerTurn = 'O'
		} else {
			playerTurn = 'X'
		}

		printBoard()

		testWin()

		time.Sleep(time.Second)

		if winner != "" {
			fmt.Println("The winner is:", winner)
			break
		}
	}

}
