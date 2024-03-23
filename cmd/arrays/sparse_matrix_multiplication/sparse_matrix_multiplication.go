package main

import "fmt"

// O(n * k * m) time | O(n * m) space
func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	if len(mat1[0]) != len(mat2) {
		return [][]int{}
	}

	res := make([][]int, len(mat1))

	for i := 0; i < len(mat1); i++ {
		res[i] = make([]int, len(mat2[0]))
		for j := 0; j < len(mat1[i]); j++ {
			for k := 0; k < len(mat2); k++ {
				res[i][k] += mat1[i][j] * mat2[j][k]
			}
		}
	}
	return res
}

// O(n * k * m) time | O(n * k + k * m) space
func multiply1(mat1 [][]int, mat2 [][]int) [][]int {
	if len(mat1[0]) != len(mat2) {
		return [][]int{}
	}

	res := make([][]int, len(mat1))

	m1 := compress(mat1)
	m2 := compress(mat2)

	for i := 0; i < len(mat1); i++ {
		res[i] = make([]int, len(mat2[0]))
		for _, info1 := range m1[i] {
			for _, info2 := range m2[info1.col] {
				res[i][info2.col] += info1.val * info2.val
			}
		}
	}
	return res
}

func compress(matrix [][]int) [][]Info {
	res := make([][]Info, len(matrix))

	for i := 0; i < len(matrix); i++ {
		res[i] = make([]Info, 0)
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				res[i] = append(res[i], Info{j, matrix[i][j]})
			}
		}
	}
	return res
}

type Info struct {
	col int
	val int
}

func main() {

	mat1 := [][]int{
		{1, 0, 0},
		{-1, 0, 3},
	}

	mat2 := [][]int{
		{7, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	}

	fmt.Println(multiply(mat1, mat2)) // [[7, 0, 0], [-7, 0, 3]]
	fmt.Println("=================")
	fmt.Println(multiply1(mat1, mat2)) // [[7, 0, 0], [-7, 0, 3]]
}
