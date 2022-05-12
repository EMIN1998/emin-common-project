package sublist

import (
	"common/utils"
	"math"
	"sort"
	"strings"
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

// link:https://leetcode-cn.com/problems/word-break/submissions/
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	hasMap := make(map[string]bool)
	for _, ss := range wordDict {
		hasMap[ss] = true
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			str := s[i:j]
			if dp[i] && hasMap[str] {
				dp[j] = true
			}
		}
	}

	return dp[len(s)]
}

// 找出经过排列后可以将整个数组变为有序的无序数组
// link:https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/solution/zui-duan-wu-xu-lian-xu-zi-shu-zu-by-leet-yhlf/
func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}

	arr := append([]int{}, nums...)
	sort.Ints(arr)

	left, right := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		if arr[i] != nums[i] {
			break
		}

		left++
	}

	for arr[right] == nums[right] {
		right--
	}

	return right - left + 1
}

func findUnsortedSubarrayV1(nums []int) int {
	max := math.MinInt64
	min := math.MaxInt64

	nLen := len(nums)
	left := -1
	right := -1
	for idx, num := range nums {
		if num >= max {
			max = num
		} else {
			right = idx
		}

		if nums[nLen-idx-1] <= min {
			min = nums[nLen-idx-1]
		} else {
			left = nLen - idx - 1
		}
	}

	if right == -1 {
		return 0
	}

	return right - left + 1
}

// ===================================================================================================
// 回溯
// 根据数字和它对应的字母进行排列
// link: https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/submissions/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitList := strings.Split(digits, "")
	resp := make([]string, 0)

	tmpRes := ""
	var backTracking = func(index int) {}
	backTracking = func(index int) {
		if index == len(digits) {
			resp = append(resp, tmpRes)
			return
		}

		digit := digitList[index]
		for _, digit := range strings.Split(getCharactor(digit), "") {
			tmpRes += digit
			backTracking(index + 1)
			tmpRes = tmpRes[:len(tmpRes)-1]
		}
	}

	backTracking(0)

	return resp
}

func getCharactor(in string) string {

	switch in {
	case "1":
		return ""
	case "2":
		return "abc"
	case "3":
		return "def"
	case "4":
		return "ghi"
	case "5":
		return "jkl"
	case "6":
		return "mno"
	case "7":
		return "pqrs"
	case "8":
		return "tuv"
	case "9":
		return "wxyz"
	case "0":
		return ""
	}

	return ""
}

// ===================================================================================================
// link:https://leetcode-cn.com/problems/next-permutation/submissions/
// 找出元素排列紧邻的下一个排列
// 1. 倒序找出最近的紧邻升序组合
// 2. 找出nums[j:]中最小的，比nums[i]大的元素nums[target]
// 3. 交换nums[i]和nums[target]
// 4. 将nums[j:]进行从小到大排序
func nextPermutation(nums []int) {
	var i, j int
	i = len(nums) - 2
	j = i + 1
	for i >= 0 {
		if nums[i] >= nums[j] { // 找出第一个升序的相邻元素， 如果没找到说明已是最大的排列了
			i--
			j--
			continue
		}

		break
	}

	sortIndex := j
	if j == 0 { // 如果排列中的最后一个，下一个紧邻的就是第一个排列组合
		i = 0
		j = len(nums) - 1
		sortIndex = i + 1
	}

	minIdx := j
	for idx, v := range nums[j:] {
		if v < nums[minIdx] && v > nums[i] {
			minIdx = idx + j
		}
	}

	nums[i], nums[minIdx] = nums[minIdx], nums[i]

	sort.Ints(nums[sortIndex:])
}

// ===================================================================================================
// link:https://leetcode-cn.com/problems/combination-sum/
// 组合总和
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	resp := make([][]int, 0)
	tmpRes := make([]int, 0)
	var backtracking = func(surplus int, currCandidates []int) {}
	backtracking = func(surplus int, currCandidates []int) {
		if surplus <= 0 {
			if surplus == 0 {
				resp = append(resp, append([]int{}, tmpRes...))
			}

			return
		}

		for idx, v := range currCandidates {
			tmpRes = append(tmpRes, v)
			backtracking(surplus-v, currCandidates[idx:])
			tmpRes = tmpRes[:len(tmpRes)-1]
		}
	}

	backtracking(target, candidates)

	return resp
}
