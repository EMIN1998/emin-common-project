package tree

import (
	"reflect"
	"testing"
)

func TestCreateTreeNode(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTreeNode(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeTreeByMiddleAndLast(t *testing.T) {
	type args struct {
		mid  []int
		last []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{args: args{
			mid: []int{6, 3, 5, 8, 4, 2}, last: []int{6, 5, 3, 4, 2, 8},
		}, name: "emin-test", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.mid, tt.args.last); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeTreeByMiddleAndLast() = %v, want %v", got, tt.want)
				got.PrintTree()
			}

		})
	}
}

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{root: CreateTreeNode([]int{5, 4, 6, -1, -1, 3, 7})}, name: "是否为合格二叉搜索树", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
