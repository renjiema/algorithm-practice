package jump_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))
	assert.Equal(t, 3, jump([]int{1, 2, 1, 1, 1}))
	assert.Equal(t, 2, jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))
	assert.Equal(t, 0, jump([]int{0}))
}
