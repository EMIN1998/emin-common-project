package sublist

import (
	"common/utils"
)

// 最长连续子数组长度
func longestConsecutive(nums []int) int {
	valueSet := make(map[int]struct{})
	for _, v := range nums {
		valueSet[v] = struct{}{}
	}

	longestLen := 0
	for k, _ := range valueSet {
		if _, ok := valueSet[k-1]; ok {
			continue
		}

		cur, currLen := k, 0
		for {
			if _, ok := valueSet[cur]; !ok {
				if longestLen > len(valueSet)/2 {
					return longestLen
				}

				break
			}

			cur++
			currLen++

			if currLen > longestLen {
				longestLen = currLen
			}
		}
	}

	return longestLen
}

// 最长递增子序列
// dp 为以第 i 个数字结尾的最长上升子序列的长度
func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return 1
	}

	var max int
	stateMap := make([]int, len(nums), 1)
	for idx, i := range nums {
		for idj, j := range nums[0 : idx+1] {
			if j < i {
				stateMap[idx] = utils.IF(stateMap[idx] < stateMap[idj], stateMap[idj]+1, stateMap[idx]).(int)
			}
		}

		if stateMap[idx] > max {
			max = stateMap[idx]
		}
	}

	return max
}

// LengthOfLongestSubstring 最长不重复的子序列的长度--滑动窗口
func LengthOfLongestSubstring(s string) int {
	maxLen := 0
	var left, right int
	windowMap := make(map[byte]int)
	n := len(s)
	for right < n {
		// rightValueIndex := -1
		if val, ok := windowMap[s[right]]; ok {
			if val >= left {
				left = val + 1
			}
		}

		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}

		windowMap[s[right]] = right
		right++
	}

	return maxLen
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		ans = append(ans, track)
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return ans
}

// 子集和全排列对比：
// 子集需要记录每个节点的结果
// 全排列只在叶子节点进行记录
// 因此track在子集函数中是全局变量
// 在全排列中是传递的参数
func subses(nums []int) [][]int {
	resp := make([][]int, 0)
	track := make([]int, 0)
	var backtracking func(start int)

	backtracking = func(start int) {
		// 每次进行回溯前，当前节点就已经是一个子集了，先收集起来
		// 这里防止指针传递
		resp = append(resp, append([]int{}, track...))
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtracking(i + 1)
			track = track[:len(track)-1]
		}
	}

	return resp
}

// 全排列
func permute(nums []int) [][]int {
	resp := make([][]int, 0)
	used := make([]bool, len(nums))

	var backtracking func(track []int)
	backtracking = func(track []int) {
		if len(track) == len(nums) {
			resp = append(resp, append([]int{}, track...))
			return
		}

		for i, v := range nums {
			if used[i] {
				continue
			}

			track = append(track, v)
			used[i] = true
			backtracking(track)
			track = track[:len(track)-1]
			used[i] = false // 剪枝后该接节点已不再当前的路径中，但是下一条路径依然可以用
		}
	}

	backtracking([]int{})
	return resp
}

// 生成括号
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var df func(left, right int, item string)
	df = func(left, right int, item string) {
		if left == right && left == n {
			res = append(res, item)
			return
		}

		if left < right || left > n || right > n {
			return
		}

		df(left+1, right, item+"(")
		df(left, right+1, item+")")
	}

	df(0, 0, "")
	return res
}
