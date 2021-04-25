package minimum_number_of_arrows_to_burst_balloons

import "sort"

// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	res := 1
	x := points[0]
	for i := 1; i < len(points); i++ {
		if points[i][0] > x[1] {
			res++
			x = points[i]
		}
	}
	return res
}
