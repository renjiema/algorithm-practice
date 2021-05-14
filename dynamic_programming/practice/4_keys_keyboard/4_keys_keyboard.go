package keyskeyboard

// https://leetcode-cn.com/problems/4-keys-keyboard
// 题目描述：键盘上存在四个键:1.A屏幕上输出A、2.Ctrl+A全选、3.Ctrl+C复制、4.Ctrl+V粘贴，每次按一个键，按N次最多输出多少个A
func maxA(N int) int {
	//定义dp，按i次最多输出的A的个数
	dp := make([]int, N+1)
	for i := 1; i <= N; i++ {
		// 第i次按A时输出的A的个数
		dp[i] = dp[i-1] + 1
		for j := 2; j <= i; j++ {
			// 全选 & 复制 dp[j-2]，连续粘贴 i - j 次
			dp[i] = max(dp[i], dp[i-2]*(i-j+1))
		}
	}
	return dp[N]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
