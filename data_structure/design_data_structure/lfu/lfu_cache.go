package lfu

// https://leetcode-cn.com/problems/lfu-cache
// 也可通过双向链表实现，node中存放freq
type LFUCache struct {
	// 存放key,value
	kv map[int]int
	// 存放key,freq
	kf map[int]int
	// 存放freq,keys
	fks map[int][]int
	// 最小freq
	minFreq int
	cap     int
	size    int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{kv: make(map[int]int), kf: make(map[int]int), fks: make(map[int][]int), cap: capacity}
}

func (this *LFUCache) Get(key int) int {
	value, ok := this.kv[key]
	if !ok {
		return -1
	}
	this.addFrequently(key)
	return value
}

func (this *LFUCache) Put(key int, value int) {
	// 判断是否存在
	if _, ok := this.kv[key]; ok {
		//覆盖值
		this.kv[key] = value
		this.addFrequently(key)
		return
	}
	// 判断是否超出capacity
	if this.size == this.cap {
		// 删除最不常使用的key
		this.removeInfrequently()
	} else {
		this.size++
	}
	// 先添加
	this.kv[key] = value
	this.kf[key]++
	this.minFreq = 1
	this.fks[this.minFreq] = append(this.fks[this.minFreq], key)

}

func (this *LFUCache) addFrequently(key int) {
	freq := this.kf[key]
	// 删除原fks中的值
	for i, k := range this.fks[freq] {
		if k == key {
			this.fks[freq] = append(this.fks[freq][0:i], this.fks[freq][i+1:]...)
			break
		}
	}
	if freq == this.minFreq && len(this.fks[freq]) < 1 {
		// 唯一一个最小
		this.minFreq++
	}
	this.kf[key]++
	freq++
	this.fks[freq] = append(this.fks[freq], key)
}

func (this *LFUCache) removeInfrequently() {
	key := this.fks[this.minFreq][0]
	this.fks[this.minFreq] = this.fks[this.minFreq][1:]
	if len(this.fks[this.minFreq]) == 0 {
		delete(this.fks, this.minFreq)
	}
	delete(this.kv, key)
	delete(this.kf, key)
}
