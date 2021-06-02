/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var dfs func(i int, temp []int)
	dfs = func(i int, temp []int) {
		// 过滤长度不够的情况
		if len(temp)+n-i < k {
			return
		}
		if len(temp) == k {
			res = append(res, append([]int{}, temp...))
			return
		}
		dfs(i+1, append(temp, i+1))
		dfs(i+1, temp)
	}
	dfs(0, []int{})
	return res
}

// @lc code=end

