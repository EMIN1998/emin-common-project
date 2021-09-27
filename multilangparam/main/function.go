package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func mainPrint(actId int64, status int64) {
	fmt.Print(actId, "\t\t", status, "\n")
}

func ReadData(data string, actIds []int64)  {
	datList := make([]*Data, 0)
	err := json.Unmarshal([]byte(data), &datList)
	if err != nil {
		print("json unmarshal err ", err.Error())
		return
	}
	dataMap := make(map[int64]*Data)
	for _, v := range datList {
		dataMap[v.ActivityId] = v
	}
	num := 0
	for _, id := range actIds {
		if _, ok := dataMap[id]; !ok {
			mainPrint(id, 0)
		} else {
			tmp := dataMap[id]
			mainPrint(tmp.ActivityId, tmp.Status)
		}
		num++
	}
	print("num:", num)
}

func JoinInt642(sep string, ints ...int64) string {
	switch len(ints) {
	case 0:
		return ""
	case 1:
		return strconv.FormatInt(ints[0], 10)
	default:
		buf := make([]byte, 0, len(ints)*5)
		for _, v := range ints {
			buf = strconv.AppendInt(buf, v, 10)
			buf = append(buf, sep...)
		}
		return string(buf[:len(buf)-len(sep)])
	}
}

func JoinInt64(sep string, ints ...int64) string {
	return JoinInt642(sep, ints...)
}


func splitArray(arr []int64, num int64) [][]int64 {
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
