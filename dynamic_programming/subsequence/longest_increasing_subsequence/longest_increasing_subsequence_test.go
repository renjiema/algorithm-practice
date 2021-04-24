package longest_increasing_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	assert.Equal(t, 4, lengthOfLISByDP([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Equal(t, 4, lengthOfLISByDP([]int{0, 1, 0, 3, 2, 3}))
	assert.Equal(t, 1, lengthOfLISByDP([]int{7, 7, 7, 7, 7}))
	assert.Equal(t, 1, lengthOfLISByDP([]int{0}))
	assert.Equal(t, 6, lengthOfLISByDP([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	assert.Equal(t, 4, lengthOfLISByBinarySearch([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Equal(t, 4, lengthOfLISByBinarySearch([]int{0, 1, 0, 3, 2, 3}))
	assert.Equal(t, 1, lengthOfLISByBinarySearch([]int{7, 7, 7, 7, 7}))
	assert.Equal(t, 1, lengthOfLISByBinarySearch([]int{0}))
	assert.Equal(t, 6, lengthOfLISByBinarySearch([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}
