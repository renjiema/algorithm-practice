package maximum_subarray

// https://leetcode-cn.com/problems/maximum-subarray
func maxSubArray(nums []int) int {
	sumMax, sum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		sumMax = max(sumMax, sum)
		if sum < 0 {
			// 跳过之前的子序列
			sum = 0
			continue
		}
	}
	return sumMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划解法
func maxSubArrayByDP(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	dp_1 := nums[0]
	res := nums[0]
	for i := 1; i < n; i++ {
		// 状态方程
		dp := max(nums[i], dp_1+nums[i])
		dp_1 = dp
		res = max(res, dp)
	}
	return res
}
