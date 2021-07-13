/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */

// @lc code=start
var void struct{}

func openLock(deadends []string, target string) int {
	// 记录deadends
	deads := make(map[string]struct{}, len(deadends))
	for i := 0; i < len(deadends); i++ {
		deads[deadends[i]] = void
	}
	// 记录visited
	visited := make(map[string]struct{})
	visited["0000"] = void
	// 记录节点
	queue := []string{}
	queue = append(queue, "0000")
	step := 0
	for len(queue) > 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 判断是否为deadends
			if _, ok := deads[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}
			// 遍历当前节点的相邻节点
			for j := 0; j < 8; j++ {
				next := rotateOne(cur, j)
				if _, ok := visited[next]; ok {
					continue
				}
				queue = append(queue, next)
				visited[next] = void
			}

		}
		step++
	}
	return -1

}
func rotateOne(cur string, i int) string {
	bytes := []byte(cur)
	if i < 4 {
		// 左转
		if bytes[i] == '9' {
			bytes[i] = '0'
		} else {
			bytes[i] = bytes[i] + 1
		}
	} else {
		//右转
		i = i - 4
		if bytes[i] == '0' {
			bytes[i] = '9'
		} else {
			bytes[i] = bytes[i] - 1
		}
	}
	return string(bytes)
}

// @lc code=end

