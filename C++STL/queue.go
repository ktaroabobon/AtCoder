package main

import "fmt"

func main() {
	q := make([]int, 0)

	q = append(q, 1)
	q = append(q, 2)
	q = append(q, 3)

	v1 := q[0]
	q = q[1:]
	fmt.Println(v1)

	v2 := q[0]
	q = q[1:]
	fmt.Println(v2)
}
