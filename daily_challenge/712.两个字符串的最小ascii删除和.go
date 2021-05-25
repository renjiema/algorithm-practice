/*
 * @lc app=leetcode.cn id=712 lang=golang
 *
 * [712] 两个字符串的最小ASCII删除和
 */

// @lc code=start
func minimumDeleteSum(s1 string, s2 string) int {
	n, m := len(s1)+1, len(s2)+1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	// base case
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
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
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

