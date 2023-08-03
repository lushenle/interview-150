package leetcode0137

/*
对于每个数字 num，分别更新变量 a 和 b 的值。具体来说，更新 a 的值为 (a ^ num) & ^b，更新 b 的值为 (b ^ num) & ^a
这两个式子的含义是，将 num 与 a 进行异或运算，然后将结果与 b 取反的结果进行按位与运算，得到新的 a 的值
将 num 与 b 进行异或运算，然后将结果与 a 取反的结果进行按位与运算，得到新的 b 的值
这样可以保证 a 和 b 中只存储出现一次的数字，而不存储出现多次的数字
*/

func singleNumber(nums []int) int {
	a, b := 0, 0

	for _, num := range nums {
		a = (a ^ num) & ^b
		b = (b ^ num) & ^a
	}

	return a
}
