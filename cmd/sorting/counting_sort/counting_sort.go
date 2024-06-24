package main

import (
	"fmt"
	"math/rand"
)

func countingSort(arr []int) {

}

func main() {

	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	fmt.Println(arr)

	countingSort(arr)

	fmt.Println(arr)
}
