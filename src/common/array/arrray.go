package array

import (
	"fmt"
	"sort"
)

//link：https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	resp := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//mid := nums[i+1]
		k := len(nums) - 1
		target := nums[i] * -1
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for j < k && nums[j]+nums[k] > target {
				k--
			}

			if j == k {
				break
			}

			if nums[j]+nums[k] == target {
				resp = append(resp, []int{nums[i], nums[j], nums[k]})
			}

		}
	}

	return resp
}

//link:https://leetcode-cn.com/problems/product-of-array-except-self/solution/
func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	leftRes := make([]int, len(nums))
	rightRes := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		tmpRes := 1
		for j := 0; j < i; j++ {
			tmpRes *= nums[j]
		}

		leftRes[i] = tmpRes
	}

	for i := len(nums) - 1; i >= 0; i-- {
		tmpRes := 1
		for j := len(nums) - 1; j > i; j-- {
			tmpRes *= nums[j]
		}

		rightRes[i] = tmpRes
	}

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = leftRes[i] * rightRes[i]
	}

	return ans
}

// link:https://leetcode-cn.com/problems/maximal-square/
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	tmpMax := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix))

		for j := 0; j < len(matrix); j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == '1' {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}

				tmpMax = max(tmpMax, dp[i][j])
				continue
			}

			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i][j-1])) + 1
			} else {
				dp[i][j] = 0
			}

			tmpMax = max(tmpMax, dp[i][j])
		}
	}

	return tmpMax * tmpMax
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func topKFrequent(nums []int, k int) []int {
	tmpMap := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = 1
		} else {
			tmpMap[v]++
		}
	}

	resp := make([]int, 0)
	if len(tmpMap) <= k {
		for k, _ := range tmpMap {
			resp = append(resp, k)
		}
		return resp
	}

	arr := make([]int, 0)
	for k, _ := range tmpMap {
		arr = append(arr, k)
	}

	var quickSourt func(arr []int, k int) []int
	quickSourt = func(arr []int, k int) []int {
		fmt.Println(arr)
		if k <= 0 {
			return nil
		}

		if len(arr) < 2 {
			return arr
		}

		mid := arr[0]
		pre := make([]int, 0)
		tail := make([]int,0)
		for _, v := range arr[1:] {
			if tmpMap[v] > tmpMap[mid] {
				pre = append(pre, v)
			} else {
				tail = append(tail, v)
			}
		}

		forw := quickSourt(pre, len(pre))
		if len(forw) > k {
			return forw[0:k]
		}

		back := quickSourt(tail, k-len(forw)-1)
		quickResp := make([]int, 0)
		quickResp = append(quickResp, forw...)
		quickResp = append(quickResp, mid)
		quickResp = append(quickResp, back...)
		return quickResp[0:k]

		//return append(append(append([]int{}, quickSourt(pre, k)...), mid), quickSourt(tail, k- len(pre))...)
	}

	return quickSourt(arr, k)
}

