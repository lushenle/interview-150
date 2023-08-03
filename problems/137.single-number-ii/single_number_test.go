package leetcode0137

import "testing"

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
		{[]int{1}, 1},
		{[]int{1, 1, 1, 2}, 2},
		{[]int{1, 1, 1, 2}, 2},
	}

	for _, tc := range testCases {
		if got := singleNumber(tc.in); got != tc.want {
			t.Errorf("Input: %v, want: %v, got: %v", tc.in, tc.want, got)
		}
	}
}
