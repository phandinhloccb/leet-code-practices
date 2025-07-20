package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func goodNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxSoFar int) int {
    if node == nil {
        return 0
    }

    count := 0
    if node.Val >= maxSoFar {
        count = 1
    }

    maxSoFar = max(node.Val, maxSoFar)

    count += dfs(node.Left, maxSoFar)
    count += dfs(node.Right, maxSoFar)

    return count
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	result := goodNodes(root)
	fmt.Println("Result:", result)
}
