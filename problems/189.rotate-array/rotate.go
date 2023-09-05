package leetcode0189

// 先将整个数组翻转，然后将前 k 个元素翻转，再将后 n-k 个元素翻转。这样就可以将数组向右旋转 k 步
//
// 计算数组的长度 n，并将 k 取模 n，以确保 k 在数组索引范围内
// 调用 reverse 函数三次，分别对整个数组、前 k 个元素和后 n-k 个元素进行翻转
// 时间复杂度为 O(n)，其中 n 是数组的长度。空间复杂度为 O(1)
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

// 使用双指针将数组的起始位置和结束位置的元素交换
// 将起始位置向右移动一位，将结束位置向左移动一位，重复上述操作，直到两个指针相遇为止
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
