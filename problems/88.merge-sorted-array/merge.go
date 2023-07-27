package leetcode0088

// 首先初始化三个变量：i 和 j 分别设置为 m-1 和 n-1，k 设置为 m+n-1。这些变量用于跟踪被比较和合并元素的索引

// 然后，函数进入一个循环，从 nums1 的最后一个索引开始，迭代 m+n 次。在每次迭代中，
// 函数都会检查 j 是否小于 0，或者 nums1 中索引 i 处的元素是否大于 nums2 中索引 j 处的元素
// 如果其中一个条件为真，函数会将 nums1 中索引 k 处的值设置为 nums1 中索引 i 处的值，并递减 i
// 否则，函数会将 nums1 中索引 k 处的值设置为 nums2 中索引 j 处的值，并递减 j

// 循环完成后，nums1 将包含经过合并和排序的 nums1 和 nums2 的元素

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 初始化变量 i 和 j 分别为 m-1 和 n-1，变量 k 为 m+n-1
	// 这些变量用于跟踪被比较和合并元素的索引
	i, j := m-1, n-1
	for k := m + n - 1; k >= 0; k-- {
		// 如果 j 小于 0 或者 nums1 中索引 i 处的元素大于 nums2 中索引 j 处的元素
		// 则将 nums1 中索引 k 处的值设置为 nums1 中索引 i 处的值，并递减 i
		// 否则，将 nums1 中索引 k 处的值设置为 nums2 中索引 j 处的值，并递减 j
		if j < 0 || (i >= 0 && nums1[i] > nums2[j]) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
