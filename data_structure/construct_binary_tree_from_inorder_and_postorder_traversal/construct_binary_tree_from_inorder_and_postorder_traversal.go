package construct_binary_tree_from_inorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	rootIndex := 0
	for i, v := range inorder {
		if v == rootVal {
			rootIndex = i
		}
	}
	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
	return root
}
