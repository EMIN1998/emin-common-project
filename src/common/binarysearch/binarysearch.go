package binarysearch

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，
//如果目标值存在返回下标，否则返回 -1。
//
//
//示例 1:
//
//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//示例 2:
//
//输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
//
//
//提示：
//
//你可以假设 nums 中的所有元素是不重复的。
//n 将在 [1, 10000]之间。
//nums 的每个元素都将在 [-9999, 9999]之间。
func findIndex(array []int, target int) int {
	pre := 0
	sub := len(array) - 1

	for pre <= sub {
		midIndex := (pre + sub) / 2
		if array[midIndex] == target {
			return midIndex
		}

		if array[midIndex] > target {
			sub = midIndex - 1
		} else {
			pre = midIndex + 1
		}
	}

	return -1
}

// 找出旋转数组中的中位数下标
// eg：[3,4,5,6,1,2] 返回2
func findMidNum(array []int) int {

	pre := 0
	sub := len(array) - 1

	var mid = 0
	for pre <= sub {
		mid = (pre + sub) / 2
		if array[mid] < array[mid-1] { // 找到最小值的下标
			if mid < len(array)-1 {
				if array[mid] < array[mid+1] {
					break
				}
			} else {
				if array[mid] < array[0] {
					break
				}
			}
		}

		if array[pre] > array[mid] { // 前半段无序，最小值在这里
			sub = mid - 1
		} else {
			pre = mid + 1
		}
	}

	var resp int
	if len(array)/2+mid >= len(array) {
		resp = len(array)/2 + mid - len(array)
	} else {
		resp = len(array)/2 + mid
	}

	return resp

}