package leetcode0071

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	testCases := []struct {
		path string
		want string
	}{
		{"", "/"},
		{"/", "/"},
		{"/home/", "/home"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
	}

	for _, tc := range testCases {
		if got := simplifyPath(tc.path); got != tc.want {
			t.Errorf("simplifyPath(%v) return %v, want %v", tc.path, got, tc.want)
		}
	}
}
