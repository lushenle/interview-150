package leetcode0392

// isSubsequence checks if string s is a subsequence of string t
// 如果 s 是一个空字符串，它被认为是任何字符串的子序列
// 如果 s 的当前字符与 t 的当前字符匹配，i 将加 1
// 如果 i 到达 s 的末尾，s是 t 的子序列，函数返回 `true`
// 如果 j 到达了 t 的末尾，而我没有到达 s 的末尾，那么 s 不是 t 的子序列，函数返回 `false`
func isSubsequence(s, t string) bool {
	// 如果 s 为空字符串，则它是任何字符串的子序列，返回 true
	if len(s) == 0 {
		return true
	}

	// 用两个指针 i 和 j 分别指向字符串 s 和 t 的开头
	i := 0
	for j := 0; j < len(t); j++ {
		// 如果当前字符与 s 中的字符相同，则将 i 向后移动一位
		if s[i] == t[j] {
			i++
			// 如果 i 移动到了 s 的末尾，则说明 s 是 t 的子序列，返回 true
			if i == len(s) {
				return true
			}
		}
	}

	// 如果遍历完 t 后 i 还没有移动到 s 的末尾，则说明 s 不是 t 的子序列，返回 false
	return false
}
