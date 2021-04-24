package russian_doll_envelopes

import (
	"sort"
)

// https://leetcode-cn.com/problems/russian-doll-envelopes
// 解题思路：将信封宽按升序排列，相同宽的信封按高降序排列，最后计算高的最长递增子序列
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 1 {
		return 1
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})
	hs := make([]int, n)
	for i := 0; i < n; i++ {
		hs[i] = envelopes[i][1]
	}
	return lengthOfLISByBinarySearch(hs)
}

func lengthOfLISByBinarySearch(nums []int) int {
	top := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		poker := nums[i]

		// 二分查找取得左边界
		left, right := 0, res
		for left < right {
			mod := left + (right-left)/2
			if top[mod] > poker {
				right = mod
			} else if top[mod] < poker {
				left = mod + 1
			} else {
				right = mod
			}
		}

		// 没有合适的堆,新建一个
		if left == res {
			res++
		}
		top[left] = poker
	}
	return res
}
