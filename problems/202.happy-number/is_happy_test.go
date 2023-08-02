package leetcode0202

import "testing"

func TestIsHappy(t *testing.T) {
	testCases := []struct {
		in   int
		want bool
	}{
		{19, true},
		{2, false},
		{3, false},
		{7, true},
		{1111111, true},
	}

	for _, tc := range testCases {
		if got := isHappy(tc.in); got != tc.want {
			t.Errorf("isHappy(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}
}
