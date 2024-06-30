package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int, l, r int) {

}

func main() {

	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	fmt.Println(arr)

	quickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			panic(fmt.Sprintf("array is not sorted: %d != %d", arr[i], arr[i+1]))
		}
	}

	fmt.Println("done")
}
