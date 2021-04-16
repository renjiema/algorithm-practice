package binary_heap

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	var mh MaxHeap
	heap.Push(&mh, 3)
	heap.Push(&mh, 4)
	heap.Push(&mh, 5)
	heap.Push(&mh, 1)
	heap.Push(&mh, 2)
	t.Log(mh)
	t.Log(heap.Pop(&mh))
	t.Log(mh)
}

type MaxHeap []int

func (mh *MaxHeap) Push(v interface{}) {
	*mh = append(*mh, v.(int))
}
func (mh *MaxHeap) Pop() interface{} {
	res := (*mh)[mh.Len()-1]
	*mh = (*mh)[:mh.Len()-1]
	return res
}

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
