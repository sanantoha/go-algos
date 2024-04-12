package main

import "fmt"

func oneEdit(s1, s2 string) bool {

	return false
}

func main() {

	stringOne := "hello"
	stringTwo := "helo"
	actual := oneEdit(stringOne, stringTwo)
	fmt.Println(actual)
}
