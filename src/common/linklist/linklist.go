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
