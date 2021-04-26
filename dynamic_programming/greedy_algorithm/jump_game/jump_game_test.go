package jump_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	assert.Equal(t, true, canJump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, false, canJump([]int{3, 2, 1, 0, 4}))
	assert.Equal(t, true, canJumpByDP([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, false, canJumpByDP([]int{3, 2, 1, 0, 4}))
}
