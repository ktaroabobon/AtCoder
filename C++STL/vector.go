package main

import "fmt"

func main() {
	s1 := make([]int, 3)
	fmt.Println(s1)
	s1 = append(s1, 10)
	fmt.Println(s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)

	s2 := make([]int, 3)
	fmt.Println(s2)
	s2 = append(s2, s1...)
	fmt.Println(s2)

	fmt.Println(s2[0])
	fmt.Println(s2[len(s2)-1])

}
