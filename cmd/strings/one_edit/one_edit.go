package main

import (
	"fmt"
	"math"
)

// O(n) time | O(1) space - where n is the length of s1 or s2 because should differentiate in 1 character.
func oneEdit(s1, s2 string) bool {
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return false
	}

	isOneEdit := false

	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			if isOneEdit {
				return false
			}
			isOneEdit = true

			if len(s1) < len(s2) {
				i++
			} else if len(s1) > len(s2) {
				j++
			} else {
				i++
				j++
			}
		} else {
			i++
			j++
		}
	}
	return isOneEdit
}

func main() {

	stringOne := "hello"
	stringTwo := "helo"
	actual := oneEdit(stringOne, stringTwo)
	fmt.Println(actual)
}
