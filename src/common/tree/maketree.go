package tree

type TreeNode struct {
	Val   int
	index int
	Left  *TreeNode
	Right *TreeNode
}

// BFS的数组翻译为二叉树
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
