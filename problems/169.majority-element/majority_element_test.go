package leeetcode0169

import "testing"

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
	}

	for _, tc := range testCases {
		got := majorityElement(tc.input)
		if got != tc.want {
			t.Errorf("input: %v, want: %v, got: %v", tc.input, tc.want, got)
		}
	}
}
