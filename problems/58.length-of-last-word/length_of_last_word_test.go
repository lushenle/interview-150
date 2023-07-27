package leeetcode0058

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		s    string
		want int
	}{
		{"Hello World", 5},
		{"a ", 1},
		{" ", 0},
		{"", 0},
	}

	for _, tc := range testCases {
		got := lengthOfLastWord(tc.s)
		if got != tc.want {
			t.Errorf("s: %v, want: %v, got: %v", tc.s, tc.want, got)
		}
	}
}
