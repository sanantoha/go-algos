package main

import (
	"fmt"
	"reflect"
)

func maxSumIncreasingSubsequence(arr []int) [][]int {
	return nil
}

func main() {

	arr := []int{10, 70, 20, 30, 50, 11, 30}
	expected := [][]int{
		{110},
		{10, 20, 30, 50},
	}

	actual := maxSumIncreasingSubsequence(arr)
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))
}
