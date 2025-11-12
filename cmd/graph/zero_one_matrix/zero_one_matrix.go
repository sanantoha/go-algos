package main

import (
	"fmt"
	"math"
)

// O(w * h) time | O(w * h) space
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	replaceOneOnMaxVal(matrix)

	queue := findAllZeros(matrix)

	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		row := t[0]
		col := t[1]
		level := t[2]

		for _, p := range getNeighborns(matrix, row, col) {
			nrow := p[0]
			ncol := p[1]
			if matrix[nrow][ncol] == math.MaxInt {
				newLevel := level + 1
				matrix[nrow][ncol] = newLevel
				queue = append(queue, []int{nrow, ncol, newLevel})
			}
		}
	}

	return matrix
}

func findAllZeros(matrix [][]int) [][]int {
	queue := make([][]int, 0)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 0 {
				queue = append(queue, []int{row, col, 0})
			}
		}
	}
	return queue
}

func replaceOneOnMaxVal(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 1 {
				matrix[row][col] = math.MaxInt
			}
		}
	}
}

func getNeighborns(matrix [][]int, row int, col int) [][]int {
	res := make([][]int, 0)

	if row > 0 {
		res = append(res, []int{row - 1, col})
	}
	if col > 0 {
		res = append(res, []int{row, col - 1})
	}
	if row+1 < len(matrix) {
		res = append(res, []int{row + 1, col})
	}
	if col+1 < len(matrix[row]) {
		res = append(res, []int{row, col + 1})
	}

	return res
}

func main() {

	input0 := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	input1 := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	fmt.Println(updateMatrix(input0))

	fmt.Println(updateMatrix(input1))
}
