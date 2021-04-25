package delete_operation_for_two_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	assert.Equal(t, 2, minDistance("sea", "eat"))
	assert.Equal(t, 2, minDistanceByIteration("sea", "eat"))
}
