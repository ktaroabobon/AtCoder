package main

import "fmt"

type intStack []int

func (stack *intStack) push(i int) {
	*stack = append(*stack, i)
}

func (stack *intStack) pop() (result int) {
	result = (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return
}

func main() {
	var stack intStack = make([]int, 0)
	stack.push(4)
	fmt.Println(stack)
	stack.push(3)
	fmt.Println(stack.pop())
	stack.push(8)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
