package leetcode0136

import "testing"

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
		{[]int{1, 1, 2}, 2},
	}

	for _, tc := range testCases {
		if got := singleNumber(tc.in); got != tc.want {
			t.Errorf("Input: %v, want: %v, got: %v", tc.in, tc.want, got)
		}
	}
}
