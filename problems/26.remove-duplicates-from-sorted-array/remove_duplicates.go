package leetcode0026

// removeDuplicates 函数接收一个有序整数数组 nums，返回去重后的数组长度。
// 去重后的数组前 i 个元素即为去重后的数组。
// 使用双指针法，i 指向去重后的数组的最后一个元素，j 指向当前遍历的元素。
// 如果 nums[i] != nums[j]，则将 nums[j] 赋值给 nums[i+1]，i 自增 1。
// 最后返回 i，即为去重后的数组长度。
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
