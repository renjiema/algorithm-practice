package serialize_and_deserialize_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	root := genNode(0, []int{1, 2, 3, -1, -1, 4, 5})
	c := Constructor()
	ser := c.serialize(root)
	t.Log(ser)
	res := c.deserialize(ser)
	assert.EqualValues(t, root, res)
}

// 生成树节点
func genNode(i int, path []int) *TreeNode {
	t := &TreeNode{Val: path[i]}
	k := 0
	for j := i - 1; j > 0; j = j - 2 {
		if path[j] == -1 {
			k++
		}
	}
	li := 2*(i-k) + 1
	ri := 2*(i-k) + 2
	if i < len(path) && li < len(path) && path[li] != -1 {
		t.Left = genNode(li, path)
	}

	if i < len(path) && ri < len(path) && path[ri] != -1 {
		t.Right = genNode(ri, path)
	}
	return t
}
