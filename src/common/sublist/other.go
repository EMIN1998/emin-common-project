package sublist

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
