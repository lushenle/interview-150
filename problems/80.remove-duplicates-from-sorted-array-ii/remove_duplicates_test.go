package leeetcode0080

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}

	for _, tc := range testCases {
		got := removeDuplicates(tc.nums)
		if got != tc.want {
			t.Errorf("nums: %v, want: %d, got: %d", tc.nums, tc.want, got)
		}
	}
}
