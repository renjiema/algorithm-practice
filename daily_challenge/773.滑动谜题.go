/*
 * @lc app=leetcode.cn id=773 lang=golang
 *
 * [773] 滑动谜题
 */

// @lc code=start
var void struct{}

func slidingPuzzle(board [][]int) int {
	b := []byte{}
	target := "123450"
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b = append(b, byte(board[i][j]+'0'))

		}
	}
	start := string(b)

	// 记录一维字符串的相邻索引
	neighbor := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}
	// 定义dfs队列，visited
	queue := []string{}
	visited := make(map[string]struct{})
	queue = append(queue, start)
	visited[start] = void
	step := 0
	for len(queue) > 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 判断终止条件
			if cur == target {
				return step
			}
			// 找到0索引
			index := strings.Index(cur, "0")
			// 将数字 0 和相邻的数字交换位置
			for _, j := range neighbor[index] {
				by := []byte(cur)
				by[index], by[j] = by[j], by[index]
				newStr := string(by)
				if _, ok := visited[newStr]; !ok {
					queue = append(queue, newStr)
					visited[newStr] = void
				}
			}
		}
		step++
	}
	return -1
}

// @lc code=end

