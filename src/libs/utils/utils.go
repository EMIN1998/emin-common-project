package utils

import (
	"fmt"
	"sort"
	"strings"
)

// UniqueInt64Slice int64 数组去重
func UniqueInt64Slice(slice *[]int64) {
	found := make(map[int64]bool)
	total := 0

	for idx, val := range *slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slice)[total] = (*slice)[idx]
			total++
		}
	}

	*slice = (*slice)[:total]
}

func ToIdString(ids []int64) string {
	stList := make([]string, 0)
	for _, id := range ids {
		st := fmt.Sprintf("%d", id)
		stList = append(stList, st)
	}

	str := strings.Join(stList, ",")
	return str
}

// 分割数组，根据传入的数组和分割大小，将数组分割为大小等于指定大小的多个数组，如果不够分，则最后一个数组元素小于其他数组
func SplitArray(arr []int64, num int64) [][]int64 {
	max := int64(len(arr))
	// 判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]int64{arr}
	}

	// 获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}

	// 声明分割好的二维数组
	var segments = make([][]int64, 0)
	// 声明分割数组的截止下标
	var start, end, i int64

	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

func InInt64Slice(item int64, slice ...int64) bool {
	for _, v := range slice {
		if item == v {
			return true
		}
	}

	return false
}

// 获取补集
func GetSupplementarySet(child, parent []int64) []int64 {
	resp := make([]int64, 0)
	for _, id := range parent {
		if !InInt64Slice(id, child...) {
			resp = append(resp, id)
		}
	}

	return resp
}

func Int64Sort(data []int64) []int64 {
	sort.Slice(data, func(i, j int) bool {
		if data[i] < data[j] {
			return true
		}
		return false
	})

	return data
}
