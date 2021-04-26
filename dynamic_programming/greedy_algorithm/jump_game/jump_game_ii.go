package jump_game

// https://leetcode-cn.com/problems/jump-game-ii
func jump(nums []int) int {
	fast := 0
	res := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > fast {
			fast = i + nums[i]
		}
		if fast >= len(nums)-1 {
			res++
			return res
		}
		if end == i {
			end = fast
			res++
		}
	}
	return res
}
