package main

import (
	"fmt"
	"math"
)

// O(n ^ 2) time | O(1) space
func validateStartingCity(distances []int, fuel []int, mpg int) int {
	if distances == nil || len(distances) == 0 || fuel == nil || len(fuel) == 0 {
		return -1
	}

	for city := 0; city < len(distances); city++ {
		remainingDistance := 0
		for idx := city; idx < city+len(distances); idx++ {
			ridx := idx % len(distances)
			remainingDistance += fuel[ridx]*mpg - distances[ridx]

			if remainingDistance < 0 {
				break
			}
		}
		if remainingDistance >= 0 {
			return city
		}
	}
	return -1
}

// O(n) time | O(1) space
func validateStartingCity1(distance []int, fuel []int, mpg int) int {

	remainingDistance := 0
	minDistance := math.MaxInt
	resCity := 0

	for city := 1; city < len(distance); city++ {
		remainingDistance += fuel[city-1]*mpg - distance[city-1]

		if remainingDistance < minDistance {
			minDistance = remainingDistance
			resCity = city
		}
	}
	return resCity
}

func main() {
	distances := []int{5, 25, 15, 10, 15}
	fuel := []int{1, 2, 1, 0, 3}
	mpg := 10
	expected := 4

	actual := validateStartingCity(distances, fuel, mpg)
	fmt.Println(actual)
	fmt.Println(expected == actual)

	actual = validateStartingCity1(distances, fuel, mpg)
	fmt.Println(actual)
	fmt.Println(expected == actual)
}
