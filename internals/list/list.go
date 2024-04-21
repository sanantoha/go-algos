package list

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	builder := strings.Builder{}

	curr := head

	for curr != nil {
		builder.WriteString(fmt.Sprintf("%d", curr.Val))
		if curr.Next != nil {
			builder.WriteString(" -> ")
		}

		curr = curr.Next
	}

	return builder.String()
}
