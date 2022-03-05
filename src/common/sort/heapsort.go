package sort

// MaxHeapify 堆排序
// 数组中的某个节点的左子节点：2i+1
// 右子节点：2i+2
// 父节点：(i-1)/2
// 最后一个非叶子节点：最后一个节点的父节点

func swap(arr *[]int, i, j int) {
	if i < 0 || j < 0 || i == j {
		return
	}

	tmp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tmp
	return
}
func MaxHeapify(nums []int, i int) []int {
	if len(nums) < 2 {
		return nums
	}

	left := 2*i + 1
	right := 2*i + 2

	max := i
	if left < len(nums) && nums[max] < nums[left] {
		max = left
	}

	if right < len(nums) && nums[max] < nums[right] {
		max = right
	}

	if max != i {
		swap(&nums, i, max)
		nums = MaxHeapify(nums, max)
	}

	return nums
}

// 构建大顶堆数组
func buildHeap(nums []int) []int {
	start := (len(nums) - 1) / 2
	for start > -1 {
		nums = MaxHeapify(nums, start)
		start--
	}

	return nums
}

// HeapSort 堆排序
func HeapSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	resp := make([]int, 0)

	numsLen := len(nums)
	for numsLen != 0 {
		nums = buildHeap(nums)
		resp = append(resp, nums[0])
		nums = nums[1:]

		numsLen = len(nums)
	}

	return resp
}

// FindKthLargest 找出第K大的数
func FindKthLargest(nums []int, k int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	nums = buildHeap(nums)
	var i int
	for i < k-1 {
		nums = nums[1:]
		nums = buildHeap(nums)
		i++
	}

	return nums[0]
}
