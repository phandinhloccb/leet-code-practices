package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := getLeaves(root1, []int{})
	leaves2 := getLeaves(root2, []int{})

	fmt.Println("leaves1", leaves1)
	fmt.Println("leaves2", leaves2)

	return reflect.DeepEqual(leaves1, leaves2)
}

func getLeaves(root *TreeNode, leaves []int) []int {
	if root == nil {
		return leaves
	}

	if root.Left == nil && root.Right == nil {
		leaves = append(leaves, root.Val)
	}

	leaves = getLeaves(root.Left, leaves)
	leaves = getLeaves(root.Right, leaves)

	return leaves
}


func main() {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	result := leafSimilar(root1, root2)
	fmt.Println("Result:", result)
}
