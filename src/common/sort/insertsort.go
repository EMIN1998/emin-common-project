package sort

// InsertSort ζε₯ζεΊ
func InsertSort(arr []int) []int {
	for idx, v := range arr {
		preIndex := idx - 1
		curr := v
		for preIndex >= 0 && curr < arr[preIndex] {
			arr[preIndex+1] = arr[preIndex]
			arr[preIndex] = v
			preIndex--
		}
	}

	return arr
}
