package main

import "fmt"

func igcs(v1, v2 int) int {
	if v1 > v2 {
		v1, v2 = v2, v1
	}
	for v1 != 0 {
		v1, v2 = v2%v1, v1
	}
	return v2
}

func main() {
	a := 426480
	b := 971735
	fmt.Println(igcs(a, b))
}
