package permutations

import "testing"

func TestPermute(t *testing.T) {
	nums := []int{5,4,6,2}
	res := Permute(nums)
	t.Log(res)
}
