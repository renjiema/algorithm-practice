package minimum_ascii_delete_sum_for_two_strings

// https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings
func minimumDeleteSum(s1 string, s2 string) int {
	n, m := len(s1), len(s2)
	cache := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = make([]int, m+1)
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if cache[i][j] != 0 {
			return cache[i][j]
		}
		// base case
		if i == n {
			for k := j; k < m; k++ {
				cache[i][j] += int(s2[k])
			}
			return cache[i][j]
		}
		if j == m {
			for k := i; k < n; k++ {
				cache[i][j] += int(s1[k])
			}
			return cache[i][j]
		}
		if s1[i] == s2[j] {
			cache[i][j] = dp(i+1, j+1)
		} else {
			cache[i][j] = min(dp(i+1, j)+int(s1[i]), dp(i, j+1)+int(s2[j]))
		}
		return cache[i][j]
	}
	return dp(0, 0)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minimumDeleteSumByRecursive(s1 string, s2 string) int {
	n, m := len(s1)+1, len(s2)+1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	// base case
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return dp[n-1][m-1]
}
