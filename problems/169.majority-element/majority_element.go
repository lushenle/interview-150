package leeetcode0169

import "sort"

// import "golang.org/x/exp/constraints"
// func bubbleSort[T constraints.Ordered](nums []T) []T {
// 	swapped := true
// 	for swapped {
// 		swapped = false
// 		for i := 1; i < len(nums); i++ {
// 			if nums[i-1] > nums[i] {
// 				nums[i-1], nums[i] = nums[i], nums[i-1]
// 				swapped = true
// 			}
// 		}
// 	}

// 	return nums
// }

// majorityElement 函数接收一个整数数组 nums，返回该数组中出现次数超过 ⌊ n/2 ⌋ 的元素。
// 该函数使用快速排序算法对 nums 进行排序，然后返回排序后的中间元素，即为出现次数超过 ⌊ n/2 ⌋ 的元素。
func majorityElement(nums []int) int {
	// nums = bubbleSort(nums)
	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] > nums[j]
	// })
	sort.Ints(nums)
	return nums[len(nums)>>1]
}

/*
// majorityElement 函数接收一个整数数组 nums，返回其中出现次数超过数组长度一半的元素。
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}
*/
