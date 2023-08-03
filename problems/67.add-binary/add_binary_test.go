package leetcode0067

import "testing"

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		a      string
		b      string
		expect string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"0", "0", "0"},
		{"0", "11", "11"},
	}

	for _, tc := range testCases {
		if output := addBinary(tc.a, tc.b); output != tc.expect {
			t.Fatalf("input: %v %v, output: %v, expect: %v", tc.a, tc.b, output, tc.expect)
		}
	}
}
