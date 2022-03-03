package main

import (
	"common/lru"
	"encoding/json"
	"fmt"
)

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

	op := "[\"LRUCache\",\"put\",\"put\",\"get\",\"put\",\"get\",\"put\",\"get\",\"get\",\"get\"]"
	pa := "[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]"

	opList := make([]string, 0)
	paList := make([][]int, 0)

	err := json.Unmarshal([]byte(op), &opList)
	err = json.Unmarshal([]byte(pa), &paList)
	if err != nil {
		fmt.Println(err)
	}

	var cache lru.LRUCache
	resp := make([]int, 0)
	for idx, o := range opList {
		switch o {
		case cacheInit:
			cache = lru.Constructor(paList[idx][0])
		case put:
			cache.Put(paList[idx][0], paList[idx][1])
		case get:
			resp = append(resp, cache.Get(paList[idx][0]))
		}
	}

	for _, v := range resp {
		if v == -1 {
			fmt.Print("null")
		} else {
			fmt.Print("%d", v)
		}
	}
	return
}
