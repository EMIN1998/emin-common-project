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
