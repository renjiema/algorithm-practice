package coin_change_ii

// https://leetcode-cn.com/problems/coin-change-2
func change(amount int, coins []int) int {
	n := len(coins) + 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, amount+1)
		// base case
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][amount]
}

// 状态压缩
func changeStateOfCompression(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i] {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
