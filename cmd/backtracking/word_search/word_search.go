package main

import "fmt"

// O(n * 3 ^ l) time | O(l) space - where n total amount of cells, and l length of word
func exist(board [][]rune, word string) bool {
	if board == nil || len(board) == 0 || word == "" {
		return false
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == rune(word[0]) && dfs(board, word, row, col, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]rune, word string, row int, col int, idx int) bool {
	if idx == len(word) {
		return true
	}
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[row]) || board[row][col] != rune(word[idx]) {
		return false
	}

	tmp := board[row][col]
	board[row][col] = ' '

	found := dfs(board, word, row-1, col, idx+1) || dfs(board, word, row, col-1, idx+1) ||
		dfs(board, word, row+1, col, idx+1) || dfs(board, word, row, col+1, idx+1)

	board[row][col] = tmp
	return found
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
