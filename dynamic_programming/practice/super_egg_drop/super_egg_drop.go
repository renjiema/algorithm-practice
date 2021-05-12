package super_egg_drop

// https://leetcode-cn.com/problems/super-egg-drop/
// 暴力解法，会超时
func superEggDrop1(k int, n int) int {
	cache := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		cache[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			cache[i][j] = -1
		}
	}
	var dp func(K, N int) int
	dp = func(K, N int) int {
		// base case
		if N == 0 {
			return 0
		}
		if cache[K][N] != -1 {
			return cache[K][N]
		}
		if K == 1 {
			cache[K][N] = N
			return cache[K][N]
		}
		res := N
		for i := 1; i <= N; i++ {
			res = min(res, max(dp(K-1, i-1), dp(K, N-i))+1)
		}
		cache[K][N] = res
		return res
	}
	return dp(k, n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 二分优化
func superEggDrop2(k int, n int) int {
	cache := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		cache[i] = make([]int, n+1)
	}
	var dp func(K, N int) int
	dp = func(K, N int) int {
		if N == 0 {
			return 0
		}
		if cache[K][N] != 0 {
			return cache[K][N]
		}
		if K == 1 {
			cache[K][N] = N
			return N
		}
		res := N
		left, right := 1, N
		for left <= right {
			mod := (right + left) / 2
			broken := dp(K-1, mod-1)
			not_broken := dp(K, N-mod)
			if broken > not_broken {
				right = mod - 1
				res = min(res, broken+1)
			} else {
				left = mod + 1
				res = min(res, not_broken+1)
			}
		}
		cache[K][N] = res
		return res
	}
	return dp(k, n)
}

// 修改状态转移方程：d[k][m]，k为鸡蛋数，m为运行扔的次数，值为最高的楼层
func superEggDrop3(k int, n int) int {
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}
	m := 0
	for dp[k][m] < n {
		m++
		for i := 1; i <= k; i++ {
			dp[i][m] = dp[i-1][m-1] + dp[i][m-1] + 1
		}
	}
	return m
}
