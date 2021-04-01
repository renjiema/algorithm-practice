package house_robber

// 题2：https://leetcode-cn.com/problems/house-robber-ii
func rob_2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(robRange(nums[:n-1]), robRange(nums[1:]))
}

func robRange(nums []int) int {
	var dp, dp_i_1, dp_i_2 int
	for i := len(nums) - 1; i >= 0; i-- {
		dp = max(dp_i_1, nums[i]+dp_i_2)
		dp_i_1, dp_i_2 = dp, dp_i_1
	}
	return dp
}
