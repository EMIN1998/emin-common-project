package sublist

import "common/utils"

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
