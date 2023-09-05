package leetcode0189

import "testing"

func TestRotate(t *testing.T) {
	testCases := []struct {
		in1 []int
		in2 int
		out []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			2,
			[]int{5, 6, 1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		rotate(tc.in1, tc.in2)
		for i := 0; i < len(tc.in1); i++ {
			if tc.in1[i] != tc.out[i] {
				t.Errorf("for %v, expected %v, got %v", tc.out, tc.in1, tc.out)
			}
		}
	}
}
