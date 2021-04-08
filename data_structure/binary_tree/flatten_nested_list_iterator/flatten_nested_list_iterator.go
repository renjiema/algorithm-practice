package flatten_nested_list_iterator

// https://leetcode-cn.com/problems/flatten-nested-list-iterator
type NestedInteger struct {
	value     int
	isInteger bool
	sub       []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return this.isInteger
}

func (this NestedInteger) GetInteger() int {
	return this.value
}

func (this *NestedInteger) SetInteger(value int) {
	this.value = value
}

func (this *NestedInteger) Add(elem NestedInteger) {
	this.sub = append(this.sub, &elem)
}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.sub
}

func NewNestedInteger(values []interface{}) []*NestedInteger {
	var res []*NestedInteger
	for _, v := range values {
		cur := &NestedInteger{}
		switch v.(type) {
		case int:
			cur.SetInteger(v.(int))
			cur.isInteger = true
		case []interface{}:
			cur.isInteger = false
			nexts := NewNestedInteger(v.([]interface{}))
			for _, next := range nexts {
				cur.Add(*next)
			}
		}
		res = append(res, cur)
	}
	return res
}

/**
*一次遍历出所有值
type NestedIterator struct {
	values []int
	len    int
	index  int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	n := &NestedIterator{}
	n.ConstructorHelper(nestedList)
	n.len = len(n.values)
	return n
}
func (this *NestedIterator) ConstructorHelper(nestedList []*NestedInteger) {
	for _, v := range nestedList {
		if v.IsInteger() {
			this.values = append(this.values, v.GetInteger())
		} else {
			this.ConstructorHelper(v.GetList())
		}
	}
}

func (this *NestedIterator) Next() int {
	v := this.values[this.index]
	this.index++
	return v
}

func (this *NestedIterator) HasNext() bool {
	return this.index < this.len
}
*/

/**
// 使用队列
type NestedIterator struct {
	it []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{it: nestedList}
}

func (this *NestedIterator) Next() int {
	n := (this.it[0]).GetInteger()
	this.it = this.it[1:]
	return n
}

func (this *NestedIterator) HasNext() bool {
	// 第一个元素不是Integer，则将第一个Integer压入栈顶
	for len(this.it) > 0 && !this.it[0].IsInteger() {
		first := this.it[0].GetList()
		this.it = this.it[1:]
		for i := len(first) - 1; i >= 0; i-- {
			this.it = append([]*NestedInteger{first[i]}, this.it...)
		}
	}
	return len(this.it) > 0
}
*/
// 使用栈
type NestedIterator struct {
	// 栈中存放队列
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (this *NestedIterator) Next() int {
	queue := this.stack[len(this.stack)-1]
	val := queue[0].GetInteger()
	this.stack[len(this.stack)-1] = queue[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		// 栈顶队列
		queue := this.stack[len(this.stack)-1]
		if len(queue) < 1 {
			//出栈
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}
		n := queue[0]
		if n.IsInteger() {
			return true
		}
		// 将队列压入栈
		this.stack[len(this.stack)-1] = queue[1:]
		this.stack = append(this.stack, n.GetList())
	}
	return false
}
