package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(h) time | O(1) space
func findInOrderSuccessor(root *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}

	if root.Right != nil {
		return findMostLeftChild(root.Right)
	}
	return findMostParentRight(root)
}

func findMostParentRight(node *tree.Node) *tree.Node {
	curr := node

	for curr.Parent != nil && curr.Parent.Right.Key == curr.Key {
		curr = curr.Parent
	}

	return curr.Parent
}

func findMostLeftChild(node *tree.Node) *tree.Node {
	curr := node

	for curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

func main() {
	node5 := &tree.Node{Key: 5}
	node9 := &tree.Node{Key: 9}
	node11 := &tree.Node{Key: 11}
	node12 := &tree.Node{Key: 12}
	node14 := &tree.Node{Key: 14}
	node20 := &tree.Node{Key: 20}
	node25 := &tree.Node{Key: 25}

	node5.Parent = node9

	node9.Left = node5
	node9.Right = node12
	node9.Parent = node20

	node11.Parent = node12

	node12.Left = node11
	node12.Right = node14
	node12.Parent = node9

	node14.Parent = node12

	node20.Left = node9
	node20.Right = node25

	node25.Parent = node20

	//         20
	//       /    \
	//      9      25
	//    /   \
	//   5     12
	//        /   \
	//       11   14

	fmt.Println(node20)
	fmt.Println(findInOrderSuccessor(node5).Key == 9)
	fmt.Println(findInOrderSuccessor(node9).Key == 11)
	fmt.Println(findInOrderSuccessor(node11).Key == 12)
	fmt.Println(findInOrderSuccessor(node12).Key == 14)
	fmt.Println(findInOrderSuccessor(node14).Key == 20)
	fmt.Println(findInOrderSuccessor(node20).Key == 25)
	fmt.Println(findInOrderSuccessor(node25) == nil)
	fmt.Println(findInOrderSuccessor(node20).Key == 25)
}
