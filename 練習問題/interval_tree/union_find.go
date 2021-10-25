package main

// 参考URL: https://qiita.com/haru1843/items/2295d0ec1f5002bd5c33
// 参考URL: https://atcoder.jp/contests/abc217/submissions/25627954

type UnionFind struct {
	par  []int // parent numbers
	rank []int // height of tree
	size []int
}

func NewUnionFind(N int) *UnionFind {
	uf := new(UnionFind)
	uf.par = make([]int, N)
	uf.rank = make([]int, N)
	uf.size = make([]int, N)
	for i := range uf.par {
		uf.par[i] = -1
		uf.rank[i] = 0
		uf.size[i] = 1
	}
	return uf
}

func (u UnionFind) InitSize(s []int) {
	for i := 1; i < len(u.size); i++ {
		u.size[i] = s[i] - s[i-1]
	}
}

func (u UnionFind) Root(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.Root(u.par[x])
	return u.par[x]
}

func (u UnionFind) Unite(x, y int) {
	xr := u.Root(x)
	yr := u.Root(y)

	if xr == yr {
		return
	}
	// rank
	if u.rank[xr] < u.rank[yr] {
		u.par[xr] = yr
		u.size[yr] += u.size[xr]
	} else {
		u.par[yr] = xr
		u.size[xr] += u.size[yr]
		if u.rank[xr] == u.rank[yr] {
			u.rank[xr]++
		}
	}
}

func (u UnionFind) Same(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

func (u UnionFind) Size(x int) int {
	return u.size[u.Root(x)]
}

func main() {
	N := 10
	t := *NewUnionFind(N)

}
