package sort

// MergeSort 归并排序
func MergeSort(request []int) []int {
	if len(request) == 0 || len(request) == 1 {
		return request
	}

	index := len(request) / 2
	req1 := request[0:index]
	req2 := request[index:]

	res1 := MergeSort(req1)
	res2 := MergeSort(req2)

	result := make([]int, 0)
	var i, j int
	for {

		if res1[i] < res2[j] {
			result = append(result, res1[i])
			i++
		} else {
			result = append(result, res2[j])
			j++
		}

		if i == len(res1) {
			result = append(result, res2[j:]...)
			break
		}

		if j == len(res2) {
			result = append(result, res1[i:]...)
			break
		}
	}

	return result
}
