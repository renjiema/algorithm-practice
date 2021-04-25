package minimum_ascii_delete_sum_for_two_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumDeleteSum(t *testing.T) {
	assert.Equal(t, 231, minimumDeleteSum("sea", "eat"))
	assert.Equal(t, 403, minimumDeleteSum("delete", "leet"))
	assert.Equal(t, 231, minimumDeleteSumByRecursive("sea", "eat"))
	assert.Equal(t, 403, minimumDeleteSumByRecursive("delete", "leet"))
}
