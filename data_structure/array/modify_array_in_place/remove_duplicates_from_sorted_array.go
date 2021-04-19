package modify_array_in_place

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	lens := 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[lens] = nums[i]
			lens++
		}
	}
	return lens
}
