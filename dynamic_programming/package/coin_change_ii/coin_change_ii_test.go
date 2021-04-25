package coin_change_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChange(t *testing.T) {
	assert.Equal(t, 4, change(5, []int{1, 2, 5}))
	assert.Equal(t, 0, change(3, []int{2}))
	assert.Equal(t, 1, change(10, []int{10}))
	assert.Equal(t, 4, changeStateOfCompression(5, []int{1, 2, 5}))
	assert.Equal(t, 0, changeStateOfCompression(3, []int{2}))
	assert.Equal(t, 1, changeStateOfCompression(10, []int{10}))
}
