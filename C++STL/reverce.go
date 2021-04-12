package main

import (
	"fmt"
	"sort"
)

func isReverce(data []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	return data
}

func main() {
	s := make([]int, 0)
	s = append(s, 2, 3, 4, 5)
	fmt.Println(s)
	fmt.Println(isReverce(s))
}
