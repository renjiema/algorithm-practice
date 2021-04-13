package lfu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFUCache(t *testing.T) {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	assert.Equal(t, 1, lfu.Get(1))
	lfu.Put(3, 3)
	assert.Equal(t, -1, lfu.Get(2))
	assert.Equal(t, 3, lfu.Get(3))
	lfu.Put(4, 4)
	assert.Equal(t, -1, lfu.Get(1))
	assert.Equal(t, 3, lfu.Get(3))
	assert.Equal(t, 4, lfu.Get(4))
}
