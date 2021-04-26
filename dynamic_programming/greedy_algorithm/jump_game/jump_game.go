package jump_game

// https://leetcode-cn.com/problems/jump-game
// 解题思路：让每一步都跳最远，记录最终能跳的最远值，只要最远值大于最的所有就能达到
func canJump(nums []int) bool {
	n := len(nums) - 1
	maxIndex := 0
	for i := 0; i < n; i++ {
		if i+nums[i] > maxIndex {
			maxIndex = i + nums[i]
		}
		//最大距离到不了下一个索引
		if maxIndex <= i {
			return false
		}
	}
	return maxIndex >= n
}

// 动态规划压缩状态后就和贪心算法一致了
func canJumpByDP(nums []int) bool {
	n := len(nums) - 1
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], i+nums[i])
		if dp[i] <= i {
			return false
		}
	}
	return dp[n-1] >= n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
