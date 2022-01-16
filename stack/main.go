package main

import "fmt"

type Stack struct {
	stack *[]int
}

func (s *Stack) push(data int) {
	var stack = s.stack
	*stack = append(*stack, data)
}

func main() {
	stack := make([]int, 0)
	s := Stack{&stack}
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(stack)
}
