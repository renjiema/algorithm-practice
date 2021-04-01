package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	res := lengthOfLongestSubstring("abcabcbb")
	assert.Equal(t, 3, res)
	res = lengthOfLongestSubstring("bbbbb")
	assert.Equal(t, 1, res)
	res = lengthOfLongestSubstring("pwwkew")
	assert.Equal(t, 3, res)
	res = lengthOfLongestSubstring("")
	assert.Equal(t, 0, res)
}
