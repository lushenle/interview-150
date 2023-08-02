package leetcode0020

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		in  string
		out bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"))", false},
		{"]]", false},
		{"}}[]", false},
		{"{{}", false},
	}

	for _, tc := range testCases {
		if actual := isValid(tc.in); actual != tc.out {
			t.Errorf("input: %s, expected %t, actual %t", tc.in, tc.out, actual)
		}
	}
}
