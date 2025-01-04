package main

import "fmt"

var gameboard [3][3]rune = [3][3]rune{
	{' ', ' ', ' '},
	{' ', ' ', ' '},
	{' ', ' ', ' '},
}

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

func main() {
	gameboard[1][1] = 'X'

	gameStart()

	printBoard()
}
