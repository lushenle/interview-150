package leeetcode0028

// func strStr(haystack string, needle string) int {
// 	return strings.Index(haystack, needle)
// }

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
