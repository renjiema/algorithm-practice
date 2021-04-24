package russian_doll_envelopes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxEnvelopes(t *testing.T) {
	assert.Equal(t, 3, maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	assert.Equal(t, 1, maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}}))
	assert.Equal(t, 7, maxEnvelopes([][]int{{1, 2}, {2, 3}, {3, 4}, {3, 5}, {4, 5}, {5, 5}, {5, 6}, {6, 7}, {7, 8}}))
}
