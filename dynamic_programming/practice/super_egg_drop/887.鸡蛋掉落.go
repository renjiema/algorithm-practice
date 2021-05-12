/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */
package super_egg_drop

// @lc code=start
func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}
	m := 0
	for dp[k][m] < n {
		m++
		for i := 1; i <= k; i++ {
			dp[i][m] = dp[i-1][m-1] + dp[i][m-1] + 1
		}
	}
	return m
}

// @lc code=end
