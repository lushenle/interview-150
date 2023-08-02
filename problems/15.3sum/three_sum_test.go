package leetcode0015

import (
	"reflect"
	"testing"
)

func TestRThreeSum(t *testing.T) {
	testCases := []struct {
		in  []int
		out [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			[]int{0, 0, 0},
			[][]int{
				{0, 0, 0},
			},
		},
		{
			[]int{-2, 0, 1, 1, 2},
			[][]int{
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
		{
			[]int{-2, 0, 0, 2, 2},
			[][]int{
				{-2, 0, 2},
			},
		},
	}

	for _, tc := range testCases {
		if actual := threeSum(tc.in); !reflect.DeepEqual(actual, tc.out) {
			t.Errorf("input: %v, expected %v, actual %v", tc.in, tc.out, actual)
		}
	}
}
