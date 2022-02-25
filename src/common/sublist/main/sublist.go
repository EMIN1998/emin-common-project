package main

import "fmt"

// 最长回文字串
func palindromeMax(s string) (string, int) {
	stateMap := make(map[int]map[int]bool)
	for t := 0; t <= len(s); t++ {
		stateMap[t] = make(map[int]bool)
	}

	var maxLen, index int
	for j := 0; j < len(s); j++ {
		for i := 0; i <= len(s)-1 && i < j; i++ {
			if s[i] != s[j] {
				stateMap[i][j] = false
			} else {
				if j-i < 3 {
					stateMap[i][j] = true
				} else {
					// 这里会用到[i+1][j-1],所以把j放在外循环的话使用[i+1][j-1]肯定是有值的，因为外循环在j-1是就已经赋值过了
					stateMap[i][j] = stateMap[i+1][j-1]
				}
			}

			if stateMap[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				index = i
			}
		}
	}

	return s[index : index+maxLen], maxLen
}

func main() {
	var s = "babcaaccabab"
	fmt.Print(palindromeMax(s))
}
