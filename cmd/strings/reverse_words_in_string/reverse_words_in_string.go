package main

import "fmt"

// O(n) time | O(n) space
func reverse(str string) string {
	arr := []rune(str)
	rev(arr, 0, len(arr)-1)

	startIdx := 0

	for startIdx < len(arr) {
		endIdx := startIdx

		for endIdx < len(arr) && arr[endIdx] != ' ' {
			endIdx++
		}
		rev(arr, startIdx, endIdx-1)
		startIdx = endIdx + 1
	}

	return string(arr)
}

func rev(arr []rune, l, r int) {

	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func main() {

	s := "the sky is blue"

	fmt.Println(reverse(s))

	fmt.Printf("|%s|\n", reverse(" a good   example     "))

	fmt.Println(reverse("this      string     has a     lot of   whitespace"))
}
