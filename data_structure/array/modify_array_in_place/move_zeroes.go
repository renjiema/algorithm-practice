package modify_array_in_place

// https://leetcode-cn.com/problems/move-zeroes
func moveZeroes(nums []int) {
	slow, fast, n := 0, 0, len(nums)
	for fast < n {
		if nums[fast] != 0 {
			if slow != fast {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++
		}
		fast++
	}
}
