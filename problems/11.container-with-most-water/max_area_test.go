package leetcode0011

import "testing"

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
	}

	for _, tc := range testCases {
		if actual := maxArea(tc.height); actual != tc.expected {
			t.Errorf("expected: %d, actual: %d", tc.expected, actual)
		}
	}
}
