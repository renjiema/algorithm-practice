package permutations

// 全排列
func Permute(nums []int) [][]int {
	res := [][]int{}
	track := make([]int, 0)
	backtrack(nums, track, &res)
	return res
}

func backtrack(nums []int, track []int, res *[][]int) {
	// 结束条件
	if len(nums) == len(track) {
		c := make([]int, len(track))
		copy(c, track)
		*res = append(*res, c)
		return
	}
	// 循环
	for _, num := range nums {
		// 排除路径
		if contains(track, num) {
			continue
		}
		// 做选择
		track = append(track, num)
		// 进入下一层决策树
		backtrack(nums, track, res)
		// 撤销选择
		track = track[:len(track)-1]
	}
}
func contains(track []int, num int) bool {
	for _, v := range track {
		if v == num {
			return true
		}
	}
	return false
}
