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

// 迭代法链表反转
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

// 精简迭代反转链表
func reverseLinkList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	last := reverseLinkList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return last
}

// =======================================================
// 判断链表元素是否对称
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

// 快慢指针找出中间节点
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

// 链表反转
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

// =======================================================
// 反转链表的left 到right 位置的节点
// link:https://leetcode-cn.com/problems/reverse-linked-list-ii/submissions/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 { // 当left == 1时就成了反转前N个节点
		return reverseN(head.Next, right)
	}

	head.Next = reverseBetween(head, left-1, right-1)
	return head
}

// 反转链表的前N个元素
var succ *ListNode // 用来记录所有的剩余节点，反转后原来的头节点需要指向succ
func reverseN(head *ListNode, m int) *ListNode {
	if m == 1 {
		succ = head.Next
		return head
	}

	last := reverseN(head.Next, m-1)
	head.Next.Next = head
	head.Next = succ // 按递归的思路，每次递归都当作最后一次，因此每次递归的原头节点象征性指向succ
	return last
}
