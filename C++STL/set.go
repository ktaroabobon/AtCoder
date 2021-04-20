package main

import "fmt"

type set map[interface{}]bool

func (h set) add(s interface{})         { h[s] = true }
func (h set) erase(s interface{})       { delete(h, s) }
func (h set) exists(s interface{}) bool { return h[s] }
func (h *set) clear()                   { *h = make(set) }

func main() {
	set1 := set{}
	fmt.Println(set1)

	set1.add("A")
	set1.add("B")
	set1.add("C")
	fmt.Println(set1)

	set1.erase("A")
	fmt.Println(set1)

	fmt.Println(set1.exists("B"))
	fmt.Println(set1.exists("D"))

	set1.clear()
	fmt.Println(set1)

	set2 := set{}
	set2.add("D")
	set2.add("E")
	set2.add("F")

}
