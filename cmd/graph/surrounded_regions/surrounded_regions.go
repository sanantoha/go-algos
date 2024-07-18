package main

import "fmt"

// O(w * h) time | O(w * h) space
func solve(board [][]rune) {
	if len(board) == 0 {
		return
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			isBorder := row == 0 || col == 0 || row == len(board)-1 || col == len(board[row])-1

			if !isBorder || board[row][col] != 'O' {
				continue
			}

			dfs(board, row, col)
		}
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			} else if board[row][col] == '*' {
				board[row][col] = 'O'
			}
		}
	}
}

func dfs(board [][]rune, startRow int, startCol int) {
	stack := make([][]int, 1)
	stack[0] = []int{startRow, startCol}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		row := p[0]
		col := p[1]

		if board[row][col] != 'O' {
			continue
		}
		board[row][col] = '*'

		neighbors := getNeighbors(board, row, col)
		stack = append(stack, neighbors...)
	}
}

func getNeighbors(board [][]rune, row int, col int) [][]int {
	res := make([][]int, 0)

	if row > 0 {
		res = append(res, []int{row - 1, col})
	}
	if col > 0 {
		res = append(res, []int{row, col - 1})
	}
	if row+1 < len(board) {
		res = append(res, []int{row + 1, col})
	}
	if col+1 < len(board[row]) {
		res = append(res, []int{row, col + 1})
	}

	return res
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
