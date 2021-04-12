package union_find

import "testing"

func TestSolve(t *testing.T) {
	board := [][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}}
	solve(board)
	t.Log(board)
}
func TestSolveUF(t *testing.T) {
	board := [][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}
	solveUF(board)
	t.Log(board)
}
