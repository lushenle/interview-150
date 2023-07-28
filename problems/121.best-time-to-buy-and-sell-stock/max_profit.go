package leeetcode0121

// maxProfit returns the maximum profit that can be achieved by buying and selling a stock.
// It takes a slice of integers representing the prices of the stock on different days.
// The function returns an integer representing the maximum profit that can be achieved.
func maxProfit(prices []int) int {
	maxPrice, minPrice := 0, 1<<31-1

	for _, price := range prices {
		maxPrice = max(maxPrice, price-minPrice)
		minPrice = min(minPrice, price)
	}

	return maxPrice
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// min returns the minimum of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
