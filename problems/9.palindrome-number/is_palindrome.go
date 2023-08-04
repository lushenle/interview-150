package leetcode0009

// isPalindrome 检查给定的整数是否是回文数
// 回文数是指正着读和倒着读都一样的数字
// 例如，121是回文数，而123不是
func isPalindrome(x int) bool {
	// 如果x是负数或者x的个位数是0但x不是0，那么x不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 将x的数字反转，直到反转的数字大于等于x
	var reversed int
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// 如果x等于反转的数字或者x等于反转的数字除以10（去掉中间的数字），那么x是回文数
	return x == reversed || x == reversed/10
}
