/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	sum, minSum := 0, math.MaxInt32
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		if sum < minSum {
			minSum = sum
			//sum最小时，下一个就是起点
			start = i + 1
		}
	}
	if sum < 0 {
		return -1
	}
	if start == len(gas) {
		return 0
	}
	return start
}

// @lc code=end

