package leeetcode0121

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{1, 2},
			want:   1,
		},
		{
			prices: []int{2, 1},
			want:   0,
		},
	}

	for _, tc := range testCases {
		got := maxProfit(tc.prices)
		if got != tc.want {
			t.Errorf("prices: %v, want: %v, got: %v", tc.prices, tc.want, got)
		}
	}
}
