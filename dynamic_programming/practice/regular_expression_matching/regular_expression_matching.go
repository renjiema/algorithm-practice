package regular_expression_matching

import "strconv"

// https://leetcode-cn.com/problems/regular-expression-matching
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	cache := make(map[string]bool)
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		key := strconv.Itoa(i) + strconv.Itoa(j)
		if v, ok := cache[key]; ok {
			return v
		}
		if j == m {
			return i == n
		}
		first := i < n && (s[i] == p[j] || p[j] == '.')
		ans := first && dp(i+1, j+1)
		if j <= m-2 && p[j+1] == '*' {
			ans = dp(i, j+2) || (first && dp(i+1, j))
		}
		cache[key] = ans
		return cache[key]
	}

	return dp(0, 0)
}
