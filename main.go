package main

import (
	"fmt"
	"strconv"
)

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
			fmt.Println("\n\nError reading input, please repeat typing the coordinates.")
		} else {
			break
		}
	}

	return coords
}

func main() {
	gameStart()

	move := getMove()

	fmt.Println(move)

	printBoard()
}
