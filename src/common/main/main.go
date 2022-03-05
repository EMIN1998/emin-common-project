package main

import "common/linklist"

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

const (
	cacheInit = "LRUCache"
	put       = "put"
	get       = "get"
)

func main() {
	//var req = []int{9, -12, 1, 4, -1, 5, -2}
	//res := FindGreatestSumOfSubArray(req)
	//fmt.Print(res)

	//res := countBits(5)
	//fmt.Println(res)
	//defer timeCost()()
	//
	//var s = "cacdcae"
	//str := longestPalindrome1(s)
	//fmt.Println(str)

	head := linklist.CreateLinkList([]int{4, 2, 1, 3})
	head.PrintLinkList()
	resp := linklist.SortList(head)
	resp.PrintLinkList()
	return
}

// 最长子序列
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
