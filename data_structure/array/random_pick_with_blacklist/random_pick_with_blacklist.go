package random_pick_with_blacklist

import "math/rand"

// https://leetcode-cn.com/problems/random-pick-with-blacklist
type Solution struct {
	// 将[0,size)中黑名单的值映射为非黑名单值
	blackMap map[int]int
	size     int
}

func Constructor(N int, blacklist []int) Solution {
	// end := len(blacklist) - 1
	n := N - len(blacklist)
	s := Solution{blackMap: make(map[int]int), size: n}
	for _, v := range blacklist {
		s.blackMap[v] = -1
	}
	last := N - 1
	for _, v := range blacklist {
		if v >= n {
			continue
		}
		for _, ok := s.blackMap[last]; ok; {
			last--
			_, ok = s.blackMap[last]
		}
		s.blackMap[v] = last
		last--
	}
	return s
}

func (this *Solution) Pick() int {
	r := rand.Int() % this.size
	if v, ok := this.blackMap[r]; ok {
		return v
	}
	return r
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
