/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package besttimetobuyandsellstock

// @lc code=start
func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			res = max(res, prices[j]-prices[i])
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
