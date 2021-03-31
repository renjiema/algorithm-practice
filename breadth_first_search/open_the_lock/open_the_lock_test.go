package open_the_lock

import (
	"testing"
)

func TestOpenLock(t *testing.T) {
	res := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	t.Logf("解锁最少步骤: %d", res)
	res = openLock([]string{"8888"}, "0009")
	t.Logf("解锁最少步骤: %d", res)
	res = openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	t.Logf("解锁最少步骤: %d", res)
}

func TestTwoWayOpenLock(t *testing.T) {
	res := twoWayOpenLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	t.Logf("解锁最少步骤: %d", res)
	res = twoWayOpenLock([]string{"8888"}, "0009")
	t.Logf("解锁最少步骤: %d", res)
	res = twoWayOpenLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	t.Logf("解锁最少步骤: %d", res)
}

func TestRotateOne(t *testing.T) {
	s := rotateOne("0000", 1)
	t.Log(s)
}
