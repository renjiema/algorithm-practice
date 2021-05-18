/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 动态规划解法
	n := len(prices)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	// base case
	dp[0][1] = math.MinInt32
	for i := 1; i <= n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1])
	}
	return dp[n][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

