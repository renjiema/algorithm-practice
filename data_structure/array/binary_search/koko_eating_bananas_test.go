package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	assert.Equal(t, 4, minEatingSpeed([]int{3, 6, 7, 11}, 8))
	assert.Equal(t, 30, minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	assert.Equal(t, 23, minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}
