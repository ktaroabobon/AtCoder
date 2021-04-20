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

func designatedLowerBound(intTarget []int, x int) (returnIndex int, f bool) {
	returnIndex = lowerBound(intTarget, x)
	if returnIndex < len(intTarget) && intTarget[returnIndex] == x {
		f = true
	} else {
		f = false
		returnIndex = -1
	}
	return
}
func upperBound(intTarget []int, x int) (returnIndex int) {
	returnIndex = sort.Search(len(intTarget), func(i int) bool { return intTarget[i] > x }) - 1
	return
}

func designatedUpperBound(intTarget []int, x int) (returnIndex int, f bool) {
	returnIndex = upperBound(intTarget, x)
	if returnIndex < len(intTarget) && intTarget[returnIndex] == x {
		f = true
	} else {
		f = false
		returnIndex = -1
	}
	return
}

func main() {
	a := []int{1, 3, 6, 6, 6, 6, 6, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 7

	i11 := lowerBound(a, x)
	fmt.Printf("found %d at index %d in %v\n", x, i11, a)

	i1, ok := designatedLowerBound(a, x)
	if ok {
		fmt.Printf("found %d at index %d in %v\n", x, i1, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}

	i22 := upperBound(a, x)
	fmt.Printf("found %d at index %d in %v\n", x, i22, a)

	i2, ok := designatedUpperBound(a, x)

	if ok {
		fmt.Printf("found %d at index %d in %v\n", x, i2, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}
