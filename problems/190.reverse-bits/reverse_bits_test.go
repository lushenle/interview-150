package leetcode0190

import "testing"

func TestReverseBits(t *testing.T) {
	testCases := []struct {
		in  uint32
		out uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
		{0, 0},
		{1, 2147483648},
		{2147483648, 1},
	}

	for _, tc := range testCases {
		if got := reverseBits(tc.in); got != tc.out {
			t.Errorf("reverseBits(%v) = %v; want %v", tc.in, got, tc.out)
		}
	}
}
