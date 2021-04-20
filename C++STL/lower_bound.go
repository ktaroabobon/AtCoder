package main

import (
	"fmt"
	"sort"
)

// 二分探索の関数
// AtCoderの問題はABC143D("https://atcoder.jp/contests/abc143/tasks/abc143_d")

func lowerBound(intTarget []int, x int) (returnIndex int, f bool) {
	returnIndex = sort.Search(len(intTarget), func(i int) bool { return intTarget[i] >= x })
	if returnIndex < len(intTarget) && intTarget[returnIndex] == x {
		f = true
	} else {
		f = false
		x = len(intTarget)
	}

	return
}
func upperBound(intTarget []int, x int) (returnIndex int, f bool) {
	returnIndex = sort.Search(len(intTarget), func(i int) bool { return intTarget[i] > x }) - 1
	if returnIndex < len(intTarget) && intTarget[returnIndex] == x {
		f = true
	} else {
		f = false
		x = len(intTarget)
	}

	return
}

func main() {
	a := []int{1, 3, 6, 6, 6, 6, 6, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i1, ok := lowerBound(a, x)

	if ok {
		fmt.Printf("found %d at index %d in %v\n", x, i1, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}

	i2, ok := upperBound(a, x)

	if ok {
		fmt.Printf("found %d at index %d in %v\n", x, i2, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}
