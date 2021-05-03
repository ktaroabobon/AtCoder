package main

import "fmt"

type udGraph = [][]int

type Edge struct {
	to     int
	weight int
}

type wGraph = [][]Edge

func main() {
	g := make(udGraph, 3)

	g[0] = []int{1}
	g[1] = []int{2, 3}
	g = append(g, []int{3, 6, 9})

	fmt.Println(g)

	wg := make(wGraph, 4)

	wg[1] = []Edge{{2, 4}}
	wg[0] = []Edge{{3, 4}, {2, 5}}

	fmt.Println(wg)
}
