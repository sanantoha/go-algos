package main

import "fmt"

func balancedBrackets(str string) bool {
	return false
}

func main() {

	input := "([])(){}(())()()"

	fmt.Println(balancedBrackets(input))

	fmt.Println(!balancedBrackets("}{"))
}
