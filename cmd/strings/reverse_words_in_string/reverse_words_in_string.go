package main

import "fmt"

func reverse(str string) string {
	return ""
}

func main() {

	s := "the sky is blue"

	fmt.Println(reverse(s))

	fmt.Printf("|%s|", reverse(" a good   example     "))

	fmt.Println(reverse("this      string     has a     lot of   whitespace"))
}
