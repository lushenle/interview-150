package leeetcode0013

// 当小值在大值的左边，则减小值，如 IV=5-1=4
// 当小值在大值的右边，则加小值，如 VI=5+1=6
// 由上可知，右值永远为正，因此最后一位必然为正
// 从左往右遍历，如果当前值小于右边的值，则减去当前值，否则加上当前值
// 由于最后一位必然为正，因此最后一位必然会加上

var m = map[rune]int{
	'I': 1,    // 1
	'V': 5,    // 5
	'X': 10,   // 10
	'L': 50,   // 50
	'C': 100,  // 100
	'D': 500,  // 500
	'M': 1000, // 1000
}

func romanToInt(s string) int {
	var ans int
	var prev int

	for _, r := range s {
		curr := m[r]
		if curr > prev {
			ans += curr - 2*prev // 如果当前值比前一个值大，则减去前一个值的两倍
		} else {
			ans += curr // 否则加上当前值
		}
		prev = curr // 更新前一个值
	}

	return ans
}

/*
func romanToInt(s string) int {
	var ans int

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && m[rune(s[i])] < m[rune(s[i+1])] {
			ans -= m[rune(s[i])]
		} else {
			ans += m[rune(s[i])]
		}
	}

	return ans
}
*/

/*
func romanToInt(s string) int {
	s = strings.ReplaceAll(s, "IV", "a")
	s = strings.ReplaceAll(s, "IX", "b")
	s = strings.ReplaceAll(s, "XL", "c")
	s = strings.ReplaceAll(s, "XC", "d")
	s = strings.ReplaceAll(s, "CD", "e")
	s = strings.ReplaceAll(s, "CM", "f")

	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
		'a': 4,
		'b': 9,
		'c': 40,
		'd': 90,
		'e': 400,
		'f': 900,
	}

	var ans int

	for _, r := range s {
		if v, ok := m[r]; ok {
			ans += v
		}
	}

	return ans
}
*/
