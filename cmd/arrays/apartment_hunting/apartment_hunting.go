package main

import (
	"fmt"
	"math"
)

// O(b * r) time | O(b * r) space
func apartmentHunting(blocks []map[string]bool, reqs []string) int {

	distances := make([]map[string]int, len(blocks))
	for i := 0; i < len(distances); i++ {
		distances[i] = make(map[string]int)
		for _, req := range reqs {
			distances[i][req] = math.MaxInt
		}
	}

	for name, val := range blocks[0] {
		if val {
			distances[0][name] = 0
		}
	}

	for i := 1; i < len(blocks); i++ {
		for _, req := range reqs {
			if blocks[i][req] {
				distances[i][req] = 0
			} else {
				newDistance := distances[i-1][req]
				if distances[i-1][req] != math.MaxInt {
					newDistance = distances[i-1][req] + 1
				}
				distances[i][req] = min(newDistance, distances[i][req])
			}
		}
	}

	for i := len(blocks) - 2; i >= 0; i-- {
		for _, req := range reqs {
			distances[i][req] = min(distances[i][req], 1+distances[i+1][req])
		}
	}

	min_distance_id := 0
	min_distance := math.MaxInt

	for idx, distance := range distances {
		max_val := 0

		for _, req := range reqs {
			max_val = max(max_val, distance[req])
		}

		if min_distance > max_val {
			min_distance = max_val
			min_distance_id = idx
		}
	}

	return min_distance_id
}

func main() {

	blocks := make([]map[string]bool, 0)

	blocks = append(blocks, map[string]bool{"gym": false, "school": true, "store": false})
	blocks = append(blocks, map[string]bool{"gym": true, "school": false, "store": false})
	blocks = append(blocks, map[string]bool{"gym": true, "school": true, "store": false})
	blocks = append(blocks, map[string]bool{"gym": false, "school": true, "store": false})
	blocks = append(blocks, map[string]bool{"gym": false, "school": true, "store": true})

	reqs := []string{"gym", "school", "store"}

	res := apartmentHunting(blocks, reqs)
	fmt.Println(res) // 3
}
