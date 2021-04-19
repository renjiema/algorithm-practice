package modify_array_in_place

// https://leetcode-cn.com/problems/remove-element
func removeElement(nums []int, val int) int {
	n := len(nums)
	slow := 0
	for fast := 0; fast < n; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
