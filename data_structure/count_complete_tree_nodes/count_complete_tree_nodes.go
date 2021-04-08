package count_complete_tree_nodes

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(i int, path []int) *TreeNode {
	t := &TreeNode{Val: path[i]}
	k := 0
	for j := i - 1; j > 0; j = j - 2 {
		if path[j] == 0 {
			k++
		}
	}
	li := 2*(i-k) + 1
	ri := 2*(i-k) + 2
	if i < len(path) && li < len(path) && path[li] != 0 {
		t.Left = NewTreeNode(li, path)
	}

	if i < len(path) && ri < len(path) && path[ri] != 0 {
		t.Right = NewTreeNode(ri, path)
	}
	return t
}

// https://leetcode-cn.com/problems/count-complete-tree-nodes
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 左右树高度
	lh, rh := 1, 1
	left, rigth := root.Left, root.Right
	// 左数高
	for left != nil {
		lh++
		left = left.Left
	}
	for rigth != nil {
		rh++
		rigth = rigth.Right
	}
	if lh == rh {
		// 满树
		return int(math.Pow(2, float64(lh))) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
