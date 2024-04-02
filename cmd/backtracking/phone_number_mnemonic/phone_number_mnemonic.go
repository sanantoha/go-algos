package main

import "fmt"

func phoneNumberMnemonics(str string) []string {
	letters := []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	var _ = letters

	return nil
}

func main() {
	str := "1905"

	fmt.Println(phoneNumberMnemonics(str))
}
