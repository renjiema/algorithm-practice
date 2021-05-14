package stonegame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoneGame2(t *testing.T) {
	assert.Equal(t, true, stoneGame2([]int{3, 9, 1, 2}))
	assert.Equal(t, false, stoneGame2([]int{1, 100, 2}))
	assert.Equal(t, true, stoneGame3([]int{3, 9, 1, 2}))
	assert.Equal(t, false, stoneGame3([]int{1, 100, 2}))
}
