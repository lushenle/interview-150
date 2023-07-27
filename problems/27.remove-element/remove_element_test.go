package leetcode0027

import "testing"

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		nums   []int
		val    int
		result int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, tc := range testCases {
		actual := removeElement(tc.nums, tc.val)
		if actual != tc.result {
			t.Errorf("nums: %v, val: %d, expected: %d, actual: %d", tc.nums, tc.val, tc.result, actual)
		}
	}
}
