package leetcode0918

import "testing"

func TestMaxSubarraySumCircular(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, -2, 3, -2},
			expected: 3,
		},
		{
			input:    []int{5, -3, 5},
			expected: 10,
		},
		{
			input:    []int{3, -1, 2, -1},
			expected: 4,
		},
		{
			input:    []int{3, -2, 2, -3},
			expected: 3,
		},
		{
			input:    []int{-2, -3, -1},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		if actual := maxSubarraySumCircular(tc.input); actual != tc.expected {
			t.Errorf("input: %v, expected: %v, actual: %v", tc.input, tc.expected, actual)
		}
	}
}
