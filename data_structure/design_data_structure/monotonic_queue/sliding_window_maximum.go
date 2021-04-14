package monotonic_queue

// https://leetcode-cn.com/problems/sliding-window-maximum
func maxSlidingWindow(nums []int, k int) []int {
	mq := []int{}
	res := make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		for len(mq) > 0 && mq[len(mq)-1] < nums[i] {
			mq = mq[:len(mq)-1]
		}
		mq = append(mq, nums[i])
		if i >= k-1 {
			res[i-k+1] = mq[0]
			if mq[0] == nums[i-k+1] {
				mq = mq[1:]
			}
		}
	}
	return res
}
