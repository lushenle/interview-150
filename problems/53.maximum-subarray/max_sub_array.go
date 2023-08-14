package leetcode0053

// 初始化 max 和 sum 为数组的第一个元素，
// 然后从数组的第二个元素开始遍历，如果当前的和 sum 小于 0，则将 sum 更新为当前元素的值，否则将 sum 加上当前元素的值
// 然后判断 sum 是否大于 max，如果是，则将 max 更新为 sum。最后返回 max
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for _, num := range nums {
		sum += num
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
