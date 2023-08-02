package leetcode0020

// isValid 函数接受一个字符串s作为输入，并返回一个布尔值，表示输入的字符串是否有有效的括号
// 使用栈来跟踪左括号，并检查右括号是否与栈中的最后一个左括号匹配
// 如果栈为空，并且遇到右括号，则返回false
// 如果匹配了所有的括号，并且栈为空，则返回true
func isValid(s string) bool {
	// 如果字符串长度为奇数，直接返回 false
	if len(s)&1 == 1 {
		return false
	}

	// 创建一个栈，用于存储左括号
	stack := make([]byte, len(s))
	top := 0

	// 遍历字符串中的每一个字符
	for i := 0; i < len(s); i++ {
		// 如果是左括号，将其入栈
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack[top] = s[i]
			top++
		} else {
			// 如果是右括号，判断栈是否为空
			if top == 0 {
				return false
			}
			// 如果栈不为空，判断栈顶元素是否与当前右括号匹配
			if s[i] == ')' && stack[top-1] != '(' {
				return false
			}
			if s[i] == '}' && stack[top-1] != '{' {
				return false
			}
			if s[i] == ']' && stack[top-1] != '[' {
				return false
			}
			// 如果匹配成功，将栈顶元素出栈
			top--
		}
	}

	// 如果栈为空，说明所有括号都匹配成功，返回 true
	return top == 0
}

/*
// isValid 函数接受一个字符串s作为输入，并返回一个布尔值，表示输入的字符串是否有有效的括号
// 使用栈来跟踪左括号，并检查右括号是否与栈中的最后一个左括号匹配
// 如果栈为空，并且遇到右括号，则返回false
// 如果匹配了所有的括号，并且栈为空，则返回true
func isValid(s string) bool {
    stack := make([]byte, 0)

    mapping := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for i := 0; i < len(s); i++ {
        switch s[i] {
        case '(', '{', '[':
            stack = append(stack, s[i])
        case ')', '}', ']':
            if len(stack) == 0 || stack[len(stack)-1] != mapping[s[i]] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}
*/
