package leetcode0009

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		in   int
		want bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{0, true},
		{123454321, true},
		{123456789, false},
	}

	for _, tc := range testCases {
		if got := isPalindrome(tc.in); got != tc.want {
			t.Errorf("isPalindrome(%d) = %v, want %v", tc.in, got, tc.want)
		}
	}
}
