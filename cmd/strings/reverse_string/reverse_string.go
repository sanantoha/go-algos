package main

import "fmt"

func reverseString(arr []rune) {

}

func main() {

	s := []rune{'h', 'e', 't', 'l', 'l', 'o'}

	reverseString(s)

	for _, c := range s {
		fmt.Printf("%v ", string(c))
	}
}
