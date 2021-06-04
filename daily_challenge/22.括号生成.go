/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var dfs func(left, right int, s []byte)
	dfs = func(left, right int, s []byte) {
		// 过滤不成对括号
		if left > right || left < 0 || right < 0 {
			return
		}
		if right == 0 {
			res = append(res, string(s))
		}
		// 添加左括号
		s = append(s, '(')
		dfs(left-1, right, s)
		s = s[:len(s)-1]
		// 添加左括号
		s = append(s, ')')
		dfs(left, right-1, s)
		s = s[:len(s)-1]

	}
	dfs(n, n, []byte{})
	return res
}

// @lc code=end

