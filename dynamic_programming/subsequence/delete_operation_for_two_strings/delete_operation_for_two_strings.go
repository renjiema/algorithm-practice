package delete_operation_for_two_strings

// https://leetcode-cn.com/problems/delete-operation-for-two-strings
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if cache[i][j] != 0 {
			return cache[i][j]
		}
		if word1[i] == word2[j] {
			cache[i][j] = dp(i-1, j-1)
		} else {
			cache[i][j] = min(dp(i-1, j)+1, dp(i, j-1)+1)
		}
		return cache[i][j]
	}
	return dp(n-1, m-1)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 迭代解法
func minDistanceByIteration(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}
	// base case
	for i := 0; i < m+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1)
			}
		}
	}

	return dp[n][m]
}
