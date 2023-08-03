package leetcode0191

func hammingWeight(num uint32) int {
	count := 0

	// 循环遍历二进制位，直到 num 变成 0
	for num != 0 {
		count++        // 将计数器 count 加 1
		num &= num - 1 // 将 num 与 num-1 进行按位与运算，将结果存储到 num 中
		// 按位与运算的结果是将 num 的最后一位 1 变成 0，因此每次循环都会将一个 1 变成 0
	}

	return count
}
