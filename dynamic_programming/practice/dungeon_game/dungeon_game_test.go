package dungeon_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMinimumHP(t *testing.T) {
	assert.Equal(t, 7, calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
	assert.Equal(t, 1, calculateMinimumHP([][]int{{0, 0}}))
}
