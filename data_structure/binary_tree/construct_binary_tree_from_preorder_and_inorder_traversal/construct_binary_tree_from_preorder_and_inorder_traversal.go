package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	rootVal := preorder[0]
	rootIndex := 0
	for i, v := range inorder {
		if rootVal == v {
			rootIndex = i
		}
	}
	// create root
	root := &TreeNode{Val: rootVal}

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}
