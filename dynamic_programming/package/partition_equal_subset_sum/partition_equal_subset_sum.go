package partition_equal_subset_sum

// https://leetcode-cn.com/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	n := len(nums)
	// dp[i][j]:i表示数组长,j表示背包容量,值为是否能填满,状态压缩,dp[i][j]=dp[i-1]...,因此删除i
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				// 背包容量不足
				dp[i][j] = dp[i-1][j]
			} else {
				// 不装或装
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][sum]
}
func canPartitionStateOfCompression(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	n := len(nums)
	// dp[i][j]:i表示数组长,j表示背包容量,值为是否能填满,状态压缩,dp[i][j]=dp[i-1]...,因此删除i
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		// 倒序是防止当前i 计算的dp[j]影响到后续dp[j]计算,倒序计算dp[j]时,dp[j-nums[i-1]]的值一定时小于i时计算的结果
		for j := sum; j >= 0; j-- {
			if j-nums[i-1] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i-1]]
			}
		}
	}
	return dp[sum]
}
