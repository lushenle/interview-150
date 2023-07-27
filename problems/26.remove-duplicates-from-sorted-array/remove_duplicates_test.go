package leetcode0026

import "testing"

func TestRemoveDupicates(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2}, 3},
	}

	for _, tc := range testCases {
		got := removeDuplicates(tc.nums)
		if tc.want != got {
			t.Errorf("nums: %v, want: %d, got: %d", tc.nums, tc.want, got)
		}
	}
}
