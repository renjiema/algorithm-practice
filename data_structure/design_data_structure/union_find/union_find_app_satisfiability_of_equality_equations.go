package union_find

// https://leetcode-cn.com/problems/satisfiability-of-equality-equations
func equationsPossible(equations []string) bool {
	uf := NewUF(26)
	// 处理相等
	for _, equation := range equations {
		if equation[1] == '=' {
			uf.Union(int(equation[0]-'a'), int(equation[3]-'a'))
		}
	}
	for _, equation := range equations {
		if equation[1] == '!' && uf.Connected(int(equation[0]-'a'), int(equation[3]-'a')) {
			return false
		}
	}
	return true
}
