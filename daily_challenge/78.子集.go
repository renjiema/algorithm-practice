/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	var dfs func(i int)
	set := make([]int, 0)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int{}, set...))
			return
		}
		set = append(set, nums[i])
		dfs(i + 1)
		set = set[:len(set)-1]
		dfs(i + 1)
	}
	dfs(0)
	return res
}

// @lc code=end

