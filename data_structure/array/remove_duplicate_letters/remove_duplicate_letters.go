package remove_duplicate_letters

// https://leetcode-cn.com/problems/remove-duplicate-letters
func removeDuplicateLetters(s string) string {
	//记录每个字符个数
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	res := []byte{}
	// 记录字符是否以添加
	recode := make([]bool, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']--
		if recode[s[i]-'a'] {
			continue
		}
		// 保证字典序最小
		for len(res) > 0 && res[len(res)-1] > s[i] {
			if count[res[len(res)-1]-'a'] == 0 {
				break
			}
			recode[res[len(res)-1]-'a'] = false
			res = res[:len(res)-1]
		}
		res = append(res, s[i])
		recode[s[i]-'a'] = true
	}
	return string(res)
}
