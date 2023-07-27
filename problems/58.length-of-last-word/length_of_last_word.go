package leeetcode0058

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if length > 0 {
				break
			}
		} else {
			length++
		}
	}
	return length
}

// func lengthOfLastWord(s string) int {
// 	s = strings.TrimSpace(s)
// 	strs := strings.Split(s, " ")
// 	return len(strs[len(strs)-1])
// }
