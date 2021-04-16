package double_pointer

// https://leetcode-cn.com/problems/linked-list-cycle-ii
// 找环节点解析：慢节点走了k步，快节点走了2k步，设环节点到相遇节点为m步，那么head到环节点相距k-m步，相遇节点再走k-m步的位置一定是环节点。
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
