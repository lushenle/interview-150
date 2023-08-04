package leetcode0050

import (
	"testing"
)

func TestMyPow(t *testing.T) {
	testCases := []struct {
		x    float64
		n    int
		want float64
	}{
		{2.00000, 10, 1024.00000},
		// {2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
		{0.00001, 2147483647, 0.00000},
		{2.00000, -2147483648, 0.00000},
	}

	for _, tc := range testCases {
		if got := myPow(tc.x, tc.n); got != tc.want {
			t.Errorf("x: %f, n: %d, want: %f, got: %f", tc.x, tc.n, tc.want, got)
		}
	}
}
