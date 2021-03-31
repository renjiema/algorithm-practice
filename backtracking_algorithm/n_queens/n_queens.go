package n_queens

// https://leetcode-cn.com/problems/n-queens/
// n皇后问题
func solveNQueens(n int) [][]string {
	res := [][]string{}
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		b := make([]byte, n)
		for i := 0; i < n; i++ {
			b[i] = '.'
		}
		board[i] = b
	}
	backtrack(0, &board, &res)
	return res
}

func backtrack(row int, board *[][]byte, res *[][]string) {
	// 结束条件
	if row == len(*board) {
		path := make([]string, 0, len(*board))
		for _, v := range *board {
			path = append(path, string(v))
		}
		*res = append(*res, path)
		return
	}
	// 循环
	n := len(*board)
	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValid(*board, row, col) {
			continue
		}
		// 选择
		(*board)[row][col] = 'Q'
		// 进入下一决策
		backtrack(row+1, board, res)
		// 撤销选择
		(*board)[row][col] = '.'
	}
}

func isValid(board [][]byte, row int, col int) bool {
	// 判断列
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 判断右上角
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 判断左上角
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
