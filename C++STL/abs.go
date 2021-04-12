package main

import "fmt"

func iabs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
func main() {
	a := -65
	fmt.Println(iabs(a))
}
