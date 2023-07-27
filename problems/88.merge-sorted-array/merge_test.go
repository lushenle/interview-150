package leetcode0088

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tastCases := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		{
			name:   "TestMerge01",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			expect: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:   "TestMerge02",
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			expect: []int{1},
		},
		{
			name:   "TestMerge03",
			nums1:  []int{0},
			m:      0,
			nums2:  []int{1},
			n:      1,
			expect: []int{1},
		},
		{
			name:   "TestMerge04",
			nums1:  []int{2, 0},
			m:      1,
			nums2:  []int{1},
			n:      1,
			expect: []int{1, 2},
		},
	}

	for _, tc := range tastCases {
		t.Run(tc.name, func(t *testing.T) {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
			if !reflect.DeepEqual(tc.nums1, tc.expect) {
				t.Errorf("nums1: %v, expect: %v", tc.nums1, tc.expect)
			}
		})
	}
}
