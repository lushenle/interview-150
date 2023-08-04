package leetcode0172

import "testing"

func TestTrailingZeroes(t *testing.T) {
	testCases := []struct {
		in   int
		want int
	}{
		{0, 0},
		{3, 0},
		{5, 1},
		{10, 2},
		{15, 3},
		{20, 4},
		{25, 6},
		{30, 7},
	}

	for _, tc := range testCases {
		if got := trailingZeroes(tc.in); got != tc.want {
			t.Errorf("input: %d, want: %d, got: %d", tc.in, tc.want, got)
		}
	}
}
