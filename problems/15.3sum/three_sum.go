package leetcode0015

// threeSum 三数之和
// 数组长度小于3，返回nil
// 对数组进行排序
// 遍历数组，如果当前数字大于0，则三数之和一定大于0，所以结束循环
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	// 排序
	quickSort(nums, 0, len(nums)-1)

	ans := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		// 如果当前数字大于0，则三数之和一定大于0，所以结束循环
		if nums[i] > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			switch {
			case sum == 0:
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 去重
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			case sum < 0:
				left++
			case sum > 0:
				right--
			}
		}
	}

	return ans
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot

	return left
}
