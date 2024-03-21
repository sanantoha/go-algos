package main

import "fmt"

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	return nil
}

func multiply1(mat1 [][]int, mat2 [][]int) [][]int {
	return nil
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
