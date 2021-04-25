package non_overlapping_intervals

import (
	"sort"
)

// https://leetcode-cn.com/problems/non-overlapping-intervals
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		// end升序
		return intervals[i][1] < intervals[j][1]
	})
	res := 0
	x := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if x[1] > intervals[i][0] {
			res++
			continue
		}
		x = intervals[i]
	}
	return res
}

func eraseOverlapIntervals2(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := 0
	x := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < x[1] {
			res++
			if intervals[i][0] == x[0] || intervals[i][1] < x[1] {
				x = intervals[i]
			}
			continue
		}
		x = intervals[i]
	}
	return res
}
