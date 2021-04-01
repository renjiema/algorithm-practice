package find_all_anagrams_in_a_string

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
func findAnagrams(s string, p string) []int {
	// need
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	// window
	window := make(map[byte]int)
	left, right, valid := 0, 0, 0
	res := []int{}
	for right < len(s) {
		b := s[right]
		right++
		if need[b] > 0 {
			window[b]++
			if need[b] == window[b] {
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		if right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			b := s[left]
			left++
			if need[b] > 0 {
				if need[b] == window[b] {
					valid--
				}
				window[b]--
			}
		}
	}
	return res
}
