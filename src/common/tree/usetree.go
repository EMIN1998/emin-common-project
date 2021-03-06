package tree

import (
	"math"
)

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

// LevelOrder 广度优先遍历BFS
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	resp := make([][]int, 0)
	level := []*TreeNode{root}
	for i := 0; len(level) > 0; i++ {
		resp = append(resp, []int{})
		nextLevel := make([]*TreeNode, 0)
		for _, node := range level {
			resp[i] = append(resp[i], node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		level = nextLevel
	}

	return resp
}

// 判断是否为合格的平衡二叉树
// 平衡二叉树左子树的所有节点比根节点小，右子树比根节点大
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBST(root, math.MaxInt64, math.MinInt64)
}

func isBST(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}

	left := isBST(root.Left, root.Val, min)
	right := isBST(root.Right, max, root.Val)

	if !(left && right) {
		return false
	}

	if root.Left != nil {
		if root.Left.Val > root.Val || root.Left.Val < min {
			return false
		}
	}

	if root.Right != nil {
		if root.Right.Val < root.Val || root.Right.Val > max {
			return false
		}
	}

	return true
}

// 中序遍历判断二叉搜索树
// 二叉搜索树中序遍历后每个节点都比后一个节点小
func isValidBSTV1(root *TreeNode) bool {
	nodeStack := make([]*TreeNode, 0)
	//nodeStack = append(nodeStack, root)
	min := math.MinInt64
	for len(nodeStack) != 0 || root != nil {
		for root != nil {
			nodeStack = append(nodeStack, root)
			root = root.Left
		}

		root = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[0 : len(nodeStack)-1]

		if root.Val <= min {
			return false
		}

		min = root.Val
		root = root.Right
	}

	return true
}

// 1～n 组成多少个二叉搜索树
// link：https://leetcode-cn.com/problems/unique-binary-search-trees/
// 题目要求是计算不同二叉搜索树的个数。为此，我们可以定义两个函数：
// G(n)G(n): 长度为 nn 的序列能构成的不同二叉搜索树的个数。
// F(i, n)F(i,n): 以 i 为根、序列长度为 n 的不同二叉搜索树个数 (1≤i≤n)。
// 动态规划
func numTrees(n int) int {
	g := make(map[int]int)
	g[0], g[1] = 0, 1

	for i := 2; i <= n; i++ { // i 为前i个节点
		for j := 1; j <= i; j++ { // j 为根节点的下标
			g[i] += g[j-1] * g[i-j]
		}
	}

	return g[n]
}
