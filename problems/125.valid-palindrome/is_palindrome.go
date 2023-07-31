package leetcode0125

// 判断一个字符串是否为回文字符串
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	// 左右指针相向移动，直到相遇
	for left < right {
		// 如果左指针指向的字符不是字母或数字，则左指针右移
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		// 如果右指针指向的字符不是字母或数字，则右指针左移
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}
		// 如果左右指针指向的字符不相等，则不是回文字符串
		if s[left]|32 != s[right]|32 {
			return false
		}
		// 左右指针同时向中间移动
		left++
		right--
	}

	return true
}

// 判断一个字符是否为字母或数字
func isAlphanumeric(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9'
}
