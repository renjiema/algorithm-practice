package union_find

type UF struct {
	// 连通分量
	count int
	// 记录父节点，初始化父节点为自己
	parent []int
	// 优化：平衡性,记录节点数
	size []int
}

func NewUF(n int) *UF {
	uf := &UF{count: n}
	for i := 0; i < n; i++ {
		uf.parent = append(uf.parent, i)
		uf.size = append(uf.size, 1)
	}
	return uf
}

// Union 将 p 和 q 连接
func (uf *UF) Union(p, q int) {
	pr := uf.find(p)
	qr := uf.find(q)
	if pr == qr {
		return
	}
	// 平衡性优化
	if uf.size[pr] >= uf.size[qr] {
		uf.parent[qr] = pr
		uf.size[pr] += uf.size[qr]
	} else {
		uf.parent[pr] = qr
		uf.size[qr] += uf.size[pr]
	}
	// 优化前：uf.parent[pr] = qr
	uf.count--
}

// Connected 判断 p 和 q 是否连通
func (uf *UF) Connected(p, q int) bool {
	pr := uf.find(p)
	qr := uf.find(q)
	return pr == qr
}

// Count 返回图中有多少个连通分量
func (uf *UF) Count() int {
	return uf.count
}

// 查找根节点
func (uf UF) find(p int) int {
	for p != uf.parent[p] {
		// 优化：路径压缩
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}
	return p
}
