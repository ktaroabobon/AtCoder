package main

import "fmt"

func iswap(v1, v2 int) (int, int) {
	return v2, v1
}

func main() {
	a := 10
	b := 20

	fmt.Println(iswap(a, b))
}
