package minimum_path_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	assert.Equal(t, 7, minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	assert.Equal(t, 12, minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
