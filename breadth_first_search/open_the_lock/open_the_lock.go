package open_the_lock

// https://leetcode-cn.com/problems/open-the-lock/
// 单向BFS
func openLock(deadends []string, target string) int {
	// 记录deadends
	deads := make(map[string]byte, len(deadends))
	for _, s := range deadends {
		deads[s] = 0x00
	}
	// 记录已经穷举过的密码
	visited := make(map[string]byte)
	queue := []string{}
	// 设置起点
	queue = append(queue, "0000")
	visited["0000"] = 0x00
	step := 0
	for len(queue) > 0 {
		count := len(queue)
		// 将当前队列中的所有节点向四周扩散
		for i := 0; i < count; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 判断是否到达终点
			if _, ok := deads[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}
			// 遍历当前节点的相邻节点
			for j := 0; j < 8; j++ {
				next := rotateOne(cur, j)
				if _, ok := visited[next]; !ok {
					queue = append(queue, next)
					visited[next] = 0x00
				}
			}
		}
		step++
	}
	return -1
}

// 双向BFS
func twoWayOpenLock(deadends []string, target string) int {
	// 添加dead集
	deads := make(map[string]bool)
	for _, v := range deadends {
		deads[v] = true
	}
	// 记录穷举过的密码
	visited := make(map[string]bool)
	// 记录开始队列和结束队列
	startQueue := make(map[string]bool)
	endQueue := make(map[string]bool)
	startQueue["0000"] = true
	endQueue[target] = true
	step := 0
	for len(startQueue) > 0 && len(endQueue) > 0 {
		// 将当前队列中的所有节点向四周扩散
		temp := make(map[string]bool)
		for cur := range startQueue {
			delete(startQueue, cur)
			// 判断是否到达终点
			if deads[cur] {
				continue
			}
			if endQueue[cur] {
				return step
			}
			visited[cur] = true
			// 遍历当前节点的相邻节点
			for i := 0; i < 8; i++ {
				next := rotateOne(cur, i)
				if !visited[next] {
					temp[next] = true
				}
			}
		}
		// 添加步数
		step++
		// 交换start和end
		if len(endQueue) < len(temp) {
			startQueue, endQueue = endQueue, temp
		} else {
			startQueue = temp
		}
	}
	return -1
}

func rotateOne(s string, i int) string {
	by := []byte(s)
	if i < 4 {
		// 左转:0->1
		if by[i] == '9' {
			by[i] = '0'
		} else {
			by[i] = by[i] + 1
		}
	} else {
		i = i - 4
		// 右转：0->9
		if by[i] == '0' {
			by[i] = '9'
		} else {
			by[i] = by[i] - 1
		}
	}
	return string(by)
}
