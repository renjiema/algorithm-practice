package minimum_depth_of_binary_tree

import "testing"

func TestMinDepth2(t *testing.T) {
	root := genNode(0, []int{2, 0, 3, 0, 4, 0, 5, 0, 6})
	res := minDepth(root)
	t.Logf("min depth:%v", res)
}

func genNode(i int, path []int) *TreeNode {
	t := &TreeNode{Val: path[i]}
	if i < len(path) && 2*i+1 < len(path) && path[2*i+1] != 0 {
		t.Left = genNode(2*i+1, path)
	}
	if i < len(path) && 2*i+2 < len(path) && path[2*i+2] != 0 {
		t.Right = genNode(2*i+2, path)
	}
	return t
}
