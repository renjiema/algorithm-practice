/*
 * @lc app=leetcode.cn id=810 lang=golang
 *
 * [810] 黑板异或游戏
 */

// @lc code=start
func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}
	res := 0
	for _, nums := range nums {
		res ^= nums
	}
	return res == 0
}

// @lc code=end

