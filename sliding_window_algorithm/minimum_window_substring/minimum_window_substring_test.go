package minimum_window_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	expected := "BANC"
	res := minWindow("ADOBECODEBANC", "ABC")
	assert.Equal(t, expected, res)
	expected = "a"
	res = minWindow("a", "a")
	assert.Equal(t, "a", res)
}
