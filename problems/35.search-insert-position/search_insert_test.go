package leetcode0035

import "testing"

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	expected := 2
	if output := searchInsert(nums, target); output != expected {
		t.Errorf("searchInsert(%v, %d) = %d; want %d", nums, target, output, expected)
	}

	nums = []int{1, 3, 5, 6}
	target = 2
	expected = 1
	if output := searchInsert(nums, target); output != expected {
		t.Errorf("searchInsert(%v, %d) = %d; want %d", nums, target, output, expected)
	}

	nums = []int{1, 3, 5, 6}
	target = 7
	expected = 4
	if output := searchInsert(nums, target); output != expected {
		t.Errorf("searchInsert(%v, %d) = %d; want %d", nums, target, output, expected)
	}

	nums = []int{1, 3, 5, 6}
	target = 0
	expected = 0
	if output := searchInsert(nums, target); output != expected {
		t.Errorf("searchInsert(%v, %d) = %d; want %d", nums, target, output, expected)
	}
}
