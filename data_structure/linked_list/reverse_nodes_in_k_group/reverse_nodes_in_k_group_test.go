package reverse_nodes_in_k_group

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	root := genNode([]int{1, 2, 3, 4, 5})
	k := 2
	res := reverseKGroup(root, k)
	assert.EqualValues(t, genNode([]int{2, 1, 4, 3, 5}), res)
	root = genNode([]int{1, 2})
	k = 2
	assert.EqualValues(t, genNode([]int{2, 1}), reverseKGroup(root, k))
}

func genNode(nums []int) *ListNode {
	var root *ListNode = &ListNode{Val: nums[0]}
	cur := root
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val: nums[i]}
		cur.Next = node
		cur = node
	}
	return root
}
