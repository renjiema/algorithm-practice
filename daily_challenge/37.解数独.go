/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	var dfs func(n, m int) bool
	dfs = func(n, m int) bool {
		// 处理边界
		if n > 8 {
			return true
		}
		if m > 8 {
			return dfs(n+1, 0)
		}
		if board[n][m] != '.' {
			return dfs(n, m+1)
		}
		var i byte
		row := n - n%3
		col := m - m%3
		for i = '1'; i <= '9'; i++ {
			// 过滤无效数
			if !isValid(n, m, i, row, col, board) {
				continue
			}
			// 做选择
			board[n][m] = i
			if dfs(n, m+1) {
				return true
			}
			// 取消选择
			board[n][m] = '.'

		}
		return false
	}
	dfs(0, 0)
}

func isValid(n, m int, i byte, row, col int, board [][]byte) bool {
	// 1.过滤行和列
	for j := 0; j < 9; j++ {
		if i == board[n][j] || i == board[j][m] || i == board[row+j/3][col+j%3] {
			return false
		}
	}

	// 2.过滤3x3宫
	// for j := 0; j < 3; j++ {
	// 	for k := 0; k < 3; k++ {
	// 		if i == board[row+j][col+k] {
	// 			return false
	// 		}
	// 	}
	// }
	return true
}

// @lc code=end

