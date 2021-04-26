package freedom_trail

// https://leetcode-cn.com/problems/freedom-trail
func findRotateSteps(ring string, key string) int {
	n, m := len(ring), len(key)
	cache := make([][]int, n)
	indexMap := make(map[byte][]int)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		b := ring[i]
		if _, ok := indexMap[b]; !ok {
			indexMap[b] = make([]int, 0)
		}
		indexMap[b] = append(indexMap[b], i)
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// key查找完毕
		if j == m {
			return 0
		}
		if cache[i][j] != 0 {
			return cache[i][j]
		}
		res := n * (m + 1)
		for _, v := range indexMap[key[j]] {
			dist := v - i
			if dist < 0 {
				dist = -dist
			}
			count := min(dist, n-dist)
			res = min(res, 1+count+dp(v, j+1))
		}

		cache[i][j] = res
		return cache[i][j]
	}
	return dp(0, 0)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
