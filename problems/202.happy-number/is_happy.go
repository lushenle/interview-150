package leetcode0202

// 对于任何一个大于 4 的数字，最终都会得到 1 或 4，因此可以在循环中判断当前数字是否为 4，如果是，则直接返回 false
func isHappy(n int) bool {
	for n != 1 { // 当n不等于1时，继续循环
		if n == 4 { // 如果n等于4，返回false
			return false
		}
		n = next(n) // 计算n的下一个数
	}

	return n == 1 // 如果n最终等于1，返回true
}

// next 函数返回 n 的各个位上数字的平方和
func next(n int) int {
	sum := 0
	for n > 0 { // 当n大于0时，继续循环
		digit := n % 10      // 取n的个位数字
		sum += digit * digit // 将个位数字的平方加到sum上
		n /= 10              // 将n的个位数字去掉
	}
	return sum // 返回各个位上数字的平方和
}

func isHappy1(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		if m[n] {
			return false
		}
		m[n] = true
		n = next(n)
	}

	return n == 1
}
