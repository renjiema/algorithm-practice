package populating_next_right_pointers_in_each_node

import (
	"testing"
)

func TestConnect(t *testing.T) {
	root := connect(genNode(0, []int{1, 2, 3, 4, 5, 6, 7}))
	t.Log(root)
}

func genNode(i int, path []int) *Node {
	t := &Node{Val: path[i]}
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
