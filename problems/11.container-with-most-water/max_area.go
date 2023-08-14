package leetcode0011

// 使用双指针，从两端开始向中间移动，每次移动高度较小的指针，直到两个指针相遇为止
// 在移动指针的过程中，我们计算当前的容器的面积，并更新最大值
//
// 当左指针小于右指针时，计算当前容器的面积，并更新最大值
// 如果左指针对应的高度小于右指针对应的高度，则左指针向右移动一位，否则右指针向左移动一位
// 重复上面两个步骤，直到左指针和右指针相遇为止
func maxArea(height []int) int {
	max := 0                        // 用于记录最大值
	left, right := 0, len(height)-1 // 初始化左指针和右指针

	for left < right { // 当左指针小于右指针时，循环执行以下操作
		area := (right - left) * min(height[left], height[right]) // 计算当前容器的面积
		if area > max {                                           // 如果当前面积大于最大值，则更新最大值
			max = area
		}

		if height[left] > height[right] { // 如果左指针对应的高度大于右指针对应的高度，则右指针向左移动一位
			right--
		} else { // 否则左指针向右移动一位
			left++
		}
	}

	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
