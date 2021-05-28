/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	// 计算子集和
	var sum int
	n := len(nums)
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	// 判断是否有大于target的值
	for _, num := range nums {
		if num > target {
			return false
		}
	}
	// 记录索引是否使用
	indexs := make([]bool, n)
	var bt func(currentK, currSum, index int) bool
	bt = func(currentK, currSum, index int) bool {
		// 结束条件
		if currentK == 0 {
			return true
		}
		if currSum == target {
			return bt(currentK-1, 0, 0)
		}
		for i := index; i < n; i++ {
			if indexs[i] || currSum+nums[i] > target {
				continue
			}
			// 选择
			indexs[i] = true
			currSum += nums[i]
			if bt(currentK, currSum, i+1) {
				return true
			}
			// 撤销选择
			indexs[i] = false
			currSum -= nums[i]
		}
		return false
	}
	return bt(k, 0, 0)
}

// @lc code=end
