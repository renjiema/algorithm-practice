package fibonacci_number

import "testing"

var n int = 20

func TestFib(t *testing.T) {
	res := Fib(n)
	t.Logf("暴力解法, res:%v", res)
}

func TestCacheFib(t *testing.T) {
	res := cacheFib(n)
	t.Logf("缓存解法, res:%v", res)
}

func TestToTopFib(t *testing.T) {
	res := toTopFib(n)
	t.Logf("自底而上的解法, res:%v", res)
}
