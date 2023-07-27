package leetcode0027

// removeElement 函数从一个整数数组 nums 中移除所有值为 val 的元素，并返回新数组的长度。
// 双指针法：使用两个指针 i 和 j，初始时 i 指向数组头部，j 指向数组尾部。
// 当 nums[i] == val 时，将 nums[i] 和 nums[j] 交换，j--，相当于将 val 移动到数组尾部。
// 当 nums[i] != val 时，i++，相当于将非 val 元素向数组头部移动。
// 当 i > j 时，遍历结束，返回 i，即新数组的长度。
func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	return i
}
