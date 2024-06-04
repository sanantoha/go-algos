package main

import (
	"fmt"
	"reflect"
)

func longestCommonSubsequence(str1, str2 string) []rune {
	return nil
}

func longestCommonSubsequence1(str1, str2 string) []rune {
	return nil
}

func main() {

	expected := []rune{'X', 'Y', 'Z', 'W'}

	actual := longestCommonSubsequence("ZXVVYZW", "XKYKZPW")
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))

	actual = longestCommonSubsequence1("ZXVVYZW", "XKYKZPW")
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))
}
