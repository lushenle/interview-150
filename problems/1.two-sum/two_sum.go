package leetcode0001

func twoSum(nums []int, target int) []int {
	// 创建一个 map，用于存储数组元素的值和索引
	m := make(map[int]int)

	// 遍历数组 nums
	for i, num := range nums {
		// 检查 map 中是否存在满足条件的目标元素
		if j, ok := m[target-num]; ok {
			// 如果存在，直接返回结果
			return []int{j, i}
		}

		// 如果不存在，将当前元素的值和索引存入 map
		m[num] = i
	}

	// 如果遍历结束，仍然没有找到满足条件的目标元素，则返回一个空的数组
	return []int{}
}
