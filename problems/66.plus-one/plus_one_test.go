package leetcode0066

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		in  []int
		out []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{0}, []int{1}},
		{[]int{1, 9}, []int{2, 0}},
	}

	for _, tc := range testCases {
		if actual := plusOne(tc.in); !reflect.DeepEqual(actual, tc.out) {
			t.Errorf("Input: %v, expected: %v, output: %v", tc.in, tc.out, actual)
		}
	}
}
