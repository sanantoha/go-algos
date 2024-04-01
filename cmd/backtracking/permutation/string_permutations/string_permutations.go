package main

import "fmt"

// O(n! * n ^ 2) time | O(n * n!) space
func permute(str string) []string {
	res := make([]string, 0)
	backtrack(str, "", &res)
	return res
}

func backtrack(str string, ans string, res *[]string) {
	if len(str) == 0 {
		*res = append(*res, ans)
		return
	}

	for i := 0; i < len(str); i++ {
		rem := str[:i] + str[i+1:]
		backtrack(rem, ans+string(str[i]), res)
	}
}

func main() {

	str := "abc"

	fmt.Println(permute(str))
}
