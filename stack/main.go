package main

import "fmt"

type stack []int

func (stack *stack) push(data int) {
	*stack = append(*stack, data)
}

func main() {
	var s stack = make([]int, 0)
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s)
}
