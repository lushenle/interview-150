package leetcode0014

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		in  []string
		out string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"a"}, "a"},
		{[]string{"ab", "a"}, "a"},
		{[]string{"ab", "ab"}, "ab"},
		{[]string{"ab", "abc"}, "ab"},
		{[]string{"ab", "abc", "abcd"}, "ab"},
		{[]string{"ab", "abc", "abcd", "abcde"}, "ab"},
		{[]string{"ab", "abc", "abcd", "abcde", "abcdef"}, "ab"},
		{[]string{"ab", "abc", "abcd", "abcde", "abcdef", "abcdefg"}, "ab"},
		{[]string{"ab", "abc", "abcd", "abcde", "abcdef", "abcdefg", "abcdefgh"}, "ab"},
		{[]string{"cir", "car"}, "c"},
	}

	for _, tc := range testCases {
		got := longestCommonPrefix(tc.in)
		if got != tc.out {
			t.Errorf("Input: %v, expected: %v, got: %v", tc.in, tc.out, got)
		}
	}
}
