/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	// base case
	i_2_0, i_2_1, i_1_0, i_1_1 := 0, math.MinInt32, 0, math.MinInt32
	for i := 0; i < n; i++ {
		i_2_0 = max(i_2_0, i_2_1+prices[i])
		i_2_1 = max(i_2_1, i_1_0-prices[i])
		i_1_0 = max(i_1_0, i_1_1+prices[i])
		i_1_1 = max(i_1_1, -prices[i])
	}
	return i_2_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
func maxProfit(prices []int) int {
	// 动态规划解法
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j <= 2; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	// base case
	for j := 0; j <= 2; j++ {
		dp[0][j][1] = math.MinInt32
	}
	for i := 1; i <= len(prices); i++ {
		for j := 2; j >= 1; j-- {
			// 第i天未持股：第i-1天未持股，或第i天卖出
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			// 第i天持股：第i-1天持股，或第i天买入
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}
	return dp[n][2][0]
}
