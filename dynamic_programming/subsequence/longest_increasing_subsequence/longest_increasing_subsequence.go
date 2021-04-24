package longest_increasing_subsequence

// https://leetcode-cn.com/problems/longest-increasing-subsequence/
// 动态规划解法
func lengthOfLISByDP(nums []int) int {
	dp := make([]int, len(nums))
	// base case
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

	}
	res := 0
	for _, v := range dp {
		res = max(res, v)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 原理：数小的牌压到点数比它大的牌上；如果当前牌点数较大没有可以放置的堆，则新建一个堆，把这张牌放进去；
// 如果当前牌有多个堆可供选择，则选择最左边的那一堆放置。最后牌的堆数就是最长递增子序列的长度。具体证明不懂。
func lengthOfLISByBinarySearch(nums []int) int {
	top := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		poker := nums[i]

		// 二分查找取得左边界
		left, right := 0, res
		for left < right {
			mod := left + (right-left)/2
			if top[mod] > poker {
				right = mod
			} else if top[mod] < poker {
				left = mod + 1
			} else {
				right = mod
			}
		}

		// 没有合适的堆,新建一个
		if left == res {
			res++
		}
		top[left] = poker
	}
	return res
}
