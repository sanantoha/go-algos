package main

import "fmt"

// O(n ^ 2) time | O(n ^ 2) space
func generateMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	startRow := 0
	endRow := n - 1
	startCol := 0
	endCol := n - 1

	idx := 1

	for startRow <= endRow && startCol <= endCol {
		for i := startCol; i <= endCol; i++ {
			matrix[startRow][i] = idx
			idx++
		}
		startRow++

		for i := startRow; i <= endRow; i++ {
			matrix[i][endCol] = idx
			idx++
		}
		endCol--

		if startRow <= endRow && startCol <= endCol {
			for i := endCol; i >= startCol; i-- {
				matrix[endRow][i] = idx
				idx++
			}
			endRow--

			for i := endRow; i >= startRow; i-- {
				matrix[i][startCol] = idx
				idx++
			}
			startCol++
		}

	}

	return matrix
}

func main() {
	fmt.Println(generateMatrix(3)) // [[1 2 3] [8 9 4] [7 6 5]]
	fmt.Println("===============================")
	fmt.Println(generateMatrix(1)) // [[1]]
	fmt.Println("===============================")
	fmt.Println(generateMatrix(4)) // [[1 2 3 4] [12 13 14 5] [11 16 15 6] [10 9 8 7]]
}
