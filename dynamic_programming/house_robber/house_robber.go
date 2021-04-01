package house_robber

// 打家劫舍共三道题
// 题1：https://leetcode-cn.com/problems/house-robber
var cache []int

func rob_1(nums []int) int {
	return dp_1(nums, 0)
}
func dp_1(nums []int, start int) int {

	// base case
	if start >= len(nums) {
		return 0
	}

	return max(dp_1(nums, start+1), nums[start]+dp_1(nums, start+2))
}

func rob_1_2(nums []int) int {
	cache = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = -1
	}
	cache = append(cache, 0)
	cache = append(cache, 0)
	// cache = append(cache, 0)
	return dp_1_2(nums, 0)
}
func dp_1_2(nums []int, start int) int {
	// base case
	if cache[start] != -1 {
		return cache[start]
	}
	cache[start] = max(dp_1_2(nums, start+1), nums[start]+dp_1_2(nums, start+2))
	return cache[start]
}

func rob_1_3(nums []int) int {
	dp_i_1, dp_i_2 := 0, 0
	var dp int
	for i := len(nums) - 1; i >= 0; i-- {
		dp = max(dp_i_1, nums[i]+dp_i_2)
		dp_i_2, dp_i_1 = dp_i_1, dp
	}
	return dp
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
