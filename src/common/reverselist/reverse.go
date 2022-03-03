package reverselist

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 * 递归法
 */
func ReverseList(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	tmp := ReverseList(pHead.Next)
	pHead.Next.Next = pHead
	pHead = nil

	return tmp
}

func ReverseListV1(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	var nHead *ListNode
	for {
		if pHead == nil {
			break
		}

		tmp := pHead.Next
		pHead.Next = nHead
		nHead = pHead
		pHead = tmp

	}

	return nHead
}

func isPalindrome(head *ListNode) bool {
	// 1. 找到中间节点
	node := head
	midNode := findMiddleNode(head)
	// 2. 反转后半节点
	lastHalf := reverseListNode(midNode.Next)
	node1 := lastHalf
	// 3. 迭代对比
	var res = true
	for res && node1 != nil {
		if node1.Val != node.Val {
			res = false
		}

		node1 = node1.Next
		node = node.Next
	}
	// 4. 将后半段复原
	midNode.Next = reverseListNode(lastHalf)
	return res
}

func findMiddleNode(node *ListNode) *ListNode {
	var slow, fast *ListNode
	slow = node
	fast = node

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseListNode(node *ListNode) *ListNode {
	var cur = node
	var pre *ListNode
	for cur != nil {

		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}
