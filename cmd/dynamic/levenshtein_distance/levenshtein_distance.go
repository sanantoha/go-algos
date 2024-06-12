package main

import "fmt"

// O(s1 * s2) time | O(s1 * s2) space
func levenshteinDistance(str1, str2 string) int {
	if len(str1) == 0 {
		return len(str2)
	}
	if len(str2) == 0 {
		return len(str1)
	}

	ld := make([][]int, len(str1)+1)
	arr := make([]int, len(str2)+1)
	for i, _ := range arr {
		arr[i] = 1
	}
	ld[0] = arr

	for i := 1; i <= len(str1); i++ {
		ld[i] = make([]int, len(str2)+1)
		ld[i][0] = 1
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				ld[i][j] = ld[i-1][j-1]
			} else {
				ld[i][j] = min(ld[i-1][j-1], ld[i-1][j], ld[i][j-1]) + 1
			}
		}
	}
	return ld[len(str1)][len(str2)]
}

func levenshteinDistance1(str1, str2 string) int {
	return -1
}

func main() {

	fmt.Println(levenshteinDistance("abc", "yabd") == 2)
	fmt.Println(levenshteinDistance1("abc", "yabd") == 2)
}
