package main

import "fmt"

// O(n) time | O(1) space
func reverseString(arr []rune) {
	l := 0
	r := len(arr) - 1

	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func main() {

	s := []rune{'h', 'e', 't', 'l', 'l', 'o'}

	reverseString(s)

	for _, c := range s {
		fmt.Printf("%v ", string(c))
	}
}
