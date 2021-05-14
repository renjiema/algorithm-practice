package stonegame

// 石子游戏改：石头堆数改为任意整数，石子总数也改为任意总数
func stoneGame2(piles []int) bool {
	n := len(piles)
	// 定义dp数组,i和j表示piles的下标，k两个值，0表示先手获取的最大值、1表示后手获取的最大值
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 2)
			if i == j {
				dp[i][j][0] = piles[i]
				dp[i][j][1] = 0
			}
		}
	}
	//右对角线遍历
	for i, j := 0, 1; j < n && j > i; {
		left := dp[i+1][j][1] + piles[i]
		right := dp[i][j-1][1] + piles[j]
		if left > right {
			dp[i][j][0] = left
			dp[i][j][1] = dp[i+1][j][0]
		} else {
			dp[i][j][0] = right
			dp[i][j][1] = dp[i][j-1][0]
		}
		if j == n-1 && i != 0 {
			j = j - i + 1
			i = 0
		} else {
			i++
			j++
		}
	}

	return dp[0][n-1][0] > dp[0][n-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 状态压缩
func stoneGame3(piles []int) bool {
	n := len(piles)
	// 定义dp数组,i和j表示piles的下标，k两个值，0表示先手获取的最大值、1表示后手获取的最大值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = piles[i]
	}
	pre := dp[0]
	//右对角线遍历
	for i, j := 0, 1; j < n; {
		left := dp[j][1] + piles[i]
		right := pre[1] + piles[j]
		temp := []int{dp[j][0], dp[j][1]}
		if left > right {
			dp[j][1] = dp[j][0]
			dp[j][0] = left
		} else {
			dp[j][1] = pre[0]
			dp[j][0] = right
		}
		if j == n-1 && i != 0 {
			j = j - i + 1
			i = 0
			pre = dp[j-1]
		} else {
			i++
			j++
			pre = temp
		}
	}

	return dp[n-1][0] > dp[n-1][1]
}
