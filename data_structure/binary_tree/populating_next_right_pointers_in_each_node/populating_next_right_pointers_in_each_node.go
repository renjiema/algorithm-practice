package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectHelper(root.Left, root.Right)
	return root
}

func connectHelper(left *Node, right *Node) {
	if left == nil || right == nil || left.Next == right {
		return
	}
	left.Next = right
	connectHelper(left.Left, left.Right)
	connectHelper(right.Left, right.Right)
	connectHelper(left.Right, right.Left)
}
