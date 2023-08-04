package leetcode0050

func myPow(x float64, n int) float64 {
	// 如果 n 等于 0，返回 1
	if n == 0 {
		return 1
	}

	// 如果 n 小于 0，将 x 取倒数，n 取相反数
	if n < 0 {
		x = 1 / x
		n = -n
	}

	// 如果 n 是偶数，递归调用 myPow 函数，传入 x 的平方和 n 除以 2 的结果
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}

	// 如果 n 是奇数，返回 x 乘以递归调用 myPow 函数，传入 x 的平方和 n 除以 2 的结果
	return x * myPow(x*x, n/2)
}
