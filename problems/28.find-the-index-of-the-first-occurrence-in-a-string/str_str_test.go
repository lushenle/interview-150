package leeetcode0028

import "testing"

func TestStrStr(t *testing.T) {
	testCases := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"a", "b", -1},
		{"mississippi", "issip", 4},
	}

	for _, tc := range testCases {
		got := strStr(tc.haystack, tc.needle)
		if got != tc.want {
			t.Errorf("haystack: %v, needle: %v, want: %v, got: %v", tc.haystack, tc.needle, tc.want, got)
		}
	}
}
