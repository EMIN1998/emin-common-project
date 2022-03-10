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
