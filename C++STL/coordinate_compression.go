package main

import "sort"

// 座標圧縮
// 練習問題：https://atcoder.jp/contests/abc113/tasks/abc113_c

var dist map[int]bool

func lowerBound(intTarget []int, x int) (returnIndex int) {
	returnIndex = sort.Search(len(intTarget), func(i int) bool { return intTarget[i] >= x })
	return
}

func main() {
	x := []int{10, 23, 4, 42, 58, 19, 3, 98}

	var vals []int

	for _, v := range x {
		if dist[v] {
			continue
		}
		dist[v] = true
		vals = append(vals, v)
	}

	sort.Ints(vals)
}
