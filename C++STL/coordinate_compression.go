package main

import "sort"

func lowerBound(intTarget []int, x int) (returnIndex int) {
	returnIndex = sort.Search(len(intTarget), func(i int) bool { return intTarget[i] >= x })
	return
}

func main() {
	x := []int{10, 23, 4, 42, 58, 19, 3, 98}

	vals := sort.Ints(x)

}
