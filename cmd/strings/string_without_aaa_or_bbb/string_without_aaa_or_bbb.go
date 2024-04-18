package main

import (
	"fmt"
	"strings"
)

// O(a + b) time | O(a + b) space
func stringWithoutAAAorBBB(a int, b int) string {
	var builder = strings.Builder{}

	for 0 < a && 0 < b {
		if a > b {
			builder.WriteString("aab")
			a -= 2
			b--
		} else {
			if a == b && a == 1 {
				builder.WriteString("ab")
				a--
				b--
				break
			}

			builder.WriteString("abb")
			a--
			b -= 2
		}
	}

	for 0 < a {
		builder.WriteString("a")
		a--
	}

	prefix := strings.Builder{}
	for 0 < b {
		prefix.WriteString("b")
		b--
	}
	prefix.WriteString(builder.String())
	return prefix.String()
}

func main() {

	fmt.Println(stringWithoutAAAorBBB(1, 1)) // ab
	fmt.Println(stringWithoutAAAorBBB(3, 3)) // abbaab
	fmt.Println(stringWithoutAAAorBBB(2, 5)) // babbabb

	fmt.Println(stringWithoutAAAorBBB(5, 3)) // aabaabab

	fmt.Println(stringWithoutAAAorBBB(3, 3)) // abbaab

	fmt.Println(stringWithoutAAAorBBB(1, 4)) // bbabb
}
