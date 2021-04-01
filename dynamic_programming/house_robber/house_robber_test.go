package house_robber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试house robber案例一
func TestRob_1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := 4
	assert.Equal(t, expected, rob_1(nums))
	assert.Equal(t, expected, rob_1_2(nums))
	assert.Equal(t, expected, rob_1_3(nums))
	nums = []int{2, 7, 9, 3, 1}
	expected = 12
	assert.Equal(t, expected, rob_1(nums))
	assert.Equal(t, expected, rob_1_2(nums))
	assert.Equal(t, expected, rob_1_3(nums))
	nums = []int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}
	assert.Equal(t, rob_1_2(nums), rob_1_3(nums))
}
func TestRob_2(t *testing.T) {
	nums := []int{2, 3, 2}
	expected := 3
	assert.Equal(t, expected, rob_2(nums))
	nums = []int{1, 2, 3, 1}
	expected = 4
	assert.Equal(t, expected, rob_2(nums))
	nums = []int{0}
	expected = 0
	assert.Equal(t, expected, rob_2(nums))
}

func TestRob_3(t *testing.T) {
	expected := 7
	assert.Equal(t, expected, rob(genNode(0, []int{3, 2, 3, 0, 3, 0, 1})))
	assert.Equal(t, expected, rob_3_2(genNode(0, []int{3, 2, 3, 0, 3, 0, 1})))
	expected = 9
	assert.Equal(t, expected, rob(genNode(0, []int{3, 4, 5, 1, 3, 0, 1})))
	assert.Equal(t, expected, rob_3_2(genNode(0, []int{3, 4, 5, 1, 3, 0, 1})))
}

// 生成树节点
func genNode(i int, path []int) *TreeNode {
	t := &TreeNode{Val: path[i]}
	k := 0
	for j := i - 1; j > 0; j = j - 2 {
		if path[j] == 0 {
			k++
		}
	}
	li := 2*(i-k) + 1
	ri := 2*(i-k) + 2
	if i < len(path) && li < len(path) && path[li] != 0 {
		t.Left = genNode(li, path)
	}

	if i < len(path) && ri < len(path) && path[ri] != 0 {
		t.Right = genNode(ri, path)
	}
	return t
}
