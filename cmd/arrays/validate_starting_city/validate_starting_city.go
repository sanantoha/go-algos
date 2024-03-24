package main

import "fmt"

func validateStartingCity(distances []int, fuel []int, mpg int) int {
	return -1
}

func validateStartingCity1(distance []int, fuel []int, mpg int) int {
	return -1
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
