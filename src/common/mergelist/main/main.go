package main

import "fmt"

/**
 *
 * @param A int整型一维数组
 * @param B int整型一维数组
 * @return void
 */
func merge(A []int, m int, B []int, n int) {
	// write code here
	var i, j, p int
	i = m - 1
	j = n - 1
	p = m + n - 1

	for {
		if p < 0 || j < 0 {
			break
		}

		if i < 0 {
			A = append(A, B...)
			break
		}
		if A[i] < B[j] {
			A[p] = B[j]
			p--
			j--
		} else {
			A[p] = A[i]
			p--
			i--
		}

		if i == p {
			break
		}

		if i < 0 && j >= 0 {
			for {
				A[p] = B[j]
				j--
				p--
				if p < 0 {
					break
				}
			}
		}

		if j < 0 && i >= 0 {
			for {
				A[p] = A[i]
				i--
				p--
				if p < 0 {
					break
				}
			}
		}
	}

	fmt.Print(A)
	return
}

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	cur := head
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
			}

			cur = cur.Next
		} else if list1 != nil {
			cur.Next = list1
			break
		} else {
			cur.Next = list2
			break
		}
	}

	return head.Next
}

func main() {
	a := []int{}
	b := []int{1}
	merge(a, 0, b, 1)
}
