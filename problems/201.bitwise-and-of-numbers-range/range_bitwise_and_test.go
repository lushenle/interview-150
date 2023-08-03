package leetcode0201

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {
	testCases := []struct {
		in   [2]int
		want int
	}{
		{[2]int{5, 7}, 4},
		{[2]int{0, 1}, 0},
		{[2]int{1, 2}, 0},
		{[2]int{3, 3}, 3},
		{[2]int{3, 4}, 0},
		{[2]int{4, 5}, 4},
		{[2]int{5, 6}, 4},
	}

	for _, tc := range testCases {
		if got := rangeBitwiseAnd(tc.in[0], tc.in[1]); got != tc.want {
			t.Errorf("Input: %v, want: %v, got: %v", tc.in, tc.want, got)
		}
	}

}
