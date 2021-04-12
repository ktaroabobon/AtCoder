package main

import (
	"container/heap"
	"fmt"
)

//例題
//ACB141-D

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	ih := &intHeap{}

	heap.Init(ih)
	heap.Push(ih, 3)
	heap.Push(ih, 1)
	heap.Push(ih, 10)
	heap.Push(ih, 7)

	for ih.Len() > 0 {
		v := heap.Pop(ih)
		fmt.Println(v)
	}
}
