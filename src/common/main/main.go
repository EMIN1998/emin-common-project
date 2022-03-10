package main

import (
	"common/tree"
	"fmt"
)

/**
 *
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	var tmpRes, max int
	for _, a := range array {
		if tmpRes < 0 {
			tmpRes = a
		} else {
			tmpRes += a
		}

		if max < tmpRes {
			max = tmpRes
		}

	}

	return max
}

func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

func main() {
	var param = []int{3, 2, 1, 5, 6, 4, 8}
	root := tree.CreateTreeNode(param)
	resp := tree.LevelOrder(root)
	//resp := sort.HeapSort(param)
	fmt.Print(resp)
	return
}
