package insert_delete_getrandom_o1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizedSet(t *testing.T) {
	rs := Constructor()
	assert.Equal(t, true, rs.Insert(0))
	assert.Equal(t, true, rs.Insert(1))
	assert.Equal(t, true, rs.Remove(0))
	assert.Equal(t, true, rs.Insert(2))
	assert.Equal(t, true, rs.Remove(1))
	assert.Equal(t, 2, rs.GetRandom())

}
