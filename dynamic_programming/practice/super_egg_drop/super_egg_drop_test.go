package super_egg_drop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperEggDrop(t *testing.T) {
	assert.Equal(t, 2, superEggDrop(1, 2))
	assert.Equal(t, 3, superEggDrop(2, 6))
	assert.Equal(t, 4, superEggDrop(3, 14))
	assert.Equal(t, 16, superEggDrop(5, 5000))
}
