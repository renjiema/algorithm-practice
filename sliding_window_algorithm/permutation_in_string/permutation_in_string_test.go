package permutation_in_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	res := checkInclusion("ab", "eidbaooo")
	assert.Equal(t, true, res)
	res = checkInclusion("ab", "eidboaoo")
	assert.Equal(t, false, res)
}
func TestCheckInclusion2(t *testing.T) {
	res := checkInclusion2("ab", "eidbaooo")
	assert.Equal(t, true, res)
	res = checkInclusion2("ab", "eidboaoo")
	assert.Equal(t, false, res)
}
