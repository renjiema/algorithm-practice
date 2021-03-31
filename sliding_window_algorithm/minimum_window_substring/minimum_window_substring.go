package minimum_window_substring

// https://leetcode-cn.com/problems/minimum-window-substring
func minWindow(s string, t string) string {
	// need
	need := make(map[byte]int)
	for _, v := range []byte(t) {
		need[v]++
	}
	// window
	window := make(map[byte]int)
	left, right := 0, 0
	valid := 0
	// 记录最小覆盖子串的索引范围
	start, lens := 0, len(s)+1
	for right < len(s) {
		// 移动的字符
		b := s[right]
		right++
		if need[b] > 0 {
			window[b]++
			if need[b] == window[b] {
				valid++
			}
		}
		for valid == len(need) {
			// 更新最小覆盖字串
			if right-left < lens {
				start = left
				lens = right - left
			}
			// 窗口左移
			b := s[left]
			left++
			if need[b] > 0 {
				if window[b] == need[b] {
					valid--
				}
				window[b]--
			}
		}
	}
	if lens == len(s)+1 {
		return ""
	}
	return s[start : start+lens]
}
