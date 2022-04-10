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
