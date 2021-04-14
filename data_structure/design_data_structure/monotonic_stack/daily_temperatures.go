package monotonic_stack

// https://leetcode-cn.com/problems/daily-temperatures
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := []int{}
	for i := len(T) - 1; i >= 0; i-- {
		for len(stack) > 0 && T[stack[len(stack)-1]] <= T[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < 1 {
			res[i] = 0
		} else {
			res[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return res
}
