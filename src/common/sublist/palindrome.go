package sublist

import (
	"common/utils"
	"strings"
)

// 最长回文字串
// 动态规划
func longestPalindrome(s string) string {
	str := strings.Split(s, "")
	if len(str) <= 1 {
		return s
	}

	statusMap := make([][]bool, len(s))

	for index, _ := range str {
		statusMap[index] = make([]bool, len(s))
		statusMap[index][index] = true
	}

	maxLen, preIdx := 1, 0

	for subLen := 2; subLen <= len(str); subLen++ {
		for idx, _ := range str {
			idy := subLen + idx - 1

			if idy > len(str)-1 {
				break
			}

			if str[idx] != str[idy] {
				statusMap[idx][idy] = false
			} else {
				if idy-idx < 3 {
					statusMap[idx][idy] = true
				} else {
					statusMap[idx][idy] = statusMap[idx+1][idy-1]
				}
			}

			if statusMap[idx][idy] && idy-idx+1 > maxLen {
				maxLen = idy - idx + 1
				preIdx = idx
			}
		}
	}

	return strings.Join(str[preIdx:preIdx+maxLen], "")
}

// 中心扩展算法
func longestPalindromeV1(s string) string {
	defer utils.TimeCost() // 耗时
	start, end := 0, 0
	for i, _ := range s {
		left1, right1 := expendAndIndex(s, i, i)
		left2, right2 := expendAndIndex(s, i, i+1)

		if right1-left1 > end-start {
			start, end = left1, right1
		}

		if right2-left2 > end-start {
			start, end = left2, right2
		}

	}

	return s[start : end+1]
}

func expendAndIndex(s string, i int, i2 int) (int, int) {
	l := len(s)
	for i >= 0 && i2 < l && s[i] == s[i2] {
		i--
		i2++
	}

	return i - 1, i2 + 1
}

// link:https://leetcode-cn.com/problems/palindromic-substrings/
func countSubstrings(s string) int {
	resp := 0
	strMap := make([][]bool, len(s))

	for i, _ := range s {
		strMap[i] = make([]bool, len(s))
		strMap[i][i] = true
		resp++
	}

	// idy-idx + 1 = subLen
	// idy = subLen + idx - 1

	for subLen := 2; subLen <= len(s); subLen++ {
		for idx := 0; idx < len(s); idx++ {
			idy := subLen + idx - 1

			if idy > len(s)-1 {
				break
			}

			if s[idx] != s[idy] {
				strMap[idx][idy] = false
			} else {
				if idy-idx+1 < 3 || strMap[idx+1][idy-1] {
					strMap[idx][idy] = true
					resp++
				}
			}
		}
	}

	return resp
}
