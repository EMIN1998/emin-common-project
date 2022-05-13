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
		tail := make([]int, 0)
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

// link:https://leetcode-cn.com/problems/merge-intervals/
// 合并数组区间
func merge(intervals [][]int) [][]int {
	resp := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}

		return false
	})

	resp = append(resp, intervals[0])
	for _, v := range intervals {
		last := resp[len(resp)-1]
		currLeft := v[0]
		currRight := v[1]
		if last[1] > currLeft {
			last[1] = max(currRight, last[1])
		} else {
			resp = append(resp, v)
		}
	}

	return resp
}

// 寻找旋转数组
// link: https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
func search(arr []int, target int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		return -1
	}

	pre := 0
	sub := len(arr) - 1
	for pre <= sub {
		mid := (sub + pre) / 2
		if arr[mid] == target {
			return mid
		} else if arr[pre] < arr[mid] { // 前半段有序
			if target >= arr[pre] && target < arr[mid] {
				sub = mid - 1
			} else {
				pre = mid + 1
			}
		} else { // 后半段有序
			if mid+1 >= arrLen || (target >= arr[mid+1] && target <= arr[sub]) {
				pre = mid + 1
			} else {
				sub = mid - 1
			}
		}
	}

	return -1
}

// ==========================================================
// 接雨水
// link:https://leetcode-cn.com/problems/trapping-rain-water/submissions/
func trap(height []int) int {
	sum := 0
	for idx, h := range height {
		// 找出最小左墙
		leftMax := 0
		for i := idx - 1; i >= 0; i-- {
			if height[i] > leftMax {
				leftMax = height[i]
			}
		}

		rightMax := 0
		for j := idx + 1; j < len(height); j++ {
			if height[j] > rightMax {
				rightMax = height[j]
			}
		}

		minSide := min(leftMax, rightMax)
		if h < minSide {
			cur := minSide - h
			sum += cur
		}
	}

	return sum
}

// 动态规划法
func trapV1(height []int) int {
	sum := 0
	maxLeftSide := make([]int, len(height))
	maxRightSide := make([]int, len(height))

	// 找出每个节点最高的左墙
	for i := 1; i < len(height); i++ {
		maxLeftSide[i] = max(maxLeftSide[i-1], height[i-1])
	}

	// 找出每个节点最高的右墙
	for j := len(height) - 2; j >= 0; j-- {
		maxRightSide[j] = max(maxRightSide[j+1], height[j+1])
	}

	for k := 1; k < len(height); k++ {
		minSide := min(maxLeftSide[k], maxRightSide[k])
		// 只要当前节点比最矮的墙还要低，那么就可以接到雨水
		if height[k] < minSide {
			cur := minSide - height[k]
			sum += cur
		}
	}

	return sum
}

// 双指针
func trapV2(height []int) int {
	leftMax, rightMax := 0, 0
	left, right := 0, len(height)-1
	sum := 0
	for left < right {
		if height[left] < height[right] { // 右边比左边高 left++
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				sum += leftMax - height[left]
			}
			left++
		} else { // 左边比右边高， right--
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				sum += rightMax - height[right]
			}
			right--
		}
	}

	return sum
}

// ==========================================================
// link :https://leetcode-cn.com/problems/minimum-path-sum/
// 64. 最小路径和
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	resMap := make([][]int, len(grid))
	m := len(grid) - 1    // 这是列
	n := len(grid[0]) - 1 // 这是行

	for i := 0; i < len(grid); i++ {
		resMap[i] = make([]int, len(grid[0]))
	}

	resMap[0][0] = grid[0][0]
	for i := 1; i <= m; i++ {
		resMap[0][i] = resMap[0][i-1] + grid[0][i]
	}

	if len(grid) == 1 {
		return resMap[0][n]
	}

	for j := 1; j <= n; j++ {
		resMap[j][0] = resMap[j-1][0] + grid[j][0]
	}

	if len(grid[0]) == 1 {
		return resMap[m][0]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			resMap[i][j] = min(resMap[i-1][j], resMap[i][j-1]) + grid[i][j]
		}
	}

	return resMap[m][n]
}

// ====================================================================================================================
// link : https://leetcode-cn.com/problems/unique-paths/
// 62. 不同路径
func uniquePaths(m int, n int) int {
	if m == 0 && n == 0 {
		return 0
	}

	resMap := make([][]int, m)
	for i := 0; i < m; i++ {
		resMap[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		resMap[i][0] = 1
	}

	for j := 0; j < n; j++ {
		resMap[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			resMap[i][j] = resMap[i-1][j] + resMap[i][j-1]
		}
	}

	return resMap[m-1][n-1]
}

// ====================================================================================================================
// 48. 旋转图像
// link:https://leetcode.cn/problems/rotate-image/
func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = tmp

		}
	}
}

// ====================================================================================================================
// link : https://leetcode.cn/problems/jump-game/solution/55-by-ikaruga/
// 55. 跳跃游戏
// 思路：从倒数第二个元素开始，算该元素能否到达下一个下标，如果可以就把下标往前推
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	// dp := make([]int, len(nums)-1)
	m := len(nums) - 1
	for i := m - 1; i >= 0; i-- {
		if i+nums[i] < m {
			continue
		}

		m = i
	}

	if m > 0 {
		return false
	}

	return true
}

// ====================================================================================================================
// link:https://leetcode.cn/problems/single-number/solution/zhi-chu-xian-yi-ci-de-yuan-su-by-disco-2-q634/
// 136. 只出现一次的数字, 用异或
// a^0=a。
// a^0a=0。
// 异或运算满足交换律和结合律

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
