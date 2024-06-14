package main

import (
	"fmt"
	"reflect"
)

// O(s1 * s2 * min(s1, s2)) time | O(s1 * s2 * min(s1 * s2)) space
func longestCommonSubsequence(str1, str2 string) []rune {
	if len(str1) == 0 || len(str2) == 0 {
		return []rune{}
	}

	lcs := make([][][]rune, len(str1)+1)
	for i := 0; i <= len(str1); i++ {
		lcs[i] = make([][]rune, len(str2)+1)
		for j := 0; j <= len(str2); j++ {
			lcs[i][j] = make([]rune, 0)
		}
	}

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				arr := make([]rune, len(lcs[i-1][j-1]))
				copy(arr, lcs[i-1][j-1])
				arr = append(arr, rune(str1[i-1]))
				lcs[i][j] = arr
			} else {
				up := lcs[i-1][j]
				left := lcs[i][j-1]
				if len(up) > len(left) {
					lcs[i][j] = up
				} else {
					lcs[i][j] = left
				}
			}
		}
	}

	return lcs[len(str1)][len(str2)]
}

func longestCommonSubsequence1(str1, str2 string) []rune {
	return nil
}

func main() {

	expected := []rune{'X', 'Y', 'Z', 'W'}

	actual := longestCommonSubsequence("ZXVVYZW", "XKYKZPW")
	fmt.Println(actual) // [88 89 90 87]
	fmt.Println(reflect.DeepEqual(actual, expected))

	actual = longestCommonSubsequence1("ZXVVYZW", "XKYKZPW")
	fmt.Println(actual) // [88 89 90 87]
	fmt.Println(reflect.DeepEqual(actual, expected))
}
