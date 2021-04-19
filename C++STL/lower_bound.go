package main

import (
	"fmt"
	"sort"
)

// 二分探索の関数
// AtCoderの問題はABC143D("https://atcoder.jp/contests/abc143/tasks/abc143_d")

func lowerBound(intTarget []int, x int) (returnIndex int) {
	returnIndex = sort.Search(len(intTarget), func(i int) bool { return intTarget[i] >= x })
	return
}

func main() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i1 := lowerBound(a, x)

	if i1 < len(a) && a[i1] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i1, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}
