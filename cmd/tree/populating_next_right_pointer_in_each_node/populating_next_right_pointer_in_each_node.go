package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func connect(root *tree.Node) *tree.Node {
	return nil
}

func main() {

	root := &tree.Node{
		Key: 1,
		Left: &tree.Node{
			Key: 2,
			Left: &tree.Node{
				Key: 4,
			},
			Right: &tree.Node{
				Key: 5,
			},
		},
		Right: &tree.Node{
			Key: 3,
			Left: &tree.Node{
				Key: 6,
			},
			Right: &tree.Node{
				Key: 7,
			},
		},
	}

	res := connect(root)

	fmt.Println(res)
}
