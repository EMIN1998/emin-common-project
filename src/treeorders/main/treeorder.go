package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func threeOrders(root *TreeNode) [][]int {
	resp := make([][]int, 0)
	if root == nil || root.Val == 0 {
		val := make([]int, 0)
		resp = append(resp, val)
		resp = append(resp, val)
		resp = append(resp, val)
		return resp
	}

	pre := preSort(root)
	mid := midSort(root)
	last := lastSort(root)

	resp = append(resp, pre)
	resp = append(resp, mid)
	resp = append(resp, last)

	return resp
}

func preSort(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var leftRes, rightRes []int

	if root.Left != nil {
		leftRes = preSort(root.Left)
	}

	if root.Right != nil {
		rightRes = preSort(root.Right)
	}

	resp := make([]int, 0)
	resp = append(resp, root.Val)
	resp = append(resp, leftRes...)
	resp = append(resp, rightRes...)
	return resp
}

func midSort(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var leftRes, rightRes []int

	if root.Left != nil {
		leftRes = midSort(root.Left)
	}

	if root.Right != nil {
		rightRes = midSort(root.Right)
	}

	resp := make([]int, 0)
	resp = append(resp, leftRes...)
	resp = append(resp, root.Val)
	resp = append(resp, rightRes...)
	return resp
}

func lastSort(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var leftRes, rightRes []int

	if root.Left != nil {
		leftRes = lastSort(root.Left)
	}

	if root.Right != nil {
		rightRes = lastSort(root.Right)
	}

	resp := make([]int, 0)
	resp = append(resp, leftRes...)
	resp = append(resp, rightRes...)
	resp = append(resp, root.Val)

	return resp
}

func main() {
	valLeft := &TreeNode{
		Val: 3,
	}

	valRight := &TreeNode{
		Val: 5,
	}

	valRoot := &TreeNode{
		Val:   2,
		Left:  valLeft,
		Right: valRight,
	}

	res := threeOrders(valRoot)
	fmt.Print(res)
}
