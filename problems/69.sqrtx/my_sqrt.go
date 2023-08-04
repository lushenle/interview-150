package leetcode0069

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2 // 初始化搜索区间的左右端点
	for left <= right {   // 当搜索区间的左端点小于等于右端点时，执行循环
		mid := left + (right-left)/2 // 计算搜索区间的中间点
		if mid == x/mid {            // 如果中间点的平方等于 x，返回中间点
			return mid
		} else if mid < x/mid { // 如果中间点的平方小于 x，将搜索区间的左端点移动到中间点的右侧，继续搜索
			left = mid + 1
		} else {
			right = mid - 1 // 如果中间点的平方大于 x，将搜索区间的右端点移动到中间点的左侧，继续搜索
		}
	}

	return right
}

func mySqrt1(x int) int {
	if x < 2 {
		return x
	}

	guess := x / 2        //初始化猜测值为 x 的一半
	for guess*guess > x { // 当猜测值的平方大于 x 时，执行循环
		guess = (guess + x/guess) / 2 // 更新猜测值为 (猜测值加上 x 除以猜测值的商)除以 2
	}

	return guess
}
