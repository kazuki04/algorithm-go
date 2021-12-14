package main

import "fmt"

type Node struct {
	data      int
	next_node *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) append(data int) {
	new_node := &Node{data: data, next_node: nil}

	if l.head == nil {
		l.head = new_node
		return
	}

	last_node := l.head
	for last_node.next_node != nil {
		last_node = last_node.next_node
	}
	last_node.next_node = new_node
}

func main() {
	l := LinkedList{}
	l.append(1)
	fmt.Println(l.head.data)
}
