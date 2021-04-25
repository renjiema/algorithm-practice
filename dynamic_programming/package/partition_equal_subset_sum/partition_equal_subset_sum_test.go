package partition_equal_subset_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {
	assert.Equal(t, true, canPartition([]int{1, 5, 11, 5}))
	assert.Equal(t, false, canPartition([]int{1, 2, 3, 5}))
	assert.Equal(t, false, canPartition([]int{1, 2, 5}))
	assert.Equal(t, true, canPartitionStateOfCompression([]int{1, 5, 11, 5}))
	assert.Equal(t, false, canPartitionStateOfCompression([]int{1, 2, 3, 5}))
	assert.Equal(t, false, canPartitionStateOfCompression([]int{1, 2, 5}))
}
