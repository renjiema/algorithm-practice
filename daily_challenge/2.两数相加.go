/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, node *ListNode
	nextVal := 0
	for l1 != nil || l2 != nil {
		left, right := 0, 0
		if l1 != nil {
			left = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			right = l2.Val
			l2 = l2.Next
		}
		sum := nextVal + left + right
		sum, nextVal = sum%10, sum/10
		if res == nil {
			res = &ListNode{Val: sum}
			node = res
		} else {
			node.Next = &ListNode{Val: sum}
			node = node.Next
		}
	}
	if nextVal > 0 {
		node.Next = &ListNode{Val: nextVal}
	}
	return res
}

// @lc code=end

