package leetcode0191

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	testCases := []struct {
		in  uint32
		out int
	}{
		{
			in:  0b00000000000000000000000000001011,
			out: 3,
		},
		{
			in:  0b00000000000000000000000010000000,
			out: 1,
		},
		{
			in:  0b11111111111111111111111111111101,
			out: 31,
		},
	}

	for _, tc := range testCases {
		if got := hammingWeight(tc.in); got != tc.out {
			t.Errorf("expected: %d, got: %d", tc.out, got)
		}
	}
}
