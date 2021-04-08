package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	// 左右子树都变成拉平后的右子树
	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = left
	// 将right连接到left后
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
