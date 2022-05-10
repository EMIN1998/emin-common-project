package binarysearch

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "searchRange", args: args{
			nums:   []int{1, 2, 3, 4, 5, 9, 10, 11, 1, 2, 13, 14, 15, 16, 17, 18, 19, 20},
			target: 8,
		}, want: []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
