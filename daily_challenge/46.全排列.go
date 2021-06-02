/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	var dfs func(temp []int)
	index := make([]bool, n)
	dfs = func(temp []int) {
		if len(temp) == n {
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if index[i] {
				continue
			}
			index[i] = true
			dfs(append(temp, nums[i]))
			index[i] = false
		}
	}
	dfs([]int{})
	return res

}

// @lc code=end

