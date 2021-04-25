package longest_common_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	assert.Equal(t, 3, longestCommonSubsequence("abcde", "ace"))
	assert.Equal(t, 3, longestCommonSubsequence("abc", "abc"))
	assert.Equal(t, 0, longestCommonSubsequence("abc", "def"))
	assert.Equal(t, 3, longestCommonSubsequenceByRecursive("abcde", "ace"))
	assert.Equal(t, 3, longestCommonSubsequenceByRecursive("abc", "abc"))
	assert.Equal(t, 0, longestCommonSubsequenceByRecursive("abc", "def"))
}
