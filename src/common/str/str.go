package str

import "strconv"

//link :https://leetcode-cn.com/problems/decode-string/solution/
// 字符串解码v
func decodeString(s string) string {
	nums := make([]int, 0)
	stack := make([]byte, 0)
	i := 0
	for ; i < len(s); i++ {
		if isDigit(s[i]) {
			j := i + 1
			for isDigit(s[j]) {
				j++
			}

			num, _ := strconv.Atoi(s[i:j])
			nums = append(nums, num)
			i = j - 1
		} else if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			var str, tmpStr string
			for stack[len(stack)-1] != '[' {
				str = string(stack[len(stack)-1]) + str
				stack = stack[:len(stack)-1]
			}
			num := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			stack = stack[:len(stack)-1]
			for num != 0 {
				tmpStr += str
				num--
			}

			stack = append(stack, tmpStr...)
		}
	}

	return string(stack)
}

func isDigit(v byte) bool {
	if v >= '0' && v <= '9' {
		return true
	}

	return false
}
