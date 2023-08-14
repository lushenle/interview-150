package leetcode0053

import "testing"

func TestMaxSubArray(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1}, -1},
		{[]int{-2, -1}, -1},
	}

	for _, tc := range testCases {
		if got := maxSubArray(tc.in); got != tc.want {
			t.Errorf("input: %v, want: %d, got: %d", tc.in, tc.want, got)
		}
	}
}
