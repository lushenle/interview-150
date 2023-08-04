package leetcode0104

import "testing"

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		in  *TreeNode
		out int
	}{
		{
			in: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			out: 3,
		},
		{
			in:  nil,
			out: 0,
		},

		{
			in: &TreeNode{
				Val: 1,
			},
			out: 1,
		},
	}

	for _, test := range testCases {
		if out := maxDepth(test.in); out != test.out {
			t.Fatalf("expected %d, but got %d", test.out, out)
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		a   int
		b   int
		out int
	}{
		{
			a:   1,
			b:   2,
			out: 2,
		},
		{
			a:   2,
			b:   1,
			out: 2,
		},
	}

	for _, test := range testCases {
		if out := max(test.a, test.b); out != test.out {
			t.Fatalf("expected %d, but got %d", test.out, out)
		}
	}
}
