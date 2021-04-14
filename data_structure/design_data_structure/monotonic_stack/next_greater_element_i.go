package monotonic_stack

// https://leetcode-cn.com/problems/next-greater-element-i/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	nmap := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		// 移除两个高值中的低值
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < 1 {
			nmap[nums2[i]] = -1
		} else {
			nmap[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = nmap[v]
	}
	return res
}
