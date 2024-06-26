package main

import (
	"fmt"
	"math/rand"
)

func insertSort(arr []int) {

}

func main() {

	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	fmt.Println(arr)

	insertSort(arr)

	fmt.Println(arr)
}
