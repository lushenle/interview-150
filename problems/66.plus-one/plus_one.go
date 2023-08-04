package leetcode0066

/*
从最后一个元素开始，依次向前遍历，对于每个元素，如果其小于 9，则将其加一并返回 digits 数组
否则，将其置为 0，继续向前遍历。如果遍历完整个数组后仍未返回，则说明需要在数组的最前面添加一个元素 1，表示进位
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
