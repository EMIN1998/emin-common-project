package linklist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) PrintLinkList() {
	resp := make([]int, 0)
	next := l
	for next != nil {
		resp = append(resp, next.Val)
		next = next.Next
	}

	fmt.Print(resp)
}

// CreateLinkList 创建链表
func CreateLinkList(param []int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	next := head
	for _, val := range param {
		next.Next = &ListNode{
			Val: val,
		}

		next = next.Next
	}

	return head.Next
}

// SortList 链表归并排序
func SortList(root *ListNode) *ListNode {
	return sort(root)
}

func sort(root *ListNode) *ListNode {
	if root == nil {
		return nil
	}

	left, right := findMiddleNode(root)
	if right == nil {
		return left
	}

	return merge(sort(left), sort(right))
}

func merge(left, right *ListNode) *ListNode {
	head := &ListNode{
		Next: nil,
	}

	resp := head
	for left != nil || right != nil {
		if left == nil {
			resp.Next = right
			return head.Next
		}

		if right == nil {
			resp.Next = left
			return head.Next
		}

		if left.Val < right.Val {
			resp.Next = left
			left = left.Next

		} else {
			resp.Next = right
			right = right.Next
		}

		resp = resp.Next
	}

	return head.Next
}

func findMiddleNode(root *ListNode) (*ListNode, *ListNode) {
	if root == nil {
		return nil, nil
	}

	var slow, fast *ListNode
	slow = root
	fast = root

	pre := slow
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	last := slow.Next
	slow.Next = nil

	return pre, last
}

// RemoveNthFromEnd 删除倒数第N个元素
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

// 题解：https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/142-huan-xing-lian-biao-ii-jian-hua-gong-shi-jia-2/
// 寻找链表中的环
// 为何慢指针第一圈走不完一定会和快指针相遇? 可以认为快指针和慢指针是相对运动的，假设慢指针的速度是 1节点/秒，快指针的速度是 2节点/秒，当以慢指针为参考系的话（即慢指针静止），快指针的移动速度就是 1节点/秒，所以肯定会相遇。
// 为什么在第一圈就会相遇呢？ 设环的长度为 L，当慢指针刚进入环时，慢指针需要走 L 步(即 L 秒)才能走完一圈，此时快指针距离慢指针的最大距离为 L-1，我们再次以慢指针为参考系，如上所说，快指针在按照1节点/秒的速度在追赶慢指针，所以肯定能在 L 秒内追赶到慢指针。
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			tmp := head
			for tmp != slow {
				tmp = tmp.Next
				slow = slow.Next
			}

			return tmp
		}
	}

	return nil
}

// ====================================================================================================================
// 160. 相交链表 --- 双指针法
// link:https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 思路：
// 指针a 遍历A（a+c个节点，c为共同节点的长度），如果遍历完就指向B（b+c个节点，c为共同节点的长度）继续遍历
// 指针b同理
// 当两个指针分别遍历完A和B后 a和b都遍历了a+b+c个节点
// 因此a和b相遇在相交节点或者分别到两个节点的末尾空节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}

	}

	return a
}
