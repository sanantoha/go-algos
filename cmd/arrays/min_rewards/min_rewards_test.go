package main

import (
	"fmt"
	"testing"
)

func TestMinRewards(t *testing.T) {
	arr := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}
	exp := 25

	funcs := []func([]int) int{
		minRewards,
		minRewards1,
		minRewards2,
	}

	for i, fn := range funcs {
		t.Run(fmt.Sprintf("minRewards%v", i), func(t *testing.T) {
			res := fn(arr)
			fmt.Println(res)

			if res != exp {
				t.Errorf("expected %v, but got %v", exp, res)
			}
		})
	}
}
