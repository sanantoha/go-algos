package main

import "fmt"

func isMatch(s, p string) bool {
	return false
}

func isMatchIter(s, p string) bool {
	return false
}

func main() {

	fmt.Println(!isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("abcde", ".*"))
	fmt.Println(isMatch("abcde", ".*de"))
	fmt.Println(!isMatch("abcde", ".*dk"))

	fmt.Println("======================================")

	fmt.Println(!isMatchIter("aa", "a"))
	fmt.Println(isMatchIter("aa", "a*"))
	fmt.Println(isMatchIter("abcde", ".*"))
	fmt.Println(isMatchIter("abcde", ".*de"))
	fmt.Println(!isMatchIter("abcde", ".*dk"))

}
