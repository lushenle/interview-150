package leetcode0069

import "testing"

func TestMySqrt(t *testing.T) {
	testCases := []struct {
		in  int
		out int
	}{
		{1, 1},
		{4, 2},
		{8, 2},
		{9, 3},
		{10, 3},
		{11, 3},
		{12, 3},
		{13, 3},
		{14, 3},
		{15, 3},
		{16, 4},
	}

	for _, tc := range testCases {
		if actual := mySqrt(tc.in); actual != tc.out {
			t.Errorf("in: %d, got: %d, expected: %d", tc.in, actual, tc.out)
		}
	}

	for _, tc := range testCases {
		if actual := mySqrt1(tc.in); actual != tc.out {
			t.Errorf("in: %d, got: %d, expected: %d", tc.in, actual, tc.out)
		}
	}
}
