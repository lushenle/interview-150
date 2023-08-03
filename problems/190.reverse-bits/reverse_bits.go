package leetcode0190

// reverseBits 函数接收一个 uint32 类型的参数 num，将其二进制位翻转后返回。
func reverseBits(num uint32) uint32 {
	var res uint32

	// 遍历 32 位二进制位
	for i := 0; i < 32; i++ {
		res <<= 1      // 将结果变量 res 左移一位
		res |= num & 1 // 将 num 的最后一位与 1 进行按位或运算，将结果存储到 res 中
		num >>= 1      // 将 num 右移一位
	}

	return res
}
