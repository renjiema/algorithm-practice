package reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	if head == nil {
		return head
	}
	dummyNode := &ListNode{Next: head}
	a, b := dummyNode.Next, dummyNode.Next
	for i := 0; i < k; i++ {
		if b == nil {
			return dummyNode.Next
		}
		b = b.Next
	}
	// 翻转[1,k]
	cur := dummyNode.Next
	next := dummyNode.Next
	for i := 0; i < k-1; i++ {
		next = cur.Next
		cur.Next = next.Next
		next.Next = dummyNode.Next
		dummyNode.Next = next
	}
	// 递归处理下一个k
	a.Next = reverseKGroup(b, k)
	return next
}
