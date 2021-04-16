package binary_search

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, getMax(piles)
	for left < right {
		mid := left + (right-left)/2
		if enableComplete(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 能否在k速度下吃完
func enableComplete(piles []int, h int, k int) bool {
	for _, v := range piles {
		if v%k > 0 {
			h--
		}
		h = h - v/k
	}
	return h >= 0
}

// 最大速度
func getMax(piles []int) int {
	max := piles[0]
	for i := 1; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
		}
	}
	return max
}
