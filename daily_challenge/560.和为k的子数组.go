/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	n := len(nums)
	count := 0
	// 计算前缀和
	preSum := make(map[int]int)
	// base case
	preSum[0] = 1
	sum_i := 0
	for i := 0; i < n; i++ {
		sum_i += nums[i]
		sum_j := sum_i - k
		if val, ok := preSum[sum_j]; ok {
			count += val
		}
		preSum[sum_i] = preSum[sum_i] + 1
	}
	return count
}

// @lc code=end

