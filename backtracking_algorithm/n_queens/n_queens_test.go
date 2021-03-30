package n_queens

import "testing"

func TestSolveNQueens(t *testing.T) {
	res := solveNQueens(4)
	t.Log(res)
}
