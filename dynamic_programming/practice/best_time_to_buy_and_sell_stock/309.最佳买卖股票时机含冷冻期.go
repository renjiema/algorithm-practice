/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	// base case
	dp_i_0, dp_i_1, dp_pre_0 := 0, math.MinInt32, 0
	for i := 0; i < n; i++ {
		temp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, dp_pre_0-prices[i])
		dp_pre_0 = temp
	}
	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

