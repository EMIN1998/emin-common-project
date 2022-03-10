package tree

import (
	"fmt"
)

func Test_main() {
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

/**
 * @desc 返回最近父节点
 * @param root TreeNode类
 * @param o1 int整型
 * @param o2 int整型
 * @return int整型
 */
func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	if root == nil { // 如果节点为空，返回-1
		return -1
	}

	if root.Val == o1 || root.Val == o2 { // 当o1或o2 被找到时返回当前root
		return root.Val
	}

	res1 := lowestCommonAncestor(root.Left, o1, o2)
	res2 := lowestCommonAncestor(root.Right, o1, o2)

	if res1 == -1 { // 如果当前节点左子树返回-1，则该子树下没有目标节点，目标在右子树
		return res2
	}
	if res2 == -1 { // 如果当前节点右子树返回-1，则该子树下没有目标节点，目标在左子树
		return res1
	}

	return root.Val // 找到了某个目标节点，返回左子树和右子树的公共父节点
}

// kthSmallest 中序遍历，找出第K小的节点
func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}
}
