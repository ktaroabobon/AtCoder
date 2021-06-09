package main

type UnionFind struct {
	par []int
}

func NewUnionFind(N int) *UnionFind {
	uf := new(UnionFind)
	uf.par = make([]int, N)
	for i := range uf.par {
		uf.par[i] = i
	}
	return uf
}

func (u UnionFind) root(x int) int {
	if u.par[x] == x {
		return x
	}
	u.par[x] = u.root(u.par[x])
	return u.par[x]
}

func (u UnionFind) unite(x, y int) {
	xr := u.root(x)
	yr := u.root(y)

	if xr == yr {
		return
	}
	u.par[xr] = yr
}

func (u UnionFind) same(x, y int) bool {
	rx := u.root(x)
	ry := u.root(y)

	return rx == ry
}
