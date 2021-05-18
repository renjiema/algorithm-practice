/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package besttimetobuyandsellstock

import "math"

// @lc code=start
func maxProfit(prices []int) int {
	res := 0
	minVal := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if prices[i] < minVal {
			minVal = prices[i]
		} else if prices[i]-minVal > res {
			res = prices[i] - minVal
		}
	}

	return res
}

// @lc code=end

func maxProfit2(prices []int) int {
	n := len(prices)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = math.MinInt32
	for i := 1; i <= n; i++ {
		// 第i天未持股：第i-1天未持股，或第i天卖出
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		// 第i天持股：第i-1天持股，或第i天买入
		dp[i][1] = max(dp[i-1][1], -prices[i-1])
	}

	return dp[n][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
