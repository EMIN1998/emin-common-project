package main

import "fmt"

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

func main() {
	var req = []int{9, -12, 1, 4, -1, 5, -2}
	res := FindGreatestSumOfSubArray(req)
	fmt.Print(res)
}
