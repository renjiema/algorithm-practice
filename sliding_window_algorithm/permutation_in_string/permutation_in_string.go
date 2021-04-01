package permutation_in_string

// https://leetcode-cn.com/problems/permutation-in-string
func checkInclusion(s1 string, s2 string) bool {
	// need
	need := make(map[byte]int)
	for _, v := range []byte(s1) {
		need[v]++
	}
	// window
	window := make(map[byte]int)
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		// 移动的字符
		b := s2[right]
		right++
		//修改窗口数据
		if need[b] > 0 {
			window[b]++
			if window[b] == need[b] {
				valid++
			}
		}
		// 判断左边界收缩
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			b := s2[left]
			left++
			//修改窗口数据
			if window[b] > 0 {
				if window[b] == need[b] {
					valid--
				}
				window[b]--
			}
		}
	}
	return false
}

// 优化
func checkInclusion2(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cnt := make([]int, 0, 1)
	for i, s := range s1 {
		cnt[s2[i]-'a']++
		cnt[s-'a']--
	}
	diff := 0
	for _, v := range cnt {
		if v != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	n := len(s1)
	for i := n; i < len(s2); i++ {
		r, l := s2[i]-'a', s2[i-n]-'a'
		if l == r {
			continue
		}
		if cnt[l] == 0 {
			diff++
		}
		cnt[l]--
		if cnt[l] == 0 {
			diff--
		}
		if cnt[r] == 0 {
			diff++
		}
		cnt[r]++
		if cnt[r] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}
