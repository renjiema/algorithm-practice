package insert_delete_getrandom_o1

import "math/rand"

// https://leetcode-cn.com/problems/insert-delete-getrandom-o1
type RandomizedSet struct {
	values []int
	indexs map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{values: []int{}, indexs: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexs[val]; ok {
		return false
	}
	this.indexs[val] = len(this.values)
	this.values = append(this.values, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.indexs[val]
	if !ok {
		return false
	}
	n := len(this.values) - 1
	v := this.values[n]
	this.values[index], this.values[n] = v, this.values[index]
	this.values = this.values[:n]
	this.indexs[v] = index
	delete(this.indexs, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Int() % len(this.values)
	return this.values[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
