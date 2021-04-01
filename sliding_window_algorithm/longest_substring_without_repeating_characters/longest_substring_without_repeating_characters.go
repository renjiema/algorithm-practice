package longest_substring_without_repeating_characters

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	// window
	window := make(map[byte]int)
	left, right, res := 0, 0, 0
	for right < len(s) {
		b := s[right]
		right++
		window[b]++
		for window[b] > 1 {
			window[s[left]]--
			left++
		}
		if right-left > res {
			res = right - left
		}
		if len(s)-left < res {
			break
		}
	}
	return res
}
