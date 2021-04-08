package palindrome_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	ori := genNode([]int{1, 2, 3, 4})
	assert.EqualValues(t, genNode([]int{4, 3, 2, 1}), reverse(ori))
}

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, false, isPalindrome1(genNode([]int{1, 2, 3, 4})))
	assert.Equal(t, true, isPalindrome1(genNode([]int{1, 2, 2, 1})))
	assert.Equal(t, true, isPalindrome1(genNode([]int{1, 2, 3, 2, 1})))
}
func TestIsPalindrome2(t *testing.T) {
	assert.Equal(t, false, isPalindrome2(genNode([]int{1, 2, 3, 4})))
	assert.Equal(t, true, isPalindrome2(genNode([]int{1, 2, 2, 1})))
	assert.Equal(t, true, isPalindrome2(genNode([]int{1, 2, 3, 2, 1})))
}
func TestIsPalindrome3(t *testing.T) {
	assert.Equal(t, false, isPalindrome3(genNode([]int{1, 2, 3, 4})))
	assert.Equal(t, true, isPalindrome3(genNode([]int{1, 2, 2, 1})))
	assert.Equal(t, true, isPalindrome3(genNode([]int{1, 2, 3, 2, 1})))
}

func genNode(nums []int) *ListNode {
	root := &ListNode{Val: nums[0]}
	node := root
	for _, num := range nums[1:] {
		cur := &ListNode{Val: num}
		node.Next = cur
		node = cur
	}
	return root
}
