package find_all_anagrams_in_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	res := findAnagrams("cbaebabacd", "abc")
	assert.EqualValues(t, []int{0, 6}, res)
	res = findAnagrams("abab", "ab")
	assert.EqualValues(t, []int{0, 1, 2}, res)
	res = findAnagrams("baac", "aa")
	assert.EqualValues(t, []int{1}, res)
}
