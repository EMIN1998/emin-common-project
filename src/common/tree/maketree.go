package tree

import "fmt"

type TreeNode struct {
	Val   int
	index int
	Left  *TreeNode
	Right *TreeNode
}

// CreateTreeNode BFS的数组翻译为二叉树
func CreateTreeNode(nums []int) *TreeNode {
	var root *TreeNode
	var queue []*TreeNode
	for i := 0; i < len(nums); i += 2 {
		if i == 0 {
			root = &TreeNode{Val: nums[i], index: i}
			queue = append(queue, root)
		}
		parentNode := queue[0]
		queue = queue[1:]

		// 添加左节点
		if i+1 < len(nums) && nums[i+1] != -1 {
			parentNode.Left = &TreeNode{Val: nums[i+1], index: i + 1}
			queue = append(queue, parentNode.Left)
		}

		// 添加右节点
		if i+2 < len(nums) && nums[i+2] != -1 {
			parentNode.Right = &TreeNode{Val: nums[i+2], index: i + 2}
			queue = append(queue, parentNode.Right)
		}
	}
	return root
}

// PrintTree BFS 打印
func (tree *TreeNode) PrintTree() {
	head := tree
	resp := LevelOrder(head)
	fmt.Printf("BFS of Tree :%v", resp)
}

// 中后遍历构造二叉树
// 后序遍历的最后一个节点为根节点
// 每次递归将后序遍历中的根节点剪掉
func makeTreeByMiddleAndLast(mid []int, last *[]int) *TreeNode {
	if len(mid) == 0 || len(*last) == 0 {
		return nil
	}

	rootValue := (*last)[len(*last)-1]
	var index int
	for _, v := range mid {
		if v == rootValue {
			break
		}
		index++
	}

	left := mid[0:index]
	right := mid[index+1:]
	*last = (*last)[0 : len(*last)-1]

	rightTree := makeTreeByMiddleAndLast(right, last)
	leftTree := makeTreeByMiddleAndLast(left, last)
	root := &TreeNode{Val: rootValue, Right: rightTree, Left: leftTree}
	return root
}

// 大佬题解
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}

	i := 0
	for ; i < len(inorder); i++ {
		if root.Val == inorder[i] {
			break
		}
	}

	root.Left = buildTree(inorder[:i], postorder[:i]) // 不懂
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}

// 先序中序构造二叉树
func buildTreeV1(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	idx := 0 // 找出跟节点在中序遍历中的位置，则左边是左子树，右边是右子树
	for i, v := range inorder {
		if v == root.Val {
			idx = i
			break
		}
	}

	// 中序遍历中的分布[左子树，根节点，右子树]
	// 根节点中序遍历的下标，在先序遍历中是左子树的最后一个节点，因此左子树在先序中的位置是[根节点，左子树，右子树] => 所以取preorder[1:idx+1]
	// 后序遍历同理[左子树，右子树，根节点]
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return root
}
