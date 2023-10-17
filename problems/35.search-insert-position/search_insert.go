package leetcode0035

func searchInsert(nums []int, target int) int {
	length := len(nums)

	//如果目标值大于数组中的最后一个元素，则插入到数组末尾
	if target > nums[length-1] {
		return length
	}
	//如果目标值小于数组中的第一个元素，则插入到数组开头
	if target < nums[0] {
		return 0
	}

	left, right := 0, length-1
	// 二分查找
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 如果没有找到目标值，则返回插入位置
	return left
}
