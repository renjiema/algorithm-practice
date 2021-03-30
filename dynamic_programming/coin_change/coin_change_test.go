package coin_change

import "testing"

var cs = []int{2, 5, 10, 1}
var amount = 27

func TestViolenceCoinChange(t *testing.T) {
	res := ViolenceCoinChange(cs, amount)
	t.Logf("凑零钱暴力解法, res:%v", res)
}

func TestCacheCoinChange(t *testing.T) {
	res := CacheCoinChange(cs, amount)
	t.Logf("凑零钱缓存递归解法, res:%v", res)
}

func TestToTopCoinChange(t *testing.T) {
	res := ToTopCoinChange(cs, amount)
	t.Logf("凑零钱自底而上解法, res:%v", res)
}
