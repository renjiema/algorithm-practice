package minimum_number_of_arrows_to_burst_balloons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinArrowShots(t *testing.T) {
	assert.Equal(t, 2, findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	assert.Equal(t, 4, findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	assert.Equal(t, 2, findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
}
