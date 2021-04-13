package lru

// https://leetcode-cn.com/problems/lru-cache/
type LRUCache struct {
	cm  map[int]*Node
	dl  *DoubleList
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cm: make(map[int]*Node), dl: NewDoubleList(), cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cm[key]
	if !ok {
		return -1
	}
	// 将node放到最后
	this.makeRecently(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	// 判断是否存在
	if node, ok := this.cm[key]; ok {
		node.value = value
		this.makeRecently(node)
		return
	}
	// 不存在，判断容量
	if this.cap <= len(this.cm) {
		// 容量不足
		this.deleteLast()
	}
	this.add(key, value)
}

func (this *LRUCache) makeRecently(node *Node) {
	this.dl.Remove(node)
	this.dl.AddLast(node)
}

func (this *LRUCache) deleteLast() {
	first := this.dl.RemoveFirst()
	delete(this.cm, first.key)
}

func (this *LRUCache) add(key, value int) {
	node := NewNode(key, value)
	this.cm[key] = node
	this.dl.AddLast(node)
}

// Node 双向链表节点类型
type Node struct {
	key, value int
	prev, next *Node
}

func NewNode(k, v int) *Node {
	return &Node{key: k, value: v}
}

// DoubleList 双向链表
type DoubleList struct {
	// 首尾虚节点
	head, tail *Node
	// 链表元素个数
	size int
}

func NewDoubleList() *DoubleList {
	dl := &DoubleList{}
	dl.head = NewNode(0, 0)
	dl.tail = NewNode(0, 0)
	dl.head.next = dl.tail
	dl.tail.prev = dl.head
	return dl
}

// AddLast
func (dl *DoubleList) AddLast(node *Node) {
	node.prev = dl.tail.prev
	node.next = dl.tail
	dl.tail.prev.next = node
	dl.tail.prev = node
	dl.size++
}

// Remove
func (dl *DoubleList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	dl.size--
}

// RemoveFirst
func (dl *DoubleList) RemoveFirst() *Node {
	if dl.head.next == dl.tail {
		return nil
	}
	// get first
	first := dl.head.next
	dl.Remove(first)
	return first
}
