package leetcode0001

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 2, 3}, 4, []int{}},
	}

	for _, tc := range testCases {
		got := twoSum(tc.nums, tc.target)
		for i := range got {
			if got[i] != tc.want[i] {
				t.Errorf("twoSum(%v, %v) return %v, want %v", tc.nums, tc.target, got, tc.want)
			}
		}
	}
}
