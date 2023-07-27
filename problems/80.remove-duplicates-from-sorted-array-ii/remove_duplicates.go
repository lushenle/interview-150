package leeetcode0080

// removeDuplicates 函数接受一个有序整数数组 nums 作为参数，返回一个整数。
// 先将变量 i 初始化为 0。这个变量用于跟踪切片中最后一个唯一元素的索引。
// 然后，函数进入一个循环，迭代切片中的每个元素，对于每个元素，检查 i 是否小于 2 或者该元素是否大于 i 之前两个位置的元素。
// 如果其中任意一个条件为真，函数将切片中索引 i 处的值设置为当前元素，并递增 i。
// 循环完成后，i 将包含切片中唯一元素的数量，而切片的前 i 个元素将是唯一的元素，每个元素最多出现两次。
func removeDuplicates(nums []int) int {
	i := 0
	for _, num := range nums {
		if i < 2 || num > nums[i-2] {
			nums[i] = num
			i++
		}
	}

	return i
}
