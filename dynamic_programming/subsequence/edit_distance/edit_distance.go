package edit_distance

// https://leetcode-cn.com/problems/edit-distance/
// 递归缓存解法
func minDistance(word1 string, word2 string) int {
	var dp func(int, int) int
	cache := make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		cache[i] = make([]int, len(word2))
	}
	dp = func(i, j int) int {
		// base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if cache[i][j] != 0 {
			return cache[i][j]
		}
		if word1[i] == word2[j] {
			cache[i][j] = dp(i-1, j-1)
		} else {
			cache[i][j] = min(dp(i, j-1)+1, //增
				dp(i-1, j)+1,   //删
				dp(i-1, j-1)+1, //替
			)
		}
		return cache[i][j]
	}
	return dp(len(word1)-1, len(word2)-1)
}

func min(i, j, k int) int {
	if i > j {
		i = j
	}
	if i > k {
		i = k
	}
	return i
}

// dp数组解法
func minDistance2(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}
	// base case
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[n][m]
}

// TODO:状态压缩，将二维数组压缩为一维数组
