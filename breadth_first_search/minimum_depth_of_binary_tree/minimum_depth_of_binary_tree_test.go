package minimum_depth_of_binary_tree

import "testing"

func TestMinDepth2(t *testing.T) {
	// root := genNode(0, []int{2, 0, 3, 0, 4, 0, 5, 0, 6})
	root := genNode(0, []int{1, 2, 3, 4, 5})
	res := minDepth(root)
	t.Logf("min depth:%v", res)
	res = minDepth2(root)
	t.Logf("min depth2:%v", res)
}

// 生成树节点
func genNode(i int, path []int) *TreeNode {
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
		t.Left = genNode(li, path)
	}

	if i < len(path) && ri < len(path) && path[ri] != 0 {
		t.Right = genNode(ri, path)
	}
	return t
}
