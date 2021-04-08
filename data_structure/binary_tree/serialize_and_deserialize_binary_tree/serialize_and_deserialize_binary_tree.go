package serialize_and_deserialize_binary_tree

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const SEP = ","
const NULL = "#"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 前序遍历
	if root == nil {
		return NULL
	}
	ser := fmt.Sprintf("%d,%s,%s", root.Val, this.serialize(root.Left), this.serialize(root.Right))
	return ser
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, SEP)
	return deserializeNodes(&nodes)
}

func deserializeNodes(nodes *[]string) *TreeNode {
	if len(*nodes) < 1 {
		return nil
	}
	// root
	val := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if val == NULL {
		return nil
	}
	intVal, _ := strconv.Atoi(val)
	root := &TreeNode{Val: intVal}
	root.Left = deserializeNodes(nodes)
	root.Right = deserializeNodes(nodes)
	return root
}
