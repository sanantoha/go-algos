package main

import "fmt"

// O(n) time | O(1) space
func minSwapsRequired(str string) int {
	if str == "" {
		return 0
	}

	cnt := 0
	l := 0
	r := len(str) - 1

	for l < r {
		if str[l] != str[r] {
			cnt++
		}
		l++
		r--
	}

	if len(str)%2 == 0 && cnt%2 == 1 {
		return -1
	}
	return (cnt + 1) / 2
}

func main() {

	fmt.Println(minSwapsRequired("0100101")) // 2
	fmt.Println(minSwapsRequired("1110"))    // -1
	fmt.Println(minSwapsRequired("11101"))   // 1
}
