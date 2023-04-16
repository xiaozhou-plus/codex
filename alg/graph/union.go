package graph

type UF struct {
	// 指向的父节点，一开始默认指向自己
	parent []int
	// 联通分量的个数
	count int
}

func NewUF(n int) *UF {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UF{parent, n}
}

func (u *UF) Union(p, q int) bool {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return false
	}

	u.parent[rootQ] = rootP
	u.count--
	return true
}

func (u *UF) Find(x int) int {
	// 路径压缩， 最终都指向跟节点
	if x != u.parent[x] {
		u.parent[x] = u.Find(u.parent[x])
	}

	return u.parent[x]
}

func (u *UF) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UF) Count() int {
	return u.count
}
