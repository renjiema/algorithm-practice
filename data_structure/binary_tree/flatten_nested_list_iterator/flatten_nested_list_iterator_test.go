package flatten_nested_list_iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNestedInteger(t *testing.T) {
	res := NewNestedInteger([]interface{}{[]interface{}{1, []interface{}{2, 4}}, 2, []interface{}{1, 1}})
	t.Log(res)
}
func TestNestedIterator(t *testing.T) {
	n := NewNestedInteger([]interface{}{[]interface{}{1, 1}, 2, []interface{}{1, 1}})
	iter := Constructor(n)
	res := []int{}
	for iter.HasNext() {
		res = append(res, iter.Next())
	}
	assert.EqualValues(t, []int{1, 1, 2, 1, 1}, res)
	n = NewNestedInteger([]interface{}{1, []interface{}{4, []interface{}{6}}})
	iter = Constructor(n)
	res = []int{}
	for iter.HasNext() {
		res = append(res, iter.Next())
	}
	assert.EqualValues(t, []int{1, 4, 6}, res)
}
