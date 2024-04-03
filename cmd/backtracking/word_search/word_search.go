package main

import "fmt"

func exist(board [][]rune, word string) bool {
	return false
}

func main() {
	board := [][]rune{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCE"

	fmt.Println(exist(board, word))

	fmt.Println("========================")

	board1 := [][]rune{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word1 := "ABCESEEEFS"

	fmt.Println(exist(board1, word1))

	fmt.Println("===========================")

	word2 := "ABCEFSADEESE"
	fmt.Println(exist(board1, word2))

	fmt.Println("==========================")

	fmt.Println(!exist(board1, "ABCEV"))
}
