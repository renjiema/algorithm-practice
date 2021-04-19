package random_pick_with_blacklist

import "testing"

func TestSolution(t *testing.T) {
	s := Constructor(4, []int{0, 1})
	t.Log(s.Pick())
	t.Log(s.Pick())
	t.Log(s.Pick())
	t.Log(s.Pick())
	t.Log(s.Pick())
}
