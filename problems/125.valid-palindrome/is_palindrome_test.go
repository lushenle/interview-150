package leetcode0125

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		in  string
		out bool
	}{
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
		{
			"0P",
			false,
		},
		{
			"0",
			true,
		},
		{
			" ",
			true,
		},
	}

	for _, tc := range testCases {
		if out := isPalindrome(tc.in); out != tc.out {
			t.Errorf("in: %v, out %v, expected: %v", tc.in, out, tc.out)
		}
	}
}
