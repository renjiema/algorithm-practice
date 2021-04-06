package maximum_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/maximum-binary-tree/
// 使用递归，前序遍历
func constructMaximumBinaryTree(nums []int) *TreeNode {
	n := len(nums)
	if n < 1 {
		return nil
	}
	// 查找最大值
	index, max := 0, nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	root := &TreeNode{Val: max}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

/**使用迭代：
1.设第一个数为root
2.后数小于root，设为右树
3.后数大于root，设为root节点，原树设为左数
*/
func constructMaximumBinaryTree2(nums []int) *TreeNode {
	n := len(nums)
	if n < 1 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	for i := 1; i < n; i++ {
		root = insert(root, nums[i])
	}
	return root
}

// 插入节点
func insert(root *TreeNode, val int) *TreeNode {
	if val < root.Val {
		insertRight(root, val)
		return root
	}
	return reRoot(root, val)
}

// 插入右树
func insertRight(root *TreeNode, val int) {
	if root.Right == nil {
		root.Right = &TreeNode{Val: val}
	} else {
		root.Right = insert(root.Right, val)
	}
}

// 重新插入root
func reRoot(node *TreeNode, val int) *TreeNode {
	root := &TreeNode{Val: val}
	root.Left = node
	return root
}
