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
