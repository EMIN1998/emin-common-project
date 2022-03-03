package tree

import (
	"fmt"
)

func test_main() {
	root := CreateTreeNode([]int{1, 2, 3, 4, 5})
	res := diameterOfBinaryTree(root)
	fmt.Println(res)
}

// 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0
	if root == nil {
		return maxDia
	}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lh := dfs(node.Left)
		rh := dfs(node.Right)
		maxDia = max(maxDia, lh+rh)
		return 1 + max(lh, rh)
	}
	dfs(root)
	return maxDia
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
