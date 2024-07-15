package main

import "fmt"

func solve(board [][]rune) {

}

func printRuneMatrix(matrix [][]rune) {
	for _, row := range matrix {
		for _, r := range row {
			fmt.Printf("%c ", r)
		}
		fmt.Println()
	}
}

func main() {

	board := [][]rune{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(board)
	printRuneMatrix(board)
	fmt.Println("===============================")

	board1 := [][]rune{
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
	}

	solve(board1)
	printRuneMatrix(board1)
	fmt.Println("===============================")

	board2 := [][]rune{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
	}

	solve(board2)
	printRuneMatrix(board2)
}
