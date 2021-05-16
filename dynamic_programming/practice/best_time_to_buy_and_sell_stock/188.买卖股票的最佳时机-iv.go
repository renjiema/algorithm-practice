/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	// 定义状态dp(i,k,0 or 1):i表示天数，k交易次数，0表示未持股，1表示持股，值为利润
	n := len(prices)
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	// base case
	for j := 0; j <= k; j++ {
		dp[0][j][1] = math.MinInt32
	}
	for i := 1; i <= len(prices); i++ {
		for j := k; j >= 1; j-- {
			// 第i天未持股：第i-1天未持股，或第i天卖出
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			// 第i天持股：第i-1天持股，或第i天买入
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}
	return dp[n][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

