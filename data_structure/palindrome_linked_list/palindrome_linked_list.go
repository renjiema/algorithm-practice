package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/palindrome-linked-list/
// 转换为数组，使用双指针
func isPalindrome1(head *ListNode) bool {
	values := conver2Array(head)
	n := len(values)
	for i := 0; i < n/2; i++ {
		if values[i] != values[n-1-i] {
			return false
		}
	}
	return true
}

var left *ListNode

// 使用递归特性获取尾节点
func isPalindrome2(head *ListNode) bool {
	left = head
	return traverse(head)
}

// 通过快慢指针获取链表中点
func isPalindrome3(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 判断链表长度是否为奇数
	if fast != nil {
		slow = slow.Next
	}
	// 翻转slow后面链表
	right := reverse(slow)
	for right != nil {
		if right.Val != head.Val {
			return false
		}
		right = right.Next
		head = head.Next
	}
	return true
}

func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)
	res = res && right.Val == left.Val
	left = left.Next
	return res
}

func conver2Array(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// 翻转链表
func reverse(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
