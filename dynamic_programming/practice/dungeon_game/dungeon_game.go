package dungeon_game

// https://leetcode-cn.com/problems/dungeon-game
// 解题思路：dp[i][j]表示i,j到终点时需要的最小血量
func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	// base case
	dp[n-1][m-1] = calc(1, dungeon[n-1][m-1])
	for i := n - 2; i >= 0; i-- {
		dp[i][m-1] = calc(dp[i+1][m-1], dungeon[i][m-1])
	}
	for j := m - 2; j >= 0; j-- {
		dp[n-1][j] = calc(dp[n-1][j+1], dungeon[n-1][j])
	}
	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			dp[i][j] = calc(min(dp[i+1][j], dp[i][j+1]), dungeon[i][j])
		}
	}
	return dp[0][0]
}

func calc(a, b int) int {
	c := a - b
	if c <= 0 {
		c = 1
	}
	return c
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
