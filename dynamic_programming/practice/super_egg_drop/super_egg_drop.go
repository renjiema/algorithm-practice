package super_egg_drop

// https://leetcode-cn.com/problems/super-egg-drop/
func superEggDrop(k int, n int) int {
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
