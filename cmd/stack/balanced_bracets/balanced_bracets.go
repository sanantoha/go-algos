package main

import "fmt"

// O(n) time | O(n) space
func balancedBrackets(str string) bool {

	stack := make([]rune, 0)

	openCloseMap := make(map[rune]rune, 3)
	openCloseMap['('] = ')'
	openCloseMap['{'] = '}'
	openCloseMap['['] = ']'

	closeMap := make(map[rune]bool, 3)
	closeMap[')'] = true
	closeMap['}'] = true
	closeMap[']'] = true

	for _, c := range str {
		if val, exists := openCloseMap[c]; exists {
			stack = append(stack, val)
		} else if _, exists := closeMap[c]; exists {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != c {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {

	input := "([])(){}(())()()"

	fmt.Println(balancedBrackets(input))

	fmt.Println(!balancedBrackets("}{"))
}
