package double_pointer

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/linked-list-cycle
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
