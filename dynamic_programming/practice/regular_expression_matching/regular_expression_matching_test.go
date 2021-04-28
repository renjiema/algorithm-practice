package regular_expression_matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	assert.Equal(t, false, isMatch("aa", "a"))
	assert.Equal(t, true, isMatch("a", "a*"))
	assert.Equal(t, true, isMatch("ab", ".*"))
	assert.Equal(t, true, isMatch("aab", "c*a*b"))
	assert.Equal(t, false, isMatch("mississippi", "mis*is*p*."))
	assert.Equal(t, false, isMatch("ab", ".*c"))
	assert.Equal(t, true, isMatch("aaa", "ab*a*c*a"))
	assert.Equal(t, false, isMatch("aaba", "ab*a*c*a"))
	assert.Equal(t, true, isMatch("a", "ab*"))
	assert.Equal(t, true, isMatch("bbbba", ".*a*a"))
}
