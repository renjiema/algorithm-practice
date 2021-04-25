package longest_common_subsequence

// https://leetcode-cn.com/problems/longest-common-subsequence
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequenceByRecursive(text1, text2 string) int {
	n, m := len(text1), len(text2)
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = -1
		}
	}
	var dp func(i int, j int) int
	dp = func(i int, j int) int {
		if i == n || j == m {
			return 0
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		if text1[i] == text2[j] {
			cache[i][j] = 1 + dp(i+1, j+1)
		} else {
			cache[i][j] = max(dp(i+1, j), dp(i, j+1))
		}
		return cache[i][j]
	}
	return dp(0, 0)
}
