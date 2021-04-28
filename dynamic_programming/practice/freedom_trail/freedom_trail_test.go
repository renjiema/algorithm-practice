package freedom_trail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRotateSteps(t *testing.T) {
	assert.Equal(t, 4, findRotateSteps("godding", "gd"))
	assert.Equal(t, 13, findRotateSteps("godding", "godding"))
	assert.Equal(t, 137, findRotateSteps("caotmcaataijjxi", "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx"))
}
