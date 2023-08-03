package leetcode0067

// 1. 从右往左遍历两个字符串，相同位置的字符相加，如果有进位则下一位加1
// 2. 如果两个字符串长度不一致，则较短的字符串前面补0
// 3. 如果最后一位有进位，则在最前面补1
// 4. 反转字符串

func addBinary(a, b string) string {
	// 如果a的长度小于b的长度，则交换a和b
	if len(a) < len(b) {
		a, b = b, a
	}

	// 将a转换为byte类型的切片，用于存储结果
	ans, c, i, j := []byte(a), 0, len(a)-1, len(b)-1

	// 从右往左遍历a和b，相同位置的字符相加，如果有进位则下一位加1
	for i >= 0 {
		if j >= 0 {
			c += int(a[i]-'0') + int(b[j]-'0')
		} else {
			c += int(a[i] - '0')
		}

		// 将相加的结果存储到ans中
		ans[i] = byte(c%2) + '0'
		c /= 2
		i--
		j--
	}

	// 如果最后一位有进位，则在最前面补1
	if c > 0 {
		ans = append([]byte{'1'}, ans...)
	}

	// 将byte类型的切片转换为字符串类型并返回
	return string(ans)
}
