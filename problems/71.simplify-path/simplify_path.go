package leetcode0071

import (
	"strings"
)

// simplifyPath 函数接受一个字符串参数 path，返回简化后的路径。
// 简化规则如下：
// 1. 路径以 '/' 开头，且不包含多余的 '/'。
// 2. '.' 表示当前目录，可以省略。
// 3. '..' 表示上级目录，需要返回上级目录，如果已经是根目录则忽略。
// 4. 其他字符串表示目录名，需要保留。
// 例如："/home//foo/" -> "/home/foo"，"/a/./b/../../c/" -> "/c"
func simplifyPath(path string) string {
	var stack []string // 定义一个字符串类型的栈，用于保存简化后的路径

	for _, s := range strings.Split(path, "/") { // 将路径按照 '/' 分割为各个部分，遍历每个部分
		switch s {
		case "", ".": // 如果部分为空或为 '.'，则不做任何操作
		case "..": // 如果部分为 '..'，则弹出栈顶元素（如果栈非空）
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default: // 否则，将部分压入栈中
			stack = append(stack, s)
		}
	}

	return "/" + strings.Join(stack, "/") // 将栈中剩余的元素按照 '/' 连接起来，形成简化后的路径，并在前面添加 '/'
}

// func simplifyPath(path string) string {
// 	return filepath.Clean(path)
// }
