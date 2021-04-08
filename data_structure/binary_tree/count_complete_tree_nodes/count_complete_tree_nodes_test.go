package count_complete_tree_nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNodes(t *testing.T) {
	root := NewTreeNode(0, []int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, 6, countNodes(root))
}
