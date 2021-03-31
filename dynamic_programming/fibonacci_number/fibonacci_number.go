package fibonacci_number

// 斐波那契数列并不算是动态规划，因为不满足：最优子结构，不过可以了解最优解法逐步求精的过程
// https://leetcode-cn.com/problems/fibonacci-number/
// 1.暴力解法
func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// 2.带缓存的递归解法
func cacheFib(n int) int {
	if n < 1 {
		return 0
	}
	cache := make([]int, n+1)
	return cacheHelper(cache, n)
}

func cacheHelper(cache []int, n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if cache[n] == 0 {
		cache[n] = cacheHelper(cache, n-1) + cacheHelper(cache, n-2)
	}
	return cache[n]
}

// 3.自底而上的解法
func toTopFib(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	var prev, cur int = 1, 1
	var sum = 0
	for i := 3; i <= n; i++ {
		sum = prev + cur
		prev = cur
		cur = sum
	}
	return sum
}
