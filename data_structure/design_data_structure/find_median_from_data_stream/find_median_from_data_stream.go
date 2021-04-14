package find_median_from_data_stream

// https://leetcode-cn.com/problems/find-median-from-data-stream/
// 也可通过container/heap实现,时间复杂度O(logn)
type MedianFinder struct {
	left, right         []int
	leftSize, rightSize int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.leftSize == this.rightSize {
		addNum(&this.left, num)
		this.leftSize++
		if this.leftSize == 1 {
			return
		}
	} else if this.leftSize > this.rightSize {
		addNum(&this.right, num)
		this.rightSize++
	}
	this.balance()

}

func (this *MedianFinder) FindMedian() float64 {
	if this.leftSize == 0 {
		return 0
	}
	if this.leftSize > this.rightSize {
		return float64(this.left[this.leftSize-1])
	}
	return (float64(this.left[this.leftSize-1]+this.right[0]) / 2)
}

// 平衡左右切片
func (this *MedianFinder) balance() {
	for this.left[this.leftSize-1] > this.right[0] {
		l := this.left[this.leftSize-1]
		this.left = this.left[:this.leftSize-1]
		r := this.right[0]
		this.right = this.right[1:]
		addNum(&this.left, r)
		addNum(&this.right, l)
	}
}

// 升序插入元素
func addNum(slice *[]int, value int) {
	if value >= (*slice)[len(*slice)-1] {
		*slice = append(*slice, value)
		return
	}
	if value <= (*slice)[0] {
		*slice = append([]int{value}, *slice...)
		return
	}
	end := len(*slice) - 1
	start := 0
	var mid int
	for start <= end {
		mid = (end + start) / 2
		if (*slice)[mid] > value {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if end < 0 {
		*slice = append([]int{value}, *slice...)
	} else if start > len(*slice)-1 {
		*slice = append(*slice, value)
	} else {
		*slice = append((*slice)[:start], append([]int{value}, (*slice)[start:]...)...)
	}
}
