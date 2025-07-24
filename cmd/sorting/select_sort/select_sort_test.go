package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	fmt.Println(arr)

	selectSort(arr)

	fmt.Println(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("array is not sorted: %d != %d", arr[i], arr[i+1])
		}
	}
}
