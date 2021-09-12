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
	if returnIndex == -1 {
		f = false
	} else if returnIndex < len(intTarget) && intTarget[returnIndex] == x {
		f = true
	} else {
		f = false
		returnIndex = -1
	}
	return
}

func main() {
	a := []int{0, 3, 6, 6, 6, 6, 6, 6, 10, 15, 21, 28, 36, 45, 55, 55}
	x := 20

	fmt.Printf("index %d\n", len(a))

	i11 := lowerBound(a, x)
	fmt.Printf("found %d at index %d in %v\n", x, i11, a)

	i1, ok1 := designatedLowerBound(a, x)
	if ok1 {
		fmt.Printf("found %d at index %d in %v\n", x, i1, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}

	i22 := upperBound(a, x)
	fmt.Printf("found %d at index %d in %v\n", x, i22, a)

	i2, ok2 := designatedUpperBound(a, x)

	if ok2 {
		fmt.Printf("found %d at index %d in %v\n", x, i2, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}

	if ok1 && ok2 {
		num := i2 - i1 + 1
		fmt.Printf("%dは%v個ある。\n", x, num)
	} else {
		fmt.Printf("%dはありません。", x)
	}

}
