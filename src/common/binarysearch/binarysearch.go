package binarysearch

// 二分查找模版：
//func binarySearch(nums []int, target int) int {
//	left, right := 0, len(nums)-1
//	for left<= right {
// 		// 这里是为了left和right过大直接相加造成溢出
//		mid := left + (right - left) / 2
//		if nums[mid] == target {
//			...
//		} else if nums[mid] < target {
//			left = ...
//		} else if nums[mid] > target {
//			right = ...
//		}
//	}
//
//	return ...;
//}

//参考链接：https://labuladong.github.io/algo/2/18/26/

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

// 找出target的范围
// link：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}

	pre := 0
	sub := len(nums) - 1
	first, last := len(nums), -1
	for pre <= sub {
		mid := (pre + sub) / 2

		if nums[mid] == target {
			if mid > last {
				last = mid
			}

			if mid < first {
				first = mid
			}

			for i := 1; mid-i >= pre && nums[mid-i] == target; i++ {
				first--
			}

			for i := 1; mid+i <= sub && nums[mid+i] == target; i++ {
				last++
			}

			break
		} else {
			if nums[mid] < target {
				pre = mid + 1
			} else {
				sub = mid - 1
			}
		}
	}

	if first == len(nums) {
		first = last
	}

	if last == -1 {
		last = first
	}

	return []int{first, last}
}

// ==============================================================================================
// 优化写法
// link：https://labuladong.github.io/algo/2/18/26/
func searchRangeOptimize(nums []int, target int) []int {
	return []int{findLeftBound(nums, target), findRightBound(nums, target)}
}

func findLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if left > len(nums)-1 || nums[left] != target {
		return -1
	}

	return left
}

func findRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if right < 0 || nums[right] != target {
		return -1
	}

	return right
}

// ==============================================================================================
